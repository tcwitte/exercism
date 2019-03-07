package erratum

import "fmt"

const testVersion = 2

// Use uses the resource to open a resource and call it with the given string.
func Use(o ResourceOpener, input string) (returnError error) {
	resource, err := o()
	if _, ok := err.(TransientError); ok {
		// TODO: add some sleep
		// Try again
		return Use(o, input)
	} else if err != nil {
		return err // Error is not wrapped!
	}

	defer resource.Close()
	defer func() {
		err := recover()
		if frobError, ok := err.(FrobError); ok {
			resource.Defrob(frobError.defrobTag)
		}

		if err != nil {
			returnError = fmt.Errorf("%v", err)
		}
	}()

	resource.Frob(input)
	return
}
