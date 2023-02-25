package multiplexer

import (
	"regexp"
	"strings"
)

var regInt, _ = regexp.Compile("[0-9]+")
var regFloat, _ = regexp.Compile("^[-+]?[0-9]*(\\.[0-9]+)$")
var regSlug, _ = regexp.Compile("^[a-z0-9]+(?:[_-][a-z0-9]+)*$")
var regString, _ = regexp.Compile("^(.*)")

func explodePath(path string)[]string {
  listString := strings.Split(path, "/")
  if listString[0] == "" {
    return listString[1:]
  }
  return listString
}

func isWithParams (path string) bool {
  first := path[0]
  last := path[len(path) - 1]
  if first == '{' && last == '}' {
    return true
  }
  return false
}

func acceptStringWithoutSpace(path string) bool {
  allword := strings.Split(path, " ")
  if len(allword) == 1 {
    return true
  }
  return false
}

func getKey(params string) string {
  cleanParamas := params[1:len(params)-1]
  keyValue := strings.Split(cleanParamas, ":")
  return keyValue[0]
}

func selectRegex(params string) *regexp.Regexp {
  cleanParamas := params[1:len(params)-1]
  keyValue := strings.Split(cleanParamas, ":")

  switch keyValue[1] {
    case "float" : 
      return regFloat
    case "int" : 
      return regInt
    case "slug" : 
      return regSlug
    default :
      return regString
  }
}

// func splitPath(path string) []string {
//   collection1 := strings.Split(path, "{")
//   collection2 := strings.Split(path, "}")
// }

// func regexInsert(path string) (string, *regexp.Regexp, error) {
//   cleanPath := path[1:]
//   result := strings.Split(cleanPath, ":")
//   rg, err := regexp.Compile(result[1])
// }
