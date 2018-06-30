package components

const removeComponentName = "Remove"

type Remove struct {
	Id string
	ComponentName string
}

func (r Remove) GetComponentId() string { return r.Id }

func (r Remove) GetComponentName() string { return GetRemoveComponentName() }

func (r Remove) IsUniquePerEntity() bool { return true }

func GetRemoveComponentName() string {return removeComponentName }