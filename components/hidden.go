package components

const hiddenComponentName = "Hidden"

type Hidden struct {}

func (h Hidden) GetComponentName() string { return GetHiddenComponentName() }

func GetHiddenComponentName() string {return hiddenComponentName }
