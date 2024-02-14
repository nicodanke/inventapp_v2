package validators

import "fmt"

func ValidString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n > maxLength || n < minLength {
		return fmt.Errorf("Must containen from %d to %d characters", minLength, maxLength)
	}
	return nil
}
