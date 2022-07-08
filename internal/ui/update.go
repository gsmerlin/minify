package ui

import (
	"github.com/gsmerlin/minify/internal/db"
	"github.com/rivo/tview"
)

func setUpdateMenu() {
	menu.Clear()
	menu.SetBorder(true).SetTitle("Menu")
	menu.AddItem("Identifier", "Select by identifier", 'i', func() { readId(updateResults) })
	menu.AddItem("Destination", "Select by destination", 'd', func() { readDest(updateResults) })
	menu.AddItem("All", "Select all database entries", 'a', func() { readAll(updateResults) })
	menu.AddItem("Back", "Return to main menu", 'b', func() { setMainMenu(); navigate(emptyView(), false) })
	menu.AddItem("Exit", "Close the dashboard", 'q', func() { app.Stop() })
}

func update() {
	setUpdateMenu()
	navigate(emptyView(), false)
}

func getUpdateForm(r db.Record, view tview.Primitive) tview.Primitive {
	form := tview.NewForm()

	form.AddInputField("Destination", r.Destination, 50, func(textToCheck string, lastChar rune) bool { return textToCheck != "" }, func(text string) {
		r.Destination = text
	})

	form.AddButton("Submit", func() {
		updateRecord(r, view)
	})

	return form
}

func updateRecord(r db.Record, view tview.Primitive) {
	db.UpdateLink(r)
	callModal("Success!", "Record successfully updated!", func(i int, s string) { navigate(view, true) })
}

func updateResults() {
	res := <-results

	flex := tview.NewFlex()
	flex.SetBorder(true).SetTitle("Read - Results")

	if len(res) == 1 {
		r := res[0]
		flex.AddItem(getUpdateForm(r, emptyView()), 0, 3, true)
		navigate(flex, true)
		return
	}

	list := tview.NewList()
	for _, r := range res {
		var blankRune rune
		list.AddItem("ID: "+r.ID, "", blankRune, func(r db.Record) func() {
			return func() {
				flex.Clear()
				flex.AddItem(list, 0, 1, false)
				fCopy := *flex
				fCopy.AddItem(emptyView(), 0, 5, false)
				flex.AddItem(getUpdateForm(r, &fCopy), 0, 5, true)
				navigate(flex, true)
			}
		}(r))
	}
	list.SetBorder(true)
	list.SetTitle("Results")
	flex.AddItem(list, 0, 1, true)
	flex.AddItem(emptyView(), 0, 5, false)
	navigate(flex, true)

}
