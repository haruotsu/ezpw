package playwright

import (
	"fmt"

	"github.com/haruotsu/ezpw/internal/errors"
	"github.com/haruotsu/ezpw/pkg/types"
	"github.com/playwright-community/playwright-go"
)

// Browser wraps playwright browser functionality
type Browser struct {
	pw      *playwright.Playwright
	browser playwright.Browser
	config  types.Config
}

// Page wraps playwright page functionality
type Page struct {
	page playwright.Page
}

// NewBrowser creates a new browser instance
func NewBrowser(config types.Config) (*Browser, error) {
	pw, err := playwright.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run playwright: %w", err)
	}

	// whether you want to run the browser in headed mode
	var browser playwright.Browser
	browserOptions := playwright.BrowserTypeLaunchOptions{
		Headless: &config.Headless,
	}

	switch config.Browser {
	case "firefox":
		browser, err = pw.Firefox.Launch(browserOptions)
	case "webkit":
		browser, err = pw.WebKit.Launch(browserOptions)
	default: // chromium
		browser, err = pw.Chromium.Launch(browserOptions)
	}

	if err != nil {
		pw.Stop()
		// Check if it's a browser installation error
		if errors.IsBrowserNotFoundError(err) {
			return nil, &errors.BrowserNotFoundError{
				Browser: config.Browser,
				Cause:   err,
			}
		}
		return nil, fmt.Errorf("failed to launch browser: %w", err)
	}

	return &Browser{
		pw:      pw,
		browser: browser,
		config:  config,
	}, nil
}

// NewPage creates a new page
func (b *Browser) NewPage() (*Page, error) {
	page, err := b.browser.NewPage()
	if err != nil {
		return nil, fmt.Errorf("failed to create new page: %w", err)
	}

	return &Page{page: page}, nil
}

// Close closes the browser and cleans up resources
func (b *Browser) Close() error {
	if b.browser != nil {
		if err := b.browser.Close(); err != nil {
			return fmt.Errorf("failed to close browser: %w", err)
		}
	}
	if b.pw != nil {
		if err := b.pw.Stop(); err != nil {
			return fmt.Errorf("failed to stop playwright: %w", err)
		}
	}
	return nil
}

// NavigateToURL navigates to the specified URL
func (p *Page) NavigateToURL(url string) error {
	_, err := p.page.Goto(url)
	if err != nil {
		return fmt.Errorf("failed to navigate to %s: %w", url, err)
	}
	return nil
}

// Goto is an alias for NavigateToURL to maintain backward compatibility
func (p *Page) Goto(url string) error {
	return p.NavigateToURL(url)
}

// ClickElement clicks on an element identified by selector using locator-based API
func (p *Page) ClickElement(selector string) error {
	locator := p.page.Locator(selector)
	err := locator.Click()
	if err != nil {
		return fmt.Errorf("failed to click element %s: %w", selector, err)
	}
	return nil
}

// Click is an alias for ClickElement to maintain backward compatibility
func (p *Page) Click(selector string) error {
	return p.ClickElement(selector)
}

// FillElement fills an input element with the given value using locator-based API
func (p *Page) FillElement(selector string, value string) error {
	locator := p.page.Locator(selector)
	err := locator.Fill(value)
	if err != nil {
		return fmt.Errorf("failed to fill element %s with value %s: %w", selector, value, err)
	}
	return nil
}

// Fill is an alias for FillElement to maintain backward compatibility
func (p *Page) Fill(selector string, value string) error {
	return p.FillElement(selector, value)
}

// URL returns the current URL of the page
func (p *Page) URL() string {
	return p.page.URL()
}

// SetContent sets the HTML content of the page
func (p *Page) SetContent(html string) error {
	err := p.page.SetContent(html)
	if err != nil {
		return fmt.Errorf("failed to set content: %w", err)
	}
	return nil
}

// GetElementValue gets the value of an input element using locator-based API
func (p *Page) GetElementValue(selector string) (string, error) {
	locator := p.page.Locator(selector)
	value, err := locator.InputValue()
	if err != nil {
		return "", fmt.Errorf("failed to get input value for %s: %w", selector, err)
	}
	return value, nil
}

// InputValue is an alias for GetElementValue to maintain backward compatibility
func (p *Page) InputValue(selector string) (string, error) {
	return p.GetElementValue(selector)
}
