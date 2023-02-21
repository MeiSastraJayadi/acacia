package multiplexer

type Router struct {
}

func NewRouter() *Router {
  return &Router{}
}

func (rt *Router) Methods(method ...string) *Router {
  return rt
}

func (rt *Router) SubRouter() *Router {
  return rt
}
