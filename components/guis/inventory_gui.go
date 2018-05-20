package guis

const InventoryGuiComponentName = "InventoryGui"

type InventoryGui struct {}

func (i InventoryGui) GetComponentName() string { return GetInvnetoryGuiComponentName() }

func (i InventoryGui) IsUniquePerEntity() bool { return true }

func GetInvnetoryGuiComponentName() string { return InventoryGuiComponentName }