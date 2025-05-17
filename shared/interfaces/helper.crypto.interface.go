package inf

type ICrypto interface {
	AES256Encrypt(secretKey, plainText string) (string, error)
	AES256Decrypt(secretKey string, cipherText string) (string, error)
	HMACSHA512Sign(secretKey, data string) (string, error)
	HMACSHA512Verify(secretKey, data, hash string) bool
	SHA256Sign(plainText string) (string, error)
	SHA512Sign(plainText string) (string, error)
}
