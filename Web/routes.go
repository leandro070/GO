package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		IndexHandler,
	},
	Route{
		"404",
		"GET",
		"/NotFound",
		NotFoundHandler,
	},
	Route{
		"Formulario",
		"GET",
		"/formulario",
		FormularioHandler,
	},
	Route{
		"Provincia",
		"GET",
		"/formulario/provincia",
		ProvinciaHandler,
	},
	Route{
		"Persona",
		"POST",
		"/formulario/persona",
		PersonaHandler,
	},
}
