package internal

type (
	UIManager struct {
		quit    chan struct{}
		Screens Stack
	}
)

func NewUIManager() *UIManager {
	return &UIManager{
		quit:    make(chan struct{}),
		Screens: NewStack(),
	}
}

// Render goroutine processes the input and rendering
func (ui *UIManager) Render() {}
func (ui *UIManager) WaitForExit() chan struct{} {
	return ui.quit
}
