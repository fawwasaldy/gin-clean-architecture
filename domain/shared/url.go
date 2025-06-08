package shared

import "fmt"

type URL struct {
	path string
}

func NewURL(path string) (URL, error) {
	if !isValidURL(path) {
		return URL{}, fmt.Errorf("invalid URL: %s", path)
	}
	return URL{
		path: path,
	}, nil
}

func isValidURL(url string) bool {
	return len(url) > 0
}

func (u URL) Path() string {
	return u.path
}
