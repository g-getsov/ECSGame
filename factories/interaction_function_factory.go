package factories

import (
	"BasicECS/core"
	"BasicECS/components"
)

func PickUp (entityManager *core.EntityManager, entityId string, interactingEntityId string) bool {

	ownableComponent, _ := entityManager.GetComponentOfClass(
		components.GetOwnableComponentName(),
		entityId).(*components.Ownable)

	if ownableComponent == nil { return false }

	inventoryComponent, _ := entityManager.GetComponentOfClass(
		components.GetInventoryComponentName(),
		interactingEntityId).(*components.Inventory)

	if inventoryComponent == nil { return false }

	ownableComponent.OwnerEntityId = interactingEntityId
	hasAddedItem := inventoryComponent.AddItem(entityId)

	if !hasAddedItem { return false }

	entityManager.AddComponentToEntity(entityId, CreateHiddenComponent())

	return true
}