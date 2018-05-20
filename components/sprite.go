package components

import "github.com/hajimehoshi/ebiten"

const spriteComponentName = "Sprite"

type Sprite struct {
	Image *ebiten.Image
	ThumbnailImage *ebiten.Image
}

func (s Sprite) GetComponentName() string { return GetSpriteComponentName() }

func (s Sprite) IsUniquePerEntity() bool { return true }

func GetSpriteComponentName() string {return spriteComponentName }