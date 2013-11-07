package martini

import (
	"net/http"
)

type Router interface {
	Get(string, ...Handler)
	Post(string, ...Handler)
	Put(string, ...Handler)
	Delete(string, ...Handler)

	Handle(http.ResponseWriter, *http.Request, Context)
}

type router struct {
	routes []route
}

func NewRouter() Router {
	return &router{}
}

func (r *router) Get(pattern string, h ...Handler) {
	r.addRoute("GET", pattern, h)
}

func (r *router) Post(pattern string, h ...Handler) {
	r.addRoute("POST", pattern, h)
}

func (r *router) Put(pattern string, h ...Handler) {
	r.addRoute("PUT", pattern, h)
}

func (r *router) Delete(pattern string, h ...Handler) {
	r.addRoute("DELETE", pattern, h)
}

func (r *router) Handle(res http.ResponseWriter, req *http.Request, context Context) {
	for _, route := range r.routes {
		// Be super strict for now. Eventually we will have some
		// super awesome pattern matching here. But not today
		if route.method == req.Method && req.URL.Path == route.pattern {
			err := context.Invoke(route.handle)
			if err != nil {
				panic(err)
			}
			return
		}
	}

	// no routes exist, 404
	res.WriteHeader(http.StatusNotFound)
}

func (r *router) addRoute(method string, pattern string, handlers []Handler) {
	route := route{method, pattern, handlers}
	if route.validate() == nil {
		r.routes = append(r.routes, route)
	}
}

type route struct {
	method   string
	pattern  string
	handlers []Handler
}

func (r route) validate() error {
	for _, handler := range r.handlers {
		err := validateHandler(handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r route) handle(c Context) {
	for _, handler := range r.handlers {
		err := c.Invoke(handler)
		if err != nil {
			panic(err)
		}
		if c.written() {
			return
		}
	}
}