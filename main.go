package main

import (
	"encoding/json"
	"fmt"
	"github.com/vmihailenco/msgpack"
	"labix.org/v2/mgo/bson"
	"unsafe"
)

type Request struct {
	Version float32
	Method  string
	Params  []string
}

func main() {
	var req = Request{
		Version: 2.0,
		Method:  "echo",
		Params: []string{
			"some",
			"params",
			"here",
		},
	}

	fmt.Printf("Struct Size: %d\r\n", unsafe.Sizeof(req))

	bson_str, _ := bson.Marshal(&req)
	fmt.Printf("BSON Size: %d\r\n", unsafe.Sizeof(bson_str))
	fmt.Printf("BSON Length: %d\r\n", len(bson_str))
	fmt.Printf("%s\r\n", bson_str)

	json_str, _ := json.Marshal(&req)
	fmt.Printf("JSON Size: %d\r\n", unsafe.Sizeof(json_str))
	fmt.Printf("JSON Length: %d\r\n", len(json_str))
	fmt.Printf("%s\r\n", json_str)

	msgpack_str, _ := msgpack.Marshal(&req)
	fmt.Printf("MSGPACK Size: %d\r\n", unsafe.Sizeof(msgpack_str))
	fmt.Printf("MSGPACK Length: %d\r\n", len(msgpack_str))
	fmt.Printf("%s\r\n", msgpack_str)
}
