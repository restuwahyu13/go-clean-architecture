package inf

import (
	"crypto/rsa"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"github.com/restuwahyu13/go-clean-architecture/shared/dto"
	opt "github.com/restuwahyu13/go-clean-architecture/shared/output"
)

type IJose interface {
	JweEncrypt(publicKey *rsa.PublicKey, plainText string) ([]byte, *opt.JweEncryptMetadata, error)
	JweDecrypt(privateKey *rsa.PrivateKey, cipherText []byte) (string, error)
	ImportJsonWebKey(jwkKey jwk.Key) (*opt.JwkMetadata, error)
	ExportJsonWebKey(privateKey *rsa.PrivateKey) (*opt.JwkMetadata, error)
	JwtSign(options *dto.JwtSignOption) ([]byte, error)
	JwtVerify(prefix string, token string, redis IRedis) (*jwt.Token, error)
}
