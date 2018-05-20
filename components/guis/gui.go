package guis

const guiComponentName = "Gui"

type Gui struct {
	EntityId string
}

func (c Gui) GetComponentName() string { return GetGuiComponentName() }

func GetGuiComponentName() string {return guiComponentName }