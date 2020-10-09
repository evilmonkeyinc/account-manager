# Code Generators Tested

## Gnostic Go Generator
[gnostic-go-generator](https://github.com/googleapis/gnostic-go-generator)

### Pro's
- Generates interface for implementation
- Supplies router implementation which supports calling interface implementation that was passed to it

### Con's
- Query params are not parsed correctly so cannot reference them

### Decision
Not a viable option as will require query parameters

<br>

## Gnostic GRPC and GRPC Gateway

[gnostic-grpc](https://github.com/googleapis/gnostic-grpc) and
[grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)

### Pro's
- Generates interface for implementation
- Supplies GRPC and RESTful server

### Con's
- Query or Path parameters from reference are own object
- Generated annotations cannot map referenced param objects correctly

### Decision
Not a viable solution without customizing `.proto` file generation or adjusting how openapi.yaml is bundled to avoid references for path and query parameters

<br>

## Swagger CodeGen

[swagger-codegen](https://github.com/swagger-api/swagger-codegen)

### Pro's
### Con's
Does not create interface, need to have ignore file to not recreate implementation
### Decision
No server interface, this is a big blocker

<br>

## OAPI CodeGen

[oapi-codegen](https://github.com/deepmap/oapi-codegen)

### Pro's
Supports oneof and allof
generated interface for implementation
Supports Echo, Chi, and net.http
exposes http response which will allow for more diverse responses
### Con's
Need to implement server but has some routing built in supported via chi
POST body not transformed for interface, need to read inbound request manually
each handler handles HTTP response writing, it would be good if it was partly hidden

### Decision
Soft Maybe, nothing is blocked but it is awkward

## Swagger CodeGen CLI
[swagger-codegen](https://github.com/swagger-api/swagger-codegen/tree/3.0.0)

### Pro's
Can support all functionality in sample yaml
### Con's
Does not generate interface, requires an ignore file to be included so as to not override implementation

### Decision
Not suitable since the generated files would include the implementation files

## ---