package validators

import (
	"fmt"
	"regexp"
)

type AccountValidator struct{}

var (
	isValidCompanyName = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`).MatchString
	isValidCountry     = regexp.MustCompile(`^ARG$`).MatchString
)

func ValidateCompanyName(value string) error {
	err := ValidString(value, 3, 100)
	if err != nil {
		return err
	}
	if !isValidCompanyName(value) {
		return fmt.Errorf("Company name only can contain letters, diggits and spaces")
	}
	return nil
}

func ValidateCountry(value string) error {
	if !isValidCountry(value) {
		return fmt.Errorf("Invalid country")
	}
	return nil
}
