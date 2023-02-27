package middleware

import (
	"net/http"

	"github.com/MeiSastraJayadi/acacia/multiplexer"
)

type UseMiddleware struct {
  function Middleware
  router *multiplexer.Router
} 

type Middleware func(next http.Handler) http.Handler

func NewUseMiddleware() *UseMiddleware {
  return &UseMiddleware{
    function: nil,
    router: nil,
  }
}

func (um *UseMiddleware) AddRouter(router *multiplexer.Router) {
  um.router = router
}

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

