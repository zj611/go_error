package errors

import "fmt"

type AreaError struct {
	Err    string
	Radius float64
}

func (e *AreaError) Error() string{
	return fmt.Sprintf("radius %0.2f: %s",e.Radius,e.Err)
}

