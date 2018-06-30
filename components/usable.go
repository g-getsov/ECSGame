package components

import "BasicECS/closures"

const usableComponentName = "Usable"

type Usable struct {
	Id string
	Use closures.UsageFunc
}

func (u Usable) GetComponentId() string { return u.Id }

func (u Usable) GetComponentName() string { return GetUsableComponentName() }

func (u Usable) IsUniquePerEntity() bool { return true }

func GetUsableComponentName() string { return usableComponentName }