package middleware

import "net/http"

// CORS struct is just middleware that can be used to set 
// access control This object containe 4 field. The field that held by this struct is : 
// --> middleware : middleware field is data with type of *UseMiddleware
// --> allowOrigin : this field will held all origin that can be set into CORS. 
//                   the allowOrigin is field with []string data type
// --> allowMethod : this field will held all method that can be set into CORS. 
//                   the allowMethod is field with []string data type
// --> allowHeader : this field will held all header that can be set into CORS. 
//                   the allowHeader is field with []string data type
type CORS struct {
  middleware *UseMiddleware
  allowOrigin []string
  allowMethod []string
  allowHeader []string
}

// NewCORS function will return *CORS
func NewCORS() *CORS {
  newMiddleware := &UseMiddleware{
    function: nil,
    handler: nil,
  }

  return &CORS{
    middleware: newMiddleware,
    allowOrigin: []string{},
    allowMethod: []string{},
    allowHeader: []string{},
  }
}

// AllowOrigins function is used to set all allowed origin 
func (cors *CORS) AllowOrigins(origins ...string) {
  for _, value := range origins {
    cors.allowOrigin = append(cors.allowOrigin, value) 
  }
}

// AllowOrigins function is used to set all allowed methods 
func (cors *CORS) AllowMethods(methods ...string) {
  for _, value := range methods {
    cors.allowMethod = append(cors.allowMethod, value) 
  }
}

// AllowOrigins function is used to set all allowed headers 
func (cors *CORS) AllowHeaders(headers ...string) {
  for _, value := range headers{
    cors.allowHeader = append(cors.allowHeader, value) 
  }
}

func (cors *CORS) corsMiddleware() Middleware {
  return func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      for _, value := range cors.allowHeader {
        w.Header().Add("Access-Control-Allow-Headers", value)
      }
      for _, value := range cors.allowMethod {
        w.Header().Add("Access-Control-Allow-Methods", value)
      }
      for _, value := range cors.allowOrigin {
        w.Header().Add("Access-Control-Allow-Origin", value)
      }
      next.ServeHTTP(w, r)
    })
  }
}

// Function AddHandler will add http.Handler into the CORS 
// Before the request into the handler, the reqeust will be handler by this 
// CORS 
func (cors *CORS) AddHandler(handler http.Handler) {
  cors.middleware.handler = handler
}

func (cors *CORS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  cors.middleware.function = cors.corsMiddleware()
  cors.middleware.ServeHTTP(w, r)
}

