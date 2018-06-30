package components

const hitboxComponentName = "Hitbox"

type Hitbox struct {
	Id string
	Width int
	Height int
	IsSolid bool
}

func (h Hitbox) GetComponentId() string { return h.Id }

func (h Hitbox) GetComponentName() string { return GetHitboxComponentName() }

func (h Hitbox) IsUniquePerEntity() bool { return true }

func GetHitboxComponentName() string {return hitboxComponentName }