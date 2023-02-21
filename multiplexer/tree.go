package multiplexer

import (
	"net/http"
	"strings"
)

type node struct {
  label string
  handler map[string]handlers
  child map[string]*node
}

type handlers struct {
  handler *http.Handler
}

type tree struct {
  root *node
}


func newNode(path string) *node {
  return &node{
    label : path, 
    handler: make(map[string]handlers),
    child: make(map[string]*node),
  }
}

func newTree(basepath string) *tree {
  nd := &node{
    label: basepath,
    handler: make(map[string]handlers),
    child: make(map[string]*node),
  } 

  return &tree{
    root: nd,
  }
}

func explodePath(path string)[]string {
  listString := strings.Split(path, "/")
  return listString
}


func (tr *tree) insert(label string, handler handlers) {
  path := explodePath(label)
  currentNode := tr.root
  for _, value := range path {
    if value != currentNode.label {
      nd := newNode(value) 
      currentNode.child[value] = nd
      currentNode = nd
    }
  }
}




