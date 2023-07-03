package pin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

func (c *Context) Query(key string) (value string) {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) PostForm(key string) (value string) {
	return c.Req.FormValue(key)
}

func (c *Context) SetStatusCode(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}
func (c *Context) String(code int, format string, values ...any) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatusCode(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) WriteString(content string) {
	c.String(200, content)
}
func (c *Context) JSON(code int, obj any) {

	c.SetHeader("Content-Type", "application/json")
	c.SetStatusCode(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		c.httpError(err)
	}
}

func (c *Context) WriteJSON(obj any) {
	c.JSON(200, obj)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.SetStatusCode(code)
	c.Writer.Write([]byte(html))
}

func (c *Context) httpError(err error) {
	if err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}
