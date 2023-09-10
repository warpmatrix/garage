package main

import "fmt"

type test struct {
	name string
}

func (t test) Name() string {
	return t.name
}

func impl_func() {
	t := test{name: "test-name"}
	var i interface{} = t
	n := i.(interface {
		Name() string
	})
	fmt.Println(n.Name())
}
