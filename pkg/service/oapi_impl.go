package service

import (
	"fmt"
	"net/http"

	"github.com/evilmonkeyinc/account-manager/gen/Openapi"
)

func New() Openapi.ServerInterface {
	return &impl{}
}

type impl struct {
}

func (server *impl) ListPlugins(w http.ResponseWriter, r *http.Request, params Openapi.ListPluginsParams) {
	w.Write([]byte(fmt.Sprintf(`{"func":"ListPlugins","params":"limit:%#v page:%#v "}`, *params.Limit, *params.Page)))
}

// (GET /plugins/{name})
func (server *impl) FetchPlugin(w http.ResponseWriter, r *http.Request, name string) {
	w.Write([]byte(fmt.Sprintf(`{"func":"FetchPlugin", "params":"%s"}`, name)))
}

// (GET /users)
func (server *impl) ListUsers(w http.ResponseWriter, r *http.Request, params Openapi.ListUsersParams) {
	w.Write([]byte(fmt.Sprintf(`{"func":"ListUsers","params":"limit:%#v page:%#v "}`, *params.Limit, *params.Page)))
}

// (POST /users)
func (server *impl) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"func":"CreateUser"}`))
}
