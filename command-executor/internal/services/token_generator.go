package services

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

const (
	base64PrivateKey = "CHANGEME"
	keyID            = "CHANGEME"
	tokenLifetime    = 20 * time.Minute
)

func generateECDSAPrivateKey() (*ecdsa.PrivateKey, error) {
	// Декодирование ключа из Base64 строки
	PEMPrivateKey, err := base64.StdEncoding.DecodeString(base64PrivateKey)
	if err != nil {
		log.Fatal("Ошибка при декодировании приватного ключа:", err)
		return nil, err
	}

	// Создаем экземпляр приватного ключа
	key, err := jwt.ParseECPrivateKeyFromPEM(PEMPrivateKey)
	if err != nil {
		log.Fatal("Не удалось создать экземпляр приватного ключа:", err)
		return nil, err
	}

	return key, nil
}

func getJWT() (string, error) {
	// Создаем новый JWT токен
	token := jwt.New(jwt.SigningMethodES256)

	// Задаем заголовок
	token.Header["alg"] = "ES256"
	token.Header["kid"] = keyID

	// payload
	token.Claims = jwt.MapClaims{
		"iss": "dd4fcb21-0c2c-4b22-9ea4-179e0b51f483",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(tokenLifetime).Unix(),
		"aud": "appstoreconnect-v1",
	}

	key, err := generateECDSAPrivateKey()
	if err != nil {
		log.Fatal("Error generating ECDSA private key:", err)
		return "", err
	}

	// Подписываем токен с использованием приватного ключа
	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Fatal("Не удалось подписать токен:", err)
		return "", err
	}

	return tokenString, nil
}

func Activate(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		default:
			jwtToken, err := getJWT()
			if err != nil {
				log.Fatal("Error generate JWT")
				return
			}
			err = os.WriteFile("token.txt", []byte(jwtToken), 0644)
			if err != nil {
				log.Fatal("Error write token")
				return
			}
			fmt.Println(jwtToken)
			fmt.Println("Token was written!")
		}

		time.Sleep(tokenLifetime)
	}
}
