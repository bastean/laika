package service

import "net/url"

func ParseUrl(rawUrl string) *url.URL {
	urlParsed, err := url.Parse(rawUrl)

	if err != nil {
		panic(err)
	}

	return urlParsed
}
