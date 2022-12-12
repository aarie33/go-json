package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// gunakan map untuk json yang tidak terprediksi
// dari pada struct

func TestMapDecode(t *testing.T) {
	jsonString := `{"id:1,"name":"Product 1","description":"Description 1","image_url":"https://example.com/image1.jpg"}`
	jsonByte := []byte(jsonString)

	var product map[string]interface{}
	json.Unmarshal(jsonByte, &product)

	fmt.Println(product)
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":          1,
		"name":        "Product 1",
		"description": "Description 1",
		"image_url":   "https://example.com/image1.jpg",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
