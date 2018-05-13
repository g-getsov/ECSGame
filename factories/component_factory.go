package factories

import (
	"github.com/hajimehoshi/ebiten"
	"BasicECS/components/weapons"
	"BasicECS/enum"
	cmpt "BasicECS/components"
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
	}
}

func CreateInputComponent() cmpt.Input {
	return cmpt.Input{}
}

func CreateSpriteComponent(image *ebiten.Image) cmpt.Sprite {
	return cmpt.Sprite{
		Image: image,
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

func CreateColidableComponent() cmpt.Collidable {
	return cmpt.Collidable{}
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

func CreateExpirableComponent() cmpt.Expirable {
	return cmpt.Expirable{}
}

func CreateInventoryComponent() cmpt.Inventory {
	return cmpt.Inventory{}
}

func CreateOwnableComponent() cmpt.Ownable {
	return cmpt.Ownable{}
}

func CreateMovableComponent() cmpt.Movable {
	return cmpt.Movable{ true }
}