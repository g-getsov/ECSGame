package components

const inputComponentName = "Input"

type Input struct {
	UpArrowKey bool
	DownArrowKey bool
	LeftArrowKey bool
	RightArrowKey bool
	WKey bool
	SKey bool
	AKey bool
	DKey bool
	SpaceKey bool
	EscapeKey bool
	LeftMouseButton bool
	RightMouseButton bool
	MiddleMouseButton bool
}

func (i Input) GetComponentName() string { return inputComponentName }