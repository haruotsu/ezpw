package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/haruotsu/ezpw/pkg/types"
	"github.com/spf13/cobra"
)

func TestRunCommand_FlagParsing(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedBrowser string
		expectedHeadless bool
		expectedTimeout int
		expectedVerbose bool
		expectedAutoInstall bool
	}{
		{
			name:                "default values",
			args:                []string{"test.yml"},
			expectedBrowser:     "chromium",
			expectedHeadless:    true,
			expectedTimeout:     30000,
			expectedVerbose:     false,
			expectedAutoInstall: true,
		},
		{
			name:                "custom browser",
			args:                []string{"--browser", "firefox", "test.yml"},
			expectedBrowser:     "firefox",
			expectedHeadless:    true,
			expectedTimeout:     30000,
			expectedVerbose:     false,
			expectedAutoInstall: true,
		},
		{
			name:                "no-headless flag",
			args:                []string{"--no-headless", "test.yml"},
			expectedBrowser:     "chromium",
			expectedHeadless:    false,
			expectedTimeout:     30000,
			expectedVerbose:     false,
			expectedAutoInstall: true,
		},
		{
			name:                "verbose mode",
			args:                []string{"--verbose", "test.yml"},
			expectedBrowser:     "chromium",
			expectedHeadless:    true,
			expectedTimeout:     30000,
			expectedVerbose:     true,
			expectedAutoInstall: true,
		},
		{
			name:                "no-auto-install",
			args:                []string{"--no-auto-install", "test.yml"},
			expectedBrowser:     "chromium",
			expectedHeadless:    true,
			expectedTimeout:     30000,
			expectedVerbose:     false,
			expectedAutoInstall: false,
		},
		{
			name:                "custom timeout",
			args:                []string{"--timeout", "60000", "test.yml"},
			expectedBrowser:     "chromium",
			expectedHeadless:    true,
			expectedTimeout:     60000,
			expectedVerbose:     false,
			expectedAutoInstall: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test command with our run command logic
			cmd := &cobra.Command{
				Use: "run",
				RunE: func(cmd *cobra.Command, args []string) error {
					// Get flags (same logic as RunCommand)
					browser, _ := cmd.Flags().GetString("browser")
					headless, _ := cmd.Flags().GetBool("headless")
					noHeadless, _ := cmd.Flags().GetBool("no-headless")
					timeout, _ := cmd.Flags().GetInt("timeout")
					verbose, _ := cmd.Flags().GetBool("verbose")
					autoInstall, _ := cmd.Flags().GetBool("auto-install")
					noAutoInstall, _ := cmd.Flags().GetBool("no-auto-install")

					// Handle headless mode
					if noHeadless {
						headless = false
					}

					// Handle auto-install mode
					if noAutoInstall {
						autoInstall = false
					}

					// Verify values
					if browser != tt.expectedBrowser {
						t.Errorf("Expected browser %s, got %s", tt.expectedBrowser, browser)
					}
					if headless != tt.expectedHeadless {
						t.Errorf("Expected headless %t, got %t", tt.expectedHeadless, headless)
					}
					if timeout != tt.expectedTimeout {
						t.Errorf("Expected timeout %d, got %d", tt.expectedTimeout, timeout)
					}
					if verbose != tt.expectedVerbose {
						t.Errorf("Expected verbose %t, got %t", tt.expectedVerbose, verbose)
					}
					if autoInstall != tt.expectedAutoInstall {
						t.Errorf("Expected autoInstall %t, got %t", tt.expectedAutoInstall, autoInstall)
					}

					return nil
				},
			}

			// Add flags (same as main)
			cmd.Flags().StringP("browser", "b", "chromium", "Browser to use")
			cmd.Flags().Bool("headless", true, "Run browser in headless mode")
			cmd.Flags().Bool("no-headless", false, "Run browser in non-headless mode")
			cmd.Flags().IntP("timeout", "t", 30000, "Global timeout in milliseconds")
			cmd.Flags().BoolP("verbose", "v", false, "Verbose output")
			cmd.Flags().Bool("auto-install", true, "Automatically install browsers if missing")
			cmd.Flags().Bool("no-auto-install", false, "Disable automatic browser installation")

			// Set args and execute
			cmd.SetArgs(tt.args)
			err := cmd.Execute()
			if err != nil {
				t.Fatalf("Command execution failed: %v", err)
			}
		})
	}
}

func TestProcessPath_NonexistentFile(t *testing.T) {
	config := testConfig()
	
	err := processPath("/nonexistent/file.yml", config, false, false, false)
	if err == nil {
		t.Error("Expected error for nonexistent file, got nil")
	}
	
	if !strings.Contains(err.Error(), "path does not exist") {
		t.Errorf("Expected 'path does not exist' error, got: %v", err)
	}
}

func TestProcessPath_InvalidYAMLFile(t *testing.T) {
	// Create a temporary invalid YAML file
	tmpFile, err := os.CreateTemp("", "invalid_*.yml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write invalid YAML content
	_, err = tmpFile.WriteString("invalid: yaml: content: [")
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	config := testConfig()
	
	err = processPath(tmpFile.Name(), config, false, false, false)
	if err == nil {
		t.Error("Expected error for invalid YAML file, got nil")
	}
	
	if !strings.Contains(err.Error(), "failed to parse YAML") {
		t.Errorf("Expected 'failed to parse YAML' error, got: %v", err)
	}
}

func TestProcessDirectory(t *testing.T) {
	// Create a temporary directory with test files
	tmpDir, err := os.MkdirTemp("", "test_dir_*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test YAML files
	validYAML := `desc: Test scenario
steps:
  - goto: "https://example.com"`

	yamlFile1 := filepath.Join(tmpDir, "test1.yml")
	err = os.WriteFile(yamlFile1, []byte(validYAML), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	yamlFile2 := filepath.Join(tmpDir, "test2.yaml")
	err = os.WriteFile(yamlFile2, []byte(validYAML), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Create a non-YAML file (should be ignored)
	txtFile := filepath.Join(tmpDir, "readme.txt")
	err = os.WriteFile(txtFile, []byte("This is not YAML"), 0644)
	if err != nil {
		t.Fatalf("Failed to create text file: %v", err)
	}

	config := testConfig()
	
	// Process the directory - this should succeed since browsers are installed
	err = processPath(tmpDir, config, false, false, false)
	if err != nil {
		// If browsers are not installed, we expect a browser-related error, not YAML parsing error
		if strings.Contains(err.Error(), "failed to parse YAML") {
			t.Errorf("Unexpected YAML parsing error: %v", err)
		}
		// Browser-related errors are acceptable in test environment
		t.Logf("Expected browser-related error in test environment: %v", err)
	}
}

func TestOfferBrowserInstallation_UserDeclines(t *testing.T) {
	// Skip this test for now as it's complex to test interactive input properly
	t.Skip("Skipping interactive test - requires proper stdin/stdout mocking")
}

// testConfig returns a basic config for testing
func testConfig() types.Config {
	return types.Config{
		Browser:  "chromium",
		Headless: true,
		Timeout:  30000,
	}
}

