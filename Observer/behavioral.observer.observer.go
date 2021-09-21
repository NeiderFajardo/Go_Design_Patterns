package observer

import (
  "container/list"
  "fmt"
)
// Un obserbable tiene una lista de los objectos que están subscritos a él
type Observable struct {
  subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
  // Subscribirse es agregar el objeto observador al arreglo
  o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
  for z := o.subs.Front(); z != nil; z = z.Next() {
    if z.Value.(Observer) == x {
      o.subs.Remove(z)
    }
  }
}
// Se debe tener un método que dispare una acción en el objeto observable
// y este notifique los cambios a sus subscriptores
func (o *Observable) Fire(data interface{}) {
  for z := o.subs.Front(); z != nil; z = z.Next() {
    z.Value.(Observer).Notify(data)
  }
}
// Todo objeto observable debe implementar esta interface para poder llamar su método Notify
type Observer interface {
  Notify(data interface{})
}

// whenever a person catches a cold,
// a doctor must be called
type Person struct {
  Observable
  Name string
}

func NewPerson(name string) *Person {
  return &Person {
    Observable: Observable{new(list.List)},
    Name: name,
  }
}

func (p *Person) CatchACold() {
  p.Fire(p.Name)
}

type DoctorService struct {}

func (d *DoctorService) Notify(data interface{}) {
  fmt.Printf("A doctor has been called for %s",
    data.(string))
}

func main() {
  p := NewPerson("Boris")
  ds := &DoctorService{}
  p.Subscribe(ds)

  // let's test it!
  p.CatchACold()
}