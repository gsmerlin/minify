package pages

import (
	"fmt"

	minify "github.com/gsmerlin/minify/pkg"
	"github.com/rivo/tview"
)

type Record struct {
	minify.Record
}

var create *tview.Form = tview.NewForm()

func init() {
	create.SetBorder(true).SetTitle("Create new entry")
	destination, id := "", ""

	create.AddInputField("Identifier (Leave blank for random)", "", 255, nil, func(identifier string) { id = identifier })
	create.AddInputField("Destination URL: ", "", 255, nil, func(url string) { destination = url })

	create.AddButton("Save", func() {

		if destination == "" {
			create.SetFocus(1)
			navigate(modal("Error - Blank Destination", "Destination cannot be left blank!", func(buttonIndex int, buttonLabel string) {
				navigate(create, true)
			}), true)
			return
		}

		id := repo.Create(destination, id)
		CreateSuccess := modal("Create - Success", fmt.Sprintf("URL created with identifier %v!", id), func(buttonIndex int, buttonLabel string) {
			mainMenu()
		})

		navigate(CreateSuccess, true)

	})
}
