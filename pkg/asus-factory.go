package pkg

type AsusFactory struct{}

func (a *AsusFactory) CreateComputer() Computer {
	return &AsusComputer{
		Memory: 16,
		CPU:    3,
	}
}

func (a *AsusFactory) CreateMonitor() Monitor {
	return &AsusMonitor{
		Size: 27,
	}
}
