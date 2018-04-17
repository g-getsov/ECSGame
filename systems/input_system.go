package systems

import (
	"BasicECS/entities"
	"BasicECS/components"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten"
)

type inputSystem struct {}

func (s inputSystem) Update(dt float32, entityManager *entities.EntityManager) {

	entityIds := entityManager.GetAllEntitiesPossessingComponentsOfClass(components.GetInputComponentName())

	for _, entityId := range entityIds {

		positionComponent := *entityManager.GetComponentOfClass(
			components.GetPositionComponentName(),
			entityId)

		positionComponentt := positionComponent.(components.Position)

		speedComponent := *entityManager.GetComponentOfClass(
			components.GetSpeedComponentName(),
			entityId)

		speedComponentt := speedComponent.(components.Speed)

		if speedComponent != nil {
			if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
				positionComponentt.Y += speedComponentt.Speed
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
				positionComponentt.Y -= speedComponentt.Speed
			}
		}
	}
}

func CreateInputSystem() System { return inputSystem{} }