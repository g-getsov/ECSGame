package components

const positionComponentName = "Position"

type Position struct {
	X int
	Y int
}

func (p Position) GetComponentName() string { return GetPositionComponentName() }

func (p Position) IsUniquePerEntity() bool { return true }

func GetPositionComponentName() string {return positionComponentName }