package app

import (
	"github.com/ednailson/httping-go"
	"net/http"
)

type response struct {
	status int
	data   interface{}
}

func NewResponse(code int, data interface{}) httping.IResponse {
	return &response{
		status: code,
		data:   data,
	}
}

func (r *response) Headers() map[string][]string {
	return nil
}
func (r *response) Cookies() []*http.Cookie {
	return nil
}
func (r *response) Response() interface{} {
	return r.data
}
func (r *response) StatusCode() int {
	return r.status
}
