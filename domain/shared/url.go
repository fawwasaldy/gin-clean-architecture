package shared

import "fmt"

type URL struct {
	Path string
}

func NewURL(path string) (URL, error) {
	if !isValidURL(path) {
		return URL{}, fmt.Errorf("invalid URL: %s", path)
	}
	return URL{
		Path: path,
	}, nil
}

func NewURLFromSchema(path string) URL {
	return URL{
		Path: path,
	}
}

func isValidURL(url string) bool {
	return len(url) > 0
}
