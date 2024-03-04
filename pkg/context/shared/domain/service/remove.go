package service

import (
	"slices"
)

func RemoveHtmlComments(content string) string {
	return HtmlCommentRegexp.ReplaceAllString(content, "")
}

func RemoveDuplicates(content []string) []string {
	uniques := []string{}

	for _, value := range content {
		if !slices.Contains(uniques, value) {
			uniques = append(uniques, value)
		}
	}

	return uniques
}
