package factories

import (
	"github.com/hajimehoshi/ebiten"
	"BasicECS/components/weapons"
	"BasicECS/enum"
	cmpt "BasicECS/components"
	"BasicECS/closures"
	"BasicECS/components/guis"
)

func CreateSpeedComponent(speed int) cmpt.Speed {
	return cmpt.Speed{
		Speed: speed,
	}
}

func CreateVelocityComponent(x int, y int) cmpt.Velocity {
	return cmpt.Velocity{
		X: x,
		Y: y,
	}
}

func CreatePositionComponent(x int, y int) cmpt.Position {
	return cmpt.Position{
		X: x,
		Y: y,
	}
}

func CreateHealthComponent(health int) cmpt.Health {
	return cmpt.Health{
		Health: health,
		MaxHealth: health,
	}
}

func CreateInputComponent() cmpt.Input {
	return cmpt.Input{}
}

func CreateSpriteComponent(image *ebiten.Image, thumbnail *ebiten.Image) cmpt.Sprite {
	return cmpt.Sprite{
		Image: image,
		ThumbnailImage:thumbnail,
	}
}

func CreateAmmunitionComponent(damage int, ammunitionType weapons.AmmunitionType) weapons.Ammunition {
	return weapons.Ammunition {
		Damage: damage,
		AmmunitionType: ammunitionType,
	}
}

func CreateWeaponBaseComponent(fireRate float32, projectileSpeed int, damage int) weapons.WeaponBase {
	return weapons.WeaponBase{
		FireRate: fireRate,
		ProjectileSpeed: projectileSpeed,
		Damage: damage,
	}
}

func CreateMagazineComponent(maxCapacity int, ammunition *weapons.Ammunition) weapons.Magazine {
	return weapons.Magazine{
		MaxCapacity: maxCapacity,
		Capacity: maxCapacity,
		Ammunition: ammunition,
	}
}

func CreateDamageComponent(damageType cmpt.DamageType, damage int) cmpt.Damage {
	return cmpt.Damage {
		Type: damageType,
		Damage: damage,
	}
}

func CreateHitboxComponent(width int, height int, isSolid bool) cmpt.Hitbox {
	return cmpt.Hitbox{
		Width: width,
		Height: height,
		IsSolid: isSolid,
	}
}

func CreateEquipableComponent(equipmentType enum.EquipmentType) cmpt.Equippable {
	return cmpt.Equippable{
		EquipmentType: equipmentType,
	}
}

func CreateEquipmentSlotComponent(equipmentType enum.EquipmentType) cmpt.EquipmentSlot {
	return cmpt.EquipmentSlot{
		Type: equipmentType,
	}
}

func CreateAttachableComponent(attachmentType enum.AttachmentType) weapons.Attachable {
	return weapons.Attachable{
		AttachmentType: attachmentType,
	}
}

func CreateAttachmentSlotComponent(attachmentType enum.AttachmentType) weapons.AttachmentSlot {
	return weapons.AttachmentSlot{
		Type: attachmentType,
	}
}

func CreateExpirableComponent(expirationTime float64) cmpt.Expirable {
	return cmpt.Expirable{
		TimeLeft:expirationTime,
	}
}

func CreateInventoryComponent(maxSize int) cmpt.Inventory {
	return cmpt.Inventory{
		MaxSize: maxSize,
		ItemEntityIds: make([]string, maxSize),
	}
}

func CreateOwnableComponent() cmpt.Ownable {
	return cmpt.Ownable{}
}

func CreateMovableComponent() cmpt.Movable {
	return cmpt.Movable{ true }
}

func CreateUsableComponent(usageFunc closures.UsageFunc) cmpt.Usable {
	return cmpt.Usable{
		Use:usageFunc,
	}
}

func CreateInteractiveComponent(interactionFunc closures.InteractionFunc) cmpt.Interactive {
	return cmpt.Interactive{
		Interact:interactionFunc,
	}
}

func CreateCollidedComponent() cmpt.Collided {
	return cmpt.Collided{
		CollidedEntities: make(map[string]bool, 0),
	}
}

func CreateTooltipComponent(text string) cmpt.Tooltip {
	return cmpt.Tooltip{
		Text: text,
	}
}

func CreateRemoveComponent() cmpt.Remove {
	return cmpt.Remove{}
}

func CreateHiddenComponent() cmpt.Hidden {
	return cmpt.Hidden{}
}

//GUI

func CreateGuiComponent(entityId string) guis.Gui {
	return guis.Gui{
		EntityId:entityId,
	}
}

func CreateCharacterStatsGuiComponent() guis.CharacterStatsGui {
	return guis.CharacterStatsGui{}
}

func CreateInventoryGuiComponent() guis.InventoryGui {
	return guis.InventoryGui{}
}
