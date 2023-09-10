package main

import "fmt"

type base struct {
	base_field int
}

func (b base) base_method() {
	fmt.Println("in base method")
}

type derived struct {
	base
	derived_field int
}

func inherit_demo() {
	var d derived
	d.base_field = 1
	d.derived_field = 2
	d.base_method()
}
