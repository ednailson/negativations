package integration

import (
	. "github.com/ednailson/serasa-challenge/helper_tests"
	. "github.com/onsi/gomega"
	"testing"
)

func TestReadNegativations(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("try to read on an invalid server", func(t *testing.T) {
		negativations, err := ReadNegativations("http://localhost:INVALIDPORT/")

		g.Expect(err).Should(HaveOccurred())
		g.Expect(negativations).Should(BeNil())
	})
	t.Run("reading on a valid server", func(t *testing.T) {
		negativationsServer := MockNegativationMainframeServer(g)

		negativations, err := ReadNegativations(negativationsServer.URL)

		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(negativations).Should(ConsistOf(FakeNegativations()))
	})
}
