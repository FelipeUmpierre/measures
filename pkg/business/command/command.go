package command

type (
	Command interface {
		CommandType() string
	}
)
