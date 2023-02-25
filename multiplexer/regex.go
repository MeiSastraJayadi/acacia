package multiplexer

import (
	"regexp"
)

var regInt, _ = regexp.Compile("[0-9]+")
var regFloat, _ = regexp.Compile("^[-+]?[0-9]*(\\.[0-9]+)$")
var regSlug, _ = regexp.Compile("^[a-z0-9]+(?:[_-][a-z0-9]+)*$")

// func splitPath(path string) []string {
//   collection1 := strings.Split(path, "{")
//   collection2 := strings.Split(path, "}")
// }

// func regexInsert(path string) (string, *regexp.Regexp, error) {
//   cleanPath := path[1:]
//   result := strings.Split(cleanPath, ":")
//   rg, err := regexp.Compile(result[1])
// }
