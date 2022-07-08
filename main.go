package main

import (
	"flag"

	"github.com/gsmerlin/minify/internal/db"
	"github.com/gsmerlin/minify/internal/logger"
	"github.com/gsmerlin/minify/internal/server"
	"github.com/gsmerlin/minify/internal/ui"
)

func main() {
	view := flag.Bool("experimentalView", false, "Controls whether the server is executed normally or in dashboard mode")
	flag.Parse()
	logger.Start(*view)
	db.Start()
	if *view {
		go server.Start()
		ui.StartDashboard()
	}

	err := server.Start()
	panic(err)

}
