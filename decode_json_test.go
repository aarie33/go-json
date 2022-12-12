package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":"John","LastName":"Smith","Age":25}`
	jsonByte := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonByte, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
}
