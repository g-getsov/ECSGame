package components

import "BasicECS/closures"

const interactiveComponentName = "Interactive"

type Interactive struct {
	Interact closures.InteractionFunc
}

func (e Interactive) GetComponentName() string { return GetInteractiveComponentName() }

func GetInteractiveComponentName() string { return interactiveComponentName }