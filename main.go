package main

import (
	"flag"

	"github.com/gsmerlin/minify/internal/server"
	"github.com/gsmerlin/minify/internal/ui"
	"github.com/gsmerlin/minify/internal/utils"
)

func main() {
	view := flag.Bool("view", false, "Controls whether the server is executed normally or in dashboard mode")
	flag.Parse()
	if *view {
		go server.StartServer(&utils.Logger{})
		ui.StartDashboard()
	}
	err := server.StartServer(nil)
	panic(err)

}
