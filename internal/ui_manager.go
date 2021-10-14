package internal

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/rs/zerolog/log"
	"sync"
	"time"
)

type (
	UIManager struct {
		quit    chan struct{}
		Screens Stack
		log     []string
		logMU   sync.Mutex
		client  *GameClient
		config  config
	}
)

func NewUIManager(c *GameClient) *UIManager {
	if err := ui.Init(); err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize termui")
	}
	return &UIManager{
		quit:    make(chan struct{}),
		Screens: NewStack(),
		log:     []string{"Application started"},
		client:  c,
		config:  newConfig(),
	}
}

func (uim *UIManager) Startup() {
	// Start a goroutine for chan consumer
	uim.client.startup()
	go func() {
		for {
			select {
			case msg := <-uim.client.logChan:
				uim.AddLog(msg)
			}
		}
	}()
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
			uim.AddLog("Screen refreshed")
		case "f":
			uim.client.AddFarm()
		case "p", "<C-left>":
		}
	}

}

func (uim *UIManager) AddLog(data string) {
	uim.logMU.Lock()
	defer uim.logMU.Unlock()
	event := fmt.Sprintf("[%s](fg:blue) %s", time.Now().Local().Format("15:04:05"), data)
	logLen := len(uim.log)
	if logLen == uim.config.maxLogSize {
		// we need to trim
		temp := uim.log[1:uim.config.maxLogSize]
		temp = append(temp, event)
		uim.log = temp
	} else {
		uim.log = append(uim.log, event)
	}
	uim.clear()
}

func (uim *UIManager) clear() {
	ui.Clear()
	tabpane := widgets.NewTabPane("1", "2", "3", "4", "5", "6")
	tabpane.Title = "Beavertown"
	tabpane.SetRect(0, 0, 50, 3)
	tabpane.Border = true

	l := widgets.NewList()
	l.Rows = uim.log
	l.Title = "LOG"
	l.SetRect(90, 6, 45, 15)

	ui.Render(l, tabpane)
}

func (uim *UIManager) WaitForExit() chan struct{} {
	return uim.quit
}

// Close performs all closing operations prior to program exit
func (uim *UIManager) Close() {
	ui.Close()
}
