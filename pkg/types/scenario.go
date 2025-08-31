package types

// Scenario represents a test scenario containing multiple steps
type Scenario struct {
	Description string `yaml:"desc" json:"description"`
	Steps       []Step `yaml:"steps" json:"steps"`
}

// Step represents a single action in a test scenario
type Step struct {
	// For simple steps like "goto: url"
	Type string `yaml:"type,omitempty" json:"type,omitempty"`
	URL  string `yaml:"url,omitempty" json:"url,omitempty"`

	// For complex steps like click/fill with selector
	Selector string `yaml:"selector,omitempty" json:"selector,omitempty"`
	Value    string `yaml:"value,omitempty" json:"value,omitempty"`

	// For assertion steps
	AssertType string `yaml:"type,omitempty" json:"assert_type,omitempty"`
	Contains   string `yaml:"contains,omitempty" json:"contains,omitempty"`

	// Raw YAML data for complex parsing
	Raw map[string]interface{} `yaml:",inline" json:"-"`
}

// Config represents configuration for the test execution
type Config struct {
	Browser  string `yaml:"browser,omitempty" json:"browser,omitempty"`
	Headless bool   `yaml:"headless,omitempty" json:"headless,omitempty"`
	Timeout  int    `yaml:"timeout,omitempty" json:"timeout,omitempty"`
}
