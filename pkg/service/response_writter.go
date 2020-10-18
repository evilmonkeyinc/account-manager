package service

import (
	"encoding/json"
	"net/http"
)

func NewResponseWriterWrapper(writer http.ResponseWriter) *ResponseWriterWrapper {
	return &ResponseWriterWrapper{
		writer: writer,
	}
}

type ResponseWriterWrapper struct {
	writer http.ResponseWriter
}

func (wrapper *ResponseWriterWrapper) WriteJSONResponse(statusCode int, data interface{}) error {
	wrapper.writer.Header().Set("Content-Type", "application/json")
	wrapper.writer.WriteHeader(statusCode)
	return json.NewEncoder(wrapper.writer).Encode(data)
}
