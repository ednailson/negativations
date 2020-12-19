package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
)

type crypto struct {
	gcm   cipher.AEAD
	nonce []byte
}

func NewCrypto(key, nonce string) (Crypto, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cipher")
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create gcm")
	}
	return &crypto{
		gcm:   gcm,
		nonce: []byte(nonce),
	}, nil
}

func (c *crypto) Encrypt(data string) string {
	text := c.gcm.Seal(nil, c.nonce, []byte(data), nil)
	return fmt.Sprintf("%x", text)
}
func (c *crypto) Decrypt(data string) (string, error) {
	decoded, err := hex.DecodeString(data)
	if err != nil {
		return "", errors.Wrap(err, "failed to decode string")
	}
	text, err := c.gcm.Open(nil, c.nonce, decoded, nil)
	if err != nil {
		return "", errors.Wrap(err, "failed to decrypt data")
	}
	return string(text), nil
}
