package cli

import "fmt"

type InvalidCommandError struct {
	Name string
	ctx *Context
}

func NewInvalidCommandError(name string, ctx *Context) error {
	return &InvalidCommandError {
		Name: name,
		ctx: ctx,
	}
}

func (e *InvalidCommandError) Error() string {
	return fmt.Sprintf("'%s' is not a vaild command", e.Name)
}

func (e *InvalidCommandError) GetSuggestions() []string {
	cmd := e.ctx.command
	return cmd.GetSuggestions(e.Name)
}

type InvalidFlagError struct {
	Name string
	ctx *Context
}

func NewInvalidFlagError(name string, ctx *Context) error {
	return &InvalidFlagError{
		Name: name,
		ctx: ctx,
	}
}

func (e *InvalidFlagError) Error() string {
	return fmt.Sprintf("invalid flag --%s", e.Name)
}

func (e *InvalidFlagError) GetSuggestions() []string {
	distance := e.ctx.command.GetSuggestDistance()
	return e.ctx.Flags().GetSuggestions(e.Name, distance)
}