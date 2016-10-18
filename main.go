package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type CreateEvent struct {
	Ref        string
	RefType    string `json:"ref_type"`
	Repository Repository
}

type Repository struct {
	Name     string
	FullName string `json:"full_name"`
}

var hasAnnounced = false

func main() {
	go signalHandler()
	if len(os.Getenv("ANNOUNCE")) > 0 {
		log.Print("Service discovery announcements are activated.")
		go announce()
		defer unAnnounce()
	} else {
		log.Print("Service discovery announcements are NOT activated.")
	}

	port := os.Getenv("PORT")
	route := os.Getenv("ROUTE")
	http.HandleFunc(route, requestHandler)
	log.Panic(http.ListenAndServe(":"+port, nil))
}

func announce() {
	for {
		time.Sleep(time.Second * 10)
		log.Println("STUB: announce")
		hasAnnounced = true
	}
}

func parsePayload(body []byte) (payload CreateEvent, err error) {
	err = json.Unmarshal(body, &payload)
	return
}

func payloadAction(payload CreateEvent) (err error) {
	switch payload.RefType {
	case "tag":
		fmt.Printf("STUB: git checkout git@github.com:%s && cd %s && git checkout %s\n",
			payload.Repository.FullName, payload.Repository.Name, payload.Ref)
	case "":
		err = fmt.Errorf("Payload ref_type was empty.")
	}
	return
}

func requestHandler(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	payload, err := parsePayload(body)
	if err != nil {
		log.Fatal(err)
	}
	err = payloadAction(payload)
	if err != nil {
		log.Fatal(err)
	}
}

func signalHandler() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigs
	unAnnounce()
	log.Fatalf("%s\n", sig)
}

func unAnnounce() {
	if hasAnnounced {
		log.Println("STUB: unannounce")
	}
}
