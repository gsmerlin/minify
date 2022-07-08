package ui

import (
	"github.com/gsmerlin/minify/internal/db"
	"github.com/rivo/tview"
)

func create() {

	form := tview.NewForm()
	form.SetTitle("Create New Record").SetBorder(true)

	id, destination, email := "", "", ""

	form.AddInputField("Destination", "", 50, func(textToCheck string, lastChar rune) bool { return textToCheck != "" }, func(text string) {
		destination = text
	})

	form.AddInputField("Email", "", 50, func(textToCheck string, lastChar rune) bool { return textToCheck != "" }, func(text string) {
		email = text
	})
	form.AddCheckbox("Custom identifier? ", false, func(checked bool) {
		if checked {
			form.AddInputField("Identifier", "", 50, func(textToCheck string, lastChar rune) bool { return textToCheck != "" }, func(text string) {
				id = text
			})
			return
		}

		form.RemoveFormItem(2)
	})

	form.AddButton("Submit", func() {
		createRecord(id, email, destination)
	})

	navigate(form, true)
}

func createRecord(id string, email string, destination string) {
	if destination == "" {
		callModal("Error", "Destination cannot be left blank!", func(i int, s string) { create() })
		return
	}

	db.NewLink(id, email, destination)
	callModal("Success!", "Record successfully created!", func(i int, s string) { navigate(emptyView(), false) })
}
