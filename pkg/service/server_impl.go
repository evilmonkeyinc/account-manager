package service

import (
	"encoding/json"
	"net/http"

	"github.com/evilmonkeyinc/account-manager/gen/server"
)

func New() http.Handler {
	return server.Handler(&serverImpl{})
}

type serverImpl struct {
}

func (impl *serverImpl) ListPlugins(w http.ResponseWriter, r *http.Request, params server.ListPluginsParams) {
	wrapper := NewResponseWriterWrapper(w)

	data := make([]server.Plugin, 0)

	limit := int(*params.Limit)
	page := int(*params.Page)
	count := len(data)
	total := 100

	response := &server.ListPluginsResponse{
		PageDetails: server.PageDetails{
			Count: count,
			Limit: limit,
			Page:  page,
			Total: total,
			Links: buildPagingLinks(r.Host, r.URL, page, limit, total),
		},
		Data: data,
	}

	wrapper.WriteJSONResponse(200, response)
}

func (impl *serverImpl) FetchPlugin(w http.ResponseWriter, r *http.Request, name string) {

	wrapper := NewResponseWriterWrapper(w)

	data := &server.Plugin{
		Name: &name,
	}

	wrapper.WriteJSONResponse(200, data)
}

func (impl *serverImpl) ListUsers(w http.ResponseWriter, r *http.Request, params server.ListUsersParams) {
	wrapper := NewResponseWriterWrapper(w)

	data := make([]server.User, 0)

	limit := int(*params.Limit)
	page := int(*params.Page)
	count := len(data)
	total := 100

	response := &server.ListUsersResponse{
		PageDetails: server.PageDetails{
			Count: count,
			Limit: limit,
			Page:  page,
			Total: total,
			Links: buildPagingLinks(r.Host, r.URL, page, limit, total),
		},
		Data: data,
	}

	wrapper.WriteJSONResponse(200, response)
}

func (impl *serverImpl) CreateUser(w http.ResponseWriter, r *http.Request) {

	newUser := new(server.User)
	err := json.NewDecoder(r.Body).Decode(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	wrapper := NewResponseWriterWrapper(w)
	wrapper.WriteJSONResponse(201, newUser)
}
