package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Archive     string
}

type Routes []Route

var routes = Routes{

	Route{
		"Formulario",
		"GET",
		"/formulario",
		FormularioHandler,
		"public/formulario",
	},
}