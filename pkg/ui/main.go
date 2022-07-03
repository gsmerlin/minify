package ui

import "github.com/rivo/tview"

var app *tview.Application
var view *tview.Flex
var menu *tview.List

func setMainMenu() {
	menu.Clear()
	menu.SetBorder(true).SetTitle("Menu")
	menu.AddItem("Create", "Create new record", 'c', func() { create() })
	menu.AddItem("Read", "Read existing records", 'r', func() { read() })
	menu.AddItem("Update", "Update existing records", 'u', func() { update() })
	menu.AddItem("Delete", "Delete existing records", 'd', func() { delete() })
	menu.AddItem("Exit", "Close the dashboard", 'q', func() { app.Stop() })
}

func emptyView() tview.Primitive { return tview.NewBox() }

func init() {
	menu = tview.NewList()
	setMainMenu()
	view = tview.NewFlex()
	view.SetBorder(true).SetTitle("Minify")
	view.AddItem(menu, 0, 1, true)
	view.AddItem(emptyView(), 0, 4, false)
	app = tview.NewApplication().EnableMouse(true).SetRoot(view, true)
}

func navigate(v tview.Primitive, focus bool) {
	view.Clear()
	view.AddItem(menu, 0, 1, !focus)
	view.AddItem(v, 0, 4, focus)
	app.SetFocus(view)
}

func callModal(title string, text string, doneFunc func(int, string)) {
	modal := tview.NewModal()
	modal.SetTitle(title).SetBorder(true)
	modal.SetText(text)
	modal.AddButtons([]string{"OK"}).SetDoneFunc(doneFunc)
	navigate(modal, true)
}

func StartDashboard() error {
	return app.Run()
}
