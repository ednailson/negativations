package app

import (
	"encoding/json"
	"fmt"
	"github.com/ednailson/serasa-challenge/database"
	"github.com/ednailson/serasa-challenge/domain"
	. "github.com/ednailson/serasa-challenge/helper_tests"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestApp(t *testing.T) {
	g := NewGomegaWithT(t)
	_ = MockAndTruncateCollection(g, DBCollTest)
	mainframeServer := MockNegativationMainframeServer(g)
	defer mainframeServer.Close()
	sut, err := LoadApp(fakeConfig(mainframeServer.URL))
	g.Expect(err).ShouldNot(HaveOccurred())
	sut.Run()
	defer sut.Close()
	t.Run("reading a negativation by document", func(t *testing.T) {
		url := fmt.Sprintf("http://localhost:%d/v1", fakeConfig("").Port)
		resp, err := http.Post(url+"/update", "application/json", nil)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(resp.StatusCode).Should(BeEquivalentTo(http.StatusNoContent))

		resp, err = http.Get(fmt.Sprintf("%s/negativation?cpf=%s", url, FakeNegativations()[0].CustomerDocument))
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(resp.StatusCode).Should(BeEquivalentTo(http.StatusOK))
		body, err := ioutil.ReadAll(resp.Body)
		g.Expect(err).ShouldNot(HaveOccurred())
		var respBody []domain.Negativation
		err = json.Unmarshal(body, &respBody)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(len(respBody)).Should(BeEquivalentTo(2))
		g.Expect(respBody).Should(ConsistOf(FakeNegativations()))
	})
}

func fakeConfig(host string) Config {
	return Config{
		MainframeUrl: host,
		Port:         5000,
		Database: database.Config{
			Collection: DBCollTest,
			Host:       DBHostTest,
			Port:       DBPortTest,
			User:       DBUserTest,
			Password:   DBPassTest,
			Database:   DBNameTest,
		},
	}
}
