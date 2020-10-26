package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	defaultRequestBodyMaxBytes int64 = 1048576
)

// MalformedRequestError error that represents a malformed request error
type MalformedRequestError struct {
	Code    int
	Message string
}

func (err *MalformedRequestError) Error() string {
	return err.Message
}

// SendError replies to the request with the specified error message and HTTP code.
// It does not otherwise end the request; the caller should ensure no further writes are done to the writter.
// The error message should be plain text.
func (err *MalformedRequestError) SendError(writer http.ResponseWriter) {
	http.Error(writer, err.Message, err.Code)
}

// RequestBodyDecoder helper that helps decoding an application/json request body.
type RequestBodyDecoder struct {
	// Strict causes the Decoder to return an error when the destination is a struct and the input contains object keys which do not match any non-ignored, exported fields in the destination.
	Strict bool
	// RequestBodyMaxBytes the maximum number of bytes expected for the request body. Defaults to 1MB.
	RequestBodyMaxBytes int64
}

// Decode reads the JSON-encoded value from the request body and stores it in the value pointed to by destination.
func (decoder *RequestBodyDecoder) Decode(writer http.ResponseWriter, request *http.Request, destination interface{}) error {
	if contentType := request.Header.Get("Content-Type"); contentType != "" {
		if !strings.Contains(contentType, "application/json") {
			return &MalformedRequestError{
				Code:    http.StatusUnsupportedMediaType,
				Message: "Content-Type header is not application/json",
			}
		}
	}

	requestBodyMaxBytes := defaultRequestBodyMaxBytes
	if decoder.RequestBodyMaxBytes > 0 {
		requestBodyMaxBytes = decoder.RequestBodyMaxBytes
	}

	if requestBodyMaxBytes > 0 {
		request.Body = http.MaxBytesReader(writer, request.Body, requestBodyMaxBytes)
	}

	dec := json.NewDecoder(request.Body)
	if decoder.Strict {
		dec.DisallowUnknownFields()
	}

	err := dec.Decode(destination)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			return &MalformedRequestError{
				Message: fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset),
				Code:    http.StatusBadRequest,
			}
		case errors.Is(err, io.ErrUnexpectedEOF):
			return &MalformedRequestError{
				Message: fmt.Sprintf("Request body contains badly-formed JSON"),
				Code:    http.StatusBadRequest,
			}
		case errors.As(err, &unmarshalTypeError):
			return &MalformedRequestError{
				Message: fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset),
				Code:    http.StatusBadRequest,
			}
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return &MalformedRequestError{
				Message: fmt.Sprintf("Request body contains unknown field %s", fieldName),
				Code:    http.StatusBadRequest,
			}
		case errors.Is(err, io.EOF):
			return &MalformedRequestError{
				Message: "Request body must not be empty",
				Code:    http.StatusBadRequest,
			}
		case err.Error() == "http: request body too large":
			return &MalformedRequestError{
				Message: fmt.Sprintf("Request body must not be larger than %d bytes", requestBodyMaxBytes),
				Code:    http.StatusRequestEntityTooLarge,
			}
		default:
			// This would normally be to catch anything else and call it a server error
			// but this could catch errors during marshalling and unmarshalling due to
			// invalid user entered values
			log.Println(err.Error())
			return &MalformedRequestError{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			}
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return &MalformedRequestError{
			Message: "Request body must only contain a single JSON object",
			Code:    http.StatusBadRequest,
		}
	}

	return nil
}
