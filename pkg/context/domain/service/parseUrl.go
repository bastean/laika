package service

import "net/url"

func ParseUrl(rawUrl string) (*url.URL, error) {
	urlParsed, err := url.Parse(rawUrl)

	if err != nil {
		return nil, err
	}

	return urlParsed, nil
}
