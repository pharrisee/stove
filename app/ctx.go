package app

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type (
	Ctx struct {
		echo.Context
	}
)

func WrapCtx(c echo.Context) *Ctx {
	if ctx, ok := c.(*Ctx); ok {
		return ctx
	}
	return &Ctx{
		Context: c,
	}
}

func (c *Ctx) Templ(status int, t templ.Component) error {
	ctx := c.Request().Context()

	b := templ.GetBuffer()
	defer templ.ReleaseBuffer(b)

	if err := t.Render(ctx, b); err != nil {
		return err
	}

	return c.HTMLBlob(status, b.Bytes())
}

func (c *Ctx) IsHTMX() bool {
	return c.Request().Header.Get("HX-Request") == "true"
}
