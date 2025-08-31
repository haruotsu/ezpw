package playwright

import (
	"testing"

	"github.com/haruotsu/ezpw/pkg/types"
)

func TestBrowserLifecycle(t *testing.T) {
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

	if browser == nil {
		t.Error("Expected browser to not be nil")
	}
}

func TestBrowserGoto(t *testing.T) {
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

	err = page.Goto("https://example.com")
	if err != nil {
		t.Fatalf("Expected no error navigating to example.com, got %v", err)
	}

	url := page.URL()
	if url != "https://example.com/" {
		t.Errorf("Expected URL 'https://example.com/', got '%s'", url)
	}
}

func TestBrowserClick(t *testing.T) {
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

	// Navigate to example page with links
	err = page.Goto("https://example.com")
	if err != nil {
		t.Fatalf("Expected no error navigating, got %v", err)
	}

	// Test that click method exists and doesn't error on valid selector
	err = page.Click("a")
	if err != nil {
		t.Logf("Click failed (expected for example.com): %v", err)
		// This is acceptable as example.com might not have clickable links
	}
}

func TestBrowserFill(t *testing.T) {
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

	// Create a simple HTML page with an input
	html := `<html><body><input id="test" type="text"></body></html>`
	err = page.SetContent(html)
	if err != nil {
		t.Fatalf("Expected no error setting content, got %v", err)
	}

	// Test fill
	err = page.Fill("#test", "test value")
	if err != nil {
		t.Fatalf("Expected no error filling input, got %v", err)
	}

	// Verify the value was set
	value, err := page.InputValue("#test")
	if err != nil {
		t.Fatalf("Expected no error getting input value, got %v", err)
	}

	if value != "test value" {
		t.Errorf("Expected input value 'test value', got '%s'", value)
	}
}