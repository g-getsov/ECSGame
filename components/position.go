package components

const positionComponentName = "Position"

type Position struct {
	X int
	Y int
}

func (p Position) GetComponentName() string { return GetPositionComponentName() }

func GetPositionComponentName() string {return positionComponentName }