package json

import (
	"encoding/json"
)

type JsonTest struct {
	Version float32
	Method  string
	Params  []string
}

func JsonPack() {
	json.Marshal(&JsonTest{
		Version: 2.0,
		Method:  "echo",
		Params: []string{
			"some",
			"params",
			"here",
		},
	})
}

func JsonUnpack() {
	var test JsonTest
	json.Unmarshal([]byte(`{"Version":2.0,"Method":"echo","params":["some","params","here"]}`), &test)
}
