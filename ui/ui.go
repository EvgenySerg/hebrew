package ui

import (
	"github.com/rivo/tview"
	"hebrew/models"
	"hebrew/config"
	"hebrew/models/wordstate"
	"fmt"
)

func NewUi(config *config.Config, dict *models.Dict) *Ui {

	return &Ui{dict: dict,config:config}
}

type UiVisualiser interface {
	Show()
}

type Ui struct {
	app        *tview.Application
	MainMenu   *tview.List
	ThemesMenu *tview.List
	pages      *tview.Pages
	dict       *models.Dict
	config     *config.Config
}

func (ui *Ui) Show() {

	ui.app = tview.NewApplication()
	ui.pages = tview.NewPages()
	ui.MainMenu = ui.CreateMainMenu()
	ui.ThemesMenu = ui.CreateThemesMenu()



	ui.pages.AddPage("main", ui.MainMenu, true, true)
	ui.pages.AddPage("themes", ui.ThemesMenu, true, true)

	ui.pages.SwitchToPage("main")

	if err := ui.app.SetRoot(ui.pages, true).SetFocus(ui.pages).Run(); err != nil {
		panic(err)
	}
}

func (ui *Ui) CreateMainMenu() *tview.List {
	return tview.NewList().
		AddItem("Themes", "", 'a', ui.ShowThemesList).
		AddItem("Quiz", "", 'b', nil).
		AddItem("Import theme", "", 'c', nil).
		AddItem("Quit", "", 'q', func() { ui.app.Stop() })
}

func (ui *Ui) CreateThemesMenu() *tview.List {
	themes := ui.dict.Themes
	themesList := tview.NewList()

	for _, theme := range themes {
		menuText := fmt.Sprintf("%s  (new words - %d)", theme.Name, len(theme.GetWordsByState(wordstate.New)))
		themesList.AddItem(menuText, "", 'a', nil)
	}

	themesList.AddItem("Back", "", 'b', func() {
		ui.pages.SwitchToPage("main")
	})

	return themesList
}

func (ui *Ui) ShowThemesList() {
	ui.pages.SwitchToPage("themes")
	ui.app.SetFocus(ui.ThemesMenu)
}

/*
	Themes
		Theme1
		Theme2
			Word1 he - ru (3 of 10 )
			Word2 he - ru
	Quiz
		CurrentQuiz: Theme, 5 of 15 33%, name, started 01/01/2019
		Words
			Word1  he - ru (3 of 10 ) - new
			Word2  he - ru (3 of 10 ) - new
			Word22 he - ru (48 of 50 ) - rep
		Start/Continue
		Reset stat
		New
	ImportTheme

 */
