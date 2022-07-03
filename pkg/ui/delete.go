package ui

import (
	minify "github.com/gsmerlin/minify/pkg"
	"github.com/rivo/tview"
)

func setDeleteMenu() {
	menu.Clear()
	menu.SetBorder(true).SetTitle("Menu")
	menu.AddItem("Identifier", "Select by identifier", 'i', func() { readId(deleteResults) })
	menu.AddItem("Destination", "Select by destination", 'd', func() { readDest(deleteResults) })
	menu.AddItem("All", "Select all database entries", 'a', func() { readAll(deleteResults) })
	menu.AddItem("Back", "Return to main menu", 'b', func() { setMainMenu(); navigate(emptyView(), false) })
	menu.AddItem("Exit", "Close the dashboard", 'q', func() { app.Stop() })
}

func delete() {
	setDeleteMenu()
	navigate(emptyView(), false)
}

func getDeleteForm(r minify.Record, v tview.Primitive) tview.Primitive {
	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)
	text := tview.NewTextView()
	text.SetTextAlign(tview.AlignCenter)
	text.SetText("\n\n\nAre you sure you want to delete this record? \n" +
		"Identifier: " + r.ID + "\n" + "Destination: " + r.Destination + "\n")

	form := tview.NewForm()
	form.AddButton("OK", func(r minify.Record) func() { return func() { deleteRecord(r, v) } }(r))
	form.AddButton("Cancel", func() { navigate(v, false) })
	form.SetButtonsAlign(tview.AlignCenter)
	flex.AddItem(text, 0, 1, false)
	flex.AddItem(form, 0, 4, true)
	return flex
}

func deleteRecord(r minify.Record, v tview.Primitive) {
	minify.Repo().Delete(r.ID)
	callModal("Success!", "Record successfully deleted!", func(i int, s string) {
		navigate(v, true)
	})
}

func deleteResults() {
	res := <-results

	flex := tview.NewFlex()
	flex.SetBorder(true).SetTitle("Delete - Results")

	if len(res) == 1 {
		r := res[0]
		flex.AddItem(getDeleteForm(r, emptyView()), 0, 3, true)
		navigate(flex, true)
		return
	}

	list := tview.NewList()
	for _, r := range res {
		var blankRune rune
		list.AddItem("ID: "+r.ID, "", blankRune, func(r minify.Record) func() {
			return func() {
				flex.Clear()
				flex.AddItem(list, 0, 1, false)
				flex.AddItem(getDeleteForm(r, emptyView()), 0, 5, true)
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
