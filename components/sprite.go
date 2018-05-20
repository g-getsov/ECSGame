package components

import "github.com/hajimehoshi/ebiten"

const spriteComponentName = "Sprite"

type Sprite struct {
	Image *ebiten.Image
	ThumbnailImage *ebiten.Image
}

func (s Sprite) GetComponentName() string { return GetSpriteComponentName() }

func GetSpriteComponentName() string {return spriteComponentName }