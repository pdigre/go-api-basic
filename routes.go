package main

import (
	"net/http"

	"github.com/gilcrest/go-api-basic/handler"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func newRouter(hdl *handler.AppHandler) *mux.Router {
	// create a new mux (multiplex) router
	rtr := mux.NewRouter()

	// send Router through PathPrefix method to validate any standard
	// subroutes you may want for your APIs. e.g. I always want to be
	// sure that every request has "/api" as part of it's path prefix
	// without having to put it into every handle path in my various
	// routing functions
	rtr = rtr.PathPrefix("/api").Subrouter()

	// Match only POST requests at /api/v1/movies
	// with Content-Type header = application/json
	rtr.Handle("/v1/movies",
		alice.New(
			hdl.SetAccessToken2Context,
			hdl.AddStandardResponseHeaders,
			hdl.SetStandardResponseFields).
			Then(http.HandlerFunc(hdl.AddMovie))).
		Methods("POST").
		Headers("Content-Type", "application/json")

	// Match only PUT requests having an ID at /api/v1/movies/{id}
	// with the Content-Type header = application/json
	rtr.Handle("/v1/movies/{id}",
		alice.New(
			hdl.SetAccessToken2Context,
			hdl.AddStandardResponseHeaders,
			hdl.SetStandardResponseFields).
			Then(http.HandlerFunc(hdl.UpdateMovie))).
		Methods("PUT").
		Headers("Content-Type", "application/json")

	// Match only DELETE requests having an ID at /api/v1/movies/{id}
	rtr.Handle("/v1/movies/{id}",
		alice.New(
			hdl.SetAccessToken2Context,
			hdl.AddStandardResponseHeaders,
			hdl.SetStandardResponseFields).
			Then(http.HandlerFunc(hdl.DeleteMovie))).
		Methods("DELETE")

	// Match only GET requests having an ID at /api/v1/movies/{id}
	// with the Content-Type header = application/json
	rtr.Handle("/v1/movies/{id}",
		alice.New(
			hdl.SetAccessToken2Context,
			hdl.AddStandardResponseHeaders,
			hdl.SetStandardResponseFields).
			Then(http.HandlerFunc(hdl.FindByID))).
		Methods("GET")

	// Match only GET requests /api/v1/movies
	// with the Content-Type header = application/json
	rtr.Handle("/v1/movies",
		alice.New(
			hdl.SetAccessToken2Context,
			hdl.AddStandardResponseHeaders,
			hdl.SetStandardResponseFields).
			Then(http.HandlerFunc(hdl.FindAll))).
		Methods("GET")

	return rtr
}
