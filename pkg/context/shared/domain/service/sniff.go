package service

import (
	"strings"
)

func SniffPlainEmails(content string) []string {
	return PlainEmailRegexp.FindAllString(content, -1)
}

func SniffJsConcatenationEmails(content string) []string {
	jsConcatenationEmails := JsConcatenationRegexp.FindAllString(content, -1)
	rawEmails := []string{}

	for _, concat := range jsConcatenationEmails {
		letters := JsConcatenationRemoveRegexp.FindAllString(concat, -1)
		rawEmail := strings.Join(letters, "")
		rawEmails = append(rawEmails, rawEmail)
	}

	return rawEmails
}
