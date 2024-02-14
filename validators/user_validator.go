package validators

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString
	isValidName     = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
	isValidLastname = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateUsername(value string) error {
	err := ValidString(value, 3, 100)
	if err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("Username only can contain lower case letters, diggits or underscore")
	}
	return nil
}

func ValidateName(value string) error {
	err := ValidString(value, 3, 100)
	if err != nil {
		return err
	}
	if !isValidName(value) {
		return fmt.Errorf("Name only can contain letters or spaces")
	}
	return nil
}

func ValidateLastname(value string) error {
	err := ValidString(value, 3, 100)
	if err != nil {
		return err
	}
	if !isValidLastname(value) {
		return fmt.Errorf("Fullname only can contain letters or spaces")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidString(value, 8, 256)
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
