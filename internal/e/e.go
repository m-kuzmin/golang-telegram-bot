package e

import "fmt"

// Wraps an error into a context message
// If the error is nil then returns nil
func Wrap(ctx_msg string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf(ctx_msg, err)
}
