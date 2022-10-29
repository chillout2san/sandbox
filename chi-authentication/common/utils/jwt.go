package utils

import (
	"chi_sample/config"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 有効期限2時間のjwtを発行する
func CreateJwt(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(2 * time.Hour),
	})

	tokenString, err := token.SignedString([]byte(config.Enviroment.SecretKey))

	if err != nil {
		log.Println("CreateJwt failed:", err)
		return "", errors.New("トークン発行に失敗しました。")
	}

	return tokenString, nil
}

// jwtの正確性を検証する
func CheckJwt(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("Unexpected signing method:", token.Header["alg"])
			return nil, errors.New("トークン認証に失敗しました。")
		}

		return config.Enviroment.SecretKey, nil
	})

	if err != nil {
		return false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		log.Println("Token is invalid.")
		return false, errors.New("トークン認証に失敗しました。")
	}

	exp, _ := claims["exp"].(int64)

	if exp < time.Now().Unix() {
		return false, errors.New("トークンの有効期限を超過しています。")
	}

	return true, nil
}
