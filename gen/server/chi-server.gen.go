// Package server provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /plugins)
	ListPlugins(w http.ResponseWriter, r *http.Request, params ListPluginsParams)

	// (GET /plugins/{name})
	FetchPlugin(w http.ResponseWriter, r *http.Request, name string)

	// (GET /users)
	ListUsers(w http.ResponseWriter, r *http.Request, params ListUsersParams)

	// (POST /users)
	CreateUser(w http.ResponseWriter, r *http.Request)

	// (POST /users/token)
	CreateToken(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ListPlugins operation middleware
func (siw *ServerInterfaceWrapper) ListPlugins(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, "BearerAuth.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListPluginsParams

	// ------------- Optional query parameter "limit" -------------
	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter limit: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "page" -------------
	if paramValue := r.URL.Query().Get("page"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "page", r.URL.Query(), &params.Page)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter page: %s", err), http.StatusBadRequest)
		return
	}

	siw.Handler.ListPlugins(w, r.WithContext(ctx), params)
}

// FetchPlugin operation middleware
func (siw *ServerInterfaceWrapper) FetchPlugin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameter("simple", false, "name", chi.URLParam(r, "name"), &name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter name: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, "BearerAuth.Scopes", []string{""})

	siw.Handler.FetchPlugin(w, r.WithContext(ctx), name)
}

// ListUsers operation middleware
func (siw *ServerInterfaceWrapper) ListUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, "BearerAuth.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListUsersParams

	// ------------- Optional query parameter "limit" -------------
	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter limit: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "page" -------------
	if paramValue := r.URL.Query().Get("page"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "page", r.URL.Query(), &params.Page)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter page: %s", err), http.StatusBadRequest)
		return
	}

	siw.Handler.ListUsers(w, r.WithContext(ctx), params)
}

// CreateUser operation middleware
func (siw *ServerInterfaceWrapper) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, "BearerAuth.Scopes", []string{""})

	siw.Handler.CreateUser(w, r.WithContext(ctx))
}

// CreateToken operation middleware
func (siw *ServerInterfaceWrapper) CreateToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, "BasicAuth.Scopes", []string{""})

	siw.Handler.CreateToken(w, r.WithContext(ctx))
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerFromMux(si, chi.NewRouter())
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	r.Group(func(r chi.Router) {
		r.Get("/plugins", wrapper.ListPlugins)
	})
	r.Group(func(r chi.Router) {
		r.Get("/plugins/{name}", wrapper.FetchPlugin)
	})
	r.Group(func(r chi.Router) {
		r.Get("/users", wrapper.ListUsers)
	})
	r.Group(func(r chi.Router) {
		r.Post("/users", wrapper.CreateUser)
	})
	r.Group(func(r chi.Router) {
		r.Post("/users/token", wrapper.CreateToken)
	})

	return r
}
