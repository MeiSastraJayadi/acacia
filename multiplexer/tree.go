package multiplexer

import (
	"errors"
	"net/http"
	"regexp"
)

// The node is smallest unit or object in tree
// node held some data
// the data that held by a single node is : 
// -- re : a regular expression
// -- rgx : a map that has key of integers and value of string
//          the key is represnted a number of regular expression type
//          and the value is represented the label of path
// -- handler : a map with key of label (string data type) and with value of handlers object
// -- child : is list of node children. This data is a map with key of string and every key point to value with *node data type
type node struct {
  re *regexp.Regexp
  rgx map[int]string 
  label string
  handler map[string]*handlers
  child map[string]*node
}

// a handlers is just handler that use 
// to handling some path
type handlers struct {
  handler http.Handler
}

type tree struct {
  root *node
}


// newNode will create new node that can be used to build a tree.
// newNode need take a paremeter with string datatype. 
// this parameter is represented a path of string that will be a basepath of node
func newNode(path string) *node {
  return &node{
    label : path, 
    handler: make(map[string]*handlers),
    child: make(map[string]*node),
    rgx: make(map[int]string),
    re: nil,
  }
}


// newTree will create a tree object. This tree object can be used 
// in Router struct. The tree is critical part that will help in routing 
// the path
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

// This function is used to inserting the path, handler, and method in to the tree
// when use function Router.Handler() this function will used to inserting the 
// the node into the router tree
func (tr *tree) insert(label string, handler handlers, method string) {
  path := explodePath(label)
  if len(path) == 1 && path[0] == tr.root.label {
    tr.root.handler[method] = &handler
    return
  }

  if len(path) == 1 && path[0] == "" {
    tr.root.handler[method] = &handler
    return
  }

  currentNode := tr.root
  for i, value := range path {
    childNode, ok := currentNode.child[value] 
    if ok {
      currentNode = childNode
    } else {
      checkIsParams := false
      if value != "" {
        checkIsParams = isWithParams(value)    
      }
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


// Search is used to search path in the tree
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







