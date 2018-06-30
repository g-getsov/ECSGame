package components

const positionComponentName = "Position"

type Position struct {
	Id string
	X int
	Y int
}

func (p Position) GetComponentId() string { return p.Id }

func (p Position) GetComponentName() string { return GetPositionComponentName() }

func (p Position) IsUniquePerEntity() bool { return true }

func GetPositionComponentName() string {return positionComponentName }