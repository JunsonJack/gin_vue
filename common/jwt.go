package common

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"junsonjack.cn/go_vue/model"
)

var jwtKey = []byte("junsonjack.cn") //秘钥

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 生成token
func ReleaseToken(user model.User) (string,error){
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "junsonjack",
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString,err := token.SignedString(jwtKey)

	if err != nil {
		return "",err
	}
	return tokenString,nil
}

// 解析token
func ParseToken (tokenString string)(*jwt.Token , *Claims, error){
	claims := &Claims{}

	token , err := jwt.ParseWithClaims(tokenString,claims,func (token *jwt.Token) (i interface{},err error) {
		return jwtKey,nil
	})
	return token,claims,err
}