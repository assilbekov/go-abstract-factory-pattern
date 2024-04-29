package pkg

import "errors"

type Factory interface {
	CreateComputer() Computer
	CreateMonitor() Monitor
}

func GetFactory(factoryName string) (Factory, error) {
	switch factoryName {
	case Asus:
		return &AsusFactory{}, nil
	case HP:
		return &HPFactory{}, nil
	default:
		return nil, errors.New("invalid factory name")
	}
}
