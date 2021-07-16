package Facade

import "fmt"

type Buffer struct {
	width, height int
	buffer        []rune
}

func NewBuffer(width, height int) *Buffer {
	return &Buffer{width, height,
		make([]rune, width*height)}
}

func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

type Viewport struct {
	buffer *Buffer
	offset int
}

func NewViewport(buffer *Buffer) *Viewport {
	return &Viewport{buffer: buffer}
}

func (v *Viewport) GetCharacterAt(index int) rune {
	return v.buffer.At(v.offset + index)
}

// a facade over buffers and viewports
type Console struct {
	buffers   []*Buffer
	viewports []*Viewport
	offset    int
}

// Este objeto actua como una Facade para evitar que se tenga que trabajar con
// objetos de bajo nivel como el Buffer o el Viewport
func NewConsole() *Console {
	b := NewBuffer(10, 10)
	v := NewViewport(b)
	// Se inicia la consola que contiene apuntadores a arreglos y se agregan los objetos
	// que se acaban de crear
	return &Console{[]*Buffer{b}, []*Viewport{v}, 0}
}

/*
  En este ejemplo no se alcanza a apreciar del todo, ya que los métodos y los objetos que
  están por debajo de la fachada no son complejos, pero la idea con este método es simplificar
  cualquiera que sean las operaciones que se hagan por debajo
*/
func (c *Console) GetCharacterAt(index int) rune {
	return c.viewports[0].GetCharacterAt(index)
}

func main() {
	c := NewConsole()
	u := c.GetCharacterAt(1)
	fmt.Println(u)
}
