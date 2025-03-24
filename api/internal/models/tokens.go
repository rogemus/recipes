package models

import (
	"crypto/rand"
	"crypto/sha256"
	"time"

	"recipes.krogowski.dev/internal/validator"
)

const (
	ScopeActivation     = "activation"
	ScopeAuthentication = "authentication"
)

type Token struct {
	Hash      []byte    `json:"-"`
	PlainText string    `json:"token"`
	UserID    int64     `json:"-"`
	Expiry    time.Time `json:"expiry"`
	Scope     string    `json:"-"`
}

func GenerateToken(userID int64, ttl time.Duration, scope string) *Token {
	token := &Token{
		PlainText: rand.Text(),
		UserID:    userID,
		Expiry:    time.Now().Add(ttl),
		Scope:     scope,
	}

	hash := sha256.Sum256([]byte(token.PlainText))
	token.Hash = hash[:]

	return token
}

func ValidateTokenPlaintext(v *validator.Validator, tokenPlaintext string) {
	v.Check(tokenPlaintext != "", "token", "must be provided")
	v.Check(len(tokenPlaintext) == 26, "token", "must be 26 bytes long")
}
