package domain

import "time"

type (
	Tokens struct {
		Access  Token
		Refresh Token
	}

	Token struct {
		TokenString string
		IssuedAt    time.Time
		ExpiresAt   time.Time
	}
)
