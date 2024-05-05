package services

import (
	"crypto/ecdsa"
	"d-kv/signer/db-common/entity"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type TokenService struct {
	tokenLifetime time.Duration
}

func NewTokenService(tokenLifetime time.Duration) *TokenService {
	return &TokenService{tokenLifetime: tokenLifetime}
}

func (s *TokenService) GenerateECDSAPrivateKey(base64PrivateKey string) (*ecdsa.PrivateKey, error) {
	PEMPrivateKey, err := base64.StdEncoding.DecodeString(base64PrivateKey)
	if err != nil {
		fmt.Println("Ошибка при декодировании приватного ключа:", err)
		return nil, err
	}

	key, err := jwt.ParseECPrivateKeyFromPEM(PEMPrivateKey)
	if err != nil {
		fmt.Println("Не удалось создать экземпляр приватного ключа:", err)
		return nil, err
	}

	return key, err
}

func (s *TokenService) GetJwtToken(tokenInfo *entity.IntegrationToken) (string, error) {
	token := jwt.New(jwt.SigningMethodES256)

	token.Header["alg"] = "ES256"
	token.Header["kid"] = tokenInfo.KeyId

	token.Claims = jwt.MapClaims{
		"iss": "dd4fcb21-0c2c-4b22-9ea4-179e0b51f483",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(s.tokenLifetime).Unix(),
		"aud": "appstoreconnect-v1",
	}

	key, err := s.GenerateECDSAPrivateKey(tokenInfo.Token)
	if err != nil {
		fmt.Println("Error generating ECDSA private key:", err)
		return "", err
	}
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println("Не удалось подписать токен:", err)
		return "", err
	}
	return tokenString, err
}
