package types

import (
	"testing"
)

func TestScenarioStructure(t *testing.T) {
	// Test that Scenario struct can be created with required fields
	scenario := Scenario{
		Description: "Test scenario",
		Steps:       []Step{},
	}
	
	if scenario.Description != "Test scenario" {
		t.Errorf("Expected description 'Test scenario', got '%s'", scenario.Description)
	}
	
	if scenario.Steps == nil {
		t.Error("Expected Steps to be initialized as empty slice, got nil")
	}
	
	if len(scenario.Steps) != 0 {
		t.Errorf("Expected Steps length to be 0, got %d", len(scenario.Steps))
	}
}

func TestStepStructure(t *testing.T) {
	// Test basic step types
	gotoStep := Step{
		Type: "goto",
		URL:  "https://example.com",
	}
	
	if gotoStep.Type != "goto" {
		t.Errorf("Expected step type 'goto', got '%s'", gotoStep.Type)
	}
	
	if gotoStep.URL != "https://example.com" {
		t.Errorf("Expected URL 'https://example.com', got '%s'", gotoStep.URL)
	}
	
	// Test click step
	clickStep := Step{
		Type:     "click",
		Selector: "button",
	}
	
	if clickStep.Type != "click" {
		t.Errorf("Expected step type 'click', got '%s'", clickStep.Type)
	}
	
	if clickStep.Selector != "button" {
		t.Errorf("Expected selector 'button', got '%s'", clickStep.Selector)
	}
	
	// Test fill step
	fillStep := Step{
		Type:     "fill",
		Selector: "input",
		Value:    "test value",
	}
	
	if fillStep.Type != "fill" {
		t.Errorf("Expected step type 'fill', got '%s'", fillStep.Type)
	}
	
	if fillStep.Selector != "input" {
		t.Errorf("Expected selector 'input', got '%s'", fillStep.Selector)
	}
	
	if fillStep.Value != "test value" {
		t.Errorf("Expected value 'test value', got '%s'", fillStep.Value)
	}
}

func TestConfigStructure(t *testing.T) {
	config := Config{
		Browser:  "chromium",
		Headless: true,
		Timeout:  30000,
	}
	
	if config.Browser != "chromium" {
		t.Errorf("Expected browser 'chromium', got '%s'", config.Browser)
	}
	
	if !config.Headless {
		t.Error("Expected headless to be true")
	}
	
	if config.Timeout != 30000 {
		t.Errorf("Expected timeout 30000, got %d", config.Timeout)
	}
}