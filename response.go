package moonshine

import (
	"fmt"
	"net/http"
)

const (
	noWritten     = -1
	defaultStatus = http.StatusOK
)

type ResponseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func (w *ResponseWriter) WriteHeader(code int) {
	if code > 0 && w.status != code {
		if w.Written() {
			fmt.Printf("[WARNING] Headers were already written. Wanted to override status code %d with %d \n", w.status, code)
			return
		}
		w.status = code
	}
}

func (w *ResponseWriter) WriteHeaderNow() {
	if !w.Written() {
		w.size = 0
		w.WriteHeader(w.status)
	}
}

func (w *ResponseWriter) Status() int {
	return w.status
}

func (w *ResponseWriter) Size() int {
	return w.size
}

func (w *ResponseWriter) Written() bool {
	return w.size != noWritten
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{w, defaultStatus, noWritten}
}
