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
			mainText:      "Update",
			secondaryText: "Update existing records",
			shortcut:      'c',
			selected:      func() { startUpdate() },
		},
		{
			mainText:      "Delete",
			secondaryText: "Delete existing records",
			shortcut:      'c',
			selected:      func() { startDelete() },
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
			selected:      func() { navigate(readIdentifier(nil), true) },
		},
		{
			mainText:      "Destination",
			secondaryText: "Find records by destination",
			shortcut:      'b',
			selected:      func() { navigate(readDestination(nil), true) },
		},
		{
			mainText:      "All records",
			secondaryText: "List all records",
			shortcut:      'c',
			selected: func() {
				navigate(readAll(nil), true)
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

func updateOpts() []opts {
	return []opts{
		{
			mainText:      "Identifier",
			secondaryText: "Update via Identifier",
			shortcut:      'a',
			selected: func() {
				navigate(updateByIdentifier(), true)
			},
		},
		{
			mainText:      "Destination",
			secondaryText: "Update via Destination",
			shortcut:      'b',
			selected: func() {
				navigate(updateByDestination(nil, nil), true)
			},
		},
		{
			mainText:      "All",
			secondaryText: "Update by listing all entries",
			shortcut:      'c',
			selected: func() {
				updateByAll()
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

func deleteOpts() []opts {
	return []opts{
		{
			mainText:      "Identifier",
			secondaryText: "Delete via Identifier",
			shortcut:      'a',
			selected: func() {
				navigate(deleteByIdentifier(), true)
			},
		},
		{
			mainText:      "Destination",
			secondaryText: "Delete via Destination",
			shortcut:      'b',
			selected: func() {
				navigate(deleteByDestination(nil, nil), true)
			},
		},
		{
			mainText:      "All",
			secondaryText: "Delete by listing all entries",
			shortcut:      'c',
			selected: func() {
				deleteByAll()
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
