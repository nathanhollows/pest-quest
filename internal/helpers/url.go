package helpers

import (
	"fmt"
	"net/url"

	"github.com/nathanhollows/pest-quest/internal/config"
)

// URL constructs a URL specific to the application
func URL(patterns ...string) string {
	var u url.URL
	u.Host = fmt.Sprint(config.Cfg.Server.Host, ":", config.Cfg.Server.Port)
	if len(patterns) > 0 {
		u.Path = patterns[0]
	} else {
		u.Path = "/"
	}
	if len(patterns) > 1 {
		u.RawQuery = patterns[1]
	}
	if len(patterns) > 2 {
		u.Fragment = patterns[2]
	}
	return u.String()
}
