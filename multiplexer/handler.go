package multiplexer

import (
	"net/http"
)

type routerHandler struct {
  handlerFunction http.HandlerFunc
}

func (rh *routerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  rh.handlerFunction(w, r)
}

func (rt *Router) Handle(path string, handler http.Handler) {
  if len(rt.saver.methods) < 1 {
    rt.saver.methods = append(rt.saver.methods, http.MethodGet)
  }

  hdl := handlers{
    handler: handler,
  }

  for _, item := range rt.saver.methods {
    rt.tree.insert(path, hdl, item)
  }
  rt.saver.methods = []string{}
}

func (rt *Router) HandleFunc(path string, handlerFunc http.HandlerFunc) {
  rh := &routerHandler{
    handlerFunction: handlerFunc,
  }
  rt.Handle(path, rh)
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
  url := r.URL.RawPath 
  method := r.Method
  handler, err := rt.tree.search(url,method)
  if err != nil {
    return err
  }
  handler.handler.ServeHTTP(w, r)
  return nil
}

