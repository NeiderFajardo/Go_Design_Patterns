package main

import "fmt"

type Person struct {
	name, position string
}

// Un arreglo de funciones que construyan el objeto de forma dinámica según
// las funciones que se definen para construir la persona
// Se puede extender agregando más funciones al arreglo
type personMod func(*Person)
type PersonBuilder struct {
	actions []personMod
}

// Por ejemplo, acá se asigna un nombre a la persona.
func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

// En esta función se construye la persona y se le aplican todas las funciones
// definidas en el arreglo de funciones personMod en el objeto PersonBuilder
func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

// extend PersonBuilder
func (b *PersonBuilder) WorksAsA(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})
	return b
}

func main() {
	b := PersonBuilder{}
	p := b.Called("Dmitri").WorksAsA("dev").Build()
	fmt.Println(*p)
}
