package pkg

type HPFactory struct{}

func (h *HPFactory) CreateComputer() Computer {
	return &HPComputer{
		Memory: 8,
		CPU:    2,
	}
}

func (h *HPFactory) CreateMonitor() Monitor {
	return &HPMonitor{
		Size: 24,
	}
}
