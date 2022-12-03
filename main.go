package main

import "fmt"

type (
	Cache interface {
		Put(key string, value []byte) error
		Get(key string) ([]byte, error)
	}

	Widget struct {
		cache Cache
	}
)

func NewWidget() *Widget {
	return &Widget{}
}

func (w *Widget) Example(k string) ([]byte, error) {
	var data []byte
	var err error

	if w.cache != nil {
		if data, err = w.cache.Get(k); err != nil {
			return nil, err
		}
	}

	if data == nil {
		data = GenerateData()
	}

	return data, nil
}

func GenerateData() []byte {
	return []byte("some example data")
}

func main() {
	w := NewWidget()
	data, err := w.Example("foo")
	if err != nil {
		panic(err)
	}
	fmt.Printf("data: %s\n", data)
}
