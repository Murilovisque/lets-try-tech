package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Murilovisque/lets-try-tech/home-page-back/internal/app"
)

var (
	srv                    http.Server
	shutdownSignalReceived chan struct{} = make(chan struct{}, 1)
	shutdownFinished       chan struct{} = make(chan struct{}, 1)
)

func Setup() error {
	go func() {
		<-shutdownSignalReceived
		if err := srv.Close(); err != nil {
			log.Println(err)
		}
		shutdownFinished <- struct{}{}
	}()
	srv = http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/customers/contact-message", func(w http.ResponseWriter, r *http.Request) {
		if !methodMatch(w, r, http.MethodPost) {
			return
		}
		jsonBody := struct {
			Name    string
			Tel     uint
			Email   string
			Message string
		}{}
		if !canGetJSONBody(w, r, &jsonBody) {
			return
		}
		err := app.ProcessContactUsMessageReceived(jsonBody.Name, jsonBody.Tel, jsonBody.Email, jsonBody.Message)
		if err == nil {
			w.WriteHeader(http.StatusCreated)
		} else {
			respondErrorFromApp(err, w)
		}
	})
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("HTTP server ListenAndServe: %v", err)
		}
	}()
	return nil
}

func Shutdown() {
	shutdownSignalReceived <- struct{}{}
	<-shutdownFinished
}

func methodMatch(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method == method {
		return true
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return false
}

func canGetJSONBody(w http.ResponseWriter, r *http.Request, jsonType interface{}) bool {
	body := r.Body
	defer body.Close()
	bodyByte, err := ioutil.ReadAll(body)
	if err != nil {
		respondInternalError(w)
		return false
	}
	if err = json.Unmarshal(bodyByte, jsonType); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	return true
}

func respondInternalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "Não foi possível processar a requisição")
}

func respondErrorFromApp(err error, w http.ResponseWriter) {
	if val, ok := err.(*app.ErrApp); ok {
		w.WriteHeader(val.HTTPStatus)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, val.Error())
	} else {
		log.Println(err)
		respondInternalError(w)
	}
}
