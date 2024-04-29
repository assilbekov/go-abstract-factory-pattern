package pkg

import "errors"

type Factory interface {
	CreateComputer() Computer
	CreateMonitor() Monitor
}

func GetFactory(factoryName string) (Factory, error) {
	if factoryName == "HP" {
		return &HPFactory{}, nil
	}
	if factoryName == "Asus" {
		return &AsusFactory{}, nil
	}
	return nil, errors.New("invalid factory name")
}
