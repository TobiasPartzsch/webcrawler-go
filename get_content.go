package main

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getH1FromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	return normalizeText(doc.Find("h1").First().Text())
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	main := doc.Find("main")
	var p string
	if main.Length() > 0 {
		p = main.Find("p").First().Text()
	} else {
		p = doc.Find("p").First().Text()
	}

	return normalizeText(p)
}

var spaceRe = regexp.MustCompile(`\s+`)

func normalizeText(s string) string {
	s = strings.TrimSpace(s)
	s = spaceRe.ReplaceAllString(s, " ")
	return s
}
