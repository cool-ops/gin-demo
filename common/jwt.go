package common

import (
	"github.com/cool-ops/gin-demo/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 加密的key
var jwtKey = []byte("secret_key")

type claims struct {
	UserID uint
	jwt.StandardClaims
}

// 生成token
func GenerateToken(user model.User)(string,error){
	expireTime := time.Now().Add(7*24*time.Hour)
	newClaim := claims{
		UserID:         user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "joker",
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,newClaim)
	tokenString,err := token.SignedString(jwtKey)
	if err != nil {
		return "",err
	}
	return tokenString, nil
}

// 解析Token
func ParseToken(tokenString string)(*jwt.Token,*claims,error){
	newClaim := &claims{}
	token, err := jwt.ParseWithClaims(tokenString,newClaim, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey,err
	})
	return token,newClaim,err
}