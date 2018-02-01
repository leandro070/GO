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
