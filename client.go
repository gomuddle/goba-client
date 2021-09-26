package gobaclient

import (
	"fmt"
	"github.com/gomuddle/goba"
	"github.com/gomuddle/gobaclient/internal/client"
	"net/url"
	"strings"
)

// GetImage retrieves the image with the given type
// and name from the server at the specified url.
func GetImage(uri url.URL, creds Credentials, typ goba.DatabaseType, name string) (image *goba.Image, err error) {
	return image, client.GetJSON(&image, client.Request{
		URL: buildPath(uri, "/images/%s/%s", typ, name),
		Headers: []client.HeaderFunc{
			authHeader(creds),
		},
		CheckResponse: checkResponse,
	})
}

// GetAllImages retrieves all images with the
// given type from the server at the specified url.
func GetAllImages(uri url.URL, creds Credentials, typ goba.DatabaseType) (images []goba.Image, err error) {
	return images, client.GetJSON(&images, client.Request{
		URL: buildPath(uri, "/images/%s", typ),
		Headers: []client.HeaderFunc{
			authHeader(creds),
		},
		CheckResponse: checkResponse,
	})
}

// CreateImage sends a request to create an image with
// the given type to the server at the specified url.
func CreateImage(uri url.URL, creds Credentials, typ goba.DatabaseType) (image *goba.Image, err error) {
	return image, client.PostJSON(&image, client.Request{
		URL: buildPath(uri, "/images/%s", typ),
		Headers: []client.HeaderFunc{
			authHeader(creds),
		},
		CheckResponse: checkResponse,
	})
}

// ApplyImage sends a request to apply the image with the
// given type and name to the database with the given type
// to the server at the specified url.
func ApplyImage(uri url.URL, creds Credentials, typ goba.DatabaseType, name string) error {
	return client.Post(client.Request{
		URL: buildPath(uri, "/images/%s/%s", typ, name),
		Headers: []client.HeaderFunc{
			authHeader(creds),
		},
		CheckResponse: checkResponse,
	})
}

// DeleteImage sends a request to delete the image with the
// given type and name to the server at the specified url.
func DeleteImage(uri url.URL, creds Credentials, typ goba.DatabaseType, name string) error {
	return client.Delete(client.Request{
		URL: buildPath(uri, "/images/%s/%s", typ, name),
		Headers: []client.HeaderFunc{
			authHeader(creds),
		},
		CheckResponse: checkResponse,
	})
}

// buildPaths formats the given uri.
func buildPath(uri url.URL, format string, args ...interface{}) string {
	if path := uri.Path; !strings.HasSuffix(path, "/") {
		path += "/"
	}
	uri.Path += fmt.Sprintf(format, args...)
	return uri.String()
}
