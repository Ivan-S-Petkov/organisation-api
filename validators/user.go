package validators

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

type ValidationError struct {
    Field   string `json:"field"`
    Message string `json:"message"`
}

func IsValidEmail(email string) bool {
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return emailRegex.MatchString(email) && len(email) <= 255
}

func IsValidName(name string) bool {
    if utf8.RuneCountInString(name) < 2 || utf8.RuneCountInString(name) > 100 {
        return false
    }
    
   nameRegex := regexp.MustCompile(`^[a-zA-ZÀ-ÿ\s'-]+$`)
    return nameRegex.MatchString(strings.TrimSpace(name))
}

func IsValidRole(role string) bool {
    return role == "admin" || role == "user"
}

func NormalizeEmail(email string) string {
    return strings.ToLower(strings.TrimSpace(email))
}

func ValidateCreateUser(name, email, role string) []ValidationError {
    var errors []ValidationError

    if strings.TrimSpace(name) == "" {
        errors = append(errors, ValidationError{
            Field:   "name",
            Message: "Name is required",
        })
    } else if !IsValidName(name) {
        errors = append(errors, ValidationError{
            Field:   "name",
            Message: "Name must be 2-100 characters and contain only letters, spaces, hyphens, and apostrophes",
        })
    }

    normalizedEmail := NormalizeEmail(email)
    if !IsValidEmail(normalizedEmail) {
        errors = append(errors, ValidationError{
            Field:   "email",
            Message: "Invalid email format or too long (max 255 characters)",
        })
    }

    if !IsValidRole(role) {
        errors = append(errors, ValidationError{
            Field:   "role",
            Message: "Role must be either 'admin' or 'user'",
        })
    }
    
    return errors
}