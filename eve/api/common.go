package api

import "strings"

const (
	baseUrl = "https://esi.evepc.163.com/latest"
)

var (
	header = map[string]interface{}{
		"accept": "application/json",
	}
)

func merge(elements ...string) string {
	if len(elements) > 0 {
		return strings.Join(elements, "")
	}
	return ""
}
