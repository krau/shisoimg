package api

import (
	"net/url"
	"strings"

	"github.com/krau/shisoimg/dao"
	"github.com/krau/shisoimg/utils"
)

func applyRules(path string) (match bool, newUrl string) {
	for _, rule := range dao.Rules() {
		if strings.HasPrefix(path, rule.Path) {
			parsedUrl, err := url.JoinPath(rule.Prefix, strings.TrimPrefix(path, rule.Path))
			if err != nil {
				utils.L.Errorf("Failed to join path %s with %s: %v", rule.Prefix, path, err)
				continue
			}
			return true, parsedUrl
		}
	}
	return false, ""
}
