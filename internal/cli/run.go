package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	ezpwErrors "github.com/haruotsu/ezpw/internal/errors"
	"github.com/haruotsu/ezpw/internal/executor"
	"github.com/haruotsu/ezpw/internal/parser"
	"github.com/haruotsu/ezpw/pkg/types"
	"github.com/spf13/cobra"
)

// RunCommand handles the run command logic
func RunCommand(cmd *cobra.Command, args []string) error {
	// Get flags
	browser, _ := cmd.Flags().GetString("browser")
	headless, _ := cmd.Flags().GetBool("headless")
	noHeadless, _ := cmd.Flags().GetBool("no-headless")
	timeout, _ := cmd.Flags().GetInt("timeout")
	verbose, _ := cmd.Flags().GetBool("verbose")
	debug, _ := cmd.Flags().GetBool("debug")
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

	config := types.Config{
		Browser:  browser,
		Headless: headless,
		Timeout:  timeout,
	}

	if verbose {
		fmt.Printf("Configuration: Browser=%s, Headless=%t, Timeout=%d\n", config.Browser, config.Headless, config.Timeout)
	}

	// Process each argument (file or directory)
	for _, arg := range args {
		err := processPath(arg, config, verbose, debug, autoInstall)
		if err != nil {
			return fmt.Errorf("failed to process %s: %w", arg, err)
		}
	}

	return nil
}

func processPath(path string, config types.Config, verbose, _, autoInstall bool) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("path does not exist: %s", path)
	}

	if info.IsDir() {
		return processDirectory(path, config, verbose, false, autoInstall)
	}
	return processFile(path, config, verbose, false, autoInstall)
}

func processDirectory(dirPath string, config types.Config, verbose, _, autoInstall bool) error {
	if verbose {
		fmt.Printf("Processing directory: %s\n", dirPath)
	}

	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && (filepath.Ext(path) == ".yml" || filepath.Ext(path) == ".yaml") {
			return processFile(path, config, verbose, false, autoInstall)
		}

		return nil
	})
}

func processFile(filePath string, config types.Config, verbose, _, autoInstall bool) error {
	if verbose {
		fmt.Printf("Processing file: %s\n", filePath)
	}

	// Parse YAML file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scenario, err := parser.ParseYAML(file)
	if err != nil {
		return fmt.Errorf("failed to parse YAML: %w", err)
	}

	if verbose {
		fmt.Printf("Parsed scenario: %s with %d steps\n", scenario.Description, len(scenario.Steps))
	}

	// Create and run engine
	engine, err := executor.NewEngine(config)
	if err != nil {
		// Check if it's a browser not found error
		var browserNotFoundErr *ezpwErrors.BrowserNotFoundError
		if errors.As(err, &browserNotFoundErr) {
			if autoInstall && offerBrowserInstallation(browserNotFoundErr.Browser, verbose) {
				// Retry after installation
				engine, err = executor.NewEngine(config)
				if err != nil {
					return fmt.Errorf("failed to create engine after browser installation: %w", err)
				}
			} else {
				return fmt.Errorf("browser installation required: %w", err)
			}
		} else {
			return fmt.Errorf("failed to create engine: %w", err)
		}
	}
	defer engine.Close()

	err = engine.Execute(scenario)
	if err != nil {
		return fmt.Errorf("failed to execute scenario: %w", err)
	}

	fmt.Printf("✓ Successfully executed: %s\n", filePath)
	return nil
}

// offerBrowserInstallation offers to install missing browsers
func offerBrowserInstallation(browser string, verbose bool) bool {
	fmt.Printf("\n🚫 Browser '%s' is not installed.\n", browser)
	fmt.Println("📦 ezpw requires Playwright browsers to run tests.")
	fmt.Println()
	fmt.Println("Would you like to install the required browsers now? (y/N)")
	fmt.Print("▶ ")

	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return false
	}

	response = strings.TrimSpace(strings.ToLower(response))
	if response != "y" && response != "yes" {
		fmt.Println()
		fmt.Println("⚠️  To install browsers manually, run:")
		fmt.Println("   go run github.com/playwright-community/playwright-go/cmd/playwright@latest install")
		return false
	}

	fmt.Printf("\n📦 Installing Playwright browsers...\n")
	if verbose {
		fmt.Println("Running: go run github.com/playwright-community/playwright-go/cmd/playwright@latest install")
	}

	cmd := exec.Command("go", "run", "github.com/playwright-community/playwright-go/cmd/playwright@latest", "install")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("❌ Failed to install browsers: %v\n", err)
		fmt.Println("Please install manually using the command above.")
		return false
	}

	fmt.Println("✅ Browsers installed successfully!")
	fmt.Println()
	return true
}
