package multiplexer

import (
	"net/http"
)

// routerHandler is just a struct that use 
// http.Handler interface. Actually this struct is use to 
// save an handlerfunction and add ServeHTTP functionality
type routerHandler struct {
  handlerFunction http.HandlerFunc
}

// ServeHTTP is a function to fullfill the contract so the type of
// routerHandler can be a http.Handler type 
func (rh *routerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  rh.handlerFunction(w, r)
}

// Handle function is used to add handler in specific path to the router.  
// This function take two parameters. the first parameter is path with type of 
// string and the second paramater is handler which is a http.Handler
// The path that will be inserted to the router can have paremeter.
// To do that, the path should be written like this 
// example : 
// -- /home/user/{id:int} --> id is a parameter in path that just accept numeric value
// -- /home/{slug_name:slug} --> slug_name is a parameter in path that just accept slug 
// You can use int, float, slug, and string in parameters
// Every parameter will return a string value
func (rt *Router) Handle(path string, handler http.Handler) {
  if len(rt.saver.methods) < 1 {
    rt.saver.methods = append(rt.saver.methods, http.MethodGet)
  }

  hdl := handlers{
    handler: handler,
  }

  for _, item := range rt.saver.methods {
    rt.tree.insert(path, hdl, item)
  }
  rt.saver.methods = []string{}
}

// HandleFunc function is used to add handlerfunction in specific path to the router.  
// This function take two parameters. the first parameter is path with type of 
// string and the second paramater is a http.handlerFunc 
// The path that will be inserted to the router, can have paremeter.
// To do that, the path should be written like this 
// example : 
// -- /home/user/{id:int} --> id is a parameter in path that just accept numeric value
// -- /home/{slug_name:slug} --> slug_name is a parameter in path that just accept slug 
// You can use int, float, slug, and string in parameters
// Every parameter will return a string value
func (rt *Router) HandleFunc(path string, handlerFunc http.HandlerFunc) {
  rh := &routerHandler{
    handlerFunction: handlerFunc,
  }
  rt.Handle(path, rh)
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
  url := r.URL.RawPath 
  method := r.Method
  handler, err := rt.tree.search(url,method)
  if err != nil {
    return err
  }
  handler.handler.ServeHTTP(w, r)
  return nil
}

