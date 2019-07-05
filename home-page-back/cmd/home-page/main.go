package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Murilovisque/lets-try-tech/home-page-back/cmd/home-page/routes"
	"github.com/Murilovisque/lets-try-tech/home-page-back/internal/app"
	"github.com/Murilovisque/lets-try-tech/home-page-back/internal/platform"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	setupLog()
	log.Println("Starting Home-page-back...")
	if err := platform.SetupAll(setupLog, app.Setup, routes.Setup); err != nil {
		log.Printf("Home-page-back loading failed...\n%s", err.Error())
		os.Exit(1)
		return
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
			os.Exit(0)
			return
		default:
		}
	}
}

func setupLog() error {
	const logConfigPath = "/etc/home-page-back/log.json"
	var ljLog lumberjack.Logger
	b, err := ioutil.ReadFile(logConfigPath)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, &ljLog); err != nil {
		return err
	}
	log.SetOutput(&ljLog)
	return nil
}
