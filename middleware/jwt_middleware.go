package middleware

import "net/http"

type UseJWTMiddleware struct {
  middleware *UseMiddleware
  privateKey string
}

func veriyJWT(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { 
  }) 
} 

func NewJWTMiddleware(privateKey string) *UseJWTMiddleware {
  return &UseJWTMiddleware{
    middleware: &UseMiddleware{
    },
    privateKey: privateKey,
  }
}

