package pkg

type Factory interface {
	CreateComputer() Computer
	CreateMonitor() Monitor
}
