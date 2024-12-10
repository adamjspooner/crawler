package main

import (
	"reflect"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "http://example.com",
			expected: "example.com",
		},
		{
			name:     "keep path",
			inputURL: "http://example.com/path",
			expected: "example.com/path",
		},
		{
			name:     "remove trailing slash",
			inputURL: "http://example.com/path/",
			expected: "example.com/path",
		},
		{
			name:     "remove query",
			inputURL: "http://example.com/path?query=string",
			expected: "example.com/path",
		},
		{
			name:     "remove fragment",
			inputURL: "http://example.com/path/#/fragment",
			expected: "example.com/path",
		},
	}

	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := normalizeURL(test.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, test.name, err)
			}
			if actual != test.expected {
				t.Errorf("Test %v = '%s' FAIL: expected URL: %v, actual: %v", i, test.name, test.expected, actual)
			}
		})
	}
}

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "root edge case",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
			<body>
				<a href="/">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/", "https://other.com/path/one"},
		},
		{
			name:     "empty edge case",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
			<body>
				<a href>
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
</html>
`,
			expected: []string{"", "https://other.com/path/one"},
		},
		{
			name:     "fragment edge case",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
			<body>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="#id">
					<span>Boot.dev</span>
				</a>
			</body>
</html>
`,
			expected: []string{"https://other.com/path/one", ""},
		},
		{
			name:     "query edge case",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
			<body>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="?query=string">
					<span>Boot.dev</span>
				</a>
			</body>
</html>
`,
			expected: []string{"https://other.com/path/one", ""},
		},
	}
	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(test.inputBody, test.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, test.name, err)
			}
			if len(actual) != len(test.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected %v URLs, actual: %v", i, test.name, len(test.expected), len(actual))
			}
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Test %v = '%s' FAIL: expected URLs: %v, actual: %v", i, test.name, test.expected, actual)
			}
		})
	}
}
