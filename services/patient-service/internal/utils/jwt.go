package utils

import (
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func LoadPublicKey(path string) (*rsa.PublicKey, error){
	keyBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(keyBytes)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}