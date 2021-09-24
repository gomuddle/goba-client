package client

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"time"
)

type client struct {
	c *http.Client
}

func newClient() *client {
	c := &client{
		c: &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout: 5 * time.Second,
				}).DialContext,
			},
			Timeout: 5 * time.Second,
		},
	}
	return c
}

func (c client) do(request Request) error {
	resp, err := c.makeReqAndDo(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return c.handleResponse(request, *resp)
}

func (c client) makeReqAndDo(request Request) (*http.Response, error) {
	req, err := request.makeRequest()
	if err != nil {
		return nil, err
	}
	return c.c.Do(req)
}

func (c client) handleResponse(req Request, resp http.Response) error {
	data, err := c.safelyReadResponseBody(&resp)
	if err != nil {
		return err
	}

	if err = req.checkResponse(resp); err != nil {
		return err
	}
	return req.decode(bytes.NewReader(data))
}

func (c client) safelyReadResponseBody(resp *http.Response) ([]byte, error) {
	data, err := io.ReadAll(resp.Body)
	if err == nil {
		resp.Body = io.NopCloser(bytes.NewReader(data))
	}
	return data, err
}

func Get(req Request) error {
	req.Method = http.MethodGet
	return Do(req)
}

func GetJSON(resp interface{}, req Request) error {
	req.Method = http.MethodGet
	return DoJSON(resp, req)
}

func Post(req Request) error {
	req.Method = http.MethodPost
	return Do(req)
}

func PostJSON(resp interface{}, req Request) error {
	req.Method = http.MethodPost
	return DoJSON(resp, req)
}

func Put(req Request) error {
	req.Method = http.MethodPut
	return Do(req)
}

func PutJSON(resp interface{}, req Request) error {
	req.Method = http.MethodPut
	return DoJSON(resp, req)
}

func Patch(req Request) error {
	req.Method = http.MethodPatch
	return Do(req)
}

func PatchJSON(resp interface{}, req Request) error {
	req.Method = http.MethodPatch
	return DoJSON(resp, req)
}

func Delete(req Request) error {
	req.Method = http.MethodDelete
	return Do(req)
}

func DeleteJSON(resp interface{}, req Request) error {
	req.Method = http.MethodDelete
	return DoJSON(&resp, req)
}

func DoJSON(response interface{}, req Request) error {
	req.Decode = DecodeFnJSON(response)
	return Do(req)
}

func Do(req Request) error {
	return defaultClient.do(req)
}

var defaultClient = newClient()
