package models

import (
	"crypto/rand"
	"crypto/sha256"
	"time"
)

const (
	ScopeActivation    = "activation"
	ScopeAuthorization = "authorization"
)

type Token struct {
	Hash      []byte
	PlainText string
	UserID    int64
	Expiry    time.Time
	Scope     string
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
