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

func ParseLinks(scheme, host, html string) []string {
	href := regexp.MustCompile(`href="(` + scheme + `:\/\/` + host + `|` + host + `|\/).*?"`)

	rawLinks := href.FindAllString(html, -1)

	links := []string{}

	isFile := regexp.MustCompile(`(\.\w+)"`)

	for _, rawLink := range rawLinks {
		if isFile.MatchString(rawLink) {
			continue
		}

		route := strings.Split(rawLink, "\"")[1]

		if strings.HasPrefix(route, "/") {
			links = append(links, scheme+"://"+host+route)
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
