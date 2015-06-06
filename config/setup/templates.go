package setup

import (
	"net/http"

	"github.com/mholt/caddy/middleware"
	"github.com/mholt/caddy/middleware/templates"
)

// Templates configures a new Templates middleware instance.
func Templates(c *Controller) (middleware.Middleware, error) {
	rules, err := templatesParse(c)
	if err != nil {
		return nil, err
	}

	return func(next middleware.Handler) middleware.Handler {
		return &templates.Templates{
			Rules:   rules,
			Root:    c.Root,
			FileSys: http.Dir(c.Root),
			Next:    next,
		}
	}, nil
}

func templatesParse(c *Controller) ([]templates.Rule, error) {
	var rules []templates.Rule

	for c.Next() {
		var rule templates.Rule

		if c.NextArg() {
			// First argument would be the path
			rule.Path = c.Val()

			// Any remaining arguments are extensions
			rule.Extensions = c.RemainingArgs()
			if len(rule.Extensions) == 0 {
				rule.Extensions = defaultExtensions
			}
		} else {
			rule.Path = defaultPath
			rule.Extensions = defaultExtensions
		}

		for _, ext := range rule.Extensions {
			rule.IndexFiles = append(rule.IndexFiles, "index"+ext)
		}

		rules = append(rules, rule)
	}

	return rules, nil
}

const defaultPath = "/"

var defaultExtensions = []string{".html", ".htm", ".tmpl", ".tpl", ".txt"}
