package pkg

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"math"

	"time"

	goredis "github.com/redis/go-redis/v9"
	cons "github.com/restuwahyu13/go-clean-architecture/shared/constants"
	"github.com/restuwahyu13/go-clean-architecture/shared/dto"
	helper "github.com/restuwahyu13/go-clean-architecture/shared/helpers"
	inf "github.com/restuwahyu13/go-clean-architecture/shared/interfaces"
	opt "github.com/restuwahyu13/go-clean-architecture/shared/output"
)

type jsonWebToken struct {
	env       *dto.Environtment
	rds       inf.IRedis
	jose      inf.IJose
	cipher    inf.ICrypto
	cert      inf.ICert
	parser    inf.IParser
	transform inf.ITransform
}

func NewJsonWebToken(ctx context.Context, env *dto.Environtment, con *goredis.Client) inf.IJsonWebToken {
	jose := NewJose(ctx)

	rds, err := NewRedis(ctx, con)
	if err != nil {
		Logrus(cons.FATAL, err)
	}

	cipher := helper.NewCrypto()
	cert := helper.NewCert()
	parser := helper.NewParser()
	transform := helper.NewTransform()

	return &jsonWebToken{
		env:       env,
		rds:       rds,
		jose:      jose,
		cipher:    cipher,
		cert:      cert,
		parser:    parser,
		transform: transform,
	}
}

func (p jsonWebToken) createSecret(prefix string, body []byte) (*opt.SecretMetadata, error) {
	secretMetadataReq := new(dto.SecretMetadata)
	secretMetadataRes := new(opt.SecretMetadata)

	timeNow := time.Now().Format(time.UnixDate)
	cipherTextRandom := fmt.Sprintf("%s:%s:%s:%d", prefix, string(body), timeNow, p.env.JWT.EXPIRED)
	cipherTextData := hex.EncodeToString([]byte(cipherTextRandom))

	cipherSecretKey, err := p.cipher.SHA512Sign(cipherTextData)
	if err != nil {
		return nil, err
	}

	cipherText, err := p.cipher.SHA512Sign(timeNow)
	if err != nil {
		return nil, err
	}

	cipherKey, err := p.cipher.AES256Encrypt(cipherSecretKey, cipherText)
	if err != nil {
		return nil, err
	}

	rsaPrivateKeyPassword := []byte(cipherKey)

	privateKey, err := p.cert.GeneratePrivateKey(rsaPrivateKeyPassword)
	if err != nil {
		return nil, err
	}

	secretMetadataReq.PrivKeyRaw = privateKey
	secretMetadataReq.CipherKey = cipherKey

	if err := p.transform.ReqToRes(secretMetadataReq, secretMetadataRes); err != nil {
		return nil, err
	}

	return secretMetadataRes, nil
}

func (p jsonWebToken) createSignature(prefix string, body any) (*opt.SignatureMetadata, error) {
	var (
		signatureMetadataReq *dto.SignatureMetadata = new(dto.SignatureMetadata)
		signatureMetadataRes *opt.SignatureMetadata = new(opt.SignatureMetadata)
		signatureKey         string                 = fmt.Sprintf("CREDENTIAL:%s", prefix)
		signatureField       string                 = "signature_metadata"
	)

	bodyByte, err := p.parser.Marshal(body)
	if err != nil {
		return nil, err
	}

	secretKey, err := p.createSecret(prefix, bodyByte)
	if err != nil {
		return nil, err
	}

	rsaPrivateKey, err := p.cert.PrivateKeyRawToKey([]byte(secretKey.PrivKeyRaw), []byte(secretKey.CipherKey))
	if err != nil {
		return nil, err
	}

	cipherHash512 := sha512.New()
	cipherHash512.Write(bodyByte)
	cipherHash512Body := cipherHash512.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.SHA512, cipherHash512Body)
	if err != nil {
		return nil, err
	}

	if err := rsa.VerifyPKCS1v15(&rsaPrivateKey.PublicKey, crypto.SHA512, cipherHash512Body, signature); err != nil {
		return nil, err
	}

	signatureOutput := hex.EncodeToString(signature)

	_, jweKey, err := p.jose.JweEncrypt(&rsaPrivateKey.PublicKey, signatureOutput)
	if err != nil {
		return nil, err
	}

	signatureMetadataReq.PrivKeyRaw = secretKey.PrivKeyRaw
	signatureMetadataReq.SigKey = signatureOutput
	signatureMetadataReq.CipherKey = secretKey.CipherKey
	signatureMetadataReq.JweKey = *jweKey
	signatureMetadataReq.PrivKey = rsaPrivateKey

	signatureMetadataByte, err := p.parser.Marshal(signatureMetadataReq)
	if err != nil {
		return nil, err
	}

	jwtClaim := string(signatureMetadataByte)
	jwtExpired := time.Duration(time.Minute * time.Duration(p.env.JWT.EXPIRED))

	if err := p.rds.HSetEx(signatureKey, jwtExpired, signatureField, jwtClaim); err != nil {
		return nil, err
	}

	if err := p.transform.ReqToRes(signatureMetadataReq, signatureMetadataRes); err != nil {
		return nil, err
	}

	return signatureMetadataRes, nil
}

func (p jsonWebToken) Sign(prefix string, body any) (*opt.SignMetadata, error) {
	tokenKey := fmt.Sprintf("TOKEN:%s", prefix)
	signMetadataRes := new(opt.SignMetadata)

	tokenExist, err := p.rds.Exists(tokenKey)
	if err != nil {
		return nil, err
	}

	if tokenExist < 1 {
		signature, err := p.createSignature(prefix, body)
		if err != nil {
			return nil, err
		}

		timestamp := time.Now().Format(cons.DATE_TIME_FORMAT)
		aud := signature.SigKey[10:20]
		iss := signature.SigKey[30:40]
		sub := signature.SigKey[50:60]
		suffix := int(math.Pow(float64(p.env.JWT.EXPIRED), float64(len(aud)+len(iss)+len(sub))))

		secretKey := fmt.Sprintf("%s:%s:%s:%s:%d", aud, iss, sub, timestamp, suffix)
		secretData := hex.EncodeToString([]byte(secretKey))

		jti, err := p.cipher.AES256Encrypt(secretData, prefix)
		if err != nil {
			return nil, err
		}

		duration := time.Duration(time.Minute * time.Duration(p.env.JWT.EXPIRED))
		jwtIat := time.Now().UTC().Add(-duration)
		jwtExp := time.Now().Add(duration)

		tokenPayload := new(dto.JwtSignOption)
		tokenPayload.SecretKey = signature.CipherKey
		tokenPayload.Kid = signature.JweKey.CipherText
		tokenPayload.PrivateKey = signature.PrivKey
		tokenPayload.Aud = []string{aud}
		tokenPayload.Iss = iss
		tokenPayload.Sub = sub
		tokenPayload.Jti = jti
		tokenPayload.Iat = jwtIat
		tokenPayload.Exp = jwtExp
		tokenPayload.Claim = timestamp

		tokenData, err := p.jose.JwtSign(tokenPayload)
		if err != nil {
			return nil, err
		}

		if err := p.rds.SetEx(tokenKey, duration, string(tokenData)); err != nil {
			return nil, err
		}

		signMetadataRes.Token = string(tokenData)
		signMetadataRes.Expired = p.env.JWT.EXPIRED

		return signMetadataRes, nil
	} else {
		tokenData, err := p.rds.Get(tokenKey)
		if err != nil {
			return nil, err
		}

		signMetadataRes.Token = string(tokenData)
		signMetadataRes.Expired = p.env.JWT.EXPIRED

		return signMetadataRes, nil
	}
}
