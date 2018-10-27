package systems

import (
	"BasicECS/components"
	inp "github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten"
	"BasicECS/core"
)

type inputSystem struct {}

func (s inputSystem) Update(dt float64, entityManager *core.EntityManager) {

	entityIds := entityManager.GetAllEntitiesPossessingComponentsOfClass(components.GetInputComponentName())

	for _, entityId := range entityIds {

		inputComponent, _ := entityManager.GetComponentOfClass(
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

		if inp.IsKeyJustPressed(ebiten.KeyLeft) || inp.IsKeyJustPressed(ebiten.KeyA) {
			inputComponent.LeftKey = true
		}
		if inp.IsKeyJustReleased(ebiten.KeyLeft) || inp.IsKeyJustReleased(ebiten.KeyA) {
			inputComponent.LeftKey = false
		}

		if inp.IsKeyJustPressed(ebiten.KeyRight) || inp.IsKeyJustPressed(ebiten.KeyD) {
			inputComponent.RightKey = true
		}
		if inp.IsKeyJustReleased(ebiten.KeyRight) || inp.IsKeyJustReleased(ebiten.KeyD) {
			inputComponent.RightKey = false
		}

		if inp.IsKeyJustPressed(ebiten.KeyShift) {
			inputComponent.SprintKey = true
		}
		if inp.IsKeyJustReleased(ebiten.KeyShift) {
			inputComponent.SprintKey = false
		}

		if inp.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || inp.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			inputComponent.ShootKey = true
		}
		if inp.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || inp.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			inputComponent.ShootKey = false
		}

		if inp.IsKeyJustPressed(ebiten.KeyE) {
			inputComponent.InteractKey = true
		} else {
			inputComponent.InteractKey = false
		}

		if inp.IsKeyJustPressed(ebiten.KeyI) {
			inputComponent.InventoryToggle = !inputComponent.InventoryToggle
		}
	}
}

func CreateInputSystem() System { return inputSystem{} }