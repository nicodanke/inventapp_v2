package userValidator

import (
	"fmt"
	"regexp"

	v "github.com/nicodanke/inventapp_v2/validators"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString
	isValidName     = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
	isValidLastname = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateUsername(value string) error {
	err := v.ValidString(value, 3, 100)
	if err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("username only can contain lower case letters, diggits or underscore")
	}
	return nil
}

func ValidateName(value string) error {
	err := v.ValidString(value, 3, 100)
	if err != nil {
		return err
	}
	if !isValidName(value) {
		return fmt.Errorf("name only can contain letters or spaces")
	}
	return nil
}

func ValidateLastname(value string) error {
	err := v.ValidString(value, 3, 100)
	if err != nil {
		return err
	}
	if !isValidLastname(value) {
		return fmt.Errorf("fullname only can contain letters or spaces")
	}
	return nil
}

func ValidatePassword(value string) error {
	return v.ValidString(value, 8, 256)
}
