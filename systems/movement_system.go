package systems

import (
	"BasicECS/components"
	"BasicECS/core"
)

type movementSystem struct {}

func (s movementSystem) Update(dt float64, entityManager *core.EntityManager) {

	entityIds := entityManager.GetAllEntitiesPossessingComponentsOfClass(components.GetMovableComponentName())

	for _, entityId := range entityIds {

		movableComponent, _ := entityManager.GetComponentOfClass(
			components.GetMovableComponentName(),
			entityId).(*components.Movable)

		if !movableComponent.Movable { continue }

		positionComponent, _ := entityManager.GetComponentOfClass(
			components.GetPositionComponentName(),
			entityId).(*components.Position)

		if positionComponent == nil { continue }

		speedComponent, _ := entityManager.GetComponentOfClass(
			components.GetSpeedComponentName(),
			entityId).(*components.Speed)

		if speedComponent == nil { continue }

		staminaComponent, _ := entityManager.GetComponentOfClass(
			components.GetStaminaComponentName(),
			entityId).(*components.Stamina)

		inputComponent := entityManager.GetComponentOfClass(
			components.GetInputComponentName(),
			entityId)

		if inputComponent == nil {
			aiMovement(positionComponent, speedComponent)
		} else {
			controllerMovement(inputComponent.(*components.Input), positionComponent, speedComponent, staminaComponent)
		}
	}
}

func controllerMovement(input *components.Input, position *components.Position, speed *components.Speed, stamina *components.Stamina) {

	movementSpeed := speed.Speed
	speedBonus := 0

	if input.SprintKey && stamina != nil && stamina.Stamina > 0 {
		speedBonus = 5
	}

	hasMoved := false

	if input.ForwardKey {
		position.Y -= movementSpeed + speedBonus
		hasMoved = true
	}
	if input.BackwardsKey {
		position.Y += movementSpeed + speedBonus
		hasMoved = true
	}
	if input.LeftKey {
		position.X -= movementSpeed + speedBonus
		hasMoved = true
	}
	if input.RightKey {
		position.X += movementSpeed + speedBonus
		hasMoved = true
	}

	if hasMoved && input.SprintKey && stamina != nil && stamina.Stamina > 0 {
		stamina.Stamina -= 1
	}
}

func aiMovement(position *components.Position, speed *components.Speed) {
	position.X += speed.Speed
}

func CreateMovementSystem() System { return movementSystem{} }