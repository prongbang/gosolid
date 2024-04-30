package isp

import "fmt"

type Worker interface {
	Work()
}

type Sleeper interface {
	Sleep()
}

type Eater interface {
	Eat()
}

type HumanWorker struct{}

func (h *HumanWorker) Work() {
	fmt.Println(" > Humen wark")
}

func (h *HumanWorker) Sleep() {
	fmt.Println(" > Humen sleep")
}

func (h *HumanWorker) Eat() {
	fmt.Println(" > Humen eat")
}

type RobotWorker struct{}

func (r *RobotWorker) Work() {
	fmt.Println(" > Robot wark")
}

func Example() {
	fmt.Println("Interface Segregation Principle")

	humen := &HumanWorker{}
	humen.Work()
	humen.Eat()
	humen.Sleep()

	robot := &RobotWorker{}
	robot.Work()
}
