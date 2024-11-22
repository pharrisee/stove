package crosscut

import (
	"fmt"
	"io/fs"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// echo v4 middleware function to set a cache-control header for static files
func CachedStatic(path string, maxAge time.Duration, fsys fs.FS) echo.MiddlewareFunc {
	static := middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: http.FS(fsys),
	})
	seconds := int64(maxAge.Seconds())
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Method == http.MethodGet && strings.HasPrefix(c.Request().RequestURI, path) {
				c.Response().Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", seconds))
			}

			return static(next)(c)
		}
	}
}
