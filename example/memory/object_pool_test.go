package memory

import (
	"testing"
)

func BenchmarkWithoutPooling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := &Data{}
		obj.Values[0] = 24
		//runtime.KeepAlive(obj)
		globalSink = obj
	}
}

func BenchmarkWithPooling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := dataPool.Get().(*Data)
		obj.Values[0] = 25
		dataPool.Put(obj)
		//runtime.KeepAlive(obj)
		globalSink = obj
	}
}
