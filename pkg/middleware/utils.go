package middleware

import "net/http"

func wrapResponseWriter(wrap, with http.ResponseWriter) http.ResponseWriter {
	var (
		flusher http.Flusher
		pusher  http.Pusher
	)
	flusher, _ = wrap.(http.Flusher)
	pusher, _ = wrap.(http.Pusher)

	if flusher == nil && pusher == nil {
		return with
	}
	if flusher == nil && pusher != nil {
		return struct {
			http.ResponseWriter
			http.Pusher
		}{with, pusher}
	}
	if flusher != nil && pusher == nil {
		return struct {
			http.ResponseWriter
			http.Flusher
		}{with, flusher}
	}
	return struct {
		http.ResponseWriter
		http.Flusher
		http.Pusher
	}{with, flusher, pusher}
}
