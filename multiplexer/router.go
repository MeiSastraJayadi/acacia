package multiplexer

import (
	"errors"
)

// Router struct is router 
// that will help to routing every incoming request.
// This router used tree data structure to do request routin
type Router struct {
  prefix string
  saver saveInformation
  tree *tree
}


type saveInformation struct {
  methods []string
}

// NewRouter will return *Router. 
func NewRouter(basepath string) *Router {
  tree := newTree(basepath)
  return &Router{
    prefix: "",
    tree: tree,
    saver: saveInformation{},
  }
}

// SetPrefix will set prefix for Router object.
// Every path will be start with prefix to access it
// example : 
// the original path --> /id
// with prefix = "product" => product/id
func (rt *Router) SetPrefix(prefix string) *Router {
  if prefix == "" || prefix == "/" {
    return nil
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
  return rt
}

// Set every http method that will assign to the path 
// that will be added into Router object
func (rt *Router) Methods(methods ...string) *Router {
  for _, method := range methods {
    rt.saver.methods = append(rt.saver.methods, method)
  }
  return rt
}


// Function SubRouter will add sub router to main router
func (rt *Router) SubRouter(router *Router) error {
  if rt.tree.root.label == router.tree.root.label {
    return errors.New("Router and SubRouter has same root path, make sure router and subrouter has different root path")
  }
  rt.tree.root.child[router.tree.root.label] = router.tree.root
  return nil
}




