package bson

import (
	"labix.org/v2/mgo/bson"
)

type BsonTest struct {
	Version float32
	Method  string
	Params  []string
}

func BsonPack() {
	bson.Marshal(&BsonTest{
		Version: 2.0,
		Method:  "echo",
		Params: []string{
			"some",
			"params",
			"here",
		},
	})
}

func BsonUnpack() {
	var test BsonTest
	bson.Unmarshal([]byte(`{"Version":2.0,"Method":"echo","params":["some","params","here"]}`), &test)
}
