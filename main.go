package main

import (
	"fmt"
	"gopkg.in/go-playground/webhooks.v1"
	"gopkg.in/go-playground/webhooks.v1/github"
	"log"
	"os"
	"strings"
)

func main() {
	port := os.Getenv("HOOK_PORT")
	hook := github.New(&github.Config{})
	//hook.RegisterEvents(handlePush, github.PushEvent)
	hook.RegisterEvents(handleCreate, github.CreateEvent)
	err := webhooks.Run(hook, ":"+port, "/")
	if err != nil {
		log.Fatal(err)
	}
}

func tagTest(ref string) (result bool, err error) {
	elements := strings.Split(ref, "/")
	if len(elements) < 2 {
		return false, fmt.Errorf("Could not parse ref string")
	}
	if elements[len(elements)-2] == "tags" {
		result = true
	}
	return result, nil
}

func handlePush(payload interface{}) {
	p := payload.(github.PushPayload)
	ref := strings.Split(p.Ref, "/")
	branch := ref[len(ref)-1]
	fmt.Printf("%s ", branch)
	t, err := tagTest(p.Ref)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(p.Ref)
	fmt.Printf(" %t\n", t)
}

func handleCreate(payload interface{}) {
	p := payload.(github.CreatePayload)
	fmt.Printf("%+#v\n", p)
}
