package http

import (
	"net/http"

	"github.com/go-nt/http/request"
	"github.com/go-nt/http/response"
)

type Context struct {
	Request  *request.Driver
	Response *response.Driver
}

// Init 初始化
func (c *Context) Init(r *http.Request, w http.ResponseWriter) {

	req := new(request.Driver)
	req.Init(r)

	res := new(response.Driver)
	res.Init(w)

	c.Request = req
	c.Response = res
}
