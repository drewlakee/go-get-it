package cli

type Command struct {
	Flag         string
	DefaultValue interface{}
	Description  string
	InputValue   interface{}
}

func NewCommand(flag string, defaultValue interface{}, description string) *Command {
	return &Command{flag, defaultValue, description, nil}
}
