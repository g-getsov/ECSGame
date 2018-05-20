package components

import "BasicECS/closures"

const interactiveComponentName = "Interactive"

type Interactive struct {
	Interact closures.InteractionFunc
}

func (i Interactive) GetComponentName() string { return GetInteractiveComponentName() }

func (i Interactive) IsUniquePerEntity() bool { return true }

func GetInteractiveComponentName() string { return interactiveComponentName }