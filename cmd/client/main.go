package main

import (
	"github.com/Entrio/cfsui/internal"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	c, err := internal.NewGameClient()
	if err != nil {
		panic(err)
	}
	uim := internal.NewUIManager(c)
	uim.Startup()
	defer uim.Close()

	go uim.RenderAndUpdate()
	<-uim.WaitForExit()

}
