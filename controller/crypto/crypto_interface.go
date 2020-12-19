package crypto

type Crypto interface {
	Encrypt(data string) string
	Decrypt(data string) (string, error)
}
