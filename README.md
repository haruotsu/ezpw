# ezpw (Easy Playwright)

ezpw is a Go-based CLI tool that allows you to run Playwright E2E tests using simple YAML configuration files instead of JavaScript/TypeScript.

## Features (MVP - Phase 1)

- **Simple YAML syntax** for writing test scenarios  
- **Basic browser actions**: goto, click, fill
- **Basic assertions**: text content, URL, element existence
- **Multiple browser support**: Chromium, Firefox, WebKit
- **Headless and headed modes**
- **Cross-platform support**

## Installation

### Step 1: Install ezpw

```bash
# Build from source
git clone https://github.com/haruotsu/ezpw.git
cd ezpw
make build
```

### Step 2: Install Playwright browsers

**⚠️ IMPORTANT**: Before running any tests, you must install Playwright browsers:

```bash
# Install browsers for ezpw (required)
go run github.com/playwright-community/playwright-go/cmd/playwright@latest install

# Or install with system dependencies (recommended for CI)
go run github.com/playwright-community/playwright-go/cmd/playwright@latest install --with-deps
```

### Step 3: Verify installation

```bash
# Test with a simple example
./ezpw run testdata/basic.yml --verbose
```

**If you get a browser not found error**, make sure you ran the Playwright install command in Step 2.

### 🚀 Auto-Install Feature

ezpw can automatically install missing browsers for you! If browsers are not found, ezpw will:

1. **Ask for permission** to install browsers
2. **Automatically download** and install required browsers  
3. **Continue execution** seamlessly

```bash
# Auto-install is enabled by default
./ezpw run test.yml

# Disable auto-install (useful for CI/CD)
./ezpw run test.yml --no-auto-install
```

**Example auto-install flow:**
```
🚫 Browser 'chromium' is not installed.
📦 ezpw requires Playwright browsers to run tests.

Would you like to install the required browsers now? (y/N)
▶ y

📦 Installing Playwright browsers...
✅ Browsers installed successfully!

Executing scenario: Basic example test
...
```

## Quick Start

### 1. Create a test scenario file `basic.yml`:

```yaml
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
```

### 2. Run the test:

```bash
# Basic execution
ezpw run basic.yml

# Run with specific browser
ezpw run basic.yml --browser firefox

# Run in headed mode (show browser)
ezpw run basic.yml --no-headless

# Run with verbose output
ezpw run basic.yml --verbose

# Run all tests in a directory
ezpw run ./tests/
```

## YAML Syntax

### Basic Structure

```yaml
desc: Test description
steps:
  - step1
  - step2
  - ...
```

### Available Steps

#### Navigation

```yaml
- goto: "https://example.com"
```

#### Interactions

```yaml
# Click element
- click:
    selector: "button#submit"

# Fill input
- fill:
    selector: "input[name='username']"
    value: "testuser"
```

#### Assertions

```yaml
# Assert URL contains text
- assert:
    type: url
    contains: "/dashboard"

# Assert element text content
- assert:
    type: text_content
    selector: "h1"
    contains: "Welcome"

# Assert element exists
- assert:
    type: exists
    selector: "#success-message"
```

### Command Line Options

- `--browser`: Browser to use (chromium, firefox, webkit) - default: chromium
- `--headless`: Run in headless mode (default: true)
- `--no-headless`: Run in headed mode
- `--timeout`: Global timeout in milliseconds (default: 30000)
- `--verbose`: Verbose output
- `--debug`: Debug mode
- `--auto-install`: Automatically install browsers if missing (default: true)
- `--no-auto-install`: Disable automatic browser installation (useful for CI/CD)

### Environment Variables

ezpw respects the following environment variables for advanced configuration. **All variables are optional** - ezpw works perfectly without any environment variables.

#### `PLAYWRIGHT_DRIVER_PATH`
Specifies custom path for Playwright driver (Node.js runtime). **Priority: High**

```bash
# Use default driver cache location (recommended)
./ezpw run test.yml

# Use custom driver path (CI/Docker environments)
PLAYWRIGHT_DRIVER_PATH=/custom/driver/path ./ezpw run test.yml
```

#### `PLAYWRIGHT_BROWSERS_PATH`
Specifies custom path for Playwright browsers. **Priority: Medium** (used when `PLAYWRIGHT_DRIVER_PATH` is not set)

```bash
# Use custom browser path (legacy/compatibility)
PLAYWRIGHT_BROWSERS_PATH=/custom/browsers/path ./ezpw run test.yml
```

**Default locations (auto-detected):**
- **Driver cache**:
  - macOS: `~/Library/Caches/ms-playwright-go/`
  - Linux: `~/.cache/ms-playwright-go/`  
  - Windows: `%USERPROFILE%\AppData\Local/ms-playwright-go/`
- **Browser cache**:
  - macOS: `~/Library/Caches/ms-playwright/`
  - Linux: `~/.cache/ms-playwright/`
  - Windows: `%USERPROFILE%\AppData\Local/ms-playwright/`

**When to set environment variables:**
- ✅ **CI/CD environments** with non-standard cache locations
- ✅ **Docker containers** with custom mount points
- ✅ **Corporate environments** with restricted cache directories
- ✅ **Shared systems** with custom Playwright installations
- ❌ **Normal development** (auto-detection works perfectly)

**Priority order:**
1. `PLAYWRIGHT_DRIVER_PATH` (if set)
2. `PLAYWRIGHT_BROWSERS_PATH` (fallback)  
3. Default system cache directories (automatic)

## Development

### Requirements

- Go 1.21+
- Node.js (for Playwright)
- Playwright browsers

### Commands

```bash
# Run tests
make test

# Run with coverage
make test-coverage

# Build
make build

# Format code
make fmt

# Run linter
make lint

# Development build (format, lint, test, build)
make dev
```

### Project Structure

```
ezpw/
├── cmd/ezpw/           # CLI entry point
├── internal/
│   ├── cli/           # CLI processing
│   ├── executor/      # Test execution engine  
│   ├── parser/        # YAML parser
│   └── playwright/    # Playwright integration
├── pkg/types/         # Public type definitions
└── testdata/          # Test scenarios
```

## Contributing

1. Fork the repository
2. Create your feature branch
3. Add tests for new functionality
4. Run `make dev` to ensure quality
5. Submit a pull request

## Status

🚧 **This is the MVP (Phase 1) implementation** 🚧

Current features:
- ✅ Basic YAML parsing and execution
- ✅ Core browser actions (goto, click, fill)
- ✅ Basic assertions (URL, text content, element existence)
- ✅ Multi-browser support

Coming in future phases:
- 🔄 Advanced selectors (text, role, label-based)
- 🔄 Variables and templating
- 🔄 Wait conditions
- 🔄 Screenshot functionality
- 🔄 Advanced assertions
- 🔄 Configuration files
- 🔄 Conditional logic and loops
- 🔄 Reporting and CI/CD integration

## License

MIT License - see LICENSE file for details.
