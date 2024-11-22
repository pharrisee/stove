package root

import (
	"github.com/pharrisee/stove/app"
)

type (
	Root struct { // must implement the Controller interface (see /app/controller.go)
	}
)

func New() *Root {
	return &Root{}
}

func (c *Root) Routes(r *app.Server) {
	r.GET("/", c.Index)
	r.GET("/about", c.About)
	r.GET("/contact", c.Contact)
}
