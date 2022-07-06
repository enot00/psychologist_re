package app

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/test_server/internal/domain"
	appErrors "github.com/test_server/internal/errors"
	"github.com/test_server/internal/infra/database"
	"log"
	"time"
)

const (
	maxAvailableRefreshTokens = 5
)

type (
	TokenService interface {
		Generate(userId int64, role int) (*domain.Tokens, error)
		GenerateAndSave(userId int64, role int) (*domain.Tokens, error)
		Refresh(userId int64, role int, old string) (*domain.Tokens, error)
		Validate(t string) (*TokenClaims, error)
		Deactivate(tokens domain.Tokens) error
	}

	tokenService struct {
		tr        *database.TokenRepository
		secretKey string
		tokenTTL  time.Duration
	}

	TokenClaims struct {
		UserId int64 `json:"user_id"`
		Role   int   `json:"role"`
		jwt.RegisteredClaims
	}
)

func NewTokenService(tr *database.TokenRepository, secretKey string, tokenTTL time.Duration) TokenService {
	return &tokenService{tr: tr, secretKey: secretKey, tokenTTL: tokenTTL}
}

func (s *tokenService) Generate(userId int64, role int) (*domain.Tokens, error) {
	tokens := domain.Tokens{
		Access: domain.Token{
			IssuedAt:  time.Now(),
			ExpiresAt: time.Now().Add(s.tokenTTL),
		},
		Refresh: domain.Token{
			IssuedAt:  time.Now(),
			ExpiresAt: time.Now().Add(s.tokenTTL).Add(240 * time.Hour),
		},
	}

	access, err := s.create(TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tokens.Access.ExpiresAt),
			IssuedAt:  jwt.NewNumericDate(tokens.Access.IssuedAt),
		},
		UserId: userId,
		Role:   role,
	})
	if err != nil {
		return nil, err
	}

	refresh, err := s.create(TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tokens.Refresh.ExpiresAt),
			IssuedAt:  jwt.NewNumericDate(tokens.Refresh.IssuedAt),
		},
		UserId: userId,
		Role:   role,
	})
	if err != nil {
		return nil, err
	}

	tokens.Access.TokenString = access
	tokens.Refresh.TokenString = refresh

	return &tokens, nil
}

func (s *tokenService) GenerateAndSave(userId int64, role int) (*domain.Tokens, error) {
	tokens, err := s.Generate(userId, role)
	if err != nil {
		return nil, err
	}

	err = s.saveRefreshToken(userId, &tokens.Refresh)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (s *tokenService) Refresh(userId int64, role int, old string) (*domain.Tokens, error) {
	tokens, err := s.Generate(userId, role)
	if err != nil {
		return nil, err
	}

	err = (*s.tr).Refresh(userId, &tokens.Refresh, &domain.Token{TokenString: old})
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (s *tokenService) Validate(t string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(t, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.secretKey), nil
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, appErrors.ErrAuthenticationFailed
}

func (s *tokenService) Deactivate(tokens domain.Tokens) error {
	err := (*s.tr).AddToBlacklist(&tokens.Access)
	if err != nil {
		return err
	}

	return (*s.tr).DeleteRefreshToken(&tokens.Refresh)
}

func (s *tokenService) saveRefreshToken(userId int64, token *domain.Token) error {
	count, err := (*s.tr).CountRefreshTokensFor(userId)
	if err != nil {
		return err
	}

	if count == maxAvailableRefreshTokens {
		err = (*s.tr).DeleteOldestRefreshToken(userId)
		if err != nil {
			return err
		}
	}

	return (*s.tr).SaveRefreshToken(userId, token)
}

func (s *tokenService) create(claims TokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString, nil
}
