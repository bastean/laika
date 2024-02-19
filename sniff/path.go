package sniff

import "net/url"

func Path(path string) (string, error) {
	value, err := url.Parse(path)

	if err != nil {
		return "", err
	}

	return value.Path, nil
}
