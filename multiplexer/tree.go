package multiplexer

import (
	"errors"
	"net/http"
	"regexp"
)

type node struct {
  re *regexp.Regexp
  rgx map[int]string 
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
    rgx: make(map[int]string),
    re: nil,
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
      checkIsParams := isWithParams(value)    
      if !checkIsParams {
        // If path is not a params
        nd := newNode(value)
        currentNode.child[value] = nd
        currentNode = nd
      } else {
        // if path is a params with key and type
        re := selectRegex(value)
        lbl := getKey(value)
        nd := newNode(lbl)
        nd.re = regMap[re]
        currentNode.rgx[re] = lbl
        currentNode.child[lbl] = nd
        currentNode = nd
      }
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
    } else {
      // Check if regex exist
      for key, lbl := range currentNode.rgx {
        checkPath := regMap[key].MatchString(value)
        if checkPath {
          currentNode = currentNode.child[lbl]
          break
        }
      }
    }

    if len(listPath)-1 == i {
      if value == currentNode.label {
        _, ok = currentNode.handler[method] 
        if !ok {
          methodError := errors.New("Method not allowed")
          return nil, methodError
        }
        return currentNode.handler[method], nil
      } else if currentNode.re != nil {
        checkRe := currentNode.re.MatchString(value)
        if checkRe {
          _, ok = currentNode.handler[method] 
          if !ok {
            methodError := errors.New("Method not allowed")
            return nil, methodError
          }
          return currentNode.handler[method], nil
        }
      } 
    }
  }
  pathError := errors.New("URL path is not found")
  return nil, pathError
}







