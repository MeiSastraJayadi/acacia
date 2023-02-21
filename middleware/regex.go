package middleware

import "regexp"

type RegexURL struct {
  rgx *regexp.Regexp 
}
