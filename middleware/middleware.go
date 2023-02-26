package middleware

import (
	"net/http"
	"strings"
)

type Middleware func(next http.Handler) http.Handler

func Var(r *http.Request) map[string]string {
  splitPath := strings.Split(r.URL.RawPath, "/")  
  collection := []string{}
  for _, value := range splitPath {
    if value[0] == '{' && value[len(value)-1] == '}' {
      collection = append(collection, value)
    }
  }
 return make(map[string]string) 
}

func Query(r *http.Request) map[string]string {
  return make(map[string]string)
}
