package moonshine

import "net/http"

type Context struct {
	request *http.Request
	writer  http.ResponseWriter
}

func (c *Context) Request() *http.Request {
	return c.request
}

func (c *Context) Writer() http.ResponseWriter {
	return c.writer
}

func NewCtx(req http.Request, res http.ResponseWriter) *Context {
	return &Context{
		request: &req,
		writer:  res,
	}
}
