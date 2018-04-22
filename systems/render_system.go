package systems

import (
	"BasicECS/entities"
	"BasicECS/components"
	"github.com/hajimehoshi/ebiten"
)

type renderSystem struct {
	screen *ebiten.Image
}

func (s renderSystem) Update(dt float64, entityManager *entities.EntityManager) {

	entityIds := entityManager.GetAllEntitiesPossessingComponentsOfClass(components.GetSpriteComponentName())

	for _, entityId := range entityIds {

		positionComponent := entityManager.GetComponentOfClass(
			components.GetPositionComponentName(),
			entityId).(*components.Position)

		if &positionComponent == nil { continue }

		spriteComponent := entityManager.GetComponentOfClass(
			components.GetSpriteComponentName(),
			entityId).(*components.Sprite)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(positionComponent.X), float64(positionComponent.Y))
		s.screen.DrawImage(spriteComponent.Image, op)
	}
}

func CreateRenderSystem(screen *ebiten.Image) System { return renderSystem{
	screen,
} }
