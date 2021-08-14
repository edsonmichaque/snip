package errors

type CommandError struct {
	Code int
	Err  error
}

func (c *CommandError) Error() string {
	return c.Err.Error()
}
