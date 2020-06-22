package Errors

import "fmt"

type NotFound struct {
	Reccord string
	Inner   error
}

func (n NotFound) Error() string {
	return fmt.Sprintf("Error: Record %s not found, reason:%s", n.Reccord, n.Inner)
}
