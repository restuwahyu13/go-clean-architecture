package pkg

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"reflect"

	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwe"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jws"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"github.com/restuwahyu13/go-clean-architecture/shared/dto"
	helper "github.com/restuwahyu13/go-clean-architecture/shared/helpers"
	inf "github.com/restuwahyu13/go-clean-architecture/shared/interfaces"
	opt "github.com/restuwahyu13/go-clean-architecture/shared/output"
)

type jose struct {
	ctx       context.Context
	cert      inf.ICert
	parser    inf.IParser
	transform inf.ITransform
}

func NewJose(ctx context.Context) inf.IJose {
	jwk.Configure(jwk.WithStrictKeyUsage(true))

	cert := helper.NewCert()
	parser := helper.NewParser()
	transform := helper.NewTransform()

	return &jose{ctx: ctx, cert: cert, parser: parser, transform: transform}
}

func (p jose) JweEncrypt(publicKey *rsa.PublicKey, plainText string) ([]byte, *opt.JweEncryptMetadata, error) {
	jweEncryptMetadataReq := new(dto.JweEncryptMetadata)
	jweEncryptMetadataRes := new(opt.JweEncryptMetadata)

	headers := jwe.NewHeaders()
	headers.Set("sig", plainText)
	headers.Set("alg", jwa.RSA_OAEP_512().String())
	headers.Set("enc", jwa.A256GCM().String())

	cipherText, err := jwe.Encrypt([]byte(plainText), jwe.WithKey(jwa.RSA_OAEP_512(), publicKey), jwe.WithContentEncryption(jwa.A256GCM()), jwe.WithCompact(), jwe.WithJSON(), jwe.WithProtectedHeaders(headers))
	if err != nil {
		return nil, nil, err
	}

	if err := p.parser.Unmarshal(cipherText, jweEncryptMetadataReq); err != nil {
		return nil, nil, err
	}

	if err := p.transform.ReqToRes(jweEncryptMetadataReq, jweEncryptMetadataRes); err != nil {
		return nil, nil, err
	}

	return cipherText, jweEncryptMetadataRes, nil
}

func (p jose) JweDecrypt(privateKey *rsa.PrivateKey, cipherText []byte) (string, error) {
	jwtKey, err := jwk.Import(privateKey)
	if err != nil {
		return "", err
	}

	jwkSet := jwk.NewSet()
	if err := jwkSet.AddKey(jwtKey); err != nil {
		return "", err
	}

	plainText, err := jwe.Decrypt(cipherText, jwe.WithKey(jwa.RSA_OAEP_512(), jwtKey), jwe.WithKeySet(jwkSet, jwe.WithRequireKid(false)))
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

func (p jose) ImportJsonWebKey(jwkKey jwk.Key) (*opt.JwkMetadata, error) {
	jwkRawMetadataReq := dto.JwkMetadata{}
	jwkRawMetadataRes := opt.JwkMetadata{}

	if _, err := jwk.IsPrivateKey(jwkKey); err != nil {
		return nil, err
	}

	if err := jwk.AssignKeyID(jwkKey); err != nil {
		return nil, err
	}

	jwkKeyByte, err := p.parser.Marshal(&jwkKey)
	if err != nil {
		return nil, err
	}

	jwkRaw, err := jwk.ParseKey(jwkKeyByte)
	if err != nil {
		return nil, err
	}

	if err := p.parser.Unmarshal(jwkKeyByte, &jwkRawMetadataReq.KeyRaw); err != nil {
		return nil, err
	}

	if err := p.transform.ReqToRes(&jwkRawMetadataReq, &jwkRawMetadataRes); err != nil {
		return nil, err
	}

	jwkRawMetadataRes.Key = jwkRaw

	return &jwkRawMetadataRes, nil
}

func (p jose) ExportJsonWebKey(privateKey *rsa.PrivateKey) (*opt.JwkMetadata, error) {
	jwkRawMetadataReq := dto.JwkMetadata{}
	jwkRawMetadataRes := opt.JwkMetadata{}

	jwkRaw, err := jwk.ParseKey([]byte(p.cert.PrivateKeyToRaw(privateKey)), jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}

	jwkRawByte, err := p.parser.Marshal(&jwkRaw)
	if err != nil {
		return nil, err
	}

	if err := p.parser.Unmarshal(jwkRawByte, &jwkRawMetadataReq.KeyRaw); err != nil {
		return nil, err
	}

	if err := p.transform.ReqToRes(&jwkRawMetadataReq, &jwkRawMetadataRes); err != nil {
		return nil, err
	}

	jwkRawMetadataRes.Key = jwkRaw.(jwk.Key)

	return &jwkRawMetadataRes, nil
}

func (p jose) JwtSign(options *dto.JwtSignOption) ([]byte, error) {
	jwsHeader := jws.NewHeaders()
	jwsHeader.Set("alg", jwa.RS512)
	jwsHeader.Set("typ", "JWT")
	jwsHeader.Set("cty", "JWT")
	jwsHeader.Set("kid", options.Kid)
	jwsHeader.Set("b64", true)

	jwtBuilder := jwt.NewBuilder()
	jwtBuilder.Audience(options.Aud)
	jwtBuilder.Issuer(options.Iss)
	jwtBuilder.Subject(options.Sub)
	jwtBuilder.IssuedAt(options.Iat)
	jwtBuilder.Expiration(options.Exp)
	jwtBuilder.JwtID(options.Jti)
	jwtBuilder.Claim("timestamp", options.Claim)

	jwtToken, err := jwtBuilder.Build()
	if err != nil {
		return nil, err
	}

	token, err := jwt.Sign(jwtToken, jwt.WithKey(jwa.RS512(), options.PrivateKey, jws.WithProtectedHeaders(jwsHeader)))
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (p jose) JwtVerify(prefix string, token string, redis inf.IRedis) (*jwt.Token, error) {
	signatureKey := fmt.Sprintf("CREDENTIAL:%s", prefix)
	signatureMetadataField := "signature_metadata"

	signatureMetadata := new(dto.SignatureMetadata)
	signatureMetadataBytes, err := redis.HGet(signatureKey, signatureMetadataField)
	if err != nil {
		return nil, err
	}

	if err := p.parser.Unmarshal(signatureMetadataBytes, signatureMetadata); err != nil {
		return nil, err
	}

	if reflect.DeepEqual(signatureMetadata, dto.SignatureMetadata{}) {
		return nil, errors.New("Invalid secretkey or signature")
	}

	privateKey, err := p.cert.PrivateKeyRawToKey([]byte(signatureMetadata.PrivKeyRaw), []byte(signatureMetadata.CipherKey))
	if err != nil {
		return nil, err
	}

	exportJws, err := jws.ParseString(token)
	if err != nil {
		return nil, err
	}

	signatures := exportJws.Signatures()
	if len(signatures) < 1 {
		return nil, errors.New("Invalid signature")
	}

	jwsSignature := new(jws.Signature)
	for _, signature := range signatures {
		jwsSignature = signature
		break
	}

	jwsHeaders := jwsSignature.ProtectedHeaders()

	algorithm, ok := jwsHeaders.Algorithm()
	if !ok {
		return nil, errors.New("Invalid algorithm")
	} else if algorithm != jwa.RS512() {
		return nil, errors.New("Invalid algorithm")
	}

	kid, ok := jwsHeaders.KeyID()
	if !ok {
		return nil, errors.New("Invalid keyid")
	} else if kid != signatureMetadata.JweKey.CipherText {
		return nil, errors.New("Invalid keyid")
	}

	aud := signatureMetadata.SigKey[10:20]
	iss := signatureMetadata.SigKey[30:40]
	sub := signatureMetadata.SigKey[50:60]
	claim := "timestamp"

	jwkKey, err := jwk.Import(privateKey)
	if err != nil {
		return nil, err
	}

	_, err = jws.Verify([]byte(token), jws.WithValidateKey(true), jws.WithKey(algorithm, jwkKey), jws.WithMessage(exportJws))
	if err != nil {
		return nil, err
	}

	jwtParse, err := jwt.Parse([]byte(token),
		jwt.WithKey(algorithm, privateKey),
		jwt.WithAudience(aud),
		jwt.WithIssuer(iss),
		jwt.WithSubject(sub),
		jwt.WithRequiredClaim(claim),
	)

	if err != nil {
		return nil, err
	}

	return &jwtParse, nil
}
