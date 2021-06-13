package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

// Este método es necesario ya que sin el, al hacer una copia
// de la persona estaría copiando la referencia al objeto y no
// creando una copia del objeto como tal
func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

// Se debe crear el DeepCopy para organizar la acción de copiar la
// persona y los procesos de copia internos que se deben hacer
func (p *Person) DeepCopy() *Person {
	q := *p // copies Name
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

func main() {
	john := Person{"John",
		&Address{"123 London Rd", "London", "UK"},
		[]string{"Chris", "Matt"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
