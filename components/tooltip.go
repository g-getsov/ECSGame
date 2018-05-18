package components

const tooltipComponentName = "Tooltip"

type Tooltip struct {
	Text string
}

func (s Tooltip) GetComponentName() string { return GetTooltipComponentName() }

func GetTooltipComponentName() string {return tooltipComponentName }