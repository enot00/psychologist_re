package database

import (
	"github.com/test_server/internal/domain"
	"github.com/upper/db/v4"
	"log"
	"time"
)

const (
	tokensBlacklistTable = "tokens_blacklist"
	refreshTokensTable   = "refresh_tokens"
)

// todo use redis instead postgresql
type (
	TokenRepository interface {
		AddToBlacklist(token *domain.Token) error
		InBlacklist(token string) (bool, error)
		CountRefreshTokensFor(userId int64) (uint64, error)
		SaveRefreshToken(userId int64, token *domain.Token) error
		Refresh(userId int64, new *domain.Token, old *domain.Token) error
		DeleteRefreshToken(token *domain.Token) error
		DeleteOldestRefreshToken(userId int64) error
		FlushAllExpired() error
	}

	tokenRepository struct {
		sess *db.Session
	}

	token struct {
		UserId      int64     `db:"user_id,omitempty"`
		TokenString string    `db:"token"`
		IssuedAt    time.Time `db:"issued_at,omitempty"`
		ExpiresAt   time.Time `db:"expires_at"`
	}
)

func NewTokenRepository(sess *db.Session) TokenRepository {
	return &tokenRepository{sess: sess}
}

func (r *tokenRepository) AddToBlacklist(token *domain.Token) error {
	row := r.mapToBlacklistRow(token)

	err := (*r.sess).Collection(tokensBlacklistTable).InsertReturning(row)
	if err != nil {
		log.Printf("tokenRepository.AddToBlacklist(): %s\n", err.Error())
		return err
	}

	return nil
}

func (r *tokenRepository) InBlacklist(token string) (bool, error) {
	count, err := (*r.sess).Collection(tokensBlacklistTable).
		Find(db.Cond{"token": token}).
		Count()
	if err != nil {
		log.Printf("tokenRepository.InBlacklist(): %s\n", err.Error())
		return false, err
	}

	return count > 0, nil
}

func (r *tokenRepository) CountRefreshTokensFor(userId int64) (uint64, error) {
	count, err := (*r.sess).Collection(refreshTokensTable).
		Find(db.Cond{"user_id": userId}).
		Count()
	if err != nil {
		log.Printf("tokenRepository.CountRefreshTokensFor(): %s\n", err.Error())
		return 0, err
	}

	return count, nil
}

func (r *tokenRepository) SaveRefreshToken(userId int64, token *domain.Token) error {
	row := r.mapToRefreshRow(token)
	row.UserId = userId

	err := (*r.sess).Collection(refreshTokensTable).InsertReturning(row)
	if err != nil {
		log.Printf("tokenRepository.SaveRefreshToken(): %s\n", err.Error())
		return err
	}

	return nil
}

func (r *tokenRepository) Refresh(userId int64, new *domain.Token, old *domain.Token) error {
	err := (*r.sess).Tx(func(sess db.Session) error {
		err := r.DeleteRefreshToken(old)
		if err != nil {
			return err
		}

		err = r.SaveRefreshToken(userId, new)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Printf("tokenRepository.Refresh(): %s\n", err.Error())
		return err
	}

	return nil
}

func (r *tokenRepository) DeleteRefreshToken(token *domain.Token) error {
	err := (*r.sess).Collection(refreshTokensTable).
		Find(db.Cond{"token": token.TokenString}).
		Delete()
	if err != nil {
		log.Printf("tokenRepository.deleteRefreshToken(): %s\n", err.Error())
		return err
	}

	return nil
}

func (r *tokenRepository) DeleteOldestRefreshToken(userId int64) error {
	subQuery := (*r.sess).SQL().SelectFrom(refreshTokensTable).
		Columns("id").
		Where("user_id", userId).
		OrderBy("id").
		Limit(1).
		String()

	_, err := (*r.sess).SQL().DeleteFrom(refreshTokensTable).
		Where("id", db.In(db.Raw(subQuery, userId))).
		Exec()
	if err != nil {
		log.Printf("tokenRepository.DeleteOldestRefreshToken(): %s\n", err.Error())
		return err
	}

	return nil
}

func (r *tokenRepository) FlushAllExpired() error {
	err := (*r.sess).Collection(tokensBlacklistTable).
		Find(db.Cond{"expired_at": db.Before(time.Now())}).
		Delete()
	if err != nil {
		log.Printf("tokenRepository.FlushAllExpired(): %s\n", err.Error())
		return err
	}

	err = (*r.sess).Collection(refreshTokensTable).
		Find(db.Cond{"expired_at": db.Before(time.Now())}).
		Delete()
	if err != nil {
		log.Printf("tokenRepository.FlushAllExpired(): %s\n", err.Error())
		return err
	}

	return nil
}

func (r *tokenRepository) mapToBlacklistRow(t *domain.Token) *token {
	return &token{
		TokenString: t.TokenString,
		ExpiresAt:   t.ExpiresAt,
	}
}

func (r *tokenRepository) mapToRefreshRow(t *domain.Token) *token {
	return &token{
		TokenString: t.TokenString,
		IssuedAt:    t.IssuedAt,
		ExpiresAt:   t.ExpiresAt,
	}
}
