package main

type rect struct {
	height int
	lenght int
}

func (r rect) area() int {
	return r.height * r.lenght
}

var r = rect{
	height: 20,
	lenght: 30,
}

var t = rect{
	height: 21,
	lenght: 21,
}
