package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName: "John",
		LastName:  "Smith",
		Age:       25,
		Hobbies:   []string{"Soccer", "Basketball"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"John","LastName":"Smith","Age":25,"Hobbies":["Soccer","Basketball"]}`
	jsonByte := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonByte, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "John",
		LastName:  "Smith",
		Addresses: []Address{
			{
				Street: "123 Main St",
				City:   "Anytown",
				State:  "CA",
			}, {
				Street: "456 Main St",
				City:   "Anytown",
				State:  "CA",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"John","LastName":"Smith","Addresses":[{"Street":"123 Main St","City":"Anytown","State":"CA"},{"Street":"456 Main St","City":"Anytown","State":"CA"}]}`
	jsonByte := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonByte, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	address := []Address{
		{
			Street: "123 Main St",
			City:   "Anytown",
			State:  "CA",
		}, {
			Street: "456 Main St",
			City:   "Anytown",
			State:  "CA",
		},
	}

	bytes, _ := json.Marshal(address)
	fmt.Println(string(bytes))
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"123 Main St","City":"Anytown","State":"CA"},{"Street":"456 Main St","City":"Anytown","State":"CA"}]`
	jsonByte := []byte(jsonString)

	address := &[]Address{}
	err := json.Unmarshal(jsonByte, address)
	if err != nil {
		panic(err)
	}
	fmt.Println(address)
}
