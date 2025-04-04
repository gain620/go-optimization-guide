package example

import "sync"

type Data struct {
	Values [1024]int
}

var globalSink *Data

func createData() *Data {
	return &Data{
		Values: [1024]int{1, 2, 3, 4},
	}
}

var dataPool = sync.Pool{
	New: func() interface{} {
		return &Data{}
	},
}
