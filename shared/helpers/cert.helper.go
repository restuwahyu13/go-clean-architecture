package helper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	cons "github.com/restuwahyu13/go-clean-architecture/shared/constants"
	inf "github.com/restuwahyu13/go-clean-architecture/shared/interfaces"
)

type cert struct{}

func NewCert() inf.ICert {
	return cert{}
}

func (h cert) GeneratePrivateKey(password []byte) (string, error) {
	var pemBlock *pem.Block = new(pem.Block)

	rsaPrivateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", err
	}

	privateKeyTransform := h.PrivateKeyToRaw(rsaPrivateKey)

	if password != nil {
		encryptPemBlock, err := x509.EncryptPEMBlock(rand.Reader, "RSA PRIVATE KEY", []byte(privateKeyTransform), []byte(password), x509.PEMCipherAES256)
		if err != nil {
			return "", err
		}

		pemBlock = encryptPemBlock
	} else {
		decodePemBlock, _ := pem.Decode([]byte(privateKeyTransform))
		if pemBlock == nil {
			return "", errors.New("Invalid PrivateKey")
		}

		pemBlock = decodePemBlock
	}

	return string(pem.EncodeToMemory(pemBlock)), nil
}

func (h cert) PrivateKeyRawToKey(privateKey []byte, password []byte) (*rsa.PrivateKey, error) {
	decodePrivateKey, _ := pem.Decode(privateKey)
	if decodePrivateKey == nil {
		return nil, errors.New("Invalid PrivateKey")
	}

	if x509.IsEncryptedPEMBlock(decodePrivateKey) {
		decryptPrivateKey, err := x509.DecryptPEMBlock(decodePrivateKey, password)
		if err != nil {
			return nil, err
		}

		decodePrivateKey, _ = pem.Decode(decryptPrivateKey)
		if decodePrivateKey == nil {
			return nil, errors.New("Invalid PrivateKey")
		}
	}

	rsaPrivKey, err := x509.ParsePKCS1PrivateKey(decodePrivateKey.Bytes)
	if err != nil {
		return nil, err
	}

	return rsaPrivKey, nil
}

func (h cert) PrivateKeyToRaw(publicKey *rsa.PrivateKey) string {
	privateKeyTransform := pem.EncodeToMemory(&pem.Block{
		Type:  cons.PRIVPKCS1,
		Bytes: x509.MarshalPKCS1PrivateKey(publicKey),
	})

	return string(privateKeyTransform)
}
