package sniff

import "net/url"

func Url(route string) (*url.URL, error) {
	data, err := url.Parse(route)

	if err != nil {
		return nil, err
	}

	return data, nil
}
