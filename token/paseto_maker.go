package token

import (
	"errors"
	"fmt"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/google/uuid"
)

type PasetoMaker struct {
	symmetricKey paseto.V4SymmetricKey
	parser       paseto.Parser
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != 32 {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", 32)
	}

	key, err := paseto.V4SymmetricKeyFromBytes([]byte(symmetricKey))
	if err != nil {
		return nil, fmt.Errorf("invalid symmetric key: %w", err)
	}

	parser := paseto.NewParser()

	maker := &PasetoMaker{
		symmetricKey: key,
		parser:       parser,
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", nil, err
	}

	token := paseto.NewToken()

	token.SetIssuedAt(payload.IssuedAt)
	token.SetExpiration(payload.ExpiredAt)
	token.SetJti(payload.ID.String())

	token.SetString("username", payload.Username)

	encrypted := token.V4Encrypt(maker.symmetricKey, nil)

	return encrypted, payload, nil
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	parsedToken, err := maker.parser.ParseV4Local(maker.symmetricKey, token, nil)
	if err != nil {
		var ruleError paseto.RuleError
		if ok := errors.As(err, &ruleError); ok {
			if ruleError.Error() == "this token has expired" {
				return nil, ErrTokenExpired
			}
		}

		return nil, ErrTokenInvalid
	}

	jti, err := parsedToken.GetJti()
	if err != nil {
		return nil, ErrTokenInvalid
	}

	username, err := parsedToken.GetString("username")
	if err != nil {
		return nil, ErrTokenInvalid
	}

	issuedAt, err := parsedToken.GetIssuedAt()
	if err != nil {
		return nil, ErrTokenInvalid
	}

	expiredAt, err := parsedToken.GetExpiration()
	if err != nil {
		return nil, ErrTokenInvalid
	}

	id, err := uuid.Parse(jti)
	if err != nil {
		return nil, fmt.Errorf("1 invalid token")
	}

	payload := &Payload{
		ID:        id,
		Username:  username,
		IssuedAt:  issuedAt,
		ExpiredAt: expiredAt,
	}

	return payload, nil
}
