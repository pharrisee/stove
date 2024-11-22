package web

import (
	"embed"
	"strings"
	"time"
)

//go:embed all:public
var PublicFS embed.FS

var _cb = strings.ReplaceAll(time.Now().Format(".00000"), ".", "")

func CB(path string) string {
	return path + "?" + _cb
}
