package components

const hiddenComponentName = "Hidden"

type Hidden struct {}

func (h Hidden) GetComponentName() string { return GetHiddenComponentName() }

func (h Hidden) IsUniquePerEntity() bool { return true }

func GetHiddenComponentName() string {return hiddenComponentName }
