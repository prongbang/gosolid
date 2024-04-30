package lsp

import "fmt"

type FlyingBird interface {
	Fly()
}

type WalkingBird interface {
	Walk()
}

type Duck struct{}

func (d *Duck) Fly() {
	fmt.Println(" > Duck is flying")
}

func (d *Duck) Walk() {
	fmt.Println(" > Duck is walking")
}

type Penguin struct{}

func (p *Penguin) Walk() {
	fmt.Println(" > Penguin is walking")
}

func Example() {
	fmt.Println("Liskov Substitution Principle")

	duck := &Duck{}
	duck.Fly()
	duck.Walk()

	penguin := &Penguin{}
	penguin.Walk()
}
