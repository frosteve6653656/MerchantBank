package utils

import "regexp"

func ValidateEmail(email string) bool {
	// Regular expression pattern for email validation
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regex pattern
	regex := regexp.MustCompile(pattern)

	// Match the email against the regex pattern
	return regex.MatchString(email)
}
