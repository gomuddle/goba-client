package gobaclient

// Credentials represent the client's credentials.
type Credentials struct {
	// Username is the Credentials' password.
	Username string `json:"username,omitempty"`

	// Password is the Credentials' password.
	Password string `json:"password,omitempty"`
}
