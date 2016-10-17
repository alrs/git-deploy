package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func main() {
	port := os.Getenv("PORT")
	route := os.Getenv("ROUTE")
	http.HandleFunc(route, requestHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
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
