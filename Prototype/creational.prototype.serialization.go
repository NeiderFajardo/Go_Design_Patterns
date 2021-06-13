package prototype

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

// Uso de la desearealización y searialización para hacer una copia
// profunda del objeto, disminuyendo la cantidad de código en los
// casos en los que tengamos que copiar un objecto muy grande
func (p *Person) DeepCopy() *Person {
	// note: no error handling below
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	// peek into structure
	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)
	return &result
}

func main() {
	john := Person{"John",
		&Address{"123 London Rd", "London", "UK"},
		[]string{"Chris", "Matt", "Sam"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Jill")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)

}
