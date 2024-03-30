package validator

import "fmt"

func ValidateArgs(args []string) error {
	if len(args) > 1 {		
		return fmt.Errorf("only accepts a single argument")
	}
	if len(args) == 0 {		
		return fmt.Errorf("command requires input value")
	}
	return nil
}