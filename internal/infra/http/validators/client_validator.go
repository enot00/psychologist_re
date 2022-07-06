package validators

import (
	"encoding/json"
	"log"
	"net/http"
	"unicode"

	"github.com/go-playground/validator"
	"github.com/test_server/internal/domain"
)

type ClientValidator struct {
	validator *validator.Validate
}

func NewClientValidator() *ClientValidator {
	return &ClientValidator{
		validator: validator.New(),
	}
}

func validatePassword(Password validator.FieldLevel) bool {
	/*
		min len: 8, max len: 20
		at least one: lower case, upper case, number
		contains only: letters, numbers & ('-', '_')
	*/
	pswd := Password.Field().String()
	var lowerExists, upperExists, numberExists bool
	correctCharsCount := 0

	for _, c := range pswd {
		switch {
		case unicode.IsNumber(c):
			numberExists = true
			correctCharsCount++
		case unicode.IsLower(c):
			lowerExists = true
			correctCharsCount++
		case unicode.IsUpper(c):
			upperExists = true
			correctCharsCount++
		case string(c) == "-" || string(c) == "_":
			correctCharsCount++
		}
	}

	correctChars := len([]rune(pswd)) == correctCharsCount
	correctLen := correctCharsCount >= 8 && correctCharsCount <= 20
	return correctChars && correctLen && lowerExists && upperExists && numberExists
}

func (c ClientValidator) ValidateAndMap(r *http.Request) (*domain.Client, error) {
	var clientResource clientRequest
	err := json.NewDecoder(r.Body).Decode(&clientResource)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	c.validator.RegisterValidation("custom_validator", validatePassword)
	err = c.validator.Struct(clientResource)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return mapClientRequestToDomain(&clientResource), nil
}
