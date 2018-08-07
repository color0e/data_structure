package main

import (
	"encoding/json"
	"fmt"
)

type Fields struct {
	VisibleField   string `json:"visibleField"`
	InvisibleField string `json:"invisibleField"`
}

func main() {
	ExampleOmitFields()
}

func ExampleOmitFields() {
	f := &Fields{"a", "b"}
	b, _ := json.Marshal(struct {
		*Fields
		InvisibleField string `json:"invisibleField,omitempty"`
		Additional     string `json:"additional,omitempty"`
	}{Fields: f, Additional: "c"})

	fmt.Println(string(b))
}
