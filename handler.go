package lymon

import (
	"log"
	"net/http"
)

// HandleFunc A
func (c *Context) HandleFunc(pattern, method string, handler func(http.ResponseWriter, *http.Request, Context)) {
	// patternID contain combination of user uri pattern and given method
	// serveHTTP will generate this patternID too jsjfwjfk
	patternID := pattern + "#" + method

	// panic if patternID already exist in h.Path
	if _, ok := c.Path[patternID]; ok {
		log.Panicf("Failed to add %v : Duplicate route pattern \n", pattern)
	} else {
		// register given pattern and handler to h.Path
		c.Path[patternID] = route{
			Handler: handler,
			Method:  method,
		}
	}
}

// BeforeAll Before filters are evaluated before each request within the same context as the routes.
// They can modify the request and response.
func (c *Context) BeforeAll(handler func(http.ResponseWriter, *http.Request, Context)) {
	c.Middleware = append(c.Middleware, handler)
}

func (c Context) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if len(c.Middleware) > 0 {
		c.Middleware[0](w, r, c)
	}

	// Check is requested are registered in h.Path
	// otherwise return 404 page
	patternID := r.URL.Path + "#" + r.Method
	if val, ok := c.Path[patternID]; ok {
		val.Handler(w, r, c)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404"))
	}
}

func (c Context) Start() {
	http.ListenAndServe(c.Config.Listen, c)
}
