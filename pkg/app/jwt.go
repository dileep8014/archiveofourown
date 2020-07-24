package app

import (
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/shyptr/archiveofourown/global"
	"time"
)

type Claims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GetJwtSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(id int64, username string) (string, error) {
	now := time.Now()
	expireTime := now.Add(time.Duration(global.JWTSetting.Expire) * time.Second)
	claims := Claims{
		ID:       hex.EncodeToString([]byte(fmt.Sprint(id))),
		Username: hex.EncodeToString([]byte(username)),
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
