package sniff

import (
	"net/mail"
	"regexp"
)

var plainEmailRegex = regexp.MustCompile(`[\w\d.+~-]+@[A-Za-z-\d]+\.[a-z]{2,4}`)
var htmlCommentRegex = regexp.MustCompile(`<!--.*?-->`)

func Emails(content string) []string {
	rawEmails := []string{}

	rawEmails = append(rawEmails, plainEmailRegex.FindAllString(content, -1)...)

	contentWithoutComments := htmlCommentRegex.ReplaceAllString(content, "")

	rawEmails = append(rawEmails, plainEmailRegex.FindAllString(contentWithoutComments, -1)...)

	parsedEmails := []string{}

	for _, rawEmail := range rawEmails {
		email, err := mail.ParseAddress(rawEmail)

		if err != nil {
			continue
		}

		parsedEmails = append(parsedEmails, email.Address)
	}

	emails := []string{}
	isNotPresent := true

	for _, parsedEmail := range parsedEmails {
		isNotPresent = true

		for _, email := range emails {
			if parsedEmail == email {
				isNotPresent = false
				break
			}

		}

		if isNotPresent {
			emails = append(emails, parsedEmail)
		}
	}

	return emails
}
