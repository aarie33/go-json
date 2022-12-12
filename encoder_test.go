package gojson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := &Customer{
		FirstName: "John",
		LastName:  "Smith",
		Age:       25,
	}

	encoder.Encode(customer)
	fmt.Println(customer)
}
