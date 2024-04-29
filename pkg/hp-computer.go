package pkg

import "fmt"

type HPComputer struct {
	Memory int
	CPU    int
}

func (a *HPComputer) PrintDetails() {
	fmt.Printf("[HP] Computer, Memory: %dgb, CPU: %d\n", a.Memory, a.CPU)
}
