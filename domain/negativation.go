package domain

import "time"

type Negativation struct {
	CompanyDocument  string    `json:"companyDocument"`
	CompanyName      string    `json:"companyName"`
	CustomerDocument string    `json:"customerDocument"`
	Value            float64   `json:"value"`
	Contract         string    `json:"contract"`
	DebtDate         time.Time `json:"debtDate"`
	InclusionDate    time.Time `json:"inclusionDate"`
}

func CreateNegativation(companyDocument, companyName, customerDocument string, value float64, contract string, debtDate, inclusionDate time.Time) *Negativation {
	return &Negativation{
		CompanyDocument:  companyDocument,
		CompanyName:      companyName,
		CustomerDocument: customerDocument,
		Value:            value,
		Contract:         contract,
		DebtDate:         debtDate,
		InclusionDate:    inclusionDate,
	}
}
