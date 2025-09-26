package main

import "testing"

func TestGetH1FromHTML(t *testing.T) {
	cases := []struct {
		name string
		html string
		want string
	}{
		{"basic", "<html><h1>Test</h1></html>", "Test"},
		{"no h1", "<html><p>Hi</p></html>", ""},
		{"trim", "<h1>\n  Spaced \t</h1>", "Spaced"},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := getH1FromHTML(tt.html)
			if got != tt.want {
				t.Fatalf("want %q, got %q", tt.want, got)
			}
		})
	}
}

// go
func TestGetFirstParagraphFromHTML(t *testing.T) {
	cases := []struct {
		name string
		html string
		want string
	}{
		{
			name: "main priority",
			html: `<html><body>
                <p>Outside</p>
                <main><p>Main paragraph.</p></main>
            </body></html>`,
			want: "Main paragraph.",
		},
		{
			name: "fallback no main",
			html: `<html><body>
                <p>First outside.</p>
                <p>Second outside.</p>
            </body></html>`,
			want: "First outside.",
		},
		{
			name: "empty p in main",
			html: `<main><p>   </p><p>Second</p></main>`,
			want: "",
		},
		{
			name: "whitespace trim",
			html: `<main><p>  Trim me \n\t</p></main>`,
			want: "Trim me",
		},
		{
			name: "no p anywhere",
			html: `<html><body><div>No p</div></body></html>`,
			want: "",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := getFirstParagraphFromHTML(tt.html)
			if got != tt.want {
				t.Fatalf("want %q, got %q", tt.want, got)
			}
		})
	}
}
