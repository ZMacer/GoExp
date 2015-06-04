package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{2, 5}
	v.X = 10
	fmt.Println(v.X)
	fmt.Println(Vertex{2, 4})
	m := Vertex{1, 2}
	p := &m
	p.X = 1e9
	fmt.Println(m)
}
