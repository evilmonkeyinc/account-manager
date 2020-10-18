# Sample OpenAPI defined golang server

Project designed to create a service using:

1. Proposed standard [GoLang package structure](https://github.com/golang-standards/project-layout)
2. [GoLang plugins](https://golang.org/pkg/plugin/) inspired by [GOSH](https://github.com/vladimirvivien/gosh)
3. An API defined by [OpenAPI version 3.0.x](https://swagger.io/specification/)

## OpenAPI

The OpenAPI specification file will be written in multiple parts and combined through use of the `$ref` keywords and the [swagger-cli](https://github.com/APIDevTools/swagger-cli) npm package to bundle it into a single specification file which is accepted more by other tools than separate files.

After some extensive testing and comparisons of the various services and libraries to generate golang server stubs from openapi specification, we finally landed on [OAPI Codegen](github.com/deepmap/oapi-codegen) due to it supporting all the features we had in the sample specification at the time and is implemented through an interface unlike others that generate the implementation included in the generated files and needs additional logic to ignore it so it is not overridden
