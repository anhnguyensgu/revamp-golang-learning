# LeetCode Go Workspace

This repository contains solutions to LeetCode problems implemented in Go, following a structured approach similar to the Rust workspace pattern.

## Project Structure

- Each problem is implemented in its own package under `problems/` directory
- Each problem package contains:
  - `problem_name.go` - Implementation
  - `problem_name_test.go` - Tests
- Common utilities and helper functions are in `utils/`
- Scripts for generating new problems are in `scripts/`

## Quick Start

1. Run tests for all problems:
   ```bash
   go test ./...
   ```
   
   Or using make:
   ```bash
   make test
   ```

2. Run tests for a specific problem:
   ```bash
   go test ./problems/p0001_two_sum
   ```

3. Create a new problem:
   ```bash
   ./scripts/new.sh p0123_best_time_to_buy_and_sell_stock maxProfit
   ```
   
   Or using make:
   ```bash
   make new slug=p0123_best_time_to_buy_and_sell_stock func=maxProfit
   ```

## Example Usage

To use a solution in your code:

```go
import "github.com/yourusername/go-leetcode/problems/p0001_two_sum"

result := p0001_two_sum.TwoSum([]int{2, 7, 11, 15}, 9)
fmt.Println(result) // Output: [0 1]
```

## Available Commands

- `make test` - Run all tests
- `make test-short` - Run all tests in short mode
- `make test-cover` - Run all tests with coverage
- `make new slug=name func=FunctionName` - Create a new problem
- `make fmt` - Format all Go code
- `make vet` - Run go vet
- `make tidy` - Run go mod tidy
- `make clean` - Clean build artifacts
- `make help` - Show all available commands