package components

import "BasicECS/enum"

const equipableComponentName = "Equipable"

type Equipable struct {
	Id string
	EquipmentType enum.EquipmentType
}

func (e Equipable) GetComponentId() string { return e.Id }

func (e Equipable) GetComponentName() string { return GetEquipableComponentName() }

func (e Equipable) IsUniquePerEntity() bool { return true }

func GetEquipableComponentName() string { return equipableComponentName }