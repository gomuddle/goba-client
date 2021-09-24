package client

import (
	"net/http"
)

type CheckResponseFunc func(resp http.Response) error
