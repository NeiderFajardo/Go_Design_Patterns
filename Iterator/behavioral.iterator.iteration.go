package main

import "fmt"

type Person struct {
	FirstName, MiddleName, LastName string
}

func (p *Person) Names() []string {
	return []string{p.FirstName, p.MiddleName, p.LastName}
}

func main() {
	p := Person{"Alexander", "Graham", "Bell"}
	// range permite iterar por sobre colecciones de objetos, por lo que puede ser una aproximación para
	// implementar el patrón Iterator
	for _, name := range p.Names() {
		fmt.Println(name)
	}
}
