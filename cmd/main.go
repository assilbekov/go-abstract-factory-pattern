package main

import "abstract-factory/pkg"

func main() {
	asusFactory, err := pkg.GetFactory("Asus")
	if err != nil {
		panic(err)
	}
	asusComputer := asusFactory.CreateComputer()
	asusMonitor := asusFactory.CreateMonitor()

	hpFactory, err := pkg.GetFactory("HP")
	if err != nil {
		panic(err)
	}
	hpComputer := hpFactory.CreateComputer()
	hpMonitor := hpFactory.CreateMonitor()

	asusComputer.PrintDetails()
	asusMonitor.PrintDetails()

	hpComputer.PrintDetails()
	hpMonitor.PrintDetails()
}
