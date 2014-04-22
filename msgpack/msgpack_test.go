package msgpack

import (
	"testing"
)

func BenchmarkMsgpackPack(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MsgpackPack()
	}
}

func BenchmarkMsgpackUnpack(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MsgpackUnpack()
	}
}
