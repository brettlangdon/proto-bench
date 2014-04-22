Some redimentary benchmarks for performance + size of different over the wire protocols.


```bash
$ git clone git://github.com/brettlangdon/proto-bench
$ cd ./proto-bench
$ go run main.go
Struct Size: 48
BSON Size: 24
BSON Length: 90
Zversion@methodechoparams+0some1params2here
JSON Size: 24
JSON Length: 63
{"Version":2,"Method":"echo","Params":["some","params","here"]}
MSGPACK Size: 24
MSGPACK Length: 51
��Version�@�Method�echo�Params��some�params�here
$ go test -bench=. -benchmem ./bson ./json ./msgpack
PASS
BenchmarkBsonPack	  500000	      3950 ns/op	     621 B/op	      23 allocs/op
BenchmarkBsonUnpack	 1000000	      1206 ns/op	     261 B/op	       6 allocs/op
ok  	github.com/brettlangdon/proto-bench/bson	3.303s
PASS
BenchmarkJsonPack	 1000000	      1723 ns/op	     287 B/op	       4 allocs/op
BenchmarkJsonUnpack	  500000	      5349 ns/op	     576 B/op	      15 allocs/op
ok  	github.com/brettlangdon/proto-bench/json	4.472s
PASS
BenchmarkMsgpackPack	 1000000	      1585 ns/op	     286 B/op	       6 allocs/op
BenchmarkMsgpackUnpack	 1000000	      1440 ns/op	     390 B/op	       7 allocs/op
ok  	github.com/brettlangdon/proto-bench/msgpack	3.064s
```
