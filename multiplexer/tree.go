package multiplexer

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
)

type node struct {
  rgx *regexp.Regexp
  label string
  handler map[string]*handlers
  child map[string]*node
}

type handlers struct {
  handler http.Handler
}

type tree struct {
  root *node
}


func newNode(path string) *node {
  return &node{
    label : path, 
    handler: make(map[string]*handlers),
    child: make(map[string]*node),
    rgx: nil,
  }
}

func newTree(basepath string) *tree {
  path := explodePath(basepath)[0]
  nd := &node{
    label: path,
    handler: make(map[string]*handlers),
    child: make(map[string]*node),
  } 

  return &tree{
    root: nd,
  }
}

func explodePath(path string)[]string {
  listString := strings.Split(path, "/")
  if listString[0] == "" {
    return listString[1:]
  }
  return listString
}


func (tr *tree) insert(label string, handler handlers, method string) {
  path := explodePath(label)
  if len(path) == 1 && path[0] == tr.root.label {
    tr.root.handler[method] = &handler
    return
  }
  currentNode := tr.root
  for i, value := range path {
    childNode, ok := currentNode.child[value] 
    if ok {
      currentNode = childNode
    } else {
      nd := newNode(value)
      currentNode.child[value] = nd
      currentNode = nd
    }
    if i == len(path)-1 {
      currentNode.handler[method] = &handler
    }
  }
}

func (tr *tree) search(path string, method string) (*handlers, error) {
  listPath := explodePath(path) 
  currentNode := tr.root
  for i, value := range listPath {
    _, ok := currentNode.child[value]
    if ok {
      currentNode = currentNode.child[value]    
    }
    if len(listPath)-1 == i {
      if value == currentNode.label {
        _, ok = currentNode.handler[method] 
        if !ok {
          methodError := errors.New("Method not allowed")
          return nil, methodError
        }
        return currentNode.handler[method], nil
      }
    }
  }
  pathError := errors.New("URL path is not found")
  return nil, pathError
}






