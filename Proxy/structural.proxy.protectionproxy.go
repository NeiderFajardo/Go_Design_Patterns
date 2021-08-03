package proxy

import "fmt"

type Driven interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Car being driven")
}

type Driver struct {
	Age int
}

// Este objeto contiene un objeto de tipo Car para agregar una capa de seguridad
// que en este caso implica definir quien puede o no manejar el carro
type CarProxy struct {
	car    Car
	driver *Driver
}

// Esta función se usa para agregar la protección al Carro, si la validación lo permite
// este método llama el método del objeto original
func (c *CarProxy) Drive() {
	if c.driver.Age >= 16 {
		c.car.Drive()
	} else {
		fmt.Println("Driver too young")
	}
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{Car{}, driver}
}

func main() {
	car := NewCarProxy(&Driver{12})
	car.Drive()
}
