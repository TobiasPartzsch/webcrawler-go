package main

import (
	"net/url"
	"reflect"
	"testing"
)

// go
func TestGetImagesFromHTML(t *testing.T) {
	tests := []struct {
		name string
		base string
		body string
		want []string
	}{
		{
			name: "relative src",
			base: "https://blog.boot.dev",
			body: `<img src="/logo.png">`,
			want: []string{"https://blog.boot.dev/logo.png"},
		},
		{
			name: "absolute src",
			base: "https://blog.boot.dev",
			body: `<img src="https://cdn.boot.dev/a.png">`,
			want: []string{"https://cdn.boot.dev/a.png"},
		},
		{
			name: "missing/empty src ignored",
			base: "https://blog.boot.dev",
			body: `<img alt="x"><img src="">`,
			want: []string{},
		},
		{
			name: "multiple mixed",
			base: "https://blog.boot.dev",
			body: `<img src="/a.png"><img src="https://x/y.jpg"><img>`,
			want: []string{"https://blog.boot.dev/a.png", "https://x/y.jpg"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			u, err := url.Parse(tc.base)
			if err != nil {
				t.Fatalf("parse base: %v", err)
			}
			got, err := getImagesFromHTML(tc.body, u)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}
