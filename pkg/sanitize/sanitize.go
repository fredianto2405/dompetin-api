package sanitize

import "github.com/microcosm-cc/bluemonday"

var (
	strictPolicy = bluemonday.StrictPolicy()
	ugcPolicy    = bluemonday.UGCPolicy()
)

func SanitizeStrict(input string) string {
	return strictPolicy.Sanitize(input)
}

func SanitizeUGC(input string) string {
	return ugcPolicy.Sanitize(input)
}
