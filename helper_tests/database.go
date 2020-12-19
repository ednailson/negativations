package helper_tests

import (
	"crypto/tls"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	. "github.com/onsi/gomega"
	"strconv"
)

const DBCollTest = "collection-test"
const DBHostTest = "http://arango.service.com.br"
const DBNameTest = "serasa-challenge"
const DBPassTest = "dummyPass"
const DBUserTest = "root"
const DBPortTest = 8529

func MockAndTruncateCollection(g *GomegaWithT, collName string) driver.Collection {
	coll := MockCollection(g, collName)
	err := coll.Truncate(nil)
	g.Expect(err).ToNot(HaveOccurred())
	return coll
}

func MockCollection(g *GomegaWithT, collName string) driver.Collection {
	db, err := MockClient(g).Database(nil, DBNameTest)
	g.Expect(err).ToNot(HaveOccurred())
	coll, err := db.Collection(nil, collName)
	g.Expect(err).ToNot(HaveOccurred())
	return coll
}

func MockClient(g *GomegaWithT) driver.Client {
	dbConn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{DBHostTest + ":" + strconv.Itoa(DBPortTest)},
		TLSConfig: &tls.Config{},
	})
	g.Expect(err).ToNot(HaveOccurred())
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     dbConn,
		Authentication: driver.BasicAuthentication(DBUserTest, DBPassTest)})
	g.Expect(err).ToNot(HaveOccurred())
	return client
}
