package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// what if we want factories for specific roles?

// functional approach
func NewEmployeeFactory(position string,
	annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// structural approach
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func NewEmployeeFactory2(position string,
	annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func main() {
	// Acá estamos retornando y guardando en una variable una función, la cual podemos pasar
	// como parámetro en otra función implementando el paradigma funcional
	developerFactory := NewEmployeeFactory("Developer", 60000)
	managerFactory := NewEmployeeFactory("Manager", 80000)
	// Acá llamamos la función que almacenamos en la variable y guardamos su retorno
	// que en este caso es un pointer a un objeto de tipo Employee
	developer := developerFactory("Adam")
	fmt.Println(developer)

	manager := managerFactory("Jane")
	fmt.Println(manager)

	bossFactory := NewEmployeeFactory2("CEO", 100000)
	// can modify post-hoc
	bossFactory.AnnualIncome = 110000
	boss := bossFactory.Create("Sam")
	fmt.Println(boss)
}
