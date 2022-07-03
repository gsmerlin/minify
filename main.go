package main

import (
	minify "github.com/gsmerlin/minify/pkg"
	"github.com/gsmerlin/minify/pkg/ui"
)

func main() {
	minify.StartServer()
	ui.StartDashboard()
}
