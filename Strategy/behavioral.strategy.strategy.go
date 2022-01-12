package strategy

import (
  "fmt"
  "strings"
)

type OutputFormat int

const (
  Markdown OutputFormat = iota
  Html
)

type ListStrategy interface {
  Start(builder *strings.Builder)
  End(builder *strings.Builder)
  AddListItem(builder *strings.Builder, item string)
}
// Se define una estrategia y se implementan los métodos de la interface ListStrategy
type MarkdownListStrategy struct {}

func (m *MarkdownListStrategy) Start(builder *strings.Builder) {

}

func (m *MarkdownListStrategy) End(builder *strings.Builder) {

}
// En este caso no necesitamos iniciar y finalizar con nada, por lo que solo definimos el método AddListItem
func (m *MarkdownListStrategy) AddListItem(
  builder *strings.Builder, item string) {
  builder.WriteString(" * " + item + "\n")
}
// Para este objeto si implementamos los métodos Start y End.
type HtmlListStrategy struct {}

func (h *HtmlListStrategy) Start(builder *strings.Builder) {
  builder.WriteString("<ul>\n")
}

func (h *HtmlListStrategy) End(builder *strings.Builder) {
  builder.WriteString("</ul>\n")
}

func (h *HtmlListStrategy) AddListItem(builder *strings.Builder, item string) {
  builder.WriteString("  <li>" + item + "</li>\n")
}
// El resultado de usar este objeto cambia dependiendo de la estrategía que se este usando, no se cambia el comportamiento. 
type TextProcessor struct {
  builder strings.Builder
  listStrategy ListStrategy
}

func NewTextProcessor(listStrategy ListStrategy) *TextProcessor {
  return &TextProcessor{strings.Builder{}, listStrategy}
}
// Permite cambiar la estrategía una vez ejemplificado el objeto.
func (t *TextProcessor) SetOutputFormat(fmt OutputFormat) {
  switch fmt {
  case Markdown:
    t.listStrategy = &MarkdownListStrategy{}
  case Html:
    t.listStrategy = &HtmlListStrategy{}
  }
}

func (t *TextProcessor) AppendList(items []string) {
  t.listStrategy.Start(&t.builder)
  for _, item := range items {
    t.listStrategy.AddListItem(&t.builder, item)
  }
  t.listStrategy.End(&t.builder)
}

func (t *TextProcessor) Reset() {
  t.builder.Reset()
}

func (t *TextProcessor) String() string {
  return t.builder.String()
}

func main() {
  tp := NewTextProcessor(&MarkdownListStrategy{})
  tp.AppendList([]string{ "foo", "bar", "baz" })
  fmt.Println(tp)

  tp.Reset()
  tp.SetOutputFormat(Html)
  tp.AppendList([]string{ "foo", "bar", "baz" })
  fmt.Println(tp)
}
