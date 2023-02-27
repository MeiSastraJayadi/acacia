package middleware

import "net/http"

type CORS struct {
  middleware *UseMiddleware
  allowOrigin []string
  allowMethod []string
  allowHeader []string
}


func NewCORS() *CORS {
  newMiddleware := &UseMiddleware{
    function: nil,
    router: nil,
  }

  return &CORS{
    middleware: newMiddleware,
    allowOrigin: []string{},
    allowMethod: []string{},
    allowHeader: []string{},
  }
}

func (cors *CORS) AllowOrigins(origins ...string) {
  for _, value := range origins {
    cors.allowOrigin = append(cors.allowOrigin, value) 
  }
}

func (cors *CORS) AllowMethods(methods ...string) {
  for _, value := range methods {
    cors.allowMethod = append(cors.allowMethod, value) 
  }
}

func (cors *CORS) AllowHeaders(headers ...string) {
  for _, value := range headers{
    cors.allowHeader = append(cors.allowHeader, value) 
  }
}

func (cors *CORS) corsMiddleware() Middleware {
  return func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      for _, value := range cors.allowHeader {
        r.Header.Add("Access-Control-Allow-Headers", value)
      }
      for _, value := range cors.allowMethod {
        r.Header.Add("Access-Control-Allow-Methods", value)
      }
      for _, value := range cors.allowOrigin {
        r.Header.Add("Access-Control-Allow-Origin", value)
      }
      next.ServeHTTP(w, r)
    })
  }
}

func (cors *CORS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  cors.middleware.function = cors.corsMiddleware()
  cors.middleware.ServeHTTP(w, r)
}

