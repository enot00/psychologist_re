package app

import (
	"github.com/test_server/internal/domain"
	appErrors "github.com/test_server/internal/errors"
)

type (
	AuthenticationService interface {
		SignUp(user *domain.User) (*domain.User, *domain.Tokens, error)
		Login(email, password string) (*domain.User, *domain.Tokens, error)
		Logout(access domain.Token, refresh domain.Token) error
		RefreshTokens(token string) (*domain.Tokens, error)
	}

	authenticationService struct {
		us *UserService
		ts *TokenService
		hs *HashService
	}
)

func NewAuthenticationService(us *UserService, ts *TokenService, hs *HashService) AuthenticationService {
	return &authenticationService{us: us, ts: ts, hs: hs}
}

func (s *authenticationService) SignUp(user *domain.User) (*domain.User, *domain.Tokens, error) {
	saved, err := (*s.us).Save(user)
	if err != nil {
		return nil, nil, err

	}

	tokens, err := (*s.ts).GenerateAndSave(saved.ID, int(saved.Role))
	if err != nil {
		return nil, nil, err
	}

	return saved, tokens, nil
}

func (s *authenticationService) Login(email, password string) (*domain.User, *domain.Tokens, error) {
	user, err := (*s.us).GetOneByEmail(email)
	if err != nil {
		return nil, nil, err
	}

	if ok := (*s.hs).IsEqual(user.Password, password); !ok {
		return nil, nil, appErrors.ErrAuthenticationFailed
	}

	tokens, err := (*s.ts).GenerateAndSave(user.ID, int(user.Role))
	if err != nil {
		return nil, nil, err
	}

	return user, tokens, nil
}

func (s *authenticationService) Logout(access domain.Token, refresh domain.Token) error {
	return (*s.ts).Deactivate(domain.Tokens{Access: access, Refresh: refresh})
}

func (s *authenticationService) RefreshTokens(oldRefresh string) (*domain.Tokens, error) {
	claims, err := (*s.ts).Validate(oldRefresh)
	if err != nil {
		return nil, err
	}

	user, err := (*s.us).GetOneByID(claims.UserId)
	if err != nil {
		return nil, err
	}

	if user.IsDeleted() {
		return nil, err
	}

	tokens, err := (*s.ts).Refresh(user.ID, int(user.Role), oldRefresh)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
