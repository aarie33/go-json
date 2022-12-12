package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:          1,
		Name:        "Product 1",
		Description: "Description 1",
		ImageURL:    "https://example.com/image1.jpg",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":1,"name":"Product 1","description":"Description 1","image_url":"https://example.com/image1.jpg"}`
	jsonByte := []byte(jsonString)

	product := &Product{}

	err := json.Unmarshal(jsonByte, product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
	fmt.Println(product.Id)
	fmt.Println(product.Name)
	fmt.Println(product.Description)
	fmt.Println(product.ImageURL)
}
