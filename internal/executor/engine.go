package executor

import (
	"fmt"

	"github.com/haruotsu/ezpw/internal/browser"
	"github.com/haruotsu/ezpw/internal/playwright"
	"github.com/haruotsu/ezpw/pkg/types"
)

// Engine executes test scenarios
type Engine struct {
	config    types.Config
	browser   browser.Browser
	page      browser.Page
	assertion *playwright.Assertion
}

// NewEngine creates a new execution engine
func NewEngine(config types.Config) (*Engine, error) {
	browser, err := playwright.NewBrowser(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create browser: %w", err)
	}

	page, err := browser.NewPage()
	if err != nil {
		browser.Close()
		return nil, fmt.Errorf("failed to create page: %w", err)
	}

	assertion := playwright.NewAssertion(page)

	return &Engine{
		config:    config,
		browser:   browser,
		page:      page,
		assertion: assertion,
	}, nil
}

// Execute runs a test scenario
func (e *Engine) Execute(scenario *types.Scenario) error {
	fmt.Printf("Executing scenario: %s\n", scenario.Description)

	for i, step := range scenario.Steps {
		fmt.Printf("Step %d: %s\n", i+1, step.Type)

		err := e.executeStep(step)
		if err != nil {
			return fmt.Errorf("step %d failed: %w", i+1, err)
		}
	}

	fmt.Println("Scenario completed successfully")
	return nil
}

// executeStep executes a single step
func (e *Engine) executeStep(step types.Step) error {
	switch step.Type {
	case "goto":
		if step.URL == "" {
			return fmt.Errorf("goto step requires URL")
		}
		return e.page.Goto(step.URL)

	case "click":
		if step.Selector == "" {
			return fmt.Errorf("click step requires selector")
		}
		return e.page.Click(step.Selector)

	case "fill":
		if step.Selector == "" {
			return fmt.Errorf("fill step requires selector")
		}
		if step.Value == "" {
			return fmt.Errorf("fill step requires value")
		}
		return e.page.Fill(step.Selector, step.Value)

	case "assert":
		return e.executeAssert(step)

	default:
		return fmt.Errorf("unknown step type: %s", step.Type)
	}
}

// executeAssert handles assertion steps
func (e *Engine) executeAssert(step types.Step) error {
	switch step.AssertType {
	case "text_content":
		if step.Selector == "" {
			return fmt.Errorf("text_content assertion requires selector")
		}
		if step.Contains == "" {
			return fmt.Errorf("text_content assertion requires contains value")
		}
		return e.assertion.AssertTextContent(step.Selector, step.Contains)

	case "url":
		if step.Contains == "" {
			return fmt.Errorf("url assertion requires contains value")
		}
		return e.assertion.AssertURLContains(step.Contains)

	case "exists":
		if step.Selector == "" {
			return fmt.Errorf("exists assertion requires selector")
		}
		return e.assertion.AssertExists(step.Selector)

	default:
		return fmt.Errorf("unknown assertion type: %s", step.AssertType)
	}
}

// Close cleans up the engine resources
func (e *Engine) Close() error {
	if e.browser != nil {
		return e.browser.Close()
	}
	return nil
}
