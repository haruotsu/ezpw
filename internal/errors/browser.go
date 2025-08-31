package errors

import (
	"fmt"
	"strings"
)

// BrowserNotFoundError represents an error when browser is not installed
type BrowserNotFoundError struct {
	Cause   error
	Browser string
}

func (e *BrowserNotFoundError) Error() string {
	return fmt.Sprintf("browser '%s' is not installed: %v", e.Browser, e.Cause)
}

func (e *BrowserNotFoundError) Unwrap() error {
	return e.Cause
}

// IsBrowserNotFoundError checks if the error is due to missing browser installation
func IsBrowserNotFoundError(err error) bool {
	errStr := strings.ToLower(err.Error())
	return strings.Contains(errStr, "executable doesn't exist") ||
		strings.Contains(errStr, "browser not found") ||
		strings.Contains(errStr, "install")
}
