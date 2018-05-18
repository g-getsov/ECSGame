package components

const inventoryComponentName = "Inventory"

type Inventory struct {
	ItemEntityIds map[string]bool
}

func (i Inventory) GetComponentName() string { return GetInventoryComponentName() }

func GetInventoryComponentName() string {return inventoryComponentName }