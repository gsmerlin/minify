package pages

import (
	minify "github.com/gsmerlin/minify/pkg"
	"github.com/rivo/tview"
)

func startDelete() {

	// Set menu options for update
	setMenu(deleteOpts())

	// Blank view so user can select options
	navigate(tview.NewBox(), false)

}

func deleteByIdentifier() tview.Primitive {
	view := tview.NewForm()
	view.SetBorder(true).SetTitle("Update - Identifier")

	id := ""

	view.AddInputField("Identifier: ", "", 20, nil, func(text string) {
		if text == "" {
			navigate(modal("Error - Blank Identifier", "Identifier cannot be blank!", func(buttonIndex int, buttonLabel string) {
				navigate(updateByIdentifier(), true)
			}), true)
		}

		id = text
	})

	var result minify.Record

	if id != "" {
		result = repo.Read(id, "")[0]
	}

	view.AddButton("Submit", deleteItem(result))

	return view
}

func deleteItem(r minify.Record) func() {
	return func() {
		repo.Delete(r.ID)
		navigate(modal("Delete - Success", "Record successfully deleted!", func(buttonIndex int, buttonLabel string) {
			mainMenu()
		}), true)
	}
}

func deleteByDestination(form tview.Primitive, box tview.Primitive) tview.Primitive {
	focus := true
	if box == nil {
		focus = false
		box = tview.NewBox()
	}

	view := tview.NewFlex()
	if form == nil {
		f := tview.NewForm()
		f.SetBorder(true).SetTitle("Update - Destination")

		destination := ""

		f.AddInputField("Destination: ", "", 20, nil, func(text string) {
			if text == "" {
				navigate(modal("Error - Blank Destination", "Destination cannot be blank!", func(buttonIndex int, buttonLabel string) {
					navigate(updateByDestination(box, nil), true)
				}), true)
			}
			destination = text
		})

		f.AddButton("Submit", func() {
			results <- repo.Read("", destination)
			deleteList(updateByDestination)
		})

		form = f
	}

	view.AddItem(form, 0, 1, !focus)
	view.AddItem(box, 0, 3, focus)

	return view
}

func deleteList(view func(tview.Primitive, tview.Primitive) tview.Primitive) {
	rs := <-results
	list := tview.NewList()
	list.SetBorder(true).SetTitle("Update - Result List")

	for i, r := range rs {
		list.AddItem(r.Destination, "", rune(49+i), func() {
			updateView := tview.NewForm()
			updateView.SetBorder(true).SetTitle("Update - Record")

			updateView.AddInputField("Identifier", r.ID, len(r.ID), func(textToCheck string, lastChar rune) bool { return false }, nil)
			updateView.AddInputField("Destination", r.Destination, len(r.Destination), func(textToCheck string, lastChar rune) bool { return false }, nil)

			updateView.AddButton("Submit", func() {
				repo.Delete(r.ID)
				navigate(modal("Update - Success", "Successfully updated record!", func(buttonIndex int, buttonLabel string) {
					results <- repo.Read("", r.Destination)
					deleteList(view)
				}), true)
			})

			navigate(view(list, updateView), true)
		})
	}
	navigate(view(list, nil), true)
}

func deleteByAll() {
	results <- repo.Read("", "")

	navFunc := func(form, box tview.Primitive) tview.Primitive {
		focus := true
		if box == nil {
			focus = false
			box = tview.NewBox()
		}

		view := tview.NewFlex()
		view.AddItem(form, 0, 1, !focus)
		view.AddItem(box, 0, 3, focus)

		return view
	}

	deleteList(navFunc)
}
