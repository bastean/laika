package service

import "regexp"

var PlainEmailRegexp = regexp.MustCompile(`[\w\d.+~-]+@[A-Za-z-\d]+\.[a-z]{2,4}`)

var HtmlCommentRegexp = regexp.MustCompile(`<!--.*?-->`)

var JsConcatenationRegexp = regexp.MustCompile(`\([\w\d"'.+~-]+@[A-Za-z"'.+-]+\)`)
var JsConcatenationRemoveRegexp = regexp.MustCompile(`[^()"'+]+`)
