package middleware

import "regexp"

type RegexURL struct {
  rgx *regexp.Regexp 
}

func NewRegexURL(compile string) (*RegexURL, error) {
  result, err := regexp.Compile(compile)
  if err != nil {
    return nil, err 
  }
  newReg := &RegexURL{
    rgx: result,
  }
  return newReg, nil
}
