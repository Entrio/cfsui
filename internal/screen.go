package internal

type (
	IScreen interface {
		HandleInput()
		Render()
	}
)
