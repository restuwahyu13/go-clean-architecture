package opt

import (
	"crypto/rsa"
)

type (
	SecretMetadata struct {
		PrivKeyRaw string `json:"privKeyRaw"`
		CipherKey  string `json:"cipherKey"`
	}

	SignatureMetadata struct {
		PrivKey    *rsa.PrivateKey    `json:"privKey"`
		PrivKeyRaw string             `json:"privKeyRaw"`
		SigKey     string             `json:"sigKey"`
		CipherKey  string             `json:"cipherKey"`
		JweKey     JweEncryptMetadata `json:"jweKey"`
	}

	SignMetadata struct {
		Token   string `json:"token"`
		Expired int    `json:"expired"`
	}
)
