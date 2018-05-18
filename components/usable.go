package components

import "BasicECS/closures"

const usableComponentName = "Usable"

type Usable struct {
	Use closures.UsageFunc
}

func (e Usable) GetComponentName() string { return GetUsableComponentName() }

func GetUsableComponentName() string { return usableComponentName }