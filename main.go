package main

import (
	minify "github.com/gsmerlin/minify/pkg"
	"github.com/gsmerlin/minify/pkg/pages"
)

func main() {
	minify.StartServer()
	pages.Start()
}
