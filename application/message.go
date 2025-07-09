package application

import "fmt"

func RecoveredFromPanic(r any) error {
	return fmt.Errorf("recovered from panic: %v", r)
}
