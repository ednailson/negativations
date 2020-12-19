package helper_tests

import (
	"encoding/json"
	"github.com/ednailson/serasa-challenge/domain"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"time"
)

func FakeNegativationsJson(g *GomegaWithT) []byte {
	body, err := json.Marshal(FakeNegativations())
	g.Expect(err).ShouldNot(HaveOccurred())
	return body
}

func FakeNegativations() []domain.Negativation {
	negativations := make([]domain.Negativation, 2)
	negativations[0] = *domain.CreateNegativation(
		"59291534000167",
		"ABC S.A.",
		"51537476467",
		1235.23,
		"bc063153-fb9e-4334-9a6c-0d069a42065b",
		time.Date(2015, 11, 13, 20, 32, 51, 00, time.UTC),
		time.Date(2020, 11, 13, 20, 32, 51, 00, time.UTC),
	)
	negativations[1] = *domain.CreateNegativation(
		"77723018000146",
		"123 S.A.",
		"51537476467",
		400.00,
		"5f206825-3cfe-412f-8302-cc1b24a179b0",
		time.Date(2015, 10, 12, 20, 32, 51, 00, time.UTC),
		time.Date(2020, 10, 12, 20, 32, 51, 00, time.UTC),
	)
	return negativations
}

func MockNegativationMainframeServer(g *GomegaWithT) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			switch request.URL.Path {
			case "/":
				writer.WriteHeader(http.StatusOK)
				writer.Write(FakeNegativationsJson(g))
			default:
				writer.WriteHeader(http.StatusNotFound)
			}
		}))
}
