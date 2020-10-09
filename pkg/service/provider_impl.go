package service

// gnostic-go-generator implementation

/**
import (
	"fmt"

	"github.com/evilmonkeyinc/account-manager/gen/openapi"
)


// New returns a new instance of the openapi.Provider interface implementation
func New() openapi.Provider {
	return &serviceProvider{}
}

type serviceProvider struct {
}

// list all supported plugins
func (service *serviceProvider) ListPlugins(parameters *openapi.ListPluginsParameters, responses *openapi.ListPluginsResponse) (err error) {
	responses.Data = []openapi.Plugin{
		{
			Name: "default-plugin",
		},
	}
	responses.Limit = parameters.Limit
	responses.Page = parameters.Page

	return nil
}

// Get the specified plugin
func (service *serviceProvider) FetchPlugin(parameters *openapi.FetchPluginParameters, responses *openapi.Plugin) (err error) {

	responses.Name = fmt.Sprintf("%s-plugin", parameters.Name)

	return nil
}
**/
