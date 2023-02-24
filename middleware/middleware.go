package middleware

import "net/http"

type Middleware func(next http.Handler) http.Handler

func (md *Middleware) Var(r *http.Request) map[string]string {
 return make(map[string]string) 
}

func (md *Middleware) Query(r *http.Request) map[string]string {
  return make(map[string]string)
}
