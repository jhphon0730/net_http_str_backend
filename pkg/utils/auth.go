package utils

import (
	"time"
	"strconv"

	"image_storage_server/config"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(config.GetJWTSecretKey())

type Claims struct {
	ID        int64       `json:"id"`
	Username string 			`json:"username"`
	jwt.StandardClaims
}

// JWT 토큰 생성
func GenerateToken(id int64, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ID: id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// JWT 토큰 검증
// * 토큰이 유효하면 클레임을 반환하고, 그렇지 않으면 에러를 반환합니다.
func ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}


// string to int64
func ParseInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}
