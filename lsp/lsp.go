package lsp

import "fmt"

type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

type Bird struct{}

func (b *Bird) Fly() {
	fmt.Println(" > Flying bird is flying")
}

func (b *Bird) Walk() {
	fmt.Println(" > Working bird is walking")
}

type Duck struct {
	Bird
}

func (d *Duck) Fly() {
	d.Bird.Fly()
}

func (d *Duck) Walk() {
	fmt.Println(" > Working duck is walking")
}

type Penguin struct{}

func (p *Penguin) Walk() {
	fmt.Println(" > Working penguin is walking")
}

func Example() {
	fmt.Println("Liskov Substitution Principle")

	bird := &Bird{}
	bird.Fly()
	bird.Walk()

	duck := &Duck{}
	duck.Fly()
	duck.Walk()

	penguin := &Penguin{}
	penguin.Walk()
}
