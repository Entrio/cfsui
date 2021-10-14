package internal

type (
	config struct {
		maxLogSize int
	}
)

func newConfig() config {
	return config{
		maxLogSize: 6,
	}
}
