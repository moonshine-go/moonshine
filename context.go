package moonshine

import (
	"context"
	"net/http"
	"sync"
)

type Context struct {
	request *http.Request
	writer  http.ResponseWriter
	mu      sync.RWMutex
	errors  []error
}

func (c *Context) Request() *http.Request {
	return c.request
}

func (c *Context) Writer() http.ResponseWriter {
	return c.writer
}

func (c *Context) Context() context.Context {
	return c.request.Context()
}

func (c *Context) Error(err error) {
	if err == nil {
		panic("err is nil")
	}
	c.mu.Lock()
	c.errors = append(c.errors, err)
	c.mu.Unlock()
}

func NewCtx(req http.Request, res http.ResponseWriter) *Context {
	return &Context{
		request: &req,
		writer:  res,
	}
}
