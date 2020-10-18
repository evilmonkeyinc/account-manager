package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/evilmonkeyinc/account-manager/gen/server"
)

type errorResponseWrapper struct {
	http.ResponseWriter
	buffer bytes.Buffer
	status int
	length int
	data   []byte
}

func (wrapper *errorResponseWrapper) Header() http.Header {
	return wrapper.ResponseWriter.Header()
}

func (wrapper *errorResponseWrapper) WriteHeader(statusCode int) {
	wrapper.status = statusCode
}

func (wrapper *errorResponseWrapper) Write(bytes []byte) (int, error) {
	n, err := wrapper.buffer.Write(bytes)
	wrapper.length += n
	return wrapper.length, err
}

func (wrapper *errorResponseWrapper) Done() {
	contentType := wrapper.ResponseWriter.Header().Get("Content-Type")
	if strings.HasPrefix(contentType, "text/plain") && wrapper.status >= 400 {
		// We want to wrap any plain-text errors in our error object
		data := &server.Response500{
			Error: &server.Error{
				Code:    fmt.Sprintf("%d", wrapper.status),
				Message: strings.TrimSpace(wrapper.buffer.String()),
			},
		}

		wrapper.ResponseWriter.Header().Set("Content-Type", "application/json")
		wrapper.ResponseWriter.WriteHeader(wrapper.status)
		json.NewEncoder(wrapper.ResponseWriter).Encode(data)
		return
	}

	wrapper.ResponseWriter.WriteHeader(wrapper.status)
	wrapper.ResponseWriter.Write(wrapper.buffer.Bytes())
}

// ErrorWrapper will wrap any plain-text error messages in the default error response format
func ErrorWrapper(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		wrapper := &errorResponseWrapper{ResponseWriter: w}
		writer := wrapResponseWriter(w, wrapper)
		defer wrapper.Done()

		next.ServeHTTP(writer, r)
	}
	return http.HandlerFunc(fn)
}
