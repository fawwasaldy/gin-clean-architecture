package port

type (
	EncryptionPort interface {
		Encrypt(plainText string) (string, error)
		Decrypt(cipherText string) (string, error)
	}
)
