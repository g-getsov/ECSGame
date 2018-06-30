package guis

const InventoryGuiComponentName = "InventoryGui"

type InventoryGui struct {
	Id string
}

func (i InventoryGui) GetComponentId() string { return i.Id }

func (i InventoryGui) GetComponentName() string { return GetInvnetoryGuiComponentName() }

func (i InventoryGui) IsUniquePerEntity() bool { return true }

func GetInvnetoryGuiComponentName() string { return InventoryGuiComponentName }