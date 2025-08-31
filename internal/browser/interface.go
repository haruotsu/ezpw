package browser

// Browser represents a browser instance interface
type Browser interface {
	// NewPage creates a new page/tab in the browser
	NewPage() (Page, error)
	// Close closes the browser and cleans up resources
	Close() error
}

// Page represents a browser page interface
type Page interface {
	// Navigation
	NavigateToURL(url string) error
	Goto(url string) error // Alias for backward compatibility

	// Interactions
	ClickElement(selector string) error
	Click(selector string) error // Alias for backward compatibility
	FillElement(selector string, value string) error
	Fill(selector string, value string) error // Alias for backward compatibility

	// Getters
	URL() string
	GetElementValue(selector string) (string, error)
	InputValue(selector string) (string, error) // Alias for backward compatibility

	// Content manipulation
	SetContent(html string) error

	// Locator operations (for assertions)
	GetElementCount(selector string) (int, error)
	GetElementText(selector string) (string, error)
	ElementExists(selector string) (bool, error)
}
