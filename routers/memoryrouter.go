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
		HandlerFunc: CORSHandler(t.Create),
	},
	Route{
		Name:        "Get All Text Memories",
		Method:      "GET",
		Pattern:     "/v1/textmemories",
		HandlerFunc: CORSHandler(t.GetAll),
	},
	Route{
		Name:        "Options Handler",
		Method:      "OPTIONS",
		Pattern:     "/v1/{rest:[a-zA-Z0-9]+}",
		HandlerFunc: OptionsHandler,
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
