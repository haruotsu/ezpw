package playwright

import (
	"testing"

	"github.com/haruotsu/ezpw/pkg/types"
)

func TestGetElementCount(t *testing.T) {
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
		t.Fatalf("Failed to create browser: %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		t.Fatalf("Failed to create page: %v", err)
	}

	// Set up test HTML with multiple elements
	html := `
		<html><body>
			<div class="item">Item 1</div>
			<div class="item">Item 2</div>
			<div class="item">Item 3</div>
			<div id="unique">Unique element</div>
		</body></html>
	`
	
	playwrightPage, ok := page.(*playwrightPage)
	if !ok {
		t.Fatal("Failed to cast to playwrightPage")
	}

	err = playwrightPage.SetContent(html)
	if err != nil {
		t.Fatalf("Failed to set content: %v", err)
	}

	// Test counting multiple elements
	count, err := playwrightPage.GetElementCount(".item")
	if err != nil {
		t.Fatalf("Failed to get element count: %v", err)
	}
	if count != 3 {
		t.Errorf("Expected 3 elements with class 'item', got %d", count)
	}

	// Test counting single element
	count, err = playwrightPage.GetElementCount("#unique")
	if err != nil {
		t.Fatalf("Failed to get element count: %v", err)
	}
	if count != 1 {
		t.Errorf("Expected 1 element with id 'unique', got %d", count)
	}

	// Test counting non-existent element
	count, err = playwrightPage.GetElementCount("#nonexistent")
	if err != nil {
		t.Fatalf("Failed to get element count: %v", err)
	}
	if count != 0 {
		t.Errorf("Expected 0 elements with id 'nonexistent', got %d", count)
	}
}

func TestGetElementText(t *testing.T) {
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
		t.Fatalf("Failed to create browser: %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		t.Fatalf("Failed to create page: %v", err)
	}

	html := `
		<html><body>
			<h1 id="title">Test Title</h1>
			<p class="content">This is test content</p>
			<div id="empty"></div>
		</body></html>
	`
	
	playwrightPage, ok := page.(*playwrightPage)
	if !ok {
		t.Fatal("Failed to cast to playwrightPage")
	}

	err = playwrightPage.SetContent(html)
	if err != nil {
		t.Fatalf("Failed to set content: %v", err)
	}

	// Test getting text from heading
	text, err := playwrightPage.GetElementText("#title")
	if err != nil {
		t.Fatalf("Failed to get element text: %v", err)
	}
	if text != "Test Title" {
		t.Errorf("Expected text 'Test Title', got '%s'", text)
	}

	// Test getting text from paragraph
	text, err = playwrightPage.GetElementText(".content")
	if err != nil {
		t.Fatalf("Failed to get element text: %v", err)
	}
	if text != "This is test content" {
		t.Errorf("Expected text 'This is test content', got '%s'", text)
	}

	// Test empty element
	text, err = playwrightPage.GetElementText("#empty")
	if err != nil {
		t.Fatalf("Failed to get element text: %v", err)
	}
	if text != "" {
		t.Errorf("Expected empty text, got '%s'", text)
	}
}

func TestElementExists(t *testing.T) {
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
		t.Fatalf("Failed to create browser: %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		t.Fatalf("Failed to create page: %v", err)
	}

	html := `
		<html><body>
			<div id="existing">I exist</div>
			<input class="form-field" type="text">
		</body></html>
	`
	
	playwrightPage, ok := page.(*playwrightPage)
	if !ok {
		t.Fatal("Failed to cast to playwrightPage")
	}

	err = playwrightPage.SetContent(html)
	if err != nil {
		t.Fatalf("Failed to set content: %v", err)
	}

	// Test existing element by ID
	exists, err := playwrightPage.ElementExists("#existing")
	if err != nil {
		t.Fatalf("Failed to check element existence: %v", err)
	}
	if !exists {
		t.Error("Expected element #existing to exist")
	}

	// Test existing element by class
	exists, err = playwrightPage.ElementExists(".form-field")
	if err != nil {
		t.Fatalf("Failed to check element existence: %v", err)
	}
	if !exists {
		t.Error("Expected element .form-field to exist")
	}

	// Test non-existent element
	exists, err = playwrightPage.ElementExists("#nonexistent")
	if err != nil {
		t.Fatalf("Failed to check element existence: %v", err)
	}
	if exists {
		t.Error("Expected element #nonexistent to not exist")
	}
}

func TestGetElementValue(t *testing.T) {
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
		t.Fatalf("Failed to create browser: %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		t.Fatalf("Failed to create page: %v", err)
	}

	html := `
		<html><body>
			<input id="text-input" type="text" value="initial value">
			<textarea id="textarea">Text area content</textarea>
			<input id="empty-input" type="text">
		</body></html>
	`
	
	playwrightPage, ok := page.(*playwrightPage)
	if !ok {
		t.Fatal("Failed to cast to playwrightPage")
	}

	err = playwrightPage.SetContent(html)
	if err != nil {
		t.Fatalf("Failed to set content: %v", err)
	}

	// Test getting value from input with initial value
	value, err := playwrightPage.GetElementValue("#text-input")
	if err != nil {
		t.Fatalf("Failed to get element value: %v", err)
	}
	if value != "initial value" {
		t.Errorf("Expected value 'initial value', got '%s'", value)
	}

	// Test getting value from textarea
	value, err = playwrightPage.GetElementValue("#textarea")
	if err != nil {
		t.Fatalf("Failed to get element value: %v", err)
	}
	if value != "Text area content" {
		t.Errorf("Expected value 'Text area content', got '%s'", value)
	}

	// Test empty input
	value, err = playwrightPage.GetElementValue("#empty-input")
	if err != nil {
		t.Fatalf("Failed to get element value: %v", err)
	}
	if value != "" {
		t.Errorf("Expected empty value, got '%s'", value)
	}

	// Test after filling
	err = playwrightPage.FillElement("#empty-input", "new value")
	if err != nil {
		t.Fatalf("Failed to fill element: %v", err)
	}

	value, err = playwrightPage.GetElementValue("#empty-input")
	if err != nil {
		t.Fatalf("Failed to get element value after fill: %v", err)
	}
	if value != "new value" {
		t.Errorf("Expected value 'new value' after fill, got '%s'", value)
	}
}