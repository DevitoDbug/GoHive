package main

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type Document struct {
	content string
}

func (d *Document) Read() (outputData string) {
	outputData = d.content
	return outputData
}
func (d *Document) Write(doc string) {
	d.content = doc
}

type Printer struct {
	buffer string
}

func (p *Printer) Read() (bufferContent string) {
	bufferContent = p.buffer
	return bufferContent
}
func (p *Printer) Write(data string) {
	p.buffer = data
}

func Copy(r Reader, w Writer) {
	w.Write(r.Read())
}

func main() {
	doc1 := Document{}
	pri1 := Printer{}

	doc1.Write("This is content in doc1")
	Copy(&doc1, &pri1)

	fmt.Printf("The content of pri1 is:  %q", pri1.buffer)
}
