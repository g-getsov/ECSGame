package systems

import (
	"BasicECS/core"
	"BasicECS/components"
	"BasicECS/utils"
)

type interactionSystem struct {}

func (s interactionSystem) Update(dt float64, entityManager *core.EntityManager) {

	entityIds := entityManager.GetAllEntitiesPossessingComponentsOfClass(components.GetInteractiveComponentName())

	for _, entityId := range entityIds {

		if utils.IsEntityHidden(entityManager, entityId) { continue }

		collidedComponent, _ := entityManager.GetComponentOfClass(
			components.GetCollidedComponentName(),
			entityId).(*components.Collided)

		if collidedComponent == nil { continue }

		for collidedEntityId := range collidedComponent.CollidedEntities {

			inputComponent, _ := entityManager.GetComponentOfClass(
				components.GetInputComponentName(),
				collidedEntityId).(*components.Input)

			if inputComponent != nil && inputComponent.InteractKey {

				interactiveComponent, _ := entityManager.GetComponentOfClass(
					components.GetInteractiveComponentName(),
					entityId).(*components.Interactive)

				interactiveComponent.Interact(entityManager, entityId, collidedEntityId)
			}
		}
	}
}

func CreateInteractionSystem() System { return interactionSystem{} }