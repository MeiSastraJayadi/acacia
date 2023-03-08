package multiplexer

import (
	"net/http"
)

func (rt *Router) searchVariabel(path string) map[string]string {
  listvars := make(map[string]string)
  splitPath := explodePath(path) 
  currentNode := rt.tree.root
  for i, value := range splitPath {
    _, ok := currentNode.child[value]
    if ok {
      currentNode = currentNode.child[value]
    } else { 
      for key, label := range currentNode.rgx {
        checkPath := regMap[key].MatchString(value)
        if checkPath {
          currentNode = currentNode.child[label]
          listvars[label] = value
          break
        }
      }
    } 

    if len(splitPath)-1 == i {
      if currentNode.re != nil {
        checkRe := currentNode.re.MatchString(value)
        if checkRe {
          listvars[currentNode.label] = value
        }
      } 
    }
  }
  return listvars
}

func Vars(r *http.Request) map[string]string {
  variabel := r.Context().Value("vars")
  if variabel == nil {
    return make(map[string]string)
  }
  return variabel.(map[string]string)
}

func (rt *Router) Query(r *http.Request, key string) string {
  query := r.URL.Query()
  return query[key][0]
}


