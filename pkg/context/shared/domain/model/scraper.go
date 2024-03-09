package model

type Scraper interface {
	GetContent(source string) string
	GetLinks(source string) []string
}
