package guis

const characterStatsGuiComponentName = "CharacterStatsGui"

type CharacterStatsGui struct {}

func (c CharacterStatsGui) GetComponentName() string { return GetCharacterStatsGuiComponentName() }

func (c CharacterStatsGui) IsUniquePerEntity() bool { return true }

func GetCharacterStatsGuiComponentName() string { return characterStatsGuiComponentName }
