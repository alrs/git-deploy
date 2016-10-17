package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

const testPayload = "testdata/create_event.json"

func TestParsePayload(t *testing.T) {
	expected := map[string]string{
		"fullName": "baxterthehacker/public-repo",
		"name":     "public-repo",
		"refType":  "tag",
		"ref":      "0.0.1",
	}
	j, err := ioutil.ReadFile(testPayload)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Loaded file %s.\n", testPayload)
	}
	p, err := parsePayload(j)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Parsed JSON from %s.\n", testPayload)
	}

	result := map[string]string{
		"fullName": p.Repository.FullName,
		"name":     p.Repository.Name,
		"refType":  p.RefType,
		"ref":      p.Ref,
	}

	if reflect.DeepEqual(expected, result) {
		t.Log("Expected output and results match.")
	} else {
		t.Fatalf("Expected: %#+v\nResult: %#+v", expected, result)
	}
}

func TestValidPayloadAction(t *testing.T) {
	j, err := ioutil.ReadFile(testPayload)
	if err != nil {
		t.Fatal(err)
	}
	p, err := parsePayload(j)
	if err != nil {
		t.Fatal(err)
	}
	err = payloadAction(p)
	if err != nil {
		t.Fatalf("Payload %#+v caused error %s.\n", p, err)
	} else {
		t.Log("Succesfully triggered payload action.")
	}
}

func TestEmptyPayloadAction(t *testing.T) {
	p := CreateEvent{}
	err := payloadAction(p)
	if err != nil {
		t.Logf("Empty payload caused the following error, correctly: %s.\n", err)
	} else {
		t.Fatalf("Payload %#+v failed to cause error.\n", p)
		t.Log("Succesfully triggered payload action.")
	}
}
