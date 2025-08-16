// Level 1.1: Goroutines - The Foundation of Go Concurrency
// This file teaches the basic building blocks of goroutines

package main

import (
	"fmt"
	// "runtime" // TODO: Uncomment when implementing scheduler demos
	// "time"    // TODO: Uncomment when implementing timing/lifecycle demos
)

// TUTOR: Goroutines are lightweight threads managed by the Go runtime, not the OS.
// They're incredibly cheap - you can create thousands without worrying about memory.
// Think of them as "green threads" - the Go scheduler multiplexes them onto OS threads.
// Use goroutines when you want something to run concurrently without blocking the main flow.
// Key insight: goroutines make concurrent programming as easy as adding the `go` keyword.
// TODO: Create and run basic goroutines using the `go` keyword
func demonstrateBasicGoroutines() {
	fmt.Println("=== Basic Goroutine Creation ===")

	// TODO: Launch a simple anonymous function as a goroutine
	// TODO: Launch a named function as a goroutine
	// TODO: Show that main function doesn't wait for goroutines by default
	// TODO: Use time.Sleep to give goroutines a chance to run
}

// TUTOR: Goroutines have a lifecycle: created -> scheduled -> running -> finished.
// Unlike OS threads, creating a goroutine doesn't immediately run it - the scheduler decides when.
// Goroutines start with a small stack (2KB) that grows as needed, making them very memory efficient.
// Understanding lifecycle helps you reason about when goroutines actually execute.
// Key insight: goroutines are cooperative - they yield control at certain points.
// TODO: Demonstrate goroutine lifecycle and scheduling behavior
func demonstrateGoroutineLifecycle() {
	fmt.Println("\n=== Goroutine Lifecycle ===")

	// TODO: Create goroutines that announce their start and finish
	// TODO: Show how goroutines can run in different orders
	// TODO: Demonstrate that goroutines don't run instantly when created
	// TODO: Use runtime.Gosched() to yield control manually
}

// TUTOR: Anonymous function goroutines are extremely common and powerful for closures.
// You can capture variables from the surrounding scope, but be careful with loop variables!
// Anonymous goroutines are perfect for quick concurrent tasks without defining separate functions.
// The closure behavior lets you pass data naturally without explicit parameters.
// Key insight: goroutines + closures = elegant concurrent programming.
// TODO: Demonstrate anonymous function goroutines and closure behavior
func demonstrateAnonymousGoroutines() {
	fmt.Println("\n=== Anonymous Function Goroutines ===")

	// TODO: Create anonymous goroutines with captured variables
	// TODO: Show the classic loop variable capture problem
	// TODO: Demonstrate the correct way to pass loop variables to goroutines
	// TODO: Show how closures can share state (carefully)
}

// TUTOR: Passing parameters to goroutines avoids closure pitfalls and makes data flow explicit.
// Each goroutine gets its own copy of the parameters, preventing race conditions.
// Parameter passing is the safe way to give goroutines the data they need.
// This approach makes goroutines more like pure functions - input in, work done.
// Key insight: explicit parameters = safer concurrent programming.
// TODO: Demonstrate proper parameter passing to goroutines
func demonstrateGoroutineParameters() {
	fmt.Println("\n=== Goroutine Parameter Passing ===")

	// TODO: Create goroutines with different parameter types
	// TODO: Show how each goroutine gets its own copy of parameters
	// TODO: Demonstrate passing structs and complex data
	// TODO: Compare parameter passing vs closure capturing
}

// TUTOR: The Go scheduler automatically manages goroutines across available CPU cores.
// You can see scheduler behavior with runtime functions and understand performance.
// The scheduler uses M:N threading - many goroutines on fewer OS threads.
// Understanding scheduling helps you write efficient concurrent programs.
// Key insight: trust the scheduler, but understand what it's doing.
// TODO: Demonstrate scheduler behavior and runtime inspection
func demonstrateScheduler() {
	fmt.Println("\n=== Goroutine Scheduler ===")

	// TODO: Show current number of goroutines with runtime.NumGoroutine()
	// TODO: Display CPU count and GOMAXPROCS settings
	// TODO: Create many goroutines and watch scheduler distribute them
	// TODO: Use runtime.Gosched() to manually yield control
}

// TUTOR: Goroutines can run indefinitely if not properly managed - this is a "goroutine leak".
// Always ensure goroutines have a way to terminate, or they'll consume memory forever.
// Simple termination strategies: counters, time limits, or external stop signals.
// Proper goroutine lifecycle management is crucial for long-running programs.
// Key insight: every goroutine should have a clear way to stop.
// TODO: Demonstrate goroutine termination patterns
func demonstrateGoroutineTermination() {
	fmt.Println("\n=== Goroutine Termination ===")

	// TODO: Show a goroutine that runs for a limited time
	// TODO: Demonstrate counter-based termination
	// TODO: Show how to avoid creating infinite goroutines
	// TODO: Use simple signaling to stop goroutines cleanly
}

// TUTOR: Monitoring goroutines helps you understand your program's concurrent behavior.
// Runtime package provides tools to inspect goroutine count and behavior.
// Monitoring is essential for debugging concurrency issues and performance.
// Simple monitoring can prevent complex problems in production.
// Key insight: visibility into goroutines = better concurrent programs.
// TODO: Demonstrate basic goroutine monitoring techniques
func demonstrateBasicMonitoring() {
	fmt.Println("\n=== Basic Goroutine Monitoring ===")

	// TODO: Monitor goroutine count before, during, and after creation
	// TODO: Show how to track goroutine creation and completion
	// TODO: Demonstrate simple logging from goroutines
	// TODO: Use runtime.NumGoroutine() to track active goroutines
}

func main() {
	fmt.Println("🧵 Welcome to Goroutines - Foundation of Go Concurrency! 🧵")
	fmt.Println("Master these basics before moving to coordination and communication")

	// TODO: Implement each demonstration function
	// Start with the simplest concepts and build understanding

	demonstrateBasicGoroutines()
	// demonstrateGoroutineLifecycle()
	// demonstrateAnonymousGoroutines()
	// demonstrateGoroutineParameters()
	// demonstrateScheduler()
	// demonstrateGoroutineTermination()
	// demonstrateBasicMonitoring()

	fmt.Println("\n🎉 Congratulations! You understand goroutine fundamentals!")
	fmt.Println("Next: Learn coordination with waitgroups.go")
}

/*
🧵 Goroutine Foundation Concepts:

1. **Creation**: `go functionName()` or `go func(){...}()`
2. **Lifecycle**: Created → Scheduled → Running → Finished
3. **Scheduling**: Go runtime manages when goroutines actually run
4. **Parameters**: Pass data explicitly to avoid closure pitfalls
5. **Termination**: Always provide a way for goroutines to stop
6. **Monitoring**: Use runtime.NumGoroutine() to track active count

🎯 Key Principles:
- Goroutines are cheap - don't fear creating them
- Main function doesn't wait for goroutines automatically
- Use time.Sleep() temporarily to let goroutines run
- Pass parameters explicitly for safety
- Every goroutine should have a termination condition

🚨 Common Beginner Mistakes:
- Forgetting main function exits before goroutines finish
- Capturing loop variables in closures incorrectly
- Creating goroutines without termination conditions
- Not understanding that goroutines run concurrently, not instantly

🔗 What's Next:
After mastering basic goroutines, learn how to coordinate them with WaitGroups!
*/
