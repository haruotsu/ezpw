package parser

import (
	"strings"
	"testing"
)

func TestParseYAMLScenario(t *testing.T) {
	yamlContent := `
desc: Basic login test
steps:
  - goto: "https://example.com"
  - click:
      selector: "a[href='/login']"
  - fill:
      selector: "input[name='email']"
      value: "test@example.com"
  - fill:
      selector: "input[name='password']"
      value: "password"
  - click:
      selector: "button[type='submit']"
  - assert:
      type: url
      contains: "/dashboard"
`

	scenario, err := ParseYAML(strings.NewReader(yamlContent))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if scenario.Description != "Basic login test" {
		t.Errorf("Expected description 'Basic login test', got '%s'", scenario.Description)
	}

	if len(scenario.Steps) != 6 {
		t.Errorf("Expected 6 steps, got %d", len(scenario.Steps))
	}

	// Test goto step
	if scenario.Steps[0].Type != "goto" {
		t.Errorf("Expected first step type 'goto', got '%s'", scenario.Steps[0].Type)
	}
	if scenario.Steps[0].URL != "https://example.com" {
		t.Errorf("Expected first step URL 'https://example.com', got '%s'", scenario.Steps[0].URL)
	}

	// Test click step
	if scenario.Steps[1].Type != "click" {
		t.Errorf("Expected second step type 'click', got '%s'", scenario.Steps[1].Type)
	}
	if scenario.Steps[1].Selector != "a[href='/login']" {
		t.Errorf("Expected second step selector 'a[href='/login']', got '%s'", scenario.Steps[1].Selector)
	}

	// Test fill step
	if scenario.Steps[2].Type != "fill" {
		t.Errorf("Expected third step type 'fill', got '%s'", scenario.Steps[2].Type)
	}
	if scenario.Steps[2].Selector != "input[name='email']" {
		t.Errorf("Expected third step selector 'input[name='email']', got '%s'", scenario.Steps[2].Selector)
	}
	if scenario.Steps[2].Value != "test@example.com" {
		t.Errorf("Expected third step value 'test@example.com', got '%s'", scenario.Steps[2].Value)
	}

	// Test assert step
	if scenario.Steps[5].Type != "assert" {
		t.Errorf("Expected sixth step type 'assert', got '%s'", scenario.Steps[5].Type)
	}
	if scenario.Steps[5].AssertType != "url" {
		t.Errorf("Expected sixth step assert type 'url', got '%s'", scenario.Steps[5].AssertType)
	}
	if scenario.Steps[5].Contains != "/dashboard" {
		t.Errorf("Expected sixth step contains '/dashboard', got '%s'", scenario.Steps[5].Contains)
	}
}

func TestParseInvalidYAML(t *testing.T) {
	invalidYaml := `
desc: Test
steps:
  - invalid_syntax: [
`

	_, err := ParseYAML(strings.NewReader(invalidYaml))
	if err == nil {
		t.Error("Expected error for invalid YAML, got nil")
	}
}

func TestParseEmptyYAML(t *testing.T) {
	emptyYaml := ""

	_, err := ParseYAML(strings.NewReader(emptyYaml))
	if err == nil {
		t.Error("Expected error for empty YAML, got nil")
	}
}
