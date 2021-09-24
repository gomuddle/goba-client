package goba_client

import (
	"encoding/base64"
	"github.com/gomuddle/goba-client/internal/client"
)

// authHeader returns a client.HeaderFunc returning
// an authorization header for the given credentials.
func authHeader(creds Credentials) client.HeaderFunc {
	raw := creds.Username + ":" + creds.Password
	enc := base64.StdEncoding.EncodeToString([]byte(raw))
	return func() (string, string) {
		return "Authorization", "Basic " + enc
	}
}
