package app

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pharrisee/stove/app/crosscut"
	"github.com/pharrisee/stove/web"
	slogecho "github.com/samber/slog-echo"
)

type (
	Server struct {
		*echo.Echo
	}
)

func New() *Server {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	r := echo.New()
	r.Use(slogecho.New(logger))
	r.Use(middleware.Recover())
	r.Use(crosscut.Brotli())
	r.Use(middleware.RequestID())
	r.Use(middleware.TimeoutWithConfig(
		middleware.TimeoutConfig{
			Timeout: 10 * time.Second,
		},
	))
	r.Use(middleware.Secure())
	r.Use(crosscut.CachedStatic("/public", 24*365*time.Hour, web.PublicFS))

	return &Server{
		Echo: r,
	}
}

func (s *Server) Start() {
	fmt.Println("Starting server on port :3333")
	s.Logger.Fatal(s.Echo.Start(":3333"))
}

func (s *Server) Register(controllers ...Controller) {
	for _, c := range controllers {
		c.Routes(s)
	}
}
