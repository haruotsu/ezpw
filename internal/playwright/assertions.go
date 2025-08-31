package playwright

import (
	"fmt"
	"strings"

	"github.com/haruotsu/ezpw/internal/browser"
)

// Assertion provides assertion methods for a page
type Assertion struct {
	page browser.Page
}

// NewAssertion creates a new assertion instance
func NewAssertion(page browser.Page) *Assertion {
	return &Assertion{page: page}
}

// AssertTextContent asserts that an element contains the expected text content
func (a *Assertion) AssertTextContent(selector, expectedText string) error {
	// Check if element exists first
	exists, err := a.page.ElementExists(selector)
	if err != nil {
		return fmt.Errorf("failed to check element existence for selector %s: %w", selector, err)
	}

	if !exists {
		return fmt.Errorf("element with selector %s not found", selector)
	}

	// Get text content
	actualText, err := a.page.GetElementText(selector)
	if err != nil {
		return fmt.Errorf("failed to get text content for selector %s: %w", selector, err)
	}

	if actualText != expectedText {
		return fmt.Errorf("text content mismatch for selector %s: expected '%s', got '%s'",
			selector, expectedText, actualText)
	}

	return nil
}

// AssertURLContains asserts that the current URL contains the expected substring
func (a *Assertion) AssertURLContains(expectedSubstring string) error {
	currentURL := a.page.URL()

	if !strings.Contains(currentURL, expectedSubstring) {
		return fmt.Errorf("URL does not contain expected substring: expected URL to contain '%s', got '%s'",
			expectedSubstring, currentURL)
	}

	return nil
}

// AssertURL asserts that the current URL matches the expected URL exactly
func (a *Assertion) AssertURL(expectedURL string) error {
	currentURL := a.page.URL()

	if currentURL != expectedURL {
		return fmt.Errorf("URL mismatch: expected '%s', got '%s'", expectedURL, currentURL)
	}

	return nil
}

// AssertExists asserts that an element with the given selector exists
func (a *Assertion) AssertExists(selector string) error {
	exists, err := a.page.ElementExists(selector)
	if err != nil {
		return fmt.Errorf("failed to check element existence for selector %s: %w", selector, err)
	}

	if !exists {
		return fmt.Errorf("element with selector %s does not exist", selector)
	}

	return nil
}
