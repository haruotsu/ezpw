package playwright

import "fmt"

// elementHelper provides element query methods for assertions
type elementHelper struct {
	page *playwrightPage
}

// newElementHelper creates a new element helper
func newElementHelper(page *playwrightPage) *elementHelper {
	return &elementHelper{page: page}
}

// GetElementCount returns the count of elements matching the selector
func (h *elementHelper) GetElementCount(selector string) (int, error) {
	locator := h.page.page.Locator(selector)
	count, err := locator.Count()
	if err != nil {
		return 0, fmt.Errorf("failed to get element count for %s: %w", selector, err)
	}
	return count, nil
}

// GetElementText returns the text content of an element
func (h *elementHelper) GetElementText(selector string) (string, error) {
	locator := h.page.page.Locator(selector)
	text, err := locator.TextContent()
	if err != nil {
		return "", fmt.Errorf("failed to get text content for %s: %w", selector, err)
	}
	return text, nil
}

// ElementExists checks if an element exists
func (h *elementHelper) ElementExists(selector string) (bool, error) {
	count, err := h.GetElementCount(selector)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}