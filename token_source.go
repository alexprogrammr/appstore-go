package appstore

import (
	"crypto/ecdsa"
	"fmt"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Config struct {
	KeyID       string
	IssuerID    string
	PrivateKey  []byte
	ExpireAfter time.Duration
}

type TokenSource interface {
	Token() (string, error)
}

func NewTokenSource(config Config) (TokenSource, error) {
	pk, err := jwt.ParseECPrivateKeyFromPEM(config.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return &tokenSource{
		config: config,
		pk:     pk,
	}, nil
}

type tokenSource struct {
	sync.Mutex

	config   Config
	pk       *ecdsa.PrivateKey
	bearer   string
	expireAt time.Time
}

func (ts *tokenSource) Token() (string, error) {
	ts.Lock()
	defer ts.Unlock()

	if ts.isExpired() {
		return ts.refresh()
	}

	return ts.bearer, nil
}

func (ts *tokenSource) isExpired() bool {
	return time.Now().After(ts.expireAt)
}

func (ts *tokenSource) refresh() (string, error) {
	iat := time.Now()
	exp := iat.Add(ts.config.ExpireAfter)

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss": ts.config.IssuerID,
		"iat": iat.Unix(),
		"exp": exp.Unix(),
		"aud": "appstoreconnect-v1",
		// "scope": []string{},
	})
	token.Header["kid"] = ts.config.KeyID

	bearer, err := token.SignedString(ts.pk)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	ts.bearer = bearer
	ts.expireAt = exp

	return bearer, nil
}
