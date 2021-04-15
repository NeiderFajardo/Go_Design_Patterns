package main

import "fmt"

type Person struct {
	StreetAddress, Postcode, City string
	CompanyName, Position         string
	AnnualIncome                  int
}

type PersonBuilder struct {
	person *Person // needs to be inited
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

func (it *PersonBuilder) Build() *Person {
	return it.person
}

func (it *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*it}
}

func (it *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*it}
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (pjb *PersonJobBuilder) At(
	companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}

func (pjb *PersonJobBuilder) AsA(
	position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

func (pjb *PersonJobBuilder) Earning(
	annualIncome int) *PersonJobBuilder {
	pjb.person.AnnualIncome = annualIncome
	return pjb
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (it *PersonAddressBuilder) At(
	streetAddress string) *PersonAddressBuilder {
	it.person.StreetAddress = streetAddress
	return it
}

func (it *PersonAddressBuilder) In(
	city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}

func (it *PersonAddressBuilder) WithPostcode(
	postcode string) *PersonAddressBuilder {
	it.person.Postcode = postcode
	return it
}

func main() {
	//Inicializando un objeto PersonBuilder que retorna un apuntador a dicho objeto
	// pb apunta al objeto PersonuBuilder
	pb := NewPersonBuilder()

	pb.
		// Lives es un método definido para un objeto que apunte a un PersonBuilder y retorna un objeto
		// de tipo apunatdor a PersonAddressBuilder
		Lives().
		// El At, In, WithPostCode cada uno esta definido para un objeto de tipo apuntador a
		// PersonAddressBuilder y retorna el mismo tipo, por lo que es posible encadenar métodos
		At("123 London Road").
		In("London").
		WithPostcode("SW12BC").
		// Ya que el objeto de tipo PersonAddressBuilder es también un objeto de tipo PersonBuilder
		// El cuál es recibido por el método Works, se puede encadenar dicho método para seguir
		// con la construcción del objeto
		Works().
		At("Fabrikam").
		AsA("Programmer").
		Earning(123000)
		// Solo retorna el valor del atriburo persona que es un apuntador de Person
	person := pb.Build()
	// Se imprime el valor a donde está apuntando
	fmt.Println(*person)
	fmt.Println(person)
}
