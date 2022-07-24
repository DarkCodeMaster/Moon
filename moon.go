package moon

import (
	"fmt"
	"net/http"
)

// HandlerFunc 响应HTTP请求的Handler函数
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// Engine 实现 ServerHTTP 接口
type Engine struct {
	router map[string]HandlerFunc
}

func (engine *Engine) addRoute(method string, path string, handler HandlerFunc) {
	key := method + "-" + path
	engine.router[key] = handler
}

func (engine *Engine) GET(path string, handler HandlerFunc) {
	engine.addRoute("GET", path, handler)
}

func (engine *Engine) POST(path string, handler HandlerFunc) {
	engine.addRoute("POST", path, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}

func New() *Engine {
	engine := new(Engine)
	engine.router = make(map[string]HandlerFunc)
	return engine
}
