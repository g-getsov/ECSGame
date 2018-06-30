package components

const hiddenComponentName = "Hidden"

type Hidden struct {
	Id string
}

func (h Hidden) GetComponentId() string { return h.Id }

func (h Hidden) GetComponentName() string { return GetHiddenComponentName() }

func (h Hidden) IsUniquePerEntity() bool { return true }

func GetHiddenComponentName() string {return hiddenComponentName }
