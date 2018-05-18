package components

const hitboxComponentName = "Hitbox"

type Hitbox struct {
	Width int
	Height int
	IsSolid bool
}

func (c Hitbox) GetComponentName() string { return GetHitboxComponentName() }

func GetHitboxComponentName() string {return hitboxComponentName }