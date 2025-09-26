package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func resolveAttrBySelector(htmlBody string, base *url.URL, selector, attr string) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}
	out := []string{}
	doc.Find(selector).Each(func(_ int, s *goquery.Selection) {
		if val, ok := s.Attr(attr); ok {
			if resolved, ok := resolveAttrURL(base, val); ok {
				out = append(out, resolved)
			}
		}
	})
	return out, nil
}

func resolveAttrURL(base *url.URL, raw string) (string, bool) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", false
	}
	u, err := url.Parse(raw)
	if err != nil {
		return "", false
	}
	return base.ResolveReference(u).String(), true
}
