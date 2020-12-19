package integration

import (
	"encoding/json"
	"github.com/ednailson/serasa-challenge/domain"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func ReadNegativations(host string) ([]domain.Negativation, error) {
	resp, err := http.Get(host)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to the integration service")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read body")
	}
	var negativations []domain.Negativation
	err = json.Unmarshal(body, &negativations)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode body")
	}
	return negativations, nil
}
