package service

import (
	"github.com/bastean/laika/src/context/application/contentFromUrl"
	"github.com/bastean/laika/src/context/application/sniffEmails"
)

var ContentFromUrl = contentFromUrl.NewContentFromUrl(LocalJson)
var SniffEmails = sniffEmails.NewSniffEmails(LocalJson)
