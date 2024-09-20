package jwt

import (
	"fmt"
	"time"

	"github.com/aiagt/aiagt/app/user/conf"
	"github.com/google/uuid"

	"github.com/golang-jwt/jwt/v5"
)

const (
	issuer   = "https://aiagt.cn"
	audience = "https://aiagt.cn"
)

type Claims struct {
	jwt.RegisteredClaims
	ID int64
}

func GenerateToken(id int64) (string, *time.Time, error) {
	var (
		config    = conf.Conf().Auth
		expiresAt = time.Now().Add(config.JWTExpire * time.Hour)
	)

	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			Subject:   fmt.Sprintf("auth-%d", id),
			Audience:  jwt.ClaimStrings{audience},
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
		ID: id,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.JWTKey))
	if err != nil {
		return "", nil, err
	}

	return token, &expiresAt, nil
}

func ParseToken(token string) (int64, error) {
	var claims Claims

	t, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(conf.Conf().Auth.JWTKey), nil
	})
	if err != nil {
		return 0, err
	}

	if !t.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	return claims.ID, nil
}
