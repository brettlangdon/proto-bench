package msgpack

import (
	"github.com/vmihailenco/msgpack"
)

type MsgpackTest struct {
	Version float32
	Method  string
	Params  []string
}

func MsgpackPack() {
	msgpack.Marshal(&MsgpackTest{
		Version: 2.0,
		Method:  "echo",
		Params: []string{
			"some",
			"params",
			"here",
		},
	})
}

func MsgpackUnpack() {
	var test MsgpackTest
	msgpack.Unmarshal([]byte(`{"Version":2.0,"Method":"echo","params":["some","params","here"]}`), &test)
}
