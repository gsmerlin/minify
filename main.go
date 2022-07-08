package main

import (
	"flag"

	"github.com/gsmerlin/minify/internal/db"
	"github.com/gsmerlin/minify/internal/logger"
	"github.com/gsmerlin/minify/internal/server"
	"github.com/gsmerlin/minify/internal/ui"
)

/*
	Backend system specs:
	1. Create minify database (DONE)
		Tables:
			a) Records (ID, Destination, Email) (DONE)
			  i. (This could have also had created at and expired fields, but I'm keeping it simple)
			b) Analytics (ID, AccessedAt) (DONE)
			  i. (Just this for simple "number of clicks X day" metrics)
	2. Create CRUD database operations (DONE)
	3. Create minify server (***TO-DO***)
		Routes:
			a) / (DONE)
				i. GET;
				ii. Main redirect service. Receives an URL that contains the ID
					of the link to be redirected to, and then redirects to the
					destination stored in the database.
			b) /create (DONE)
				i. POST;
				ii. Receives a JSON object with the following fields:
					a. ID: The ID of the link. If this is not provided, a random
					ID will be generated.
					b. Email: The email address of the user who created the link
					c. Destination: The destination URL to be redirected to
				iii. Creates a new link in the database.
				iv. Returns the ID of the link.
			c) /get/?:link (DONE)
				i. GET;
				ii. Receives an ID of the link to be retrieved.
				iii. Retrieves the link from the database.
				iv. Returns the link information.
			d) /analytics/?:link (***TO-DO***)
				i. GET;
				ii. Receives an ID of the link to be retrieved.
				iii. Retrieves the analytics from the database.
				iv. Returns the analytics information.
			e) /delete/?:link (DONE)
				i. DELETE;
				ii. Receives an ID of the link to be deleted.
				iii. Deletes the link from the database.
				iv. Returns the ID of the link.
			f) /edit (DONE)
				i. POST;
				ii. Receives a JSON object with the following fields:
					a. ID: The ID of the link.
					b. Email: The email address of the user who created the link
					c. Destination: The destination URL to be redirected to
				iii. Updates the link in the database.
				iv. Returns the ID of the link.



*/

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
