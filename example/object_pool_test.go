package example

import "testing"

func BenchmarkWithoutPooling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		globalSink = &Data{}
		globalSink.Values[0] = 24
	}
}

func BenchmarkWithPooling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := dataPool.Get().(*Data)
		obj.Values[0] = 25
		dataPool.Put(obj)
		globalSink = obj
	}
}
