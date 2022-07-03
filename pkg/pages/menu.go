package pages

import (
	"github.com/rivo/tview"
)

type opts struct {
	mainText      string
	secondaryText string
	shortcut      rune
	selected      func()
}

func defaults() []opts {
	return []opts{
		{
			mainText:      "Create",
			secondaryText: "Create new record",
			shortcut:      'a',
			selected:      func() { navigate(create, true) },
		},
		{
			mainText:      "Read",
			secondaryText: "Read existing records",
			shortcut:      'b',
			selected:      func() { startRead() },
		},
		{
			mainText:      "Quit",
			secondaryText: "Press to exit",
			shortcut:      'q',
			selected: func() {
				app.Stop()
			},
		},
	}
}

func readOpts() []opts {
	return []opts{
		{
			mainText:      "Identifier",
			secondaryText: "Find record by identifier",
			shortcut:      'a',
			selected:      func() { navigate(readIdentifier(), true) },
		},
		{
			mainText:      "Destination",
			secondaryText: "Find records by destination",
			shortcut:      'b',
			selected:      func() { navigate(readDestination(), true) },
		},
		{
			mainText:      "All records",
			secondaryText: "List all records",
			shortcut:      'c',
			selected: func() {
				navigate(readAll(), true)
			},
		},
		{
			mainText:      "Back",
			secondaryText: "Return to main menu",
			shortcut:      'c',
			selected: func() {
				mainMenu()
			},
		},
		{
			mainText:      "Quit",
			secondaryText: "Exit application",
			shortcut:      'q',
			selected: func() {
				app.Stop()
			},
		},
	}
}

var Menu = tview.NewList()

func init() {
	Menu.SetBorder(true).SetTitle("Menu")
}

func setMenu(options []opts) {
	Menu.Clear()
	for _, option := range options {
		Menu.AddItem(option.mainText, option.secondaryText, option.shortcut, option.selected)
	}
}
