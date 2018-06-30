package guis

const characterStatsGuiComponentName = "CharacterStatsGui"

type CharacterStatsGui struct {
	Id string
}

func (c CharacterStatsGui) GetComponentId() string { return c.Id }

func (c CharacterStatsGui) GetComponentName() string { return GetCharacterStatsGuiComponentName() }

func (c CharacterStatsGui) IsUniquePerEntity() bool { return true }

func GetCharacterStatsGuiComponentName() string { return characterStatsGuiComponentName }
