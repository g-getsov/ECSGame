package components

const removeComponentName = "Remove"

type Remove struct {
	ComponentName string
}

func (r Remove) GetComponentName() string { return GetRemoveComponentName() }

func (r Remove) IsUniquePerEntity() bool { return true }

func GetRemoveComponentName() string {return removeComponentName }