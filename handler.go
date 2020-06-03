package lymon

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/asaskevich/govalidator"
)

// Context will hold single lymon route state
type Context struct {
	W http.ResponseWriter
	R *http.Request
	V additional
}

// handler hold func type that will used by end-user
type handler = func(Context, *Global)

type additional struct {
	// IsValidated if any validation not applied into route the value will set as true
	IsValidated bool

	// return unmarshalled r.Body or r.Form
	Map map[string]interface{}
}

type route struct {
	Handler   handler
	Method    string
	Validator map[string]interface{}
}

// HandleFunc registers the handler function for the given pattern.
func (g *Global) HandleFunc(pattern, method string, handler handler, validator ...map[string]interface{}) {
	// patternID contain combination of user uri pattern and given method
	// serveHTTP will generate this patternID
	patternID := pattern + "#" + method

	// panic if patternID already exist in h.Path
	if _, ok := g.Path[patternID]; ok {
		log.Panicf("Failed to add %v : Duplicate route pattern \n", pattern)
	} else {

		// register given pattern and handler to h.Path
		route := route{
			Handler:   handler,
			Method:    method,
			Validator: nil,
		}

		if len(validator) > 0 {
			route.Validator = validator[0]
		}

		g.Path[patternID] = route
	}
}

// BeforeAll Before filters are evaluated before each request within the same context as the routes.
// They can modify the request and response.
func (g *Global) BeforeAll(handler handler) {
	g.MiddlewareHandler = append(g.MiddlewareHandler, handler)
}

// HandleStatusCode with this middleware, You can customize the built-in error response, currently only works for 404
func (g *Global) HandleStatusCode(StatusCode int, handler handler) {
	// panic if StatusCode already exist in h.StatusCodeHandler
	if _, ok := g.StatusCodeHandler[StatusCode]; ok {
		log.Panicf("Failed to add %v handler : Status code already in use \n", StatusCode)
	} else {
		// register given pattern and handler to h.Path
		g.StatusCodeHandler[StatusCode] = handler
	}
}

// ServeHTTP called from net/http that will calls f(w, r).
func (g *Global) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// execute middleware based on array order
	for _, middleware := range g.MiddlewareHandler {
		middleware(Context{
			W: w,
			R: r,
		}, g)
	}

	// Check is requested are registered in h.Path
	// otherwise return 404 page
	patternID := r.URL.Path + "#" + r.Method
	if handlerInstance, ok := g.Path[patternID]; ok {

		private := Context{
			W: w,
			R: r,
			V: additional{
				IsValidated: true,
			},
		}

		if r.Method == "POST" {
			if handlerInstance.Validator != nil {

				body, err := ioutil.ReadAll(r.Body)
				if err == nil {
					// refill r.Body for FormToMAP uses
					r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

					var result map[string]interface{}
					err = json.Unmarshal(body, &result)
					if err != nil {
						log.Println("validator unmarshall : ", err)
						result = FormToMAP(r, g.Config.MaxMultipartMemory)
					}

					private.V.Map = result
					private.V.IsValidated, err = govalidator.ValidateMap(result, handlerInstance.Validator)
					if err != nil {
						log.Println("validation error : ", err)
					}
				}
				// refill r.Body for user uses
				r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		}

		handlerInstance.Handler(private, g)
	} else {
		// check if StatusCode already exist in h.StatusCodeHandler
		if _, ok := g.StatusCodeHandler[404]; ok {
			g.StatusCodeHandler[404](Context{
				W: w,
				R: r,
			}, g)
		} else {
			// default 404 response
			w.WriteHeader(404)
			w.Write([]byte("404"))
		}
	}
}

// Start invoke http.ListenAndServe
func (g *Global) Start() {
	http.ListenAndServe(g.Config.Listen, g)
}
