package domain

import (
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"testing"
	"time"
)

func TestNegativation(t *testing.T) {
	g := NewGomegaWithT(t)

	sut := CreateNegativation(
		"59291534000167",
		"ABC S.A.",
		"51537476467",
		1235.23,
		"bc063153-fb9e-4334-9a6c-0d069a42065b",
		time.Date(2015, 11, 13, 20, 32, 51, 00, time.UTC),
		time.Date(2020, 11, 13, 20, 32, 51, 00, time.UTC),
	)

	g.Expect(sut).Should(PointTo(MatchAllFields(Fields{
		"CompanyDocument":  BeEquivalentTo("59291534000167"),
		"CompanyName":      BeEquivalentTo("ABC S.A."),
		"CustomerDocument": BeEquivalentTo("51537476467"),
		"Value":            BeEquivalentTo(1235.23),
		"Contract":         BeEquivalentTo("bc063153-fb9e-4334-9a6c-0d069a42065b"),
		"DebtDate":         BeEquivalentTo(time.Date(2015, 11, 13, 20, 32, 51, 00, time.UTC)),
		"InclusionDate":    BeEquivalentTo(time.Date(2020, 11, 13, 20, 32, 51, 00, time.UTC)),
	})))
}
