package pkg

type AsusFactory struct{}

func (a *AsusFactory) CreateComputer() Computer {
	return &AsusComputer{}
}

func (a *AsusFactory) CreateMonitor() Monitor {
	return &AsusMonitor{}
}
