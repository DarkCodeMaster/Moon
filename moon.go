package moon

import (
	"net/http"
)

// HandlerFunc 响应HTTP请求的Handler函数
type HandlerFunc func(*Context)

// Engine 实现 ServerHTTP 接口
type Engine struct {
	router *router
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handle(c)
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	// ListenAndServe方法接收的第二个参数为Handler接口
	// Handler接口只有一个方法ServeHTTP
	// 在收到HTTP请求时会使用实现的ServeHTTP方法进行处理
	return http.ListenAndServe(addr, engine)
}

func New() *Engine {
	engine := new(Engine)
	engine.router = newRouter()
	return engine
}
