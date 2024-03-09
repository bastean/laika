package service

import (
	"html"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
)

func ParseHtml(rawHtml string) string {
	return html.UnescapeString(rawHtml)
}

func ParseLinks(scheme, host string, rawLinks []string) []string {
	links := []string{}

	isFile := regexp.MustCompile(`(\..*)`)

	for _, route := range rawLinks {
		if isFile.MatchString(route) {
			continue
		}

		switch {
		case strings.HasPrefix(route, "/"):
			links = append(links, scheme+"://"+host+route)
		case strings.HasPrefix(route, scheme) && strings.Contains(route, host):
			links = append(links, route)
		case strings.HasPrefix(route, host):
			links = append(links, scheme+"://"+route)
		}
	}

	return links
}

func ParseUrl(rawUrl string) (*url.URL, error) {
	if !strings.HasSuffix(rawUrl, "/") {
		rawUrl = rawUrl + "/"
	}

	urlParsed, err := url.Parse(rawUrl)

	if err != nil {
		return nil, err
	}

	return urlParsed, nil
}

func ParseAddresses(addresses []string) []string {
	parseAddresses := []string{}

	for _, address := range addresses {
		parsed, err := mail.ParseAddress(address)

		if err != nil {
			continue
		}

		host := strings.Split(parsed.Address, "@")[1]

		_, err = net.LookupIP(host)

		if err != nil {
			continue
		}

		parseAddresses = append(parseAddresses, parsed.Address)
	}

	return parseAddresses
}
