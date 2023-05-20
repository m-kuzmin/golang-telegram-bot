// Package e provides error handling utilities
package e

import "fmt"

// Wrap adds a context message to an error
// If the error is nil then returns nil
func Wrap(ctxMsg string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf(ctxMsg, err)
}
