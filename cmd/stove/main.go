package main

import (
	"github.com/pharrisee/stove/app"
	"github.com/pharrisee/stove/web/controllers/root"
)

func main() {

	// instantiate the server
	s := app.New()

	// Register controllers
	s.Register(root.New())

	// Start the server
	s.Start()
}
