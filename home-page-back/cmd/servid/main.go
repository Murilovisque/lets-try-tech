package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Murilovisque/lets-try-tech/home-page-back/internal/app"
)

func main() {
	log.Println("Starting Home-page-back...")
	err := app.Setup()
	if err != nil {
		log.Printf("Home-page-back loading failed...\n%s", err.Error())
		os.Exit(1)
		return
	}
	stopSignal := make(chan os.Signal)
	signal.Notify(stopSignal, syscall.SIGTERM)
	signal.Notify(stopSignal, syscall.SIGINT)
	signal.Notify(stopSignal, os.Interrupt)
	log.Println("Home-page-back started!")
	for {
		select {
		case <-stopSignal:
			app.Shutdown()
			os.Exit(0)
			return
		default:
		}
	}

}
