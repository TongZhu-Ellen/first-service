package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("VERY_SECRET_KEY!")

type Claims struct {
	IdStr string `json:"id_string"` // JSON 保持 user_id
	jwt.RegisteredClaims
}

// Issue 生成一个 token，绑定 IdStr（demo 中可理解为 boss 权限）
func IssueToken(idStr string) (string, error) {
	exp := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		IdStr: idStr,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	// 用 HS256 算法生成一个新的 JWT，里面放刚才的 claims（小纸条内容）
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 用 jwtKey 对 JWT 进行签名，返回最终的 token 字符串
	// 这个字符串就是别人拿去访问接口的“通行证”，无法篡改内容
	return t.SignedString(jwtKey)
}

func ParseID(tokenStr string) (string, bool) {
	if tokenStr == "" {
		return "", false
	}

	cp := &Claims{}
	t, err := jwt.ParseWithClaims(tokenStr, cp, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !t.Valid {
		return "", false
	}

	return cp.IdStr, true
}
