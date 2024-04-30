package srp

import "fmt"

type Employee struct {
	Name  string
	Email string
}

type EmailUpdater struct{}
type NameUpdater struct{}

func (u *NameUpdater) UpdateName(e *Employee, name string) {
	e.Name = name
}

func (u *EmailUpdater) UpdateEmail(e *Employee, email string) {
	e.Email = email
}

func Example() {
	fmt.Println("Single Responsibility Principle:")

	emp := Employee{}
	eUp := EmailUpdater{}
	nUp := NameUpdater{}
	nUp.UpdateName(&emp, "Devไปวันๆ")
	eUp.UpdateEmail(&emp, "name@email.dev")

	fmt.Println(" > ", emp)
}
