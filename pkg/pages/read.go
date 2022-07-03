package pages

import (
	"fmt"

	minify "github.com/gsmerlin/minify/pkg"
	"github.com/rivo/tview"
)

var results = make(chan []minify.Record, 1)

func startRead() {

	// Set menu options for read
	setMenu(readOpts())

	// Blank view so user can select options
	navigate(tview.NewBox(), false)

}

func readIdentifier() tview.Primitive {
	identifier := ""

	view := tview.NewForm()
	view.SetBorder(true).SetTitle("Read - Identifier")
	view.AddInputField("Identifier", "", 100, nil, func(text string) {
		if text == "" {
			m := modal("Error - Invalid args", "Identifier cannot be blank!", func(buttonIndex int, buttonLabel string) {
				navigate(readIdentifier(), true)
			})
			navigate(m, true)
		}

		identifier = text
	})

	view.AddButton("Submit", func() {
		r := repo.Read(identifier, "")
		results <- r
		navigate(readResults(tview.NewBox()), true)
	})
	return view
}

func readDestination() tview.Primitive {
	destination := ""

	view := tview.NewForm()
	view.SetBorder(true).SetTitle("Read - Destination")
	view.AddInputField("Destination", "", 100, nil, func(text string) {
		if text == "" {
			m := modal("Error - Invalid args", "Destination cannot be blank!", func(buttonIndex int, buttonLabel string) {
				navigate(readIdentifier(), true)
			})
			navigate(m, true)
		}

		destination = text
	})

	view.AddButton("Submit", func() {
		r := repo.Read("", destination)
		results <- r
		navigate(readResults(tview.NewBox()), true)
	})
	return view
}

func readAll() tview.Primitive {
	results <- repo.Read("", "")
	return readResults(tview.NewBox())
}

func readResults(box tview.Primitive) tview.Primitive {
	res := <-results
	view := tview.NewFlex()
	list := tview.NewList()
	list.SetBorder(true).SetTitle("Read - Results")
	for i, r := range res {
		list.AddItem(r.Destination, "", rune(49+i), func(id string) func() {
			return func() {
				details := repo.GetDetails(id)
				text := fmt.Sprintf("Identifier : %v \nDestination: %v \n", details.ID, details.Destination)
				for _, a := range details.Analytics {
					text += fmt.Sprintf("%v", a.AccessedAt)
				}
				detailView := tview.NewTextView().SetText(text)
				detailView.SetBorder(true).SetTitle("Read - Details")
				results <- res
				navigate(readResults(detailView), true)
			}
		}(r.ID))
	}
	view.AddItem(list, 0, 1, true)
	view.AddItem(box, 0, 3, false)
	return view
}
