package utils

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// CheckPasswordStrength checks how strong a password is.
// Returns: (ok, message)
func CheckPasswordStrength(password string) (bool, string) {
	// Minimum length
	if len(password) < 8 {
		return false, "Password must be at least 8 characters long"
	}

	// At least one uppercase
	if match, _ := regexp.MatchString(`[A-Z]`, password); !match {
		return false, "Password must contain at least one uppercase letter"
	}

	// At least one lowercase
	if match, _ := regexp.MatchString(`[a-z]`, password); !match {
		return false, "Password must contain at least one lowercase letter"
	}

	// At least one digit
	if match, _ := regexp.MatchString(`[0-9]`, password); !match {
		return false, "Password must contain at least one digit"
	}

	// At least one special character
	if match, _ := regexp.MatchString(`[!@#\$%\^&\*]`, password); !match {
		return false, "Password must contain at least one special character (!@#$%^&*)"
	}

	return true, "Password is strong"
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
