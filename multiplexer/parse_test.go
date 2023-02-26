package multiplexer

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouterVars(t *testing.T) {
  router := NewRouter("/")
  hand := &randomHandler{}
  hdl := handlers{
    handler : hand,
  }
  router.tree.insert("/product/get-all/{id:int}", hdl, http.MethodGet)
  router.tree.insert("/product/get-all/{id:int}/data/{slug:slug}", hdl, http.MethodGet)
  vr := router.searchVariabel("product/get-all/3/data/mei-sastra-jayadi")
  assert.Equal(t, "3", vr["id"])
  assert.Equal(t, "mei-sastra-jayadi", vr["slug"])
}


