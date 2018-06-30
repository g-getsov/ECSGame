package components

import "strings"

const inventoryComponentName = "Inventory"

type Inventory struct {
	Id string
	MaxSize int
	ItemEntityIds []string
}

func (i Inventory) GetComponentId() string { return i.Id }

func (i Inventory) GetComponentName() string { return GetInventoryComponentName() }

func (i Inventory) IsUniquePerEntity() bool { return true }

func (i *Inventory) AddItem(entityId string) bool {

	idx, hasEmptySpace := i.findFirstFreeSlot()

	if !hasEmptySpace { return false }

	i.ItemEntityIds[idx] = entityId

	return true
}

func (i Inventory) GetItemPosition(entityId string) (int, bool) {

	for idx, val := range i.ItemEntityIds {
		if strings.Compare(val, entityId) == 0 { return idx, true }
	}

	return 0, false
}

func (i *Inventory) RemoveItem(entityId string) {

	idx, hasFoundItem := i.GetItemPosition(entityId)

	if !hasFoundItem { return }

	i.ItemEntityIds[idx] = ""
}

func (i *Inventory) RemoveItemAtPosition() {

}

func (i Inventory) findFirstFreeSlot() (int, bool) {

	for idx, val := range i.ItemEntityIds {
		if len(val) == 0 { return idx, true }
	}

	return 0, false
}

func GetInventoryComponentName() string {return inventoryComponentName }