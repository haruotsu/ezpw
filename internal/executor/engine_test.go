package executor

import (
	"testing"

	"github.com/haruotsu/ezpw/pkg/types"
)

func TestEngineExecution(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	config := types.Config{
		Browser:  "chromium",
		Headless: true,
		Timeout:  30000,
	}

	scenario := &types.Scenario{
		Description: "Test scenario",
		Steps: []types.Step{
			{Type: "goto", URL: "https://example.com"},
		},
	}

	engine, err := NewEngine(config)
	if err != nil {
		t.Fatalf("Expected no error creating engine, got %v", err)
	}
	defer engine.Close()

	err = engine.Execute(scenario)
	if err != nil {
		t.Errorf("Expected no error executing scenario, got %v", err)
	}
}

func TestEngineExecutionWithComplexSteps(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	config := types.Config{
		Browser:  "chromium",
		Headless: true,
		Timeout:  30000,
	}

	scenario := &types.Scenario{
		Description: "Complex test scenario",
		Steps: []types.Step{
			{Type: "goto", URL: "data:text/html,<html><body><input id='test' type='text'><div id='result'>Initial</div></body></html>"},
			{Type: "fill", Selector: "#test", Value: "test input"},
			{Type: "assert", AssertType: "text_content", Selector: "#result", Contains: "Initial"},
		},
	}

	engine, err := NewEngine(config)
	if err != nil {
		t.Fatalf("Expected no error creating engine, got %v", err)
	}
	defer engine.Close()

	err = engine.Execute(scenario)
	if err != nil {
		t.Errorf("Expected no error executing complex scenario, got %v", err)
	}
}

func TestEngineExecutionWithError(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	config := types.Config{
		Browser:  "chromium",
		Headless: true,
		Timeout:  30000,
	}

	// Create scenario with invalid step
	scenario := &types.Scenario{
		Description: "Error test scenario",
		Steps: []types.Step{
			{Type: "goto", URL: "https://example.com"},
			{Type: "click", Selector: "#nonexistent-element"}, // This should fail
		},
	}

	engine, err := NewEngine(config)
	if err != nil {
		t.Fatalf("Expected no error creating engine, got %v", err)
	}
	defer engine.Close()

	err = engine.Execute(scenario)
	if err == nil {
		t.Error("Expected error executing scenario with invalid steps, got nil")
	}
}

func TestEngineStepTypes(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	config := types.Config{
		Browser:  "chromium",
		Headless: true,
		Timeout:  30000,
	}

	// Test individual step types
	testCases := []struct {
		name     string
		scenario *types.Scenario
		wantErr  bool
	}{
		{
			name: "goto step",
			scenario: &types.Scenario{
				Steps: []types.Step{{Type: "goto", URL: "https://example.com"}},
			},
			wantErr: false,
		},
		{
			name: "fill step",
			scenario: &types.Scenario{
				Steps: []types.Step{
					{Type: "goto", URL: "data:text/html,<html><body><input id='test'></body></html>"},
					{Type: "fill", Selector: "#test", Value: "test"},
				},
			},
			wantErr: false,
		},
		{
			name: "unknown step type",
			scenario: &types.Scenario{
				Steps: []types.Step{{Type: "unknown"}},
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			engine, err := NewEngine(config)
			if err != nil {
				t.Fatalf("Expected no error creating engine, got %v", err)
			}
			defer engine.Close()

			err = engine.Execute(tc.scenario)
			if tc.wantErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}
