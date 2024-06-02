package appstore

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTokenSource_Token(t *testing.T) {
	pk, b, err := generatePrivateKey()
	require.NoError(t, err)

	cfg := Config{
		KeyID:       "key-id",
		IssuerID:    "issuer-id",
		PrivateKey:  b,
		ExpireAfter: time.Hour,
	}

	source, err := NewTokenSource(cfg)
	require.NoError(t, err)

	token, err := source.Token()
	require.NoError(t, err)

	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return pk.Public(), nil
	})
	require.NoError(t, err)

	assert.Equal(t, jwt.SigningMethodES256.Alg(), parsed.Method.Alg())
	assert.Equal(t, cfg.KeyID, parsed.Header["kid"])

	claims, ok := parsed.Claims.(jwt.MapClaims)
	require.True(t, ok)

	assert.Equal(t, cfg.IssuerID, claims["iss"])
	assert.Equal(t, "appstoreconnect-v1", claims["aud"])
}

func TestTokenSource_TokenRefresh(t *testing.T) {
	_, b, err := generatePrivateKey()
	require.NoError(t, err)

	source, err := NewTokenSource(Config{
		KeyID:       "key-id",
		IssuerID:    "issuer-id",
		PrivateKey:  b,
		ExpireAfter: time.Second,
	})
	require.NoError(t, err)

	token, err := source.Token()
	require.NoError(t, err)

	same, err := source.Token()
	require.NoError(t, err)
	assert.Equal(t, token, same)

	time.Sleep(time.Second)

	new, err := source.Token()
	require.NoError(t, err)
	assert.NotEqual(t, token, new)
}

func generatePrivateKey() (*ecdsa.PrivateKey, []byte, error) {
	pk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate private key: %w", err)
	}

	keyBytes, err := x509.MarshalECPrivateKey(pk)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal private key: %w", err)
	}

	pemBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: keyBytes,
	})

	return pk, pemBytes, nil
}
