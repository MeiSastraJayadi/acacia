package multiplexer

import "net/http"

type Router struct {
  label string
  prefix string
  method []string
  child []*Router
}

func NewRouter() *Router {
  return &Router{
    label: "",
    prefix: "", 
    method: []string{},
    child: []*Router{},
  }
}

func (rt *Router) Methods(methods ...string) *Router {
  newRoute := NewRouter()
  for _, m := range methods {
    newRoute.method = append(newRoute.method, m)
  }
  rt.child = append(rt.child, newRoute)
  return rt 
}

func (rt *Router) SubRouter() *Router {
  return rt.child[len(rt.child)-1]
}

func (rt *Router) Handle(path string, handler http.Handler) {
}

func (rt *Router) HandleFunc(path string, handler http.HandlerFunc) {
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}




