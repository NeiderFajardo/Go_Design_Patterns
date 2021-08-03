package main

import "fmt"

type Image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func (b *Bitmap) Draw() {
	fmt.Println("Drawing image", b.filename)
}

func NewBitmap(filename string) *Bitmap {
	fmt.Println("Loading image from", filename)
	return &Bitmap{filename: filename}
}

func DrawImage(image Image) {
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Done drawing the image")
}

// No se crea el objeto Bitmap real hasta que vaya a ser usado
type LazyBitmap struct {
	filename string
	bitmap   *Bitmap
}

// En el momento en el que quiero dibujar la imagen hago un llamado al método Draw, de ahí
// la importancia de que implementen las mismas interfaces, de esta forma quien quiere dibujar
// no eviedencia la presencial del proxy mientras este hace su trabajo..
// En este caso inicializando la primera vez que se llama el método al objeto BitMap para dibujar
func (l *LazyBitmap) Draw() {
	if l.bitmap == nil {
		l.bitmap = NewBitmap(l.filename)
	}
	l.bitmap.Draw()
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{filename: filename}
}

func main() {
	//bmp := NewBitmap("demo.png")
	bmp := NewLazyBitmap("demo.png")
	DrawImage(bmp)
}
