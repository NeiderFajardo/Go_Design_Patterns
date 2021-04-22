package main

import "strings"

// La idea es que este objeto no sea visible para el usuario
type email struct {
	from, to, subject, body string
}

// Este sería el objeto que interactuará con el objeto email y la forma que tiene el
// usuario de poder hacer uso de email
type EmailBuilder struct {
	email email
}

// Estos métodos usan EmailBuilder para crear el email, se retorna el puntero para
// poder encadenar los llamados
func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain @")
	}
	b.email.from = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.email.to = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

func sendMailImpl(email *email) {
	// actually ends the email
}

// Se define un tipo función que recibe el apuntador a un objeto EmailBuilder
type build func(*EmailBuilder)

// Esta función recibe como parámetro la función definida para poder aplicarla al EmailBuilder
// el cual se inicia acá, por lo que nunca se tiene acceso directamente al email.
func SendEmail(action build) {
	// Se inicializa el objeto
	builder := EmailBuilder{}
	// Se aplica la función
	action(&builder)
	sendMailImpl(&builder.email)
}

func main() {
	// Dentro del llamado definimos la función en donde queremos manipular el objeto
	// EmailBuilder que se define dentro de la función que estamos llamando
	SendEmail(func(b *EmailBuilder) {
		b.From("foo@bar.com").
			To("bar@baz.com").
			Subject("Meeting").
			Body("Hello, do you want to meet?")
	})
}
