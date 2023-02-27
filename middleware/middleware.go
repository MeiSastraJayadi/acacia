package middleware

import (
	"net/http"
)


// UseMiddleware struct can create an object to apply UseMiddleware
// in specific *multiplexer.Router object. This struct has two field. 
// The first field is : ---> function, which is a Middleware Function
// The second one is : --> handler, which is a http.Handler where the middleware will be applied
type UseMiddleware struct {
  function Middleware
  handler http.Handler 
} 

// Middleware is a function that need one parameter with type of http.Handler 
// and will return a http.Handler
type Middleware func(next http.Handler) http.Handler

// NewUseMiddleware will return UseMiddleware object
func NewUseMiddleware() *UseMiddleware {
  return &UseMiddleware{
    function: nil,
    handler: nil,
  }
}

// AddRouter function will add new http.Handler into the 
// UseMiddleware object
func (um *UseMiddleware) AddHandler (handler http.Handler) {
  um.handler = handler 
}

// AddMiddleware function will add new Middleware function into the 
// UseMiddleware object
func (um *UseMiddleware) AddMiddleware(middleware Middleware) {
  um.function = middleware 
}

func (mw *UseMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if mw.function != nil && mw.handler != nil {
    hdl := mw.function(mw.handler)
    hdl.ServeHTTP(w, r)
  }
  return
}

