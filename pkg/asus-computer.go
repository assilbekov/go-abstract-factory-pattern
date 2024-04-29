package pkg

import "fmt"

type AsusComputer struct {
	Memory int
	CPU    int
}

func (a *AsusComputer) PrintDetails() {
	fmt.Printf("Memory: %dgb, CPU: %d\n", a.Memory, a.CPU)
}
