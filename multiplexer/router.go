package multiplexer

import (
	"net/http"
)

type Router struct {
  prefix string
  tree *tree
}


type saveInformation struct {
  label string
  methods []string
  handler []*handlers
}

var saver = &saveInformation{}

func NewRouter(basepath string) *Router {
  tree := newTree(basepath)
  return &Router{
    prefix: "",
    tree: tree,
  }
}

func (rt *Router) SetPrefix(prefix string) {
  rt.prefix = prefix
}

func (rt *Router) Methods(methods ...string) *Router {
  for _, method := range methods {
    saver.methods = append(saver.methods, method)
  }
  return rt
}

func (rt *Router) SubRouter(router *Router) {
  // label := router.tree.root.label
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}



