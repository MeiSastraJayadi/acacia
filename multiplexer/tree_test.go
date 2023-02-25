package multiplexer

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type randomHandler struct {
}

func (rand *randomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func TestTree(t *testing.T) {
  path := strings.Split("/product/1", "/")
  if path[0] == "" {
    fmt.Println("yes")
  }

  for i, value := range path {
    fmt.Println(i, value)
  }
  fmt.Println(path)
}

func TestNode(t *testing.T) {
  hand := &randomHandler{}
  hdl := handlers{
    handler : hand,
  }
  tr := newTree("/")
  tr.root.handler[http.MethodGet] = &hdl 
  tr.insert("/product/get-all", hdl, http.MethodGet)
  _, ok := tr.root.child["product"]
  assert.Equal(t, true, ok)
  _, ok = tr.root.child["product"].child["get-all"]
  assert.Equal(t, true, ok)
}

func TestCheckHandler(t *testing.T) {
  hand := &randomHandler{}
  hdl := handlers{
    handler : hand,
  }
  tr := newTree("/")
  tr.root.handler[http.MethodGet] = &hdl 
  tr.insert("/product/get-all", hdl, http.MethodGet)
  result := tr.root.child["product"]
  assert.Equal(t, 1, len(result.child))
  result = result.child["get-all"]
  _, ok := result.handler[http.MethodGet]
  assert.Equal(t, true, ok)
}

func TestSearch(t *testing.T) {
  hand := &randomHandler{}
  hdl := handlers{
    handler : hand,
  }
  tr := newTree("/")
  tr.root.handler[http.MethodGet] = &hdl 
  tr.insert("/product/get-all", hdl, http.MethodGet)
  result, _ := tr.search("/product/get-all", http.MethodGet)
  assert.Equal(t, &hdl, result)
}







