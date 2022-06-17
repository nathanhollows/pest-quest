package helpers

// Requires config/config to pass
// `go test -c` in this directory
// `./internal/helpers/helpers_test.go` from root project directory

import (
	"testing"

	"github.com/nathanhollows/pest-quest/internal/config"
)

// TestURLEmpty calls helpers.URL with strings, checking
// for a valid return value.
func TestURLEmpty(t *testing.T) {
	config.Cfg.Server.Host = "localhost"
	config.Cfg.Server.Port = "8803"
	url := `//localhost:8803/`
	test := URL()
	if test != url {
		t.Fatalf(`URL() = %v, want match for %q`, test, url)
	}
}

// TestURLPath calls helpers.URL with strings, checking
// for a valid return value.
func TestURLPath(t *testing.T) {
	config.Cfg.Server.Host = "localhost"
	config.Cfg.Server.Port = "8803"
	url := `//localhost:8803/testing`
	test := URL("testing")
	if test != url {
		t.Fatalf(`URL() = %v, want match for %q`, test, url)
	}
}

// TestURLRawQuery calls helpers.URL with strings, checking
// for a valid return value.
func TestURLRawQuery(t *testing.T) {
	config.Cfg.Server.Host = "localhost"
	config.Cfg.Server.Port = "8803"
	url := `//localhost:8803/testing?test=test`
	test := URL("testing", "test=test")
	if test != url {
		t.Fatalf(`URL() = %v, want match for %q`, test, url)
	}
}

// TestURLRawQueryFragment calls helpers.URL with strings, checking
// for a valid return value.
func TestURLRawQueryFragment(t *testing.T) {
	config.Cfg.Server.Host = "localhost"
	config.Cfg.Server.Port = "8803"
	url := `//localhost:8803/testing?test=test#fragment`
	test := URL("testing", "test=test", "fragment")
	if test != url {
		t.Fatalf(`URL() = %v, want match for %q`, test, url)
	}
}

// TestURLRawQueryFragment calls helpers.URL with strings, checking
// for a valid return value.
func TestURLFragment(t *testing.T) {
	config.Cfg.Server.Host = "localhost"
	config.Cfg.Server.Port = "8803"
	url := `//localhost:8803/testing#fragment`
	test := URL("testing", "", "fragment")
	if test != url {
		t.Fatalf(`URL() = %v, want match for %q`, test, url)
	}
}
