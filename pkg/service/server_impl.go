package service

import (
	"errors"
	"net/http"

	"github.com/evilmonkeyinc/account-manager/gen/server"
	"github.com/evilmonkeyinc/account-manager/pkg/service/lib"
)

func New() http.Handler {
	return server.Handler(&serverImpl{})
}

type serverImpl struct {
}

func (impl *serverImpl) ListPlugins(w http.ResponseWriter, r *http.Request, params server.ListPluginsParams) {

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
			Links: lib.BuildPagingLinks(r.Host, r.URL, page, limit, total),
		},
		Data: data,
	}
	wrapper := lib.NewResponseWriterWrapper(w)
	wrapper.WriteJSONResponse(200, response)
}

func (impl *serverImpl) FetchPlugin(w http.ResponseWriter, r *http.Request, name string) {

	wrapper := lib.NewResponseWriterWrapper(w)

	data := &server.Plugin{
		Name: &name,
	}

	wrapper.WriteJSONResponse(200, data)
}

func (impl *serverImpl) ListUsers(w http.ResponseWriter, r *http.Request, params server.ListUsersParams) {
	wrapper := lib.NewResponseWriterWrapper(w)

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
			Links: lib.BuildPagingLinks(r.Host, r.URL, page, limit, total),
		},
		Data: data,
	}

	wrapper.WriteJSONResponse(200, response)
}

func (impl *serverImpl) CreateUser(w http.ResponseWriter, r *http.Request) {

	decoder := new(lib.RequestBodyDecoder)
	decoder.Strict = true

	newUser := new(server.User)
	err := decoder.Decode(w, r, newUser)
	if err != nil {
		var malFormedError *lib.MalformedRequestError
		if errors.As(err, &malFormedError) {
			http.Error(w, malFormedError.Message, malFormedError.Code)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	wrapper := lib.NewResponseWriterWrapper(w)
	wrapper.WriteJSONResponse(201, newUser)
}

func (impl *serverImpl) CreateToken(w http.ResponseWriter, r *http.Request) {

	response := &server.CreateTokenJSONBody{}

	wrapper := lib.NewResponseWriterWrapper(w)
	wrapper.WriteJSONResponse(200, response)
}
