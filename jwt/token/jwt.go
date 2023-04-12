package token

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenExpireDuration = time.Hour * 2

// const TokenExpireDuration = time.Second * 60

var secret = []byte("jwt-666") //密钥

type MyClaims struct {
	UserName string
	jwt.StandardClaims
}

// get token
func GetToken(username string) (string, error) {
	cla := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "go-jwt",                                   // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)
	// fmt.Println("Token = ", token)
	return token.SignedString(secret) // 进行签名生成对应的token
}

// parse token
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
