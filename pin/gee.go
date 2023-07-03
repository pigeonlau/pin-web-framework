package pin

import (
	"net/http"
)

type handlerFunc = func(ctx *Context)
type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) Get(patter string, handler handlerFunc) {
	e.addRoute(patter, "GET", handler)

}

func (e *Engine) Post(patter string, handler handlerFunc) {
	e.addRoute(patter, "POST", handler)

}

func (e *Engine) Delete(patter string, handler handlerFunc) {
	e.addRoute(patter, "DELETE", handler)

}

func (e *Engine) Put(patter string, handler handlerFunc) {

	e.addRoute(patter, "PUT", handler)
}

func (e *Engine) addRoute(pattern, method string, handler handlerFunc) {
	e.router.addRoute(pattern, method, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.router.handle(newContext(w, r))
}

func (e *Engine) Run(address string) {
	http.ListenAndServe(address, e)
}
