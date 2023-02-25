package multiplexer

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRouter(t *testing.T) {
  router := NewRouter("/")
  assert.Equal(t, router.tree.root.label, "")
}

func TestSubRouter(t *testing.T) {
  router := NewRouter("/")
  hand := &randomHandler{}
  hdl := handlers{
    handler : hand,
  }
  router.tree.insert("/product/get-all", hdl, http.MethodGet)
  subrouter := NewRouter("/")
  assert.Equal(t, "", subrouter.tree.root.label)
  subrouter.SetPrefix("category")
  require.Equal(t, "category", subrouter.tree.root.label)
  assert.Equal(t, 0, len(subrouter.tree.root.child))
  subrouter.tree.insert("/", hdl, http.MethodGet)
  router.SubRouter(subrouter)
  _, ok := router.tree.root.child["category"]
  assert.Equal(t, true, ok)
  assert.Equal(t, 2, len(router.tree.root.child))
}

func TestSubRouter2(t *testing.T) {
  router := NewRouter("/")
  hand := &randomHandler{}
  hdl := handlers{
    handler : hand,
  }
  router.tree.insert("/product/get-all", hdl, http.MethodGet)
  subrouter := NewRouter("/gadget")
  subrouter.tree.insert("gadget", hdl, http.MethodGet)
  subrouter.SetPrefix("category")
  assert.Equal(t, "category", subrouter.tree.root.label)
  _, ok := subrouter.tree.root.child["gadget"]
  assert.Equal(t, true, ok)
  router.SubRouter(subrouter)
  get := router.tree.root.child["category"]
  assert.Equal(t, subrouter.tree.root, get)
  hd := get.child["gadget"].handler[http.MethodGet]
  assert.Equal(t, &hdl, hd)
}

func TestReturnRegex(t *testing.T) {
  result := selectRegex("{name:slug}")
  assert.Equal(t, regSlug, result)
  result = selectRegex("{name:string}")
  assert.Equal(t, regString, result)
}














