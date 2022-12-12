package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName string // should be Capitalized
	LastName  string
	Age       int
	Hobbies   []string
	Addresses []Address
}

type Address struct {
	Street string
	City   string
	State  string
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "John",
		LastName:  "Smith",
		Age:       25,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
