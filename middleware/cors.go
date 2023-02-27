package middleware

type CORS struct {
  *UseMiddleware
  allow []string
}
