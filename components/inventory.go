package components

const inventoryComponentName = "Inventory"

type Inventory struct {
	entityIds string
}

func (i Inventory) GetComponentName() string { return GetInventoryComponentName() }

func GetInventoryComponentName() string {return inventoryComponentName }