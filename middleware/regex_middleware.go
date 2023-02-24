package middleware

import (
	"net/http"
	"regexp"
)

var regURL = regexp.MustCompile("{([^}]*)}")

func RegexMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  })
}

  

