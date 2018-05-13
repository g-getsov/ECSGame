package components

import "BasicECS/enum"

const equippableComponentName = "Equippable"

type Equippable struct {
	EquipmentType enum.EquipmentType
}

func (e Equippable) GetComponentName() string { return GetEquippableComponentName() }

func GetEquippableComponentName() string { return equippableComponentName }