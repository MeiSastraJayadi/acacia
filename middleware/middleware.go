package middleware

import (
	"net/http"

	"github.com/MeiSastraJayadi/acacia/multiplexer"
)


// UseMiddleware struct can create an object to apply UseMiddleware
// in specific *multiplexer.Router object. This struct has two field. 
// The first field is : ---> function, which is a Middleware Function
// The second one is : --> router, which is a *multiplexer.Router where the middleware will be applied
type UseMiddleware struct {
  function Middleware
  router *multiplexer.Router
} 

// Middleware is a function that need one parameter with type of http.Handler 
// and will return a http.Handler
type Middleware func(next http.Handler) http.Handler

// NewUseMiddleware will return UseMiddleware object
func NewUseMiddleware() *UseMiddleware {
  return &UseMiddleware{
    function: nil,
    router: nil,
  }
}

// AddRouter function will add new *multiplexer.Router into the 
// UseMiddleware object
func (um *UseMiddleware) AddRouter(router *multiplexer.Router) {
  um.router = router
}

// AddMiddleware function will add new Middleware function into the 
// UseMiddleware object
func (um *UseMiddleware) AddMiddleware(middleware Middleware) {
  um.function = middleware 
}

func (mw *UseMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if mw.function != nil && mw.router != nil {
    hdl := mw.function(mw.router)
    hdl.ServeHTTP(w, r)
  }
  return
}

