package multiplexer

import (
	"errors"
)

type Router struct {
  prefix string
  saver saveInformation
  tree *tree
}


type saveInformation struct {
  methods []string
}


func NewRouter(basepath string) *Router {
  tree := newTree(basepath)
  return &Router{
    prefix: "",
    tree: tree,
    saver: saveInformation{},
  }
}

func (rt *Router) SetPrefix(prefix string) {
  if prefix == "" || prefix == "/" {
    return
  }
  rt.prefix = prefix
  if rt.tree.root.label == "" {
    rt.tree.root.label = prefix
  } else {
    currentTree := rt.tree.root
    newNode := newNode(prefix) 
    newNode.child[currentTree.label] = currentTree
    rt.tree.root = newNode
  }
}

func (rt *Router) Methods(methods ...string) *Router {
  for _, method := range methods {
    rt.saver.methods = append(rt.saver.methods, method)
  }
  return rt
}

func (rt *Router) SubRouter(router *Router) error {
  if rt.tree.root.label == router.tree.root.label {
    return errors.New("Router and SubRouter has same root path, make sure router and subrouter has different root path")
  }
  rt.tree.root.child[router.tree.root.label] = router.tree.root
  return nil
}



