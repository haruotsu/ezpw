package playwright

import (
	"fmt"

	"github.com/haruotsu/ezpw/internal/browser"
	"github.com/haruotsu/ezpw/internal/errors"
	"github.com/haruotsu/ezpw/pkg/types"
	"github.com/playwright-community/playwright-go"
)

// playwrightBrowser implements browser.Browser interface using Playwright
type playwrightBrowser struct {
	pw      *playwright.Playwright
	browser playwright.Browser
	config  types.Config
}

// playwrightPage implements browser.Page interface using Playwright
type playwrightPage struct {
	page playwright.Page
}

// NewBrowser creates a new browser instance that implements browser.Browser interface
func NewBrowser(config types.Config) (browser.Browser, error) {
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

	return &playwrightBrowser{
		pw:      pw,
		browser: browser,
		config:  config,
	}, nil
}

// NewPage creates a new page
func (b *playwrightBrowser) NewPage() (browser.Page, error) {
	page, err := b.browser.NewPage()
	if err != nil {
		return nil, fmt.Errorf("failed to create new page: %w", err)
	}

	return &playwrightPage{page: page}, nil
}

// Close closes the browser and cleans up resources
func (b *playwrightBrowser) Close() error {
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
func (p *playwrightPage) NavigateToURL(url string) error {
	_, err := p.page.Goto(url)
	if err != nil {
		return fmt.Errorf("failed to navigate to %s: %w", url, err)
	}
	return nil
}

// Goto is an alias for NavigateToURL to maintain backward compatibility
func (p *playwrightPage) Goto(url string) error {
	return p.NavigateToURL(url)
}

// ClickElement clicks on an element identified by selector using locator-based API
func (p *playwrightPage) ClickElement(selector string) error {
	locator := p.page.Locator(selector)
	err := locator.Click()
	if err != nil {
		return fmt.Errorf("failed to click element %s: %w", selector, err)
	}
	return nil
}

// Click is an alias for ClickElement to maintain backward compatibility
func (p *playwrightPage) Click(selector string) error {
	return p.ClickElement(selector)
}

// FillElement fills an input element with the given value using locator-based API
func (p *playwrightPage) FillElement(selector string, value string) error {
	locator := p.page.Locator(selector)
	err := locator.Fill(value)
	if err != nil {
		return fmt.Errorf("failed to fill element %s with value %s: %w", selector, value, err)
	}
	return nil
}

// Fill is an alias for FillElement to maintain backward compatibility
func (p *playwrightPage) Fill(selector string, value string) error {
	return p.FillElement(selector, value)
}

// URL returns the current URL of the page
func (p *playwrightPage) URL() string {
	return p.page.URL()
}

// SetContent sets the HTML content of the page
func (p *playwrightPage) SetContent(html string) error {
	err := p.page.SetContent(html)
	if err != nil {
		return fmt.Errorf("failed to set content: %w", err)
	}
	return nil
}

// GetElementValue gets the value of an input element using locator-based API
func (p *playwrightPage) GetElementValue(selector string) (string, error) {
	locator := p.page.Locator(selector)
	value, err := locator.InputValue()
	if err != nil {
		return "", fmt.Errorf("failed to get input value for %s: %w", selector, err)
	}
	return value, nil
}

// InputValue is an alias for GetElementValue to maintain backward compatibility
func (p *playwrightPage) InputValue(selector string) (string, error) {
	return p.GetElementValue(selector)
}

// GetElementCount returns the count of elements matching the selector
// This method is primarily used for assertions
func (p *playwrightPage) GetElementCount(selector string) (int, error) {
	helper := newElementHelper(p)
	return helper.GetElementCount(selector)
}

// GetElementText returns the text content of an element
// This method is primarily used for assertions
func (p *playwrightPage) GetElementText(selector string) (string, error) {
	helper := newElementHelper(p)
	return helper.GetElementText(selector)
}

// ElementExists checks if an element exists
// This method is primarily used for assertions
func (p *playwrightPage) ElementExists(selector string) (bool, error) {
	helper := newElementHelper(p)
	return helper.ElementExists(selector)
}
