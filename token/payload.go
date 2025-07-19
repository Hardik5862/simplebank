package token

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

var (
	ErrTokenExpired = errors.New("token is expired")
	ErrTokenInvalid = errors.New("token is invalid")
)

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (p Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(p.ExpiredAt), nil
}

func (p Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(p.IssuedAt), nil
}

func (p Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil
}

func (p Payload) GetIssuer() (string, error) {
	return "", nil
}

func (p Payload) GetSubject() (string, error) {
	return p.Username, nil
}

func (p Payload) GetAudience() (jwt.ClaimStrings, error) {
	return nil, nil
}
