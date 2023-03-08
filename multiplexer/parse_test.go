package multiplexer

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRouterVars(t *testing.T) {
  router := NewRouter("/")
  hand := &randomHandler{}
  hdl := handlers{
    handler : hand,
  }
  tr := router.tree
  router.tree.insert("/product/get-all/{id:int}", hdl, http.MethodGet)
  router.tree.insert("/product/get-all/{id:int}/data/{slug:slug}", hdl, http.MethodGet)
  router.tree.insert("/{slug:slug}", hdl, http.MethodGet)
  vr := router.searchVariabel("product/get-all/3/data/mei-sastra-jayadi")
  assert.Equal(t, "3", vr["id"])
  assert.Equal(t, "mei-sastra-jayadi", vr["slug"])
  han, errSearch := tr.search("/co-founder", http.MethodGet)
  require.Nil(t, errSearch)
  assert.NotNil(t, han)
}

func TestSlugRouter(t *testing.T) {
  router := NewRouter("/")
  subrouter := NewRouter("/").SetPrefix("roles")
  subrouter.Methods(http.MethodGet).HandleFunc("/", func(w http.ResponseWriter, r *http.Request){})
  subrouter.Methods(http.MethodPost).HandleFunc("/", func(w http.ResponseWriter, r *http.Request){})
  subrouter.Methods(http.MethodGet).HandleFunc("/{role_slug:slug}", func(w http.ResponseWriter, r *http.Request){})
  err := router.SubRouter(subrouter)
  require.Nil(t, err)
  handl, searchErr := router.tree.search("/roles/co-founder", http.MethodGet)
  require.Nil(t, searchErr)
  assert.NotNil(t, handl)
}

