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

	uim := internal.NewUIManager()
	defer uim.Close()

	go uim.RenderAndUpdate()
	<-uim.WaitForExit()

}
