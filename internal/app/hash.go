package app

import (
	"golang.org/x/crypto/bcrypt"
)

type (
	HashService interface {
		Hash(str string) string
		IsEqual(hash, str string) bool
	}

	hashService struct {
		salt string
	}
)

func NewHashService(salt string) HashService {
	return &hashService{salt: salt}
}

func (s *hashService) Hash(str string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(str+s.salt), 14)
	return string(bytes)
}

func (s *hashService) IsEqual(hash, str string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str+s.salt))
	if err != nil {
		return false
	}

	return true
}
