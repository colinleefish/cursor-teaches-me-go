// Week 7: Goroutine Basics
// This file demonstrates the fundamentals of goroutines in Go

package main

import (
	"fmt"
	"runtime"
	"time"
)

// TODO: Implement basic goroutine creation and execution
func demonstrateBasicGoroutines() {
	fmt.Println("=== Basic Goroutines ===")

	// TODO: Create a simple goroutine using an anonymous function
	// Hint: Use the 'go' keyword before a function call
	// Print "Hello from goroutine!" inside the goroutine

	// TODO: Create a goroutine that takes parameters
	// Create a function that prints a message with an ID
	// Launch 3 goroutines with different IDs

	// TODO: Add proper synchronization
	// Without synchronization, the main function might exit before goroutines complete
	// Use time.Sleep for now (we'll learn better ways later)

	fmt.Println("Main function continuing...")
}

// TODO: Implement goroutine lifecycle demonstration
func demonstrateGoroutineLifecycle() {
	fmt.Println("\n=== Goroutine Lifecycle ===")

	// TODO: Show goroutine creation, execution, and termination
	// Create a function that:
	// 1. Prints when it starts
	// 2. Does some work (simulate with time.Sleep)
	// 3. Prints when it finishes

	// TODO: Show the difference between:
	// - Goroutines that complete normally
	// - Main function exiting before goroutines finish
	// - Goroutines that run longer than expected
}

// TODO: Compare goroutines vs function calls
func compareGoroutinesVsFunctions() {
	fmt.Println("\n=== Goroutines vs Regular Functions ===")

	// TODO: Implement a function that takes time to execute
	slowFunction := func(name string, duration time.Duration) {
		// TODO: Print start message
		// TODO: Sleep for the given duration
		// TODO: Print completion message
	}

	// TODO: First, call the function normally (synchronous)
	fmt.Println("Calling functions synchronously:")
	start := time.Now()
	// TODO: Call slowFunction three times with different names and 1-second delays
	// TODO: Print total time taken

	// TODO: Then, call the same function as goroutines (asynchronous)
	fmt.Println("\nCalling functions as goroutines:")
	start = time.Now()
	// TODO: Launch three goroutines with the same function
	// TODO: Add appropriate waiting mechanism
	// TODO: Print total time taken
	// TODO: Compare the difference in execution time
}

// TODO: Demonstrate goroutine scheduling
func demonstrateScheduling() {
	fmt.Println("\n=== Goroutine Scheduling ===")

	// TODO: Show how goroutines are scheduled
	// Create multiple goroutines that print their progress
	// Show how they interleave execution

	// TODO: Show the effect of runtime.Gosched()
	// Create goroutines that voluntarily yield control

	// TODO: Show the effect of runtime.GOMAXPROCS()
	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))

	// TODO: Create CPU-intensive goroutines to see scheduling in action
}

// TODO: Implement anonymous function patterns
func demonstrateAnonymousFunctions() {
	fmt.Println("\n=== Anonymous Function Goroutines ===")

	// TODO: Show different ways to create goroutines with anonymous functions

	// TODO: 1. Simple anonymous function
	// go func() { /* code here */ }()

	// TODO: 2. Anonymous function with parameters
	// go func(param1, param2 type) { /* code here */ }(arg1, arg2)

	// TODO: 3. Anonymous function with closure (capturing variables)
	// Be careful about closure variable capture!

	// TODO: 4. Show the common closure pitfall
	// Loop variable capture problem and how to fix it
}

// TODO: Implement error handling in goroutines
func demonstrateErrorHandling() {
	fmt.Println("\n=== Error Handling in Goroutines ===")

	// TODO: Show that panics in goroutines crash the entire program
	// Create a goroutine that panics and show the effect

	// TODO: Show how to handle errors safely
	// Use defer and recover within goroutines

	// TODO: Show patterns for collecting errors from goroutines
	// We'll learn better ways with channels later
}

// TODO: Implement resource management
func demonstrateResourceManagement() {
	fmt.Println("\n=== Resource Management ===")

	// TODO: Show proper resource cleanup in goroutines
	// Use defer statements for cleanup

	// TODO: Demonstrate goroutine leaks
	// Show what happens when goroutines don't terminate

	// TODO: Show how to avoid goroutine leaks
	// Proper termination conditions
}

// TODO: Performance comparison
func performanceComparison() {
	fmt.Println("\n=== Performance Comparison ===")

	// TODO: Compare the cost of creating goroutines vs threads
	// Measure time and memory usage

	// TODO: Show how many goroutines you can create
	// Test with different numbers: 100, 1000, 10000, 100000

	// TODO: Use runtime.ReadMemStats to show memory usage
}

// TODO: Best practices demonstration
func demonstrateBestPractices() {
	fmt.Println("\n=== Goroutine Best Practices ===")

	// TODO: Show proper goroutine naming/identification
	// Use meaningful names in error messages

	// TODO: Show bounded parallelism
	// Don't create unlimited goroutines

	// TODO: Show graceful shutdown patterns
	// How to stop goroutines cleanly

	// TODO: Show monitoring and debugging tips
	// How to track goroutine count and status
}

// Helper function to monitor goroutines
func printGoroutineCount(label string) {
	fmt.Printf("%s - Goroutines: %d\n", label, runtime.NumGoroutine())
}

func main() {
	fmt.Println("🧵 Welcome to Goroutines! 🧵")
	fmt.Println("This file teaches you the fundamentals of Go's lightweight threads")

	// TODO: Implement each demonstration function
	// Start with basic goroutines and work your way up

	demonstrateBasicGoroutines()
	// demonstrateGoroutineLifecycle()
	// compareGoroutinesVsFunctions()
	// demonstrateScheduling()
	// demonstrateAnonymousFunctions()
	// demonstrateErrorHandling()
	// demonstrateResourceManagement()
	// performanceComparison()
	// demonstrateBestPractices()

	fmt.Println("\n🎉 Congratulations! You've learned goroutine basics!")
	fmt.Println("Next: Learn synchronization with WaitGroups in waitgroups.go")
}

/*
🔍 Key Concepts to Remember:

1. **Goroutine Creation**: Use 'go' keyword before any function call
2. **Lightweight**: Goroutines start with ~2KB stack, grow as needed
3. **Scheduling**: Go runtime schedules goroutines across OS threads
4. **Main Exit**: Program exits when main() ends, regardless of running goroutines
5. **Closure Caution**: Be careful capturing loop variables in closures
6. **Error Handling**: Panics in goroutines crash the entire program
7. **Resource Cleanup**: Use defer for proper cleanup in goroutines

🚨 Common Mistakes:
- Not waiting for goroutines to complete
- Capturing loop variables incorrectly in closures
- Creating unlimited goroutines without bounds
- Not handling panics in goroutines
- Forgetting to clean up resources

🎯 Next Steps:
- Learn WaitGroups for proper synchronization
- Understand race conditions and how to avoid them
- Master goroutine patterns and best practices
*/
