package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version = "dev"

var rootCmd = &cobra.Command{
	Use:   "ezpw",
	Short: "Easy Playwright YAML - Run Playwright tests with YAML configuration",
	Long: `ezpw (Easy Playwright) is a tool that allows you to run Playwright 
E2E tests using simple YAML configuration files instead of JavaScript/TypeScript.

This tool provides an intuitive way to write browser automation and testing scenarios.`,
	Version: Version,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("ezpw version %s\n", Version)
	},
}

var runCmd = &cobra.Command{
	Use:   "run [scenario-file-or-directory]",
	Short: "Run test scenarios",
	Long: `Run test scenarios from YAML files. You can specify:
  - A single YAML file: ezpw run test.yml
  - A directory: ezpw run ./tests/
  - Multiple files: ezpw run test1.yml test2.yml`,
	Args: cobra.MinimumNArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		fmt.Printf("Running scenarios: %v\n", args)
		// TODO: 実装する
		fmt.Println("This feature will be implemented in Phase 1")
	},
}

const defaultTimeout = 30000

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringP("browser", "b", "chromium", "Browser to use (chromium, firefox, webkit)")
	runCmd.Flags().Bool("headless", true, "Run browser in headless mode")
	runCmd.Flags().Bool("no-headless", false, "Run browser in non-headless mode")
	runCmd.Flags().IntP("parallel", "p", 1, "Number of parallel executions")
	runCmd.Flags().Int("timeout", defaultTimeout, "Global timeout in milliseconds")
	runCmd.Flags().StringP("output", "o", "./reports", "Output directory for reports")
	runCmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	runCmd.Flags().Bool("debug", false, "Debug mode")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
