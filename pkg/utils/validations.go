package utils

import (
	"errors"
	"regexp"
)

func ValidatePhilippinePhoneNumber(phoneNumber string) error {
	var validPhoneNumber = regexp.MustCompile(`^(09|\+639)\d{9}$`)
	if !validPhoneNumber.MatchString(phoneNumber) {
		return errors.New("Invalid Philippine phone number format")
	}
	return nil
}
