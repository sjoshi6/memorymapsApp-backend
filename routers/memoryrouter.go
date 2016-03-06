package routers

import (
	"memorymaps-backend/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

var t controllers.TextMemory

var textmemoryroutes = Routes{
	Route{
		Name:        "Create Memory",
		Method:      "POST",
		Pattern:     "/v1/textmemory",
		HandlerFunc: t.Create,
	},
}

// AddTextMemoryRoutes : Add Text Memory Routes
func AddTextMemoryRoutes(r *mux.Router) *mux.Router {

	for _, route := range textmemoryroutes {

		var handler http.Handler
		handler = route.HandlerFunc

		r.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return r
}
