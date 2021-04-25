package main

import (
	"fmt"
)

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hola mi nombre es %s y tengo %d años \n", p.name, p.age)
}

func (p *tiredPerson) SayHello() {
	fmt.Println("Lo siento, estoy muy cansado")
}

func NewPerson(name string, age int) Person {
	if age > 50 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

// Se puede retornar una interface, en este punto solo serán accesibles los métodos de la interface
// definidos para el tipo de objeto
func main() {
	p := NewPerson("Neider", 25)
	p.SayHello()
	oldP := NewPerson("Alejandro", 51)
	oldP.SayHello()
}
