package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

const minSecretKeyLength = 6

type MyCustomClaims struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`

	jwt.RegisteredClaims
}

type JWTMaker struct {
	secretKey string
}

func (maker JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		ID:       payload.ID,
		Username: payload.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(payload.ExpiresAt),
			IssuedAt:  jwt.NewNumericDate(payload.IssuedAt),
		},
	})

	// SignedString input must be a byte slice NOT A STRING !!!
	return claims.SignedString([]byte(maker.secretKey))
}

func (maker JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}

		return []byte(maker.secretKey), nil
	}

	claims, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, keyFunc)
	if err != nil {

		return nil, ErrInvalidToken
	}

	customClaims, ok := claims.Claims.(*MyCustomClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return &Payload{
		ID:        customClaims.ID,
		Username:  customClaims.Username,
		IssuedAt:  customClaims.IssuedAt.Time,
		ExpiresAt: customClaims.ExpiresAt.Time,
	}, nil
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeyLength {
		return nil, fmt.Errorf("secret key length must be at least %d characters", minSecretKeyLength)
	}

	return &JWTMaker{secretKey: secretKey}, nil
}
