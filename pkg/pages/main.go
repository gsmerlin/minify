package pages

import (
	minify "github.com/gsmerlin/minify/pkg"
	"github.com/rivo/tview"
)

var repo *minify.Repository
var app = tview.NewApplication()
var view = tview.NewFlex()

func init() {
	repo = minify.Repo()
}

func navigate(newView tview.Primitive, focus bool) {
	view.Clear().AddItem(Menu, 0, 1, !focus).AddItem(newView, 0, 3, focus)
	app.SetFocus(view)
}

func mainMenu() {
	// Reset to menu defaults
	setMenu(defaults())

	// Call main menu
	view.Clear().AddItem(Menu, 0, 1, true).AddItem(tview.NewBox(), 0, 3, false)
	app.SetFocus(view)
}

// Create simple modal
func modal(title string, msg string, f func(buttonIndex int, buttonLabel string)) *tview.Modal {
	modal := tview.NewModal()
	modal.SetText(msg).SetBorder(true).SetTitle(title)
	modal.AddButtons([]string{"OK"}).SetDoneFunc(f)
	return modal
}

func Start() {
	mainMenu()
	if err := app.SetRoot(view, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
