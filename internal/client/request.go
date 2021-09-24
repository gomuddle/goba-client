package client

import (
	"encoding/json"
	"io"
	"net/http"
)

type Request struct {
	Method        string
	URL           string
	Headers       []HeaderFunc
	Decode        func(r io.Reader) error
	CheckResponse CheckResponseFunc
}

func (r Request) makeRequest() (*http.Request, error) {
	req, err := http.NewRequest(r.Method, r.URL, nil)
	if err != nil {
		return nil, err
	}
	r.setHeaders(req)
	return req, nil
}

func (r Request) setHeaders(req *http.Request) {
	for _, header := range r.Headers {
		req.Header.Add(header())
	}
}

func (r Request) decode(reader io.Reader) error {
	if r.Decode == nil {
		return nil
	}
	return r.Decode(reader)
}

func (r *Request) addHeader(h HeaderFunc) {
	r.Headers = append(r.Headers, h)
}

func (r Request) checkResponse(resp http.Response) error {
	if r.CheckResponse == nil {
		return nil
	}
	return r.CheckResponse(resp)
}

func DecodeFnJSON(v interface{}) func(io.Reader) error {
	return func(r io.Reader) error {
		return json.NewDecoder(r).Decode(v)
	}
}
