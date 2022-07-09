package ui

import (
	"github.com/gsmerlin/minify/backend/internal/db"
	"github.com/rivo/tview"
)

var results = make(chan []db.Record, 1)

func setReadMenu() {
	menu.Clear()
	menu.SetBorder(true).SetTitle("Menu")
	menu.AddItem("Identifier", "Select by identifier", 'i', func() { readId(readResults) })
	menu.AddItem("Destination", "Select by destination", 'd', func() { readDest(readResults) })
	menu.AddItem("All", "Select all database entries", 'a', func() { readAll(readResults) })
	menu.AddItem("Back", "Return to main menu", 'b', func() { setMainMenu(); navigate(emptyView(), false) })
	menu.AddItem("Exit", "Close the dashboard", 'q', func() { app.Stop() })
}

func read() {
	setReadMenu()
	navigate(emptyView(), false)
}

func getReadText(r db.Record) tview.Primitive {
	text := tview.NewTextView()
	text.SetTextAlign(tview.AlignCenter)
	text.SetText(
		"Identifier: " + r.ID + "\n" +
			"Destination: " + r.Destination + "\n",
	)
	return text
}

func readId(cb func()) {
	id := ""
	form := tview.NewForm()
	form.SetTitle("Read by ID").SetBorder(true)
	form.AddInputField("Identifier", "", 50, func(textToCheck string, lastChar rune) bool { return textToCheck != "" }, func(text string) {
		id = text
	})
	form.AddButton("Submit", func() {
		findRecord(id, "", cb)
	})

	navigate(form, true)
}

func readDest(cb func()) {
	dest := ""
	form := tview.NewForm()
	form.SetTitle("Read by Destination").SetBorder(true)
	form.AddInputField("Destination", "", 50, func(textToCheck string, lastChar rune) bool { return textToCheck != "" }, func(text string) {
		dest = text
	})
	form.AddButton("Submit", func() {
		findRecord("", dest, cb)
	})

	navigate(form, true)

}

func readAll(cb func()) {
	findRecord("", "", cb)
}

func findRecord(id string, dest string, cb func()) {
	res, _ := db.GetLink(id, "", dest)
	results <- res
	cb()
}

func readResults() {
	res := <-results

	flex := tview.NewFlex()
	flex.SetBorder(true).SetTitle("Read - Results")

	if len(res) == 1 {
		r := res[0]
		flex.AddItem(getReadText(r), 0, 3, false)
		navigate(flex, false)
		return
	}

	list := tview.NewList()
	for _, r := range res {
		var blankRune rune
		list.AddItem("ID: "+r.ID, "", blankRune, func(r db.Record) func() {
			return func() {
				flex.Clear()
				flex.AddItem(list, 0, 1, true)
				flex.AddItem(getReadText(r), 0, 5, false)
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
