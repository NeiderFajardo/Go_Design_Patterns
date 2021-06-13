package factories

import "fmt"

// Plantilla del objeto que se quiere construir
type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// Numerador para cada tipo de empleado de la empresa
const (
	Developer = iota
	Manager
)

// Método que crea un empleado plantilla del tipo del rol que se le
// pasa como argumento
func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "Developer", 60000}
	case Manager:
		return &Employee{"", "Manager", 80000}
	default:
		panic("unsupported role")
	}
}

func main() {
	// Una vez creado el objeto de tipo plantilla se puede modificar
	// con los valores que se necesitan
	m := NewEmployee(Manager)
	m.Name = "Sam"
	fmt.Println(m)
}
