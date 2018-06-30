package components

import "BasicECS/closures"

const interactiveComponentName = "Interactive"

type Interactive struct {
	Id string
	Interact closures.InteractionFunc
}

func (i Interactive) GetComponentId() string { return i.Id }

func (i Interactive) GetComponentName() string { return GetInteractiveComponentName() }

func (i Interactive) IsUniquePerEntity() bool { return true }

func GetInteractiveComponentName() string { return interactiveComponentName }