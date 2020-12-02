package controller

import (
	"github.com/ednailson/serasa-challenge/database"
	"github.com/ednailson/serasa-challenge/domain"
	"github.com/ednailson/serasa-challenge/integration"
)

type Controller struct {
	db        database.Database
	mainframe string
}

func NewController(mainframe string, db database.Database) *Controller {
	return &Controller{
		mainframe: mainframe,
		db:        db,
	}
}

func (c *Controller) UpdateData() error {
	negativations, err := integration.ReadNegativations(c.mainframe)
	if err != nil {
		return err
	}
	return c.db.Save(negativations)
}

func (c *Controller) NegativationByDocument(document string) ([]domain.Negativation, error) {
	return c.db.ReadByDocument(document)
}
