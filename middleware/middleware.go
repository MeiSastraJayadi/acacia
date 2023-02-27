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

func NewUseMiddleware(router *multiplexer.Router, middleware Middleware) *UseMiddleware {
  return &UseMiddleware{
    function: middleware,
    router: router,
  }
}

func (mw *UseMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  hdl := mw.function(mw.router)
  hdl.ServeHTTP(w, r)
}

