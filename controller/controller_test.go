package controller

import (
	"github.com/ednailson/serasa-challenge/database"
	. "github.com/ednailson/serasa-challenge/helper_tests"
	. "github.com/onsi/gomega"
	"testing"
)

func TestController(t *testing.T) {
	g := NewGomegaWithT(t)
	db, err := database.NewDatabase(FakeDbConfig())
	g.Expect(err).ShouldNot(HaveOccurred())
	negativationsServer := MockNegativationMainframeServer(g)
	defer negativationsServer.Close()

	ctrl := NewController(negativationsServer.URL, db)

	t.Run("update data", func(t *testing.T) {
		client := MockClient(g)
		db, err := client.Database(nil, DBNameTest)
		g.Expect(err).ShouldNot(HaveOccurred())
		coll, err := db.Collection(nil, DBCollTest)
		g.Expect(err).ShouldNot(HaveOccurred())
		err = coll.Truncate(nil)
		g.Expect(err).ShouldNot(HaveOccurred())

		err = ctrl.UpdateData()

		g.Expect(err).ShouldNot(HaveOccurred())
		statistics, err := coll.Statistics(nil)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(statistics.Count).Should(BeEquivalentTo(2))
	})
	t.Run("read a document by document", func(t *testing.T) {
		coll := MockAndTruncateCollection(g, DBCollTest)
		negativation := FakeNegativations()[0]
		_, err := coll.CreateDocument(nil, negativation)
		g.Expect(err).ShouldNot(HaveOccurred())

		assert, err := ctrl.NegativationByDocument(negativation.CustomerDocument)

		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(len(assert)).Should(BeEquivalentTo(1))
		g.Expect(assert[0]).Should(BeEquivalentTo(negativation))
	})
}

func FakeDbConfig() database.Config {
	return database.Config{
		Collection: DBCollTest,
		Host:       DBHostTest,
		Port:       DBPortTest,
		User:       DBUserTest,
		Password:   DBPassTest,
		Database:   DBNameTest,
	}
}
