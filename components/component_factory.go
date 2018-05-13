package components

import (
	"github.com/hajimehoshi/ebiten"
	"BasicECS/components/weapons"
	"BasicECS/enum"
)

func CreateSpeedComponent(speed int) Speed {
	return Speed{
		Speed: speed,
	}
}

func CreateVelocityComponent(x int, y int) Velocity {
	return Velocity{
		X: x,
		Y: y,
	}
}

func CreatePositionComponent(x int, y int) Position {
	return Position{
		X: x,
		Y: y,
	}
}

func CreateHealthComponent(health int) Health {
	return Health{
		Health: health,
	}
}

func CreateInputComponent() Input {
	return Input{}
}

func CreateSpriteComponent(image *ebiten.Image) Sprite {
	return Sprite{
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

func CreateDamageComponent(damageType DamageType, damage int) Damage {
	return Damage {
		Type: damageType,
		Damage: damage,
	}
}

func CreateColidableComponent() Collidable {
	return Collidable{}
}

func CreateEquipableComponent(equipmentType enum.EquipmentType) Equippable {
	return Equippable{
		EquipmentType: equipmentType,
	}
}

func CreateEquipmentSlotComponent(equipmentType enum.EquipmentType) EquipmentSlot {
	return EquipmentSlot{
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

func CreateExpirableComponent() Expirable {
	return Expirable{}
}

func CreateInventoryComponent() Inventory {
	return Inventory{}
}

func CreateOwnableComponent() Ownable {
	return Ownable{}
}

func CreateMovableComponent() Movable {
	return Movable{ true }
}