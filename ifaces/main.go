package main

import "fmt"

type vehicle interface {
	move()
}

type car struct{}

type aircraft struct{}

func (c car) move() {
	fmt.Println("Auto moving")
}

func (a aircraft) move() {
	fmt.Println("Air flying")
}

func main() {

	var tesla vehicle = car{}
	var boing vehicle = aircraft{}
	tesla.move()
	boing.move()
}
