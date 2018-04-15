package components

const positionComponentName = "Position"

type Position struct {
	X int
	Y int
}

func (p Position) GetComponentName() string { return positionComponentName }