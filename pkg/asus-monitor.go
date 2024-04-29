package pkg

import "fmt"

type AsusMonitor struct {
	Size int
}

func (a *AsusMonitor) PrintDetails() {
	fmt.Printf("[ASUS] Monitor, Size: %d inch\n", a.Size)
}
