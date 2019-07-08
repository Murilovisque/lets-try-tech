package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Murilovisque/lets-try-tech/home-page-back/cmd/home-page/routes"
	"github.com/Murilovisque/lets-try-tech/home-page-back/internal/app"
	"github.com/Murilovisque/lets-try-tech/home-page-back/internal/platform"
	"github.com/Murilovisque/logs"
)

func main() {
	setupLog()
	log.Println("Starting Home-page-back...")
	if err := platform.SetupAll(app.Setup, routes.Setup); err != nil {
		log.Printf("Home-page-back loading failed...\n%s", err.Error())
		os.Exit(1)
	}
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	log.Println("Home-page-back started!")
	for {
		select {
		case <-stopSignal:
			log.Println("Shutdown signal received")
			app.Shutdown()
			routes.Shutdown()
			log.Println("Shutdown success!")
			logs.Shutdown()
			os.Exit(0)
		default:
		}
	}
}

func setupLog() {
	const logPath = "/var/log/home-page-back/home-page"
	err := logs.SetupPerDay(logPath, 10)
	if err != nil {
		log.Printf("Home-page-back loading failed...\n%s", err.Error())
		os.Exit(1)
	}

}
