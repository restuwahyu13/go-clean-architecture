package inf

import (
	"crypto/rsa"
)

type ICert interface {
	GeneratePrivateKey(password []byte) (string, error)
	PrivateKeyRawToKey(privateKey []byte, password []byte) (*rsa.PrivateKey, error)
	PrivateKeyToRaw(publicKey *rsa.PrivateKey) string
}
