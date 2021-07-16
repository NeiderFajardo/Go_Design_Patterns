package decorator

import "fmt"

// Este sería el comportamiento al que queremos que extiendan los objetos
// ya definidos
type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f",
		c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side %f", s.Side)
}

// possible, but not generic enough
type ColoredSquare struct {
	Square
	Color string
}

// Esta estructura permite exteneder el comportamiento de las figuras
// que implementen la interface Shape, permitiendo también asignarles un color
type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s",
		c.Shape.Render(), c.Color)
}

// Esta estructura permite exteneder el comportamiento de las figuras
// que implementen la interface Shape, permitiendo también asignarles trasnparencia
type TransparentShape struct {
	Shape        Shape
	Transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has %f%% transparency",
		t.Shape.Render(), t.Transparency*100.0)
}

func main() {
	circle := Circle{2}
	fmt.Println(circle.Render())
	// Podemos extender la funcionalidad del círculo porque implementa la interface shape
	redCircle := ColoredShape{&circle, "Red"}
	fmt.Println(redCircle.Render())

	rhsCircle := TransparentShape{&redCircle, 0.5}
	fmt.Println(rhsCircle.Render())
}
