package services

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	Role int `json:"role"`
	jwt.RegisteredClaims
}

func NewJwt(id uint, role int, secret string) string {
	hmacSampleSecret := []byte(secret)
	claims := MyClaims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        strconv.FormatUint(uint64(id), 10),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "localhost",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		panic(err)
	}

	return tokenString
}

func Parse(tokenString string, secret string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
