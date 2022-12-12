package gojson

import (
	"encoding/json"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJson("hello")
	logJson(123)
	logJson(123.456)
	logJson(true)
	logJson(nil)
	logJson([]interface{}{1, 2, 3})
	logJson(map[string]interface{}{"a": 1, "b": 2})
}
