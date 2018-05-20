package components

import "BasicECS/enum"

const equipmentSlotComponentName = "EquipmentSlot"

type EquipmentSlot struct {
	EntityId string
	Type enum.EquipmentType
}

func (e EquipmentSlot) GetComponentName() string { return GetEquipmentSlotComponentName() }

func (e EquipmentSlot) IsUniquePerEntity() bool { return true }

func GetEquipmentSlotComponentName() string { return equipmentSlotComponentName }