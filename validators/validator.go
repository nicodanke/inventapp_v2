package validators

import (
	"fmt"
	"net/mail"
)

func ValidString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n > maxLength || n < minLength {
		return fmt.Errorf("Must containen from %d to %d characters", minLength, maxLength)
	}
	return nil
}

func ValidateEmail(value string) error {
	err := ValidString(value, 3, 200)
	if err != nil {
		return err
	}
	_, err = mail.ParseAddress(value)
	if err != nil {
		return fmt.Errorf("%s is not a valid email address", value)
	}
	return nil
}
