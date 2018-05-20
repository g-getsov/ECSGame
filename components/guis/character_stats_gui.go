package guis

const characterStatsGuiComponentName = "CharacterStatsGui"

type CharacterStatsGui struct {}

func (c CharacterStatsGui) GetComponentName() string { return GetCharacterStatsGuiComponentName() }

func GetCharacterStatsGuiComponentName() string { return characterStatsGuiComponentName }