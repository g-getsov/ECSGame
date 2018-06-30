package guis

const guiComponentName = "Gui"

type Gui struct {
	Id string
	EntityId string
}

func (g Gui) GetComponentId() string { return g.Id }

func (g Gui) GetComponentName() string { return GetGuiComponentName() }

func (g Gui) IsUniquePerEntity() bool { return true }

func GetGuiComponentName() string {return guiComponentName }