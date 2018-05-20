package guis

const guiComponentName = "Gui"

type Gui struct {
	EntityId string
}

func (g Gui) GetComponentName() string { return GetGuiComponentName() }

func (g Gui) IsUniquePerEntity() bool { return true }

func GetGuiComponentName() string {return guiComponentName }