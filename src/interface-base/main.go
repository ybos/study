package main

import (
	"fmt"
)

type Human interface {
	SayHi()
	Sing()
}

type Student struct {
	name   string
	age    int
	school string
}

type Employee struct {
	name    string
	age     int
	company string
}

func (e *Employee) SayHi() {
	fmt.Printf("%s is an employee. How r u?\n", e.name)
}

func (e *Employee) Sing() {
	fmt.Printf("%s is %d old. He is singing a song\n", e.name, e.age)
}

func (s *Student) SayHi() {
	fmt.Printf("%s is a student. How are you?\n", s.name)
}

func (s *Student) Sing() {
	fmt.Printf("%s is %d old. He is singing a song\n", s.name, s.age)
}

func HaveFun(h Human) {
	h.SayHi()
	h.Sing()
}

func main() {
	s := Student{"John", 15, "SIIT"}
	e := Employee{"Neil", 27, "Ecovacs"}

	HaveFun(&s)
	HaveFun(&e)
}
