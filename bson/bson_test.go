package bson

import (
	"testing"
)

func BenchmarkBsonPack(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BsonPack()
	}
}

func BenchmarkBsonUnpack(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BsonUnpack()
	}
}
