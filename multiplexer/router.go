package multiplexer

import "net/http"

type Router struct {
  label string
  methods map[string]action
  child map[string]*Router
}

type action struct {
  handler *http.Handler
}

func NewRouter(prefix string) *Router {
  return &Router{
    label : prefix, 
    methods: make(map[string]action),
    child: make(map[string]*Router),
  }
}

func (rt *Router) Methods(methods ...string) *Router {
  newRouter := NewRouter("")
  rt.child = append(rt.child, newRouter)
  return rt
}

func (rt *Router) SubRouter() *Router {
  subRouter := rt.child[len(rt.child)-1]
  if len(rt.child) < 2 {
    rt.child = []*Router{}
  } else {
    rt.child = rt.child[1:]
  }
  return subRouter
}

func (rt *Router) Handle(path string, handler http.Handler) {
}

func (rt *Router) HandleFunc(path string, handler http.HandlerFunc) {
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}




