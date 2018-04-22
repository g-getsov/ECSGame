package systems

import (
	"BasicECS/entities"
	"BasicECS/components"
	inp "github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten"
)

type inputSystem struct {}

func (s inputSystem) Update(dt float64, entityManager *entities.EntityManager) {

	entityIds := entityManager.GetAllEntitiesPossessingComponentsOfClass(components.GetInputComponentName())

	for _, entityId := range entityIds {

		inputComponent := entityManager.GetComponentOfClass(
			components.GetInputComponentName(),
			entityId).(*components.Input)

		if inp.IsKeyJustPressed(ebiten.KeyUp) || inp.IsKeyJustPressed(ebiten.KeyW) {
			inputComponent.ForwardKey = true
		}
		if inp.IsKeyJustReleased(ebiten.KeyUp) || inp.IsKeyJustReleased(ebiten.KeyW) {
			inputComponent.ForwardKey = false
		}

		if inp.IsKeyJustPressed(ebiten.KeyDown) || inp.IsKeyJustPressed(ebiten.KeyS) {
			inputComponent.BackwardsKey = true
		}
		if inp.IsKeyJustReleased(ebiten.KeyDown) || inp.IsKeyJustReleased(ebiten.KeyS) {
			inputComponent.BackwardsKey = false
		}

		if inp.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || inp.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			inputComponent.ShootKey = true
		}
		if inp.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || inp.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			inputComponent.ShootKey = false
		}
	}
}

func CreateInputSystem() System { return inputSystem{} }