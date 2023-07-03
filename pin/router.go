package pin

import (
	"fmt"
	"log"
)

type router struct {
	handlers map[string]handlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]handlerFunc)}
}

func (r *router) addRoute(patter string, method string, handler handlerFunc) {
	log.Printf("add Route %s - %s", method, patter)
	key := method + "_" + patter
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "_" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		fmt.Fprintln(c.Writer, "error path , 404 not found")
	}
}
