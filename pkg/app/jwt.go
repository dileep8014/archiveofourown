package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/shyptr/archiveofourown/global"
	"time"
)

type Claims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Root     bool   `json:"root"`
	jwt.StandardClaims
}

func GetJwtSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(id int64, username string, root bool) (string, error) {
	now := time.Now()
	expireTime := now.Add(time.Duration(global.JWTSetting.Expire) * time.Second)
	claims := Claims{
		ID:       id,
		Username: username,
		Root:     root,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(GetJwtSecret())
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(*jwt.Token) (interface{}, error) {
		return GetJwtSecret(), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
