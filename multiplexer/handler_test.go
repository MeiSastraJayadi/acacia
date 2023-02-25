package multiplexer

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func SampleHandlerFunc(w http.ResponseWriter, r *http.Request) {

}

func TestRouterInsert(t *testing.T) {
  rt := NewRouter("/")
  rt.Methods(http.MethodGet).HandleFunc("/product", SampleHandlerFunc)
  nd := rt.tree.root.child["product"]
  _, ok := nd.handler[http.MethodGet]
  assert.Equal(t, "product", nd.label)
  require.Equal(t, true, ok)
}

func TestRouterInsertwithRegex(t *testing.T) {
  rt := NewRouter("/")
  rt.Methods(http.MethodGet).HandleFunc("/product/{id:int}", SampleHandlerFunc)
  nd := rt.tree.root.child["product"]
  _, ok := nd.child["id"]
  assert.NotZero(t, len(nd.child))
  assert.Equal(t, true, ok)
  lbl := nd.rgx[2]
  assert.Equal(t, "id", lbl)
}

func TestRouterSearch(t *testing.T) {
  rt := NewRouter("/")
  rt.Methods(http.MethodGet).HandleFunc("/product/{id:int}", SampleHandlerFunc)
  hdl, err := rt.tree.search("/product/1", http.MethodGet)
  assert.Nil(t, err)
  assert.NotNil(t, hdl)
  hdl, err = rt.tree.search("/product/mei-sastra", http.MethodGet)
  assert.Nil(t, hdl)
  assert.NotNil(t, err)
}






