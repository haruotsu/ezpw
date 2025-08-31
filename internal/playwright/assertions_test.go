package playwright

import (
	"testing"

	"github.com/haruotsu/ezpw/pkg/types"
)

func TestAssertTextContent(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	config := types.Config{
		Browser:  "chromium",
		Headless: true,
		Timeout:  30000,
	}

	browser, err := NewBrowser(config)
	if err != nil {
		t.Fatalf("Expected no error creating browser, got %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		t.Fatalf("Expected no error creating page, got %v", err)
	}

	// Set content with known text
	html := `<html><body><h1 id="title">Welcome</h1><p class="content">This is test content</p></body></html>`
	err = page.SetContent(html)
	if err != nil {
		t.Fatalf("Expected no error setting content, got %v", err)
	}

	// Test text content assertion
	assertion := NewAssertion(page)

	err = assertion.AssertTextContent("#title", "Welcome")
	if err != nil {
		t.Errorf("Expected no error for matching text content, got %v", err)
	}

	err = assertion.AssertTextContent(".content", "This is test content")
	if err != nil {
		t.Errorf("Expected no error for matching text content, got %v", err)
	}

	// Test failure case
	err = assertion.AssertTextContent("#title", "Wrong Text")
	if err == nil {
		t.Error("Expected error for non-matching text content, got nil")
	}
}

func TestAssertURL(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	config := types.Config{
		Browser:  "chromium",
		Headless: true,
		Timeout:  30000,
	}

	browser, err := NewBrowser(config)
	if err != nil {
		t.Fatalf("Expected no error creating browser, got %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		t.Fatalf("Expected no error creating page, got %v", err)
	}

	// Navigate to example.com
	err = page.Goto("https://example.com")
	if err != nil {
		t.Fatalf("Expected no error navigating, got %v", err)
	}

	assertion := NewAssertion(page)

	// Test URL contains assertion
	err = assertion.AssertURLContains("example.com")
	if err != nil {
		t.Errorf("Expected no error for URL contains assertion, got %v", err)
	}

	// Test failure case
	err = assertion.AssertURLContains("nonexistent.com")
	if err == nil {
		t.Error("Expected error for non-matching URL contains, got nil")
	}

	// Test exact URL assertion
	err = assertion.AssertURL("https://example.com/")
	if err != nil {
		t.Errorf("Expected no error for exact URL assertion, got %v", err)
	}

	// Test failure case for exact URL
	err = assertion.AssertURL("https://wrong.com/")
	if err == nil {
		t.Error("Expected error for non-matching exact URL, got nil")
	}
}

func TestAssertExists(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	config := types.Config{
		Browser:  "chromium",
		Headless: true,
		Timeout:  30000,
	}

	browser, err := NewBrowser(config)
	if err != nil {
		t.Fatalf("Expected no error creating browser, got %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		t.Fatalf("Expected no error creating page, got %v", err)
	}

	// Set content with known elements
	html := `<html><body><div id="existing">I exist</div><input class="form-field" type="text"></body></html>`
	err = page.SetContent(html)
	if err != nil {
		t.Fatalf("Expected no error setting content, got %v", err)
	}

	assertion := NewAssertion(page)

	// Test element exists
	err = assertion.AssertExists("#existing")
	if err != nil {
		t.Errorf("Expected no error for existing element, got %v", err)
	}

	err = assertion.AssertExists(".form-field")
	if err != nil {
		t.Errorf("Expected no error for existing element, got %v", err)
	}

	// Test element doesn't exist
	err = assertion.AssertExists("#nonexistent")
	if err == nil {
		t.Error("Expected error for non-existing element, got nil")
	}
}
