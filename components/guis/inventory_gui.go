package guis

const InventoryGuiComponentName = "InventoryGui"

type InventoryGui struct {}

func (c InventoryGui) GetComponentName() string { return GetInvnetoryGuiComponentName() }

func GetInvnetoryGuiComponentName() string { return InventoryGuiComponentName }