package database

import (
	"github.com/ednailson/serasa-challenge/domain"
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

func TestDatabase(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("test creating a new database", func(t *testing.T) {
		sut, err := NewDatabase(FakeDbConfig())

		arangoClient := MockClient(g)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(sut).ShouldNot(BeNil())
		exists, err := arangoClient.DatabaseExists(nil, DBNameTest)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(exists).Should(BeTrue())
		db, err := arangoClient.Database(nil, DBNameTest)
		g.Expect(err).ShouldNot(HaveOccurred())
		ok, err := db.CollectionExists(nil, DBCollTest)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(ok).Should(BeTrue())
	})
	t.Run("reading a nonexistent negativation by document", func(t *testing.T) {
		MockAndTruncateCollection(g, DBCollTest)
		sut, err := NewDatabase(FakeDbConfig())
		g.Expect(err).ShouldNot(HaveOccurred())

		negativation, err := sut.ReadByDocument("NONEXISTENT")

		g.Expect(err).Should(HaveOccurred())
		g.Expect(negativation).Should(BeNil())
	})
	t.Run("reading a negativation by document", func(t *testing.T) {
		coll := MockAndTruncateCollection(g, DBCollTest)
		_, err := coll.CreateDocument(nil, fakeNegativation())
		g.Expect(err).ShouldNot(HaveOccurred())
		sut, err := NewDatabase(FakeDbConfig())
		g.Expect(err).ShouldNot(HaveOccurred())

		negativation, err := sut.ReadByDocument(fakeNegativation().CustomerDocument)

		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(len(negativation)).Should(BeEquivalentTo(1))
		g.Expect(negativation[0]).Should(BeEquivalentTo(fakeNegativation()))
	})
	t.Run("save a negativation", func(t *testing.T) {
		coll := MockAndTruncateCollection(g, DBCollTest)
		sut, err := NewDatabase(FakeDbConfig())
		g.Expect(err).ShouldNot(HaveOccurred())

		key, err := sut.Save(fakeNegativation())

		g.Expect(err).ShouldNot(HaveOccurred())
		var negativation domain.Negativation
		_, err = coll.ReadDocument(nil, key, &negativation)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(negativation).Should(BeEquivalentTo(fakeNegativation()))
	})
}

func fakeNegativation() domain.Negativation {
	negativation := domain.CreateNegativation(
		"59291534000167",
		"ABC S.A.",
		"51537476467",
		1235.23,
		"bc063153-fb9e-4334-9a6c-0d069a42065b",
		time.Date(2015, 11, 13, 20, 32, 51, 00, time.UTC),
		time.Date(2020, 11, 13, 20, 32, 51, 00, time.UTC),
	)
	return *negativation
}

func FakeDbConfig() Config {
	return Config{
		Collection: DBCollTest,
		Host:       DBHostTest,
		Port:       DBPortTest,
		User:       DBUserTest,
		Password:   DBPassTest,
		Database:   DBNameTest,
	}
}
