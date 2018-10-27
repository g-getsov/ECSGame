package systems

import (
	"BasicECS/core"
)

type sprintingSystem struct {}

func (s sprintingSystem) Update(dt float64, entityManager *core.EntityManager) {

/*	entityIds := entityManager.GetAllEntitiesPossessingComponentsOfClass(components.GetStaminaComponentName())

	for _, entityId := range entityIds {

		inputComponent, _ := entityManager.GetComponentOfClass(
			components.GetInputComponentName(),
			entityId).(*components.Input)


	}*/

	//move sprinting logic from input component to here

}

func CreateSprintingSystem() System { return sprintingSystem{} }
