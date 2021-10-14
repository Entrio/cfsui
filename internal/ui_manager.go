package internal

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/rs/zerolog/log"
)

type (
	UIManager struct {
		quit    chan struct{}
		Screens Stack
		log     []string
	}
)

func NewUIManager() *UIManager {
	if err := ui.Init(); err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize termui")
	}
	return &UIManager{
		quit:    make(chan struct{}),
		Screens: NewStack(),
		log:     []string{"Application started"},
	}
}

// RenderAndUpdate goroutine processes the input and rendering
func (uim *UIManager) RenderAndUpdate() {
	uim.clear()
	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			uim.quit <- struct{}{}
			return
		case "r":
			uim.clear()
		case "p", "<C-left>":
		}
	}

}

func (uim *UIManager) AddLog(data string) {}

func (uim *UIManager) clear() {
	ui.Clear()
	tabpane := widgets.NewTabPane("Home", "Construction", "Production", "Market", "Population", "Military (10)")
	tabpane.Title = "Beavertown"
	tabpane.PaddingLeft = 2
	tabpane.SetRect(0, 0, 75, 3)
	tabpane.Border = true

	p2 := widgets.NewParagraph()
	p2.Text = "Current farms: 0\nCurrent food: 0"
	p2.Title = "Keys"
	p2.TextStyle.Fg = ui.ColorBlue
	p2.SetRect(5, 5, 10, 10)
	p2.BorderStyle.Fg = ui.ColorYellow

	l := widgets.NewList()
	l.Rows = uim.log
	l.Title = "LOG"
	l.SetRect(10, 7, 45, 15)

	ui.Render(tabpane, l)
}

func (uim *UIManager) WaitForExit() chan struct{} {
	return uim.quit
}

// Close performs all closing operations prior to program exit
func (uim *UIManager) Close() {
	ui.Close()
}
