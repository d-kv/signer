package main

import (
	"d-kv/signer/command-executor/internal/services"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ecdsaKeyExample = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIOyneQNhVqHzQCr1GcyLUyfu7vILw7P6B8v+kT53WczJoAoGCCqGSM49
AwEHoUQDQgAEwEjc3lhU/PnR0h6p0Kw1uOVuNcgAukpRyHilfvuIDjvWp0B3pdZz
kx30T6VriBo6MGAAInJ6VA9ZQdHxKunqJg==
-----END EC PRIVATE KEY-----
`

func TestGenerateECDSAPrivateKey(t *testing.T) {
	mockService := services.TokenService{}
	base64PrivateKey := base64.StdEncoding.EncodeToString([]byte(ecdsaKeyExample))
	key, err := mockService.GenerateECDSAPrivateKey(base64PrivateKey)
	assert.Nil(t, err)
	assert.NotNil(t, key)
	assert.NotNil(t, key.D)
	assert.NotNil(t, key.PublicKey.X)
	assert.NotNil(t, key.PublicKey.Y)

	// Арбитражный JWT
	tok := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{"foo": "bar"})
	str, err := tok.SignedString(key)
	assert.Nil(t, err)

	parsed, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return &key.PublicKey, nil
	})
	assert.Nil(t, err)
	assert.True(t, parsed.Valid)
}
