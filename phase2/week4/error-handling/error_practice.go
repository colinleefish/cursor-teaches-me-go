package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Exercise 1: Basic Error Handling
func exerciseBasicErrors() {
	fmt.Println("=== Exercise 1: Basic Error Handling ===")

	// TODO: Create a function parseAge(s string) (int, error)
	// Parse string to int, return error if invalid or negative
	// YOUR CODE HERE

	// Test cases
	testCases := []string{"25", "-5", "abc", "30", ""}

	for _, test := range testCases {
		age, err := parseAge(test)
		if err != nil {
			fmt.Printf("parseAge(%q) -> Error: %v\n", test, err)
		} else {
			fmt.Printf("parseAge(%q) -> Age: %d\n", test, age)
		}
	}

	// TODO: Create a function safeDivide(a, b float64) (float64, error)
	// Return error for division by zero
	// YOUR CODE HERE

	fmt.Println("\nDivision tests:")
	divisionTests := [][2]float64{{10, 2}, {15, 3}, {8, 0}, {-6, 2}}

	for _, test := range divisionTests {
		result, err := safeDivide(test[0], test[1])
		if err != nil {
			fmt.Printf("safeDivide(%.1f, %.1f) -> Error: %v\n", test[0], test[1], err)
		} else {
			fmt.Printf("safeDivide(%.1f, %.1f) -> Result: %.2f\n", test[0], test[1], result)
		}
	}
}

// TODO: Implement parseAge function
// YOUR CODE HERE
func parseAge(s string) (int, error) {
	// YOUR CODE HERE
	_ = strconv.Atoi // Remove this line when implementing
	return 0, errors.New("not implemented")
}

// TODO: Implement safeDivide function
// YOUR CODE HERE
func safeDivide(a, b float64) (float64, error) {
	// YOUR CODE HERE
	return 0, errors.New("not implemented")
}

// Exercise 2: Custom Error Types
func exerciseCustomErrors() {
	fmt.Println("\n=== Exercise 2: Custom Error Types ===")

	// TODO: Define a ValidationError struct with Field and Message
	// It should implement the error interface
	// YOUR CODE HERE

	// TODO: Define an AuthError struct with Code and Message
	// It should implement the error interface
	// YOUR CODE HERE

	// TODO: Create a function validateUser(name, email string) error
	// Return ValidationError for invalid inputs
	// YOUR CODE HERE

	// Test user validation
	users := []struct{ name, email string }{
		{"Alice", "alice@example.com"},
		{"", "bob@example.com"},
		{"Charlie", "invalid-email"},
		{"", ""},
		{"Dave", "dave@test.org"},
	}

	fmt.Println("User validation:")
	for _, user := range users {
		err := validateUser(user.name, user.email)
		if err != nil {
			fmt.Printf("User(%q, %q) -> %v\n", user.name, user.email, err)
		} else {
			fmt.Printf("User(%q, %q) -> Valid\n", user.name, user.email)
		}
	}

	// TODO: Create a function authenticateUser(token string) error
	// Return AuthError for invalid tokens
	// YOUR CODE HERE

	fmt.Println("\nAuthentication tests:")
	tokens := []string{"valid-token-123", "", "expired-abc", "invalid-xyz", "admin-token-456"}

	for _, token := range tokens {
		err := authenticateUser(token)
		if err != nil {
			fmt.Printf("authenticateUser(%q) -> %v\n", token, err)
		} else {
			fmt.Printf("authenticateUser(%q) -> Success\n", token)
		}
	}
}

// TODO: Define ValidationError struct and implement Error() method
// YOUR CODE HERE
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	// YOUR CODE HERE
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// TODO: Define AuthError struct and implement Error() method
// YOUR CODE HERE
type AuthError struct {
	Code    string
	Message string
}

func (e AuthError) Error() string {
	// YOUR CODE HERE
	return fmt.Sprintf("auth error [%s]: %s", e.Code, e.Message)
}

// TODO: Implement validateUser function
// YOUR CODE HERE
func validateUser(name, email string) error {
	// YOUR CODE HERE
	_ = strings.Contains // Remove this line when implementing
	return errors.New("not implemented")
}

// TODO: Implement authenticateUser function
// YOUR CODE HERE
func authenticateUser(token string) error {
	// YOUR CODE HERE
	return errors.New("not implemented")
}

// Exercise 3: Error Wrapping and Unwrapping
func exerciseErrorWrapping() {
	fmt.Println("\n=== Exercise 3: Error Wrapping ===")

	// TODO: Create a function readConfig(filename string) error
	// Simulate reading a config file with potential errors
	// Wrap errors with context using fmt.Errorf with %w verb
	// YOUR CODE HERE

	// TODO: Create a function processConfig() error
	// Call readConfig and wrap any errors with additional context
	// YOUR CODE HERE

	// Test error wrapping
	err := processConfig()
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)

		// TODO: Use errors.Unwrap to get the original error
		// YOUR CODE HERE

		// TODO: Use errors.Is to check for specific error types
		// YOUR CODE HERE

		// TODO: Use errors.As to extract custom error types
		// YOUR CODE HERE
	}

	// Demonstrate error chain
	fmt.Println("\nError wrapping chain:")
	// YOUR CODE HERE - create a chain of wrapped errors
}

// TODO: Implement readConfig function
// YOUR CODE HERE
func readConfig(filename string) error {
	// YOUR CODE HERE
	return errors.New("not implemented")
}

// TODO: Implement processConfig function
// YOUR CODE HERE
func processConfig() error {
	// YOUR CODE HERE
	return readConfig("config.json")
}

// Exercise 4: Error Handling Strategies
func exerciseErrorStrategies() {
	fmt.Println("\n=== Exercise 4: Error Handling Strategies ===")

	// Strategy 1: Fail Fast
	// TODO: Create a function validateInputs(data []string) error
	// Return immediately on first validation error
	// YOUR CODE HERE

	// Strategy 2: Collect All Errors
	// TODO: Create a function validateAllInputs(data []string) []error
	// Validate all inputs and return all errors found
	// YOUR CODE HERE

	// Strategy 3: Retry with Backoff
	// TODO: Create a function retryOperation(op func() error, maxRetries int) error
	// Retry an operation with increasing delay
	// YOUR CODE HERE

	// Test different strategies
	testData := []string{"valid", "", "invalid@", "good@example.com", "bad-format"}

	// Test fail fast
	fmt.Println("Fail Fast Strategy:")
	if err := validateInputs(testData); err != nil {
		fmt.Printf("First error: %v\n", err)
	}

	// Test collect all errors
	fmt.Println("\nCollect All Errors Strategy:")
	errs := validateAllInputs(testData)
	for i, err := range errs {
		if err != nil {
			fmt.Printf("Error %d: %v\n", i, err)
		}
	}

	// Test retry strategy
	fmt.Println("\nRetry Strategy:")
	attemptCount := 0
	flakyOp := func() error {
		attemptCount++
		if attemptCount < 3 {
			return errors.New("temporary failure")
		}
		return nil
	}

	err := retryOperation(flakyOp, 5)
	if err != nil {
		fmt.Printf("Operation failed after retries: %v\n", err)
	} else {
		fmt.Printf("Operation succeeded after %d attempts\n", attemptCount)
	}
}

// TODO: Implement validateInputs function
// YOUR CODE HERE
func validateInputs(data []string) error {
	// YOUR CODE HERE
	return errors.New("not implemented")
}

// TODO: Implement validateAllInputs function
// YOUR CODE HERE
func validateAllInputs(data []string) []error {
	// YOUR CODE HERE
	return []error{errors.New("not implemented")}
}

// TODO: Implement retryOperation function
// YOUR CODE HERE
func retryOperation(op func() error, maxRetries int) error {
	// YOUR CODE HERE
	return op()
}

// Exercise 5: Panic and Recover
func exercisePanicRecover() {
	fmt.Println("\n=== Exercise 5: Panic and Recover ===")

	// TODO: Create a function safeDivision(a, b float64) (result float64, err error)
	// Use panic/recover to handle division by zero gracefully
	// YOUR CODE HERE

	// TODO: Create a function processWithRecover(data []int) (result []int, err error)
	// Process data and recover from any panics, converting them to errors
	// YOUR CODE HERE

	// Test panic recovery
	fmt.Println("Testing panic recovery:")

	// Test safe division
	testCases := [][2]float64{{10, 2}, {8, 0}, {15, 3}}
	for _, test := range testCases {
		result, err := safeDivision(test[0], test[1])
		if err != nil {
			fmt.Printf("safeDivision(%.1f, %.1f) -> Error: %v\n", test[0], test[1], err)
		} else {
			fmt.Printf("safeDivision(%.1f, %.1f) -> Result: %.2f\n", test[0], test[1], result)
		}
	}

	// Test processing with recover
	fmt.Println("\nTesting process with recover:")
	testData := [][]int{{1, 2, 3}, {}, {4, 5, 6}}
	for i, data := range testData {
		result, err := processWithRecover(data)
		if err != nil {
			fmt.Printf("Process %d -> Error: %v\n", i+1, err)
		} else {
			fmt.Printf("Process %d -> Result: %v\n", i+1, result)
		}
	}

	// TODO: Demonstrate when to use panic (programming errors)
	// Create examples of appropriate panic usage
	fmt.Println("\nAppropriate panic usage:")
	// YOUR CODE HERE
}

// TODO: Implement safeDivision function with panic/recover
// YOUR CODE HERE
func safeDivision(a, b float64) (result float64, err error) {
	// YOUR CODE HERE
	return 0, errors.New("not implemented")
}

// TODO: Implement processWithRecover function
// YOUR CODE HERE
func processWithRecover(data []int) (result []int, err error) {
	// YOUR CODE HERE
	return nil, errors.New("not implemented")
}

// Exercise 6: Validation Patterns
func exerciseValidation() {
	fmt.Println("\n=== Exercise 6: Validation Patterns ===")

	// TODO: Define a User struct with validation methods
	// YOUR CODE HERE

	// TODO: Create a function NewUser(name, email string, age int) (*User, error)
	// Validate inputs and return User or error
	// YOUR CODE HERE

	// TODO: Create validation methods for User
	// ValidateName() error, ValidateEmail() error, ValidateAge() error
	// YOUR CODE HERE

	// Test user creation with validation
	fmt.Println("User creation tests:")
	testUsers := []struct {
		name  string
		email string
		age   int
	}{
		{"Alice Johnson", "alice@example.com", 25},
		{"", "bob@test.org", 30},
		{"Charlie", "invalid-email", 35},
		{"Diana", "diana@example.com", -5},
		{"Eve Smith", "eve@test.com", 150},
		{"Frank", "frank@example.org", 40},
	}

	for _, test := range testUsers {
		user, err := NewUser(test.name, test.email, test.age)
		if err != nil {
			fmt.Printf("NewUser(%q, %q, %d) -> Error: %v\n",
				test.name, test.email, test.age, err)
		} else {
			fmt.Printf("NewUser(%q, %q, %d) -> Success: %v\n",
				test.name, test.email, test.age, user)
		}
	}

	// TODO: Create a pipeline validation function
	// validatePipeline(user *User) error that runs all validations
	// YOUR CODE HERE

	// Test pipeline validation
	fmt.Println("\nPipeline validation:")
	user := &User{Name: "Test User", Email: "test@example.com", Age: 25}
	if err := validatePipeline(user); err != nil {
		fmt.Printf("Pipeline validation failed: %v\n", err)
	} else {
		fmt.Printf("Pipeline validation passed for: %v\n", user)
	}
}

// TODO: Define User struct
// YOUR CODE HERE
type User struct {
	Name  string
	Email string
	Age   int
}

// TODO: Implement NewUser function
// YOUR CODE HERE
func NewUser(name, email string, age int) (*User, error) {
	// YOUR CODE HERE
	return nil, errors.New("not implemented")
}

// TODO: Implement User validation methods
// YOUR CODE HERE
func (u *User) ValidateName() error {
	// YOUR CODE HERE
	return errors.New("not implemented")
}

func (u *User) ValidateEmail() error {
	// YOUR CODE HERE
	return errors.New("not implemented")
}

func (u *User) ValidateAge() error {
	// YOUR CODE HERE
	return errors.New("not implemented")
}

// TODO: Implement validatePipeline function
// YOUR CODE HERE
func validatePipeline(user *User) error {
	// YOUR CODE HERE
	return errors.New("not implemented")
}

// Exercise 7: Real-World Error Handling
func exerciseRealWorld() {
	fmt.Println("\n=== Exercise 7: Real-World Error Handling ===")

	// TODO: Create a FileProcessor that demonstrates comprehensive error handling
	// YOUR CODE HERE

	// TODO: Implement processFile(filename string) error
	// Handle various error scenarios: file not found, permission denied, invalid format
	// YOUR CODE HERE

	// TODO: Create a batch processor that handles partial failures
	// processBatch(filenames []string) ([]string, []error)
	// YOUR CODE HERE

	// TODO: Implement error aggregation and reporting
	// createErrorReport(errors []error) string
	// YOUR CODE HERE

	// Simulate file processing
	files := []string{"config.json", "data.csv", "missing.txt", "invalid.xml", "good.json"}

	fmt.Println("Processing individual files:")
	for _, filename := range files {
		err := processFile(filename)
		if err != nil {
			fmt.Printf("processFile(%s) -> Error: %v\n", filename, err)
		} else {
			fmt.Printf("processFile(%s) -> Success\n", filename)
		}
	}

	// Test batch processing
	fmt.Println("\nBatch processing:")
	successful, errors := processBatch(files)

	fmt.Printf("Successfully processed: %v\n", successful)
	if len(errors) > 0 {
		fmt.Println("Errors encountered:")
		for _, err := range errors {
			if err != nil {
				fmt.Printf("  - %v\n", err)
			}
		}

		report := createErrorReport(errors)
		fmt.Printf("\nError Report:\n%s\n", report)
	}

	// TODO: Demonstrate circuit breaker pattern
	// Create a service that fails fast after too many errors
	fmt.Println("\nCircuit breaker pattern:")
	// YOUR CODE HERE
}

// TODO: Implement processFile function
// YOUR CODE HERE
func processFile(filename string) error {
	// YOUR CODE HERE
	return errors.New("not implemented")
}

// TODO: Implement processBatch function
// YOUR CODE HERE
func processBatch(filenames []string) ([]string, []error) {
	// YOUR CODE HERE
	return nil, []error{errors.New("not implemented")}
}

// TODO: Implement createErrorReport function
// YOUR CODE HERE
func createErrorReport(errors []error) string {
	// YOUR CODE HERE
	return "Error report not implemented"
}

// Exercise 8: Error Testing and Debugging
func exerciseErrorDebugging() {
	fmt.Println("\n=== Exercise 8: Error Testing and Debugging ===")

	// TODO: Create functions that demonstrate error testing patterns
	// testErrorConditions() - show how to test error conditions
	// YOUR CODE HERE

	// TODO: Create error tracing utilities
	// traceError(err error) - print error chain
	// YOUR CODE HERE

	// TODO: Demonstrate error categorization
	// categorizeError(err error) string - categorize errors by type
	// YOUR CODE HERE

	// Test error tracing
	fmt.Println("Error tracing example:")
	err := createComplexError()
	if err != nil {
		traceError(err)

		category := categorizeError(err)
		fmt.Printf("Error category: %s\n", category)
	}

	// Test error conditions
	fmt.Println("\nTesting error conditions:")
	testErrorConditions()
}

// TODO: Implement createComplexError function (creates nested error chain)
// YOUR CODE HERE
func createComplexError() error {
	// YOUR CODE HERE
	return errors.New("not implemented")
}

// TODO: Implement traceError function
// YOUR CODE HERE
func traceError(err error) {
	// YOUR CODE HERE
	fmt.Printf("Error trace not implemented: %v\n", err)
}

// TODO: Implement categorizeError function
// YOUR CODE HERE
func categorizeError(err error) string {
	// YOUR CODE HERE
	return "unknown"
}

// TODO: Implement testErrorConditions function
// YOUR CODE HERE
func testErrorConditions() {
	// YOUR CODE HERE
	fmt.Println("Error testing not implemented")
}

// Main function to run all exercises
func main() {
	fmt.Println("🛡️ Go Error Handling Practice")
	fmt.Println("==============================")

	exerciseBasicErrors()
	exerciseCustomErrors()
	exerciseErrorWrapping()
	exerciseErrorStrategies()
	exercisePanicRecover()
	exerciseValidation()
	exerciseRealWorld()
	exerciseErrorDebugging()

	fmt.Println("\n✅ Error handling exercises completed!")
	fmt.Println("\n💡 Key Takeaways:")
	fmt.Println("- Always check and handle errors explicitly")
	fmt.Println("- Use custom error types for domain-specific errors")
	fmt.Println("- Wrap errors to add context while preserving original")
	fmt.Println("- Choose appropriate error handling strategies")
	fmt.Println("- Use panic/recover sparingly for exceptional cases")
	fmt.Println("- Validate inputs early and return errors promptly")
	fmt.Println("- Test error conditions as thoroughly as success cases")
	fmt.Println("- Go's explicit error handling leads to more robust code")
}
