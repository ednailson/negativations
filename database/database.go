package database

import (
	"crypto/tls"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/ednailson/serasa-challenge/domain"
	"strconv"
)

type dbDriver struct {
	db   driver.Database
	coll driver.Collection
}

func NewDatabase(config Config) (Database, error) {
	dbConn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{config.Host + ":" + strconv.Itoa(config.Port)},
		TLSConfig: &tls.Config{},
	})
	if err != nil {
		return nil, ErrConnecting
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     dbConn,
		Authentication: driver.BasicAuthentication(config.User, config.Password)})
	if err != nil {
		return nil, ErrConnecting
	}
	dbExists, err := client.DatabaseExists(nil, config.Database)
	if err != nil {
		return nil, ErrInitDatabase
	}
	var db driver.Database
	if !dbExists {
		db, err = client.CreateDatabase(nil, config.Database, nil)
		if err != nil {
			return nil, ErrInitDatabase
		}
	}
	db, err = client.Database(nil, config.Database)
	if err != nil {
		return nil, ErrInitDatabase
	}
	exist, err := db.CollectionExists(nil, config.Collection)
	if err != nil {
		return nil, ErrInitCollection
	}
	var coll driver.Collection
	if !exist {
		coll, err = db.CreateCollection(nil, config.Collection, nil)
		if err != nil {
			return nil, ErrInitCollection
		}
	}
	coll, err = db.Collection(nil, config.Collection)
	return &dbDriver{
		db:   db,
		coll: coll,
	}, nil
}

func (d *dbDriver) ReadByDocument(document string) ([]domain.Negativation, error) {
	query := `FOR n IN @@collection FILTER n.customerDocument == @document RETURN n`
	var bindVars = make(map[string]interface{})
	bindVars["@collection"] = d.coll.Name()
	bindVars["document"] = document
	cursor, err := d.db.Query(nil, query, bindVars)
	if err != nil {
		return nil, ErrReadByDocument
	}
	var negativations []domain.Negativation
	for cursor.HasMore() {
		var negativation domain.Negativation
		_, err = cursor.ReadDocument(nil, &negativation)
		if err != nil {
			return nil, ErrReadResult
		}
		negativations = append(negativations, negativation)
	}
	return negativations, nil
}

func (d *dbDriver) Save(negativations []domain.Negativation) error {
	query :=
		`FOR n IN @negativation UPSERT { contract: n.contract }
INSERT n
UPDATE n IN @@collection
RETURN NEW
`
	var bindVars = make(map[string]interface{})
	bindVars["@collection"] = d.coll.Name()
	bindVars["negativation"] = negativations
	_, err := d.db.Query(nil, query, bindVars)
	if err != nil {
		return ErrSaveDocuments
	}
	return nil
}
