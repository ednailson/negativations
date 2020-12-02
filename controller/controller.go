package controller

import (
	"github.com/ednailson/serasa-challenge/controller/crypto"
	"github.com/ednailson/serasa-challenge/database"
	"github.com/ednailson/serasa-challenge/domain"
	"github.com/ednailson/serasa-challenge/integration"
)

type Controller struct {
	db           database.Database
	mainframe    string
	cryptoModule crypto.Crypto
}

func NewController(mainframe string, db database.Database, cryptoModule crypto.Crypto) *Controller {
	return &Controller{
		mainframe:    mainframe,
		db:           db,
		cryptoModule: cryptoModule,
	}
}

func (c *Controller) UpdateData() error {
	negativations, err := integration.ReadNegativations(c.mainframe)
	if err != nil {
		return err
	}
	encrypted := c.encryptNegativations(negativations)
	return c.db.Save(encrypted)
}

func (c *Controller) NegativationByDocument(document string) ([]domain.Negativation, error) {
	negativations, err := c.db.ReadByDocument(c.cryptoModule.Encrypt(document))
	if err != nil {
		return nil, err
	}
	return c.decryptNegativations(negativations)
}

func (c *Controller) decryptNegativations(data []domain.Negativation) ([]domain.Negativation, error) {
	var negativations []domain.Negativation
	for _, negativation := range data {
		customerDocument, err := c.cryptoModule.Decrypt(negativation.CustomerDocument)
		if err != nil {
			return nil, err
		}
		companyName, err := c.cryptoModule.Decrypt(negativation.CompanyName)
		if err != nil {
			return nil, err
		}
		companyDocument, err := c.cryptoModule.Decrypt(negativation.CompanyDocument)
		if err != nil {
			return nil, err
		}
		contract, err := c.cryptoModule.Decrypt(negativation.Contract)
		if err != nil {
			return nil, err
		}
		newNegativation := domain.CreateNegativation(companyDocument, companyName, customerDocument, negativation.Value, contract, negativation.DebtDate, negativation.InclusionDate)
		negativations = append(negativations, *newNegativation)
	}
	return negativations, nil
}

func (c *Controller) encryptNegativations(data []domain.Negativation) []domain.Negativation {
	var negativations []domain.Negativation
	for _, negativation := range data {
		customerDocument := c.cryptoModule.Encrypt(negativation.CustomerDocument)
		companyName := c.cryptoModule.Encrypt(negativation.CompanyName)
		companyDocument := c.cryptoModule.Encrypt(negativation.CompanyDocument)
		contract := c.cryptoModule.Encrypt(negativation.Contract)
		newNegativation := domain.CreateNegativation(companyDocument, companyName, customerDocument, negativation.Value, contract, negativation.DebtDate, negativation.InclusionDate)
		negativations = append(negativations, *newNegativation)
	}
	return negativations
}
