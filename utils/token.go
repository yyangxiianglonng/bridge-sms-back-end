package utils

import (
	"main/config"

	jwt "github.com/dgrijalva/jwt-go"

	"time"
)

var jwtSecret = []byte(config.InitConfig().JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Leave    string `json:"leave"`
	jwt.StandardClaims
}

// 产生token的函数
func GenerateToken(username, password, leave string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(8 * time.Hour)

	claims := Claims{
		username,
		password,
		leave,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "bridge",
		},
	}
	//
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 验证token的函数
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	//
	return nil, err
}
