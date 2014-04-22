package json

import (
	"testing"
)

func BenchmarkJsonPack(b *testing.B) {
	for n := 0; n < b.N; n++ {
		JsonPack()
	}
}

func BenchmarkJsonUnpack(b *testing.B) {
	for n := 0; n < b.N; n++ {
		JsonUnpack()
	}
}
