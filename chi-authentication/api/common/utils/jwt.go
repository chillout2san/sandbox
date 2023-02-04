package utils

import (
	"chi_sample/config"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// cookieの有効期間より1分長い有効期限11分のjwtを発行する
func CreateJwt(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(11 * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.Environment.SecretKey))

	if err != nil {
		log.Println("CreateJwt failed:", err)
		return "", errors.New("トークン発行に失敗しました。")
	}

	return tokenString, nil
}

// jwtの正確性を検証する
func CheckJwt(id string, tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("Unexpected signing method:", token.Header["alg"])
			return nil, errors.New("トークン認証に失敗しました。")
		}

		return []byte(config.Environment.SecretKey), nil
	})

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		log.Println("Token is invalid.")
		return errors.New("トークン認証に失敗しました。")
	}

	if id != claims["id"] {
		return errors.New("トークンの所有者ではありません。")
	}

	return nil
}
