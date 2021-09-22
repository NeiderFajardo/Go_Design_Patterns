package state

import "fmt"

type Switch struct {
  // Atributo que implementa la interface State.. el estado del objeto depende del
  // valor(tipo) de este atributo
  State State
}

func NewSwitch() *Switch {
  // El objeto se inicializa en estado apagado
  return &Switch{NewOffState()}
}

/*
  El switch puede pasar de on a off, pero en cualquier momento se puede invocar cualquier
  operación, esa operación será tomada dependiendo del estado en el que se encuentre.
*/
func (s *Switch) On() {
  s.State.On(s)
}

func (s *Switch) Off() {
  s.State.Off(s)
}

type State interface {
  On(sw *Switch)
  Off(sw *Switch)
}
// Cuando no se le define una operación a un objeto estado, esta estructura define el comportamiento
// de los métodos no implementados
type BaseState struct {}

func (s *BaseState) On(sw *Switch) {
  fmt.Println("Light is already on")
}

func (s *BaseState) Off(sw *Switch) {
  fmt.Println("Light is already off")
}

type OnState struct {
  BaseState
}

func NewOnState() *OnState {
  fmt.Println("Light turned on")
  return &OnState{BaseState{}}
}
// Cuando esta on, solo queremos manejar cuando se apague
func (o *OnState) Off(sw *Switch) {
  fmt.Println("Turning light off...")
  sw.State = NewOffState()
}

type OffState struct {
  BaseState
}

func NewOffState() *OffState {
  fmt.Println("Light turned off")
  return &OffState{BaseState{}}
}

func (o *OffState) On(sw *Switch) {
  fmt.Println("Turning light on...")
  sw.State = NewOnState()
}

func main() {
  sw := NewSwitch()
  sw.On()
  sw.Off()
  sw.Off()
}