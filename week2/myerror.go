package myerror

import (
	"fmt"
)

type ErrNotFound struct {
	Name string
	Type string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("%s '%s' not found", e.Type, e.Name)
}

func NewErrNotFound(t string, name string) ErrNotFound {
	return ErrNotFound{
		Name: name,
		Type: t,
	}
}
