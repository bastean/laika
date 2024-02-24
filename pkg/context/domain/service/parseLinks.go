package service

import (
	"regexp"
	"strings"
)

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
