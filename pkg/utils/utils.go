package utils

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// FormatTime formats a time.Time to a specific layout
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// ParseTime parses a string to time.Time
func ParseTime(s string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", s)
}

// GenerateUUID generates a new UUID
func GenerateUUID() string {
	return uuid.New().String()
}

// StructToJSON converts a struct to JSON string
func StructToJSON(v interface{}) (string, error) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return "", fmt.Errorf("failed to marshal struct to JSON: %w", err)
	}
	return string(jsonBytes), nil
}

// JSONToStruct converts a JSON string to a struct
func JSONToStruct(jsonStr string, v interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON to struct: %w", err)
	}
	return nil
}

// ValidateEmail checks if an email is valid
func ValidateEmail(email string) bool {
	// This is a very simple validation, you might want to use a more robust method
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// TruncateString truncates a string to a specified length
func TruncateString(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}
	return s[:maxLength] + "..."
}
