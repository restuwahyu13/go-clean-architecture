package dto

import "github.com/lestrrat-go/jwx/v3/jwk"

type (
	JweEncryptMetadata struct {
		CipherText   string         `json:"ciphertext"`
		EncryptedKey string         `json:"encrypted_key"`
		Header       map[string]any `json:"header"`
		IV           string         `json:"iv"`
		Protected    string         `json:"protected"`
		Tag          string         `json:"tag"`
	}

	JwkRawMetadata struct {
		D   string `json:"d"`
		Dp  string `json:"dp"`
		Dq  string `json:"dq"`
		E   string `json:"e"`
		Kty string `json:"kty"`
		N   string `json:"n"`
		P   string `json:"p"`
		Q   string `json:"q"`
		Qi  string `json:"qi"`
	}

	JwkMetadata struct {
		KeyRaw JwkRawMetadata
		Key    jwk.Key
	}
)
