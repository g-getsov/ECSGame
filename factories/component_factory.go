package factories

import (
	"github.com/hajimehoshi/ebiten"
	"BasicECS/components/weapons"
	"BasicECS/enum"
	cmpt "BasicECS/components"
	"BasicECS/closures"
	"BasicECS/components/guis"
	"BasicECS/utils/uidutils"
)

func CreateSpeedComponent(speed int) cmpt.Speed {
	return cmpt.Speed{
		Id: uidutils.GenerateNewUniqueId(),
		Speed: speed,
	}
}

func CreateVelocityComponent(x int, y int) cmpt.Velocity {
	return cmpt.Velocity{
		Id: uidutils.GenerateNewUniqueId(),
		X: x,
		Y: y,
	}
}

func CreatePositionComponent(x int, y int) cmpt.Position {
	return cmpt.Position{
		Id: uidutils.GenerateNewUniqueId(),
		X: x,
		Y: y,
	}
}

func CreateHealthComponent(health int) cmpt.Health {
	return cmpt.Health{
		Id: uidutils.GenerateNewUniqueId(),
		Health: health,
		MaxHealth: health,
	}
}

func CreateInputComponent() cmpt.Input {
	return cmpt.Input{
		Id: uidutils.GenerateNewUniqueId(),
	}
}

func CreateSpriteComponent(image *ebiten.Image, thumbnail *ebiten.Image) cmpt.Sprite {
	return cmpt.Sprite{
		Id: uidutils.GenerateNewUniqueId(),
		Image: image,
		ThumbnailImage:thumbnail,
	}
}

func CreateAmmunitionComponent(damage int, ammunitionType weapons.AmmunitionType) weapons.Ammunition {
	return weapons.Ammunition {
		Id: uidutils.GenerateNewUniqueId(),
		Damage: damage,
		AmmunitionType: ammunitionType,
	}
}

func CreateWeaponBaseComponent(fireRate float32, projectileSpeed int, damage int) weapons.WeaponBase {
	return weapons.WeaponBase{
		Id: uidutils.GenerateNewUniqueId(),
		FireRate: fireRate,
		ProjectileSpeed: projectileSpeed,
		Damage: damage,
	}
}

func CreateMagazineComponent(maxCapacity int, ammunition *weapons.Ammunition) weapons.Magazine {
	return weapons.Magazine{
		Id: uidutils.GenerateNewUniqueId(),
		MaxCapacity: maxCapacity,
		Capacity: maxCapacity,
		Ammunition: ammunition,
	}
}

func CreateDamageComponent(damageType cmpt.DamageType, damage int) cmpt.Damage {
	return cmpt.Damage {
		Id: uidutils.GenerateNewUniqueId(),
		Type: damageType,
		Damage: damage,
	}
}

func CreateHitboxComponent(width int, height int, isSolid bool) cmpt.Hitbox {
	return cmpt.Hitbox {
		Id: uidutils.GenerateNewUniqueId(),
		Width: width,
		Height: height,
		IsSolid: isSolid,
	}
}

func CreateEquipableComponent(equipmentType enum.EquipmentType) cmpt.Equipable {
	return cmpt.Equipable {
		Id: uidutils.GenerateNewUniqueId(),
		EquipmentType: equipmentType,
	}
}

func CreateEquipmentSlotComponent(equipmentType enum.EquipmentType) cmpt.EquipmentSlot {
	return cmpt.EquipmentSlot {
		Id: uidutils.GenerateNewUniqueId(),
		Type: equipmentType,
	}
}

func CreateAttachableComponent(attachmentType enum.AttachmentType) weapons.Attachable {
	return weapons.Attachable {
		Id: uidutils.GenerateNewUniqueId(),
		AttachmentType: attachmentType,
	}
}

func CreateAttachmentSlotComponent(attachmentType enum.AttachmentType) weapons.AttachmentSlot {
	return weapons.AttachmentSlot {
		Id: uidutils.GenerateNewUniqueId(),
		Type: attachmentType,
	}
}

func CreateExpirableComponent(expirationTime float64) cmpt.Expirable {
	return cmpt.Expirable {
		Id: uidutils.GenerateNewUniqueId(),
		TimeLeft:expirationTime,
	}
}

func CreateInventoryComponent(maxSize int) cmpt.Inventory {
	return cmpt.Inventory {
		Id: uidutils.GenerateNewUniqueId(),
		MaxSize: maxSize,
		ItemEntityIds: make([]string, maxSize),
	}
}

func CreateOwnableComponent() cmpt.Ownable {
	return cmpt.Ownable {
		Id: uidutils.GenerateNewUniqueId(),
	}
}

func CreateMovableComponent() cmpt.Movable {
	return cmpt.Movable {
		Id: uidutils.GenerateNewUniqueId(),
		Movable:true,
	}
}

func CreateUsableComponent(usageFunc closures.UsageFunc) cmpt.Usable {
	return cmpt.Usable {
		Id: uidutils.GenerateNewUniqueId(),
		Use:usageFunc,
	}
}

func CreateInteractiveComponent(interactionFunc closures.InteractionFunc) cmpt.Interactive {
	return cmpt.Interactive {
		Id: uidutils.GenerateNewUniqueId(),
		Interact:interactionFunc,
	}
}

func CreateCollidedComponent() cmpt.Collided {
	return cmpt.Collided {
		Id: uidutils.GenerateNewUniqueId(),
		CollidedEntities: make(map[string]bool, 0),
	}
}

func CreateTooltipComponent(text string) cmpt.Tooltip {
	return cmpt.Tooltip {
		Id: uidutils.GenerateNewUniqueId(),
		Text: text,
	}
}

func CreateRemoveComponent() cmpt.Remove {
	return cmpt.Remove {
		Id: uidutils.GenerateNewUniqueId(),
	}
}

func CreateHiddenComponent() cmpt.Hidden {
	return cmpt.Hidden {
		Id: uidutils.GenerateNewUniqueId(),
	}
}

//*********************** GUI **************************

func CreateGuiComponent(entityId string) guis.Gui {
	return guis.Gui {
		Id: uidutils.GenerateNewUniqueId(),
		EntityId:entityId,
	}
}

func CreateCharacterStatsGuiComponent() guis.CharacterStatsGui {
	return guis.CharacterStatsGui {
		Id: uidutils.GenerateNewUniqueId(),
	}
}

func CreateInventoryGuiComponent() guis.InventoryGui {
	return guis.InventoryGui {
		Id: uidutils.GenerateNewUniqueId(),
	}
}
