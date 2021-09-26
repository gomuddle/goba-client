package gobaclient

import (
	"encoding/json"
	"errors"
	"net/http"
)

// ErrorResponse represents the server's response to a failed request.
type ErrorResponse struct {
	Error error `json:"error,omitempty"`
}

// UnmarshalJSON unmarshals data into resp.
func (resp *ErrorResponse) UnmarshalJSON(data []byte) error {
	var raw struct {
		Error string `json:"error,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if err := raw.Error; err != "" {
		resp.Error = errors.New(err)
	}
	return nil
}

// checkResponse verifies the server's response.
func checkResponse(resp http.Response) error {
	var response ErrorResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}
	return response.Error
}
