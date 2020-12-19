package crypto

import (
	. "github.com/ednailson/serasa-challenge/helper_tests"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCrypto(t *testing.T) {
	g := NewGomegaWithT(t)

	sut, err := NewCrypto(Key, Nonce)
	g.Expect(err).ShouldNot(HaveOccurred())

	t.Run("encrypt data", func(t *testing.T) {
		data := "test-data"

		assert := sut.Encrypt(data)

		g.Expect(assert).Should(BeEquivalentTo("97c88c297cd237a1a1ecc98b03cee550b2e60fdee26099204f"))
	})
	t.Run("decrypt data", func(t *testing.T) {
		data := "97c88c297cd237a1a1ecc98b03cee550b2e60fdee26099204f"

		assert, err := sut.Decrypt(data)

		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(assert).Should(BeEquivalentTo("test-data"))
	})
}
