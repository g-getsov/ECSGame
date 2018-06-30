package components

import "BasicECS/enum"

const equipmentSlotComponentName = "EquipmentSlot"

type EquipmentSlot struct {
	Id string
	EntityId string
	Type enum.EquipmentType
}

func (e EquipmentSlot) GetComponentId() string { return e.Id }

func (e EquipmentSlot) GetComponentName() string { return GetEquipmentSlotComponentName() }

func (e EquipmentSlot) IsUniquePerEntity() bool { return false }

func GetEquipmentSlotComponentName() string { return equipmentSlotComponentName }