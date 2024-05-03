package services

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	tokenLifetime = 20 * time.Minute
)

func (s *ProcessorService) generateECDSAPrivateKey(base64PrivateKey string) (*ecdsa.PrivateKey, error) {
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

	return key, nil
}

func (s *ProcessorService) getJWT(ctx context.Context, IntegrationId string) (string, error) {
	token := jwt.New(jwt.SigningMethodES256)
	err, tokenInfo := s.vault.FindTokenByIntegrationId(ctx, IntegrationId)

	token.Header["alg"] = "ES256"
	token.Header["kid"] = tokenInfo.KeyId

	token.Claims = jwt.MapClaims{
		"iss": "dd4fcb21-0c2c-4b22-9ea4-179e0b51f483",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(tokenLifetime).Unix(),
		"aud": "appstoreconnect-v1",
	}

	key, err := s.generateECDSAPrivateKey(tokenInfo.Token)
	if err != nil {
		fmt.Println("Error generating ECDSA private key:", err)
		return "", err
	}
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println("Не удалось подписать токен:", err)
		return "", err
	}
	return tokenString, nil
}

func (s *ProcessorService) GetTokenByIntegrationID(ctx context.Context, IntegrationId string) (string, error) {
	jwtToken, err := s.getJWT(ctx, IntegrationId)
	if err != nil {
		fmt.Println("Error generate JWT")
		return "", err
	}
	return jwtToken, err
}
