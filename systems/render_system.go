package systems

import (
	"BasicECS/components"
	"github.com/hajimehoshi/ebiten"
	"BasicECS/core"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/go-mplusbitmap"
	"image/color"
)

var tooltipColor = color.RGBA{0x80, 0x80, 0x80, 0x80}

type renderSystem struct {
	screen *ebiten.Image
}

func (s renderSystem) Update(dt float64, entityManager *core.EntityManager) {
	s.renderSprites(entityManager)
	s.renderTooltips(entityManager)
}

func (s renderSystem) renderSprites(entityManager *core.EntityManager) {

	spriteEntityIds := entityManager.GetAllEntitiesPossessingComponentsOfClass(components.GetSpriteComponentName())

	for _, entityId := range spriteEntityIds {

		positionComponent, _ := entityManager.GetComponentOfClass(
			components.GetPositionComponentName(),
			entityId).(*components.Position)

		if &positionComponent == nil { continue }

		spriteComponent, _ := entityManager.GetComponentOfClass(
			components.GetSpriteComponentName(),
			entityId).(*components.Sprite)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(positionComponent.X), float64(positionComponent.Y))
		s.screen.DrawImage(spriteComponent.Image, op)
	}
}

func (s renderSystem) renderTooltips(entityManager *core.EntityManager) {

	tooltipEntityIds := entityManager.GetAllEntitiesPossessingComponentsOfClass(components.GetTooltipComponentName())

	// Show tooltip
	for _, entityId := range tooltipEntityIds {

		collidedComponent, _ := entityManager.GetComponentOfClass(
			components.GetCollidedComponentName(),
			entityId).(*components.Collided)

		if collidedComponent == nil { continue }

		positionComponent, _ := entityManager.GetComponentOfClass(
			components.GetPositionComponentName(),
			entityId).(*components.Position)

		if &positionComponent == nil { continue }

		tooltipComponent, _ := entityManager.GetComponentOfClass(
			components.GetTooltipComponentName(),
			entityId).(*components.Tooltip)

		text.Draw(s.screen, tooltipComponent.Text, mplusbitmap.Gothic12r, positionComponent.X, positionComponent.Y, tooltipColor)
	}
}

func CreateRenderSystem(screen *ebiten.Image) System {
	return renderSystem{
	screen,
	}
}
