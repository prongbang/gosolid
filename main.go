package main

import (
	"github.com/prongbang/gosolid/dip"
	"github.com/prongbang/gosolid/isp"
	"github.com/prongbang/gosolid/lsp"
	"github.com/prongbang/gosolid/ocp"
	"github.com/prongbang/gosolid/srp"
)

func main() {
	// 1. Single Responsibility Principle
	srp.Example()

	// 2. Open/Closed Principle
	ocp.Example()

	// 3. Liskov Substitution Principle
	lsp.Example()

	// 4. Interface Segregation Principle
	isp.Example()

	// 5. Dependency Inversion Principle
	dip.Example()
}
