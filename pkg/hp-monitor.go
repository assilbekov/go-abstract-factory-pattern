package pkg

import "fmt"

type HPMonitor struct {
	Size int
}

func (h *HPMonitor) PrintDetails() {
	fmt.Printf("[HP] Monitor, Size: %d inch\n", h.Size)
}
