package services

import (
	"crypto/ecdsa"
	"encoding/base64"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

const (
	base64PrivateKey = "CHANGEME"
	keyID            = "CHANGEME"
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

func GetJWT(scope []string, tokenLifetime time.Duration) (string, error) {
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
