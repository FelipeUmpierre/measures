package command

type (
	MeasureCommand struct {
		Neck       float64 `json:"neck"`
		UpperChest float64 `json:"upper_chest"`
		Chest      float64 `json:"chest"`
		UnderChest float64 `json:"under_chest"`
		Waist      float64 `json:"waist"`
		Abs        float64 `json:"abs"`
		Hip        float64 `json:"hip"`
		Thigh      float64 `json:"thigh"`
	}
)

func (m MeasureCommand) CommandType() string {
	return `measure`
}
