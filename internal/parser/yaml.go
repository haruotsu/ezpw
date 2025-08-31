package parser

import (
	"fmt"
	"io"

	"github.com/haruotsu/ezpw/pkg/types"
	"gopkg.in/yaml.v3"
)

const (
	stepTypeGoto   = "goto"
	stepTypeClick  = "click"
	stepTypeFill   = "fill"
	stepTypeAssert = "assert"
)

// ParseYAML parses YAML content and returns a Scenario
func ParseYAML(reader io.Reader) (*types.Scenario, error) {
	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read YAML content: %w", err)
	}

	if len(content) == 0 {
		return nil, fmt.Errorf("empty YAML content")
	}

	var rawScenario map[string]interface{}
	err = yaml.Unmarshal(content, &rawScenario)
	if err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	scenario, err := convertToScenario(rawScenario)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to scenario: %w", err)
	}

	return scenario, nil
}

func convertToScenario(rawScenario map[string]interface{}) (*types.Scenario, error) {
	scenario := &types.Scenario{}

	// Parse description
	if desc, ok := rawScenario["desc"].(string); ok {
		scenario.Description = desc
	}

	// Parse steps
	if stepDataList, ok := rawScenario["steps"].([]interface{}); ok {
		steps, err := convertSteps(stepDataList)
		if err != nil {
			return nil, err
		}
		scenario.Steps = steps
	}

	return scenario, nil
}

// convertSteps converts YAML step data list to Step structs
func convertSteps(stepDataList []interface{}) ([]types.Step, error) {
	var steps []types.Step

	for _, stepData := range stepDataList {
		stepMap, ok := stepData.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid step format")
		}

		step, err := convertStep(stepMap)
		if err != nil {
			return nil, err
		}

		steps = append(steps, step)
	}

	return steps, nil
}

func convertStep(stepMap map[string]interface{}) (types.Step, error) {
	step := types.Step{}

	foundValidType := false
	// Handle different step types
	for key, value := range stepMap {
		switch key {
		case stepTypeGoto:
			handleGotoStep(&step, value)
			foundValidType = true
		case stepTypeClick:
			handleClickStep(&step, value)
			foundValidType = true
		case stepTypeFill:
			handleFillStep(&step, value)
			foundValidType = true
		case stepTypeAssert:
			handleAssertStep(&step, value)
			foundValidType = true
		}
	}

	if !foundValidType {
		return step, fmt.Errorf("no valid step type found in step")
	}

	// Set raw data for complex parsing if needed
	step.Raw = stepMap

	return step, nil
}

func handleGotoStep(step *types.Step, value interface{}) {
	step.Type = stepTypeGoto
	if url, ok := value.(string); ok {
		step.URL = url
	}
}

func handleClickStep(step *types.Step, value interface{}) {
	step.Type = stepTypeClick
	if clickData, ok := value.(map[string]interface{}); ok {
		if selector, ok := clickData["selector"].(string); ok {
			step.Selector = selector
		}
	}
}

func handleFillStep(step *types.Step, value interface{}) {
	step.Type = stepTypeFill
	if fillData, ok := value.(map[string]interface{}); ok {
		if selector, ok := fillData["selector"].(string); ok {
			step.Selector = selector
		}
		if val, ok := fillData["value"].(string); ok {
			step.Value = val
		}
	}
}

func handleAssertStep(step *types.Step, value interface{}) {
	step.Type = stepTypeAssert
	if assertData, ok := value.(map[string]interface{}); ok {
		if assertType, ok := assertData["type"].(string); ok {
			step.AssertType = assertType
		}
		if contains, ok := assertData["contains"].(string); ok {
			step.Contains = contains
		}
		if selector, ok := assertData["selector"].(string); ok {
			step.Selector = selector
		}
	}
}
