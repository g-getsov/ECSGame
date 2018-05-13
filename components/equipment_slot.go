package components

import "BasicECS/enum"

const equipmentSlotComponentName = "EquipmentSlot"

type EquipmentSlot struct {
	EntityId string
	Type enum.EquipmentType
}

func (e EquipmentSlot) GetComponentName() string { return GetEquipmentSlotComponentName() }

func GetEquipmentSlotComponentName() string { return equipmentSlotComponentName }