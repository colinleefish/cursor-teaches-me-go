// Week 7: WaitGroups for Goroutine Synchronization
// This file demonstrates how to coordinate multiple goroutines using sync.WaitGroup

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// TODO: Implement basic WaitGroup usage
func demonstrateBasicWaitGroup() {
	fmt.Println("=== Basic WaitGroup Usage ===")

	// TODO: Create a WaitGroup variable
	// var wg sync.WaitGroup

	// TODO: Create 5 goroutines that do some work
	// For each goroutine:
	// 1. Call wg.Add(1) before launching
	// 2. Use 'defer wg.Done()' inside the goroutine
	// 3. Do some work (simulate with time.Sleep and print messages)

	// TODO: Wait for all goroutines to complete
	// Call wg.Wait() to block until all are done

	fmt.Println("All goroutines completed!")
}

// TODO: Implement dynamic WaitGroup management
func demonstrateDynamicWaitGroups() {
	fmt.Println("\n=== Dynamic WaitGroup Management ===")

	// TODO: Show how to add goroutines dynamically
	// Start with a slice of work items
	jobs := []string{"task1", "task2", "task3", "task4", "task5"}

	// TODO: Create a WaitGroup and process each job
	// Use a goroutine for each job
	// Show how to add to WaitGroup based on slice length

	processJob := func(jobName string, wg *sync.WaitGroup) {
		// TODO: Implement job processing
		// 1. Use defer wg.Done()
		// 2. Simulate work with random duration
		// 3. Print start and completion messages
	}

	// TODO: Launch all jobs and wait for completion
}

// TODO: Implement WaitGroup with error handling
func demonstrateWaitGroupWithErrors() {
	fmt.Println("\n=== WaitGroup with Error Handling ===")

	// TODO: Create a worker function that might fail
	worker := func(id int, wg *sync.WaitGroup, errCh chan error) {
		// TODO: Implement worker with error handling
		// 1. Use defer wg.Done()
		// 2. Simulate work that might fail (use random chance)
		// 3. Send errors to error channel if they occur
		// 4. Print success/failure messages
	}

	// TODO: Create WaitGroup and error channel
	// TODO: Launch multiple workers
	// TODO: Use a separate goroutine to close error channel after all workers complete
	// TODO: Collect and handle errors
}

// TODO: Implement nested WaitGroups
func demonstrateNestedWaitGroups() {
	fmt.Println("\n=== Nested WaitGroups ===")

	// TODO: Show how to use WaitGroups for hierarchical work
	// Main WaitGroup for high-level tasks
	// Nested WaitGroups for subtasks within each main task

	processMainTask := func(taskId int, mainWg *sync.WaitGroup) {
		// TODO: Implement main task processing
		// 1. Use defer mainWg.Done()
		// 2. Create subtasks with their own WaitGroup
		// 3. Process subtasks concurrently
		// 4. Wait for all subtasks before completing main task
	}

	// TODO: Launch multiple main tasks, each with subtasks
}

// TODO: Implement WaitGroup patterns
func demonstrateWaitGroupPatterns() {
	fmt.Println("\n=== WaitGroup Patterns ===")

	// TODO: Pattern 1: Producer-Consumer with WaitGroup
	// Create producers that generate work
	// Create consumers that process work
	// Use WaitGroups to coordinate shutdown

	// TODO: Pattern 2: Bounded Worker Pool
	// Limit the number of concurrent workers
	// Use WaitGroup to wait for all work to complete

	// TODO: Pattern 3: Batch Processing
	// Process items in batches using WaitGroups
	// Show how to control batch size and parallelism
}

// TODO: Implement WaitGroup with timeouts
func demonstrateWaitGroupTimeout() {
	fmt.Println("\n=== WaitGroup with Timeout ===")

	// TODO: Show how to add timeout functionality to WaitGroup
	// Use a separate goroutine with time.After
	// Demonstrate both successful completion and timeout scenarios

	waitWithTimeout := func(wg *sync.WaitGroup, timeout time.Duration) bool {
		// TODO: Implement timeout mechanism
		// Return true if completed within timeout, false otherwise
		return false
	}

	// TODO: Test with fast and slow goroutines
}

// TODO: Implement WaitGroup best practices
func demonstrateBestPractices() {
	fmt.Println("\n=== WaitGroup Best Practices ===")

	// TODO: Show common mistakes and how to avoid them

	// TODO: Mistake 1: Adding to WaitGroup inside goroutine
	// Show the race condition this can cause

	// TODO: Mistake 2: Forgetting defer wg.Done()
	// Show what happens when goroutines don't signal completion

	// TODO: Mistake 3: Reusing WaitGroup incorrectly
	// Show proper WaitGroup lifecycle

	// TODO: Best Practice 1: Pass WaitGroup by pointer
	// Show why value passing doesn't work

	// TODO: Best Practice 2: Add before launching goroutine
	// Always call wg.Add() before 'go func()'

	// TODO: Best Practice 3: Use defer for Done()
	// Ensures Done() is called even if goroutine panics
}

// TODO: Implement complex coordination scenario
func demonstrateComplexCoordination() {
	fmt.Println("\n=== Complex Coordination Scenario ===")

	// TODO: Implement a multi-stage processing pipeline
	// Stage 1: Data preparation (multiple goroutines)
	// Stage 2: Data processing (different set of goroutines)
	// Stage 3: Result aggregation (single goroutine)
	// Use WaitGroups to coordinate between stages

	// TODO: Show how to pass data between stages
	// Use channels for data flow, WaitGroups for synchronization
}

// TODO: Implement performance comparison
func performanceComparison() {
	fmt.Println("\n=== Performance Comparison ===")

	// TODO: Compare different synchronization methods
	// 1. No synchronization (incorrect but fast)
	// 2. WaitGroup synchronization
	// 3. Channel-based synchronization
	// 4. Mutex-based synchronization

	// TODO: Measure execution time and resource usage
	// Show the overhead of different approaches
}

// Helper function to simulate work
func simulateWork(workerId int, duration time.Duration) {
	fmt.Printf("Worker %d: Starting work...\n", workerId)
	time.Sleep(duration)
	fmt.Printf("Worker %d: Work completed!\n", workerId)
}

// Helper function to simulate work that might fail
func simulateWorkWithError(workerId int) error {
	fmt.Printf("Worker %d: Starting risky work...\n", workerId)

	// Simulate work time
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	// 20% chance of failure
	if rand.Float32() < 0.2 {
		return fmt.Errorf("worker %d failed", workerId)
	}

	fmt.Printf("Worker %d: Risky work completed successfully!\n", workerId)
	return nil
}

func main() {
	fmt.Println("🔄 Welcome to WaitGroups! 🔄")
	fmt.Println("This file teaches you how to coordinate goroutines properly")

	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// TODO: Implement each demonstration function
	// Start with basic usage and progress to advanced patterns

	demonstrateBasicWaitGroup()
	// demonstrateDynamicWaitGroups()
	// demonstrateWaitGroupWithErrors()
	// demonstrateNestedWaitGroups()
	// demonstrateWaitGroupPatterns()
	// demonstrateWaitGroupTimeout()
	// demonstrateBestPractices()
	// demonstrateComplexCoordination()
	// performanceComparison()

	fmt.Println("\n🎉 Congratulations! You've mastered WaitGroup synchronization!")
	fmt.Println("Next: Learn about race conditions in race_conditions.go")
}

/*
🔍 Key Concepts to Remember:

1. **WaitGroup Purpose**: Coordinate completion of multiple goroutines
2. **Three Methods**:
   - Add(n): Add n goroutines to wait for
   - Done(): Signal that one goroutine is complete
   - Wait(): Block until all goroutines are done
3. **Pass by Pointer**: Always pass WaitGroup as pointer (*sync.WaitGroup)
4. **Add Before Launch**: Call wg.Add() before launching goroutine
5. **Use Defer**: Always use 'defer wg.Done()' to ensure completion signal
6. **No Reuse**: Don't reuse WaitGroup after Wait() returns

📋 WaitGroup Pattern:
```go
var wg sync.WaitGroup

for i := 0; i < n; i++ {
    wg.Add(1)  // Add before launching
    go func(id int) {
        defer wg.Done()  // Always use defer
        // Do work here
    }(i)
}

wg.Wait()  // Wait for all to complete
```

🚨 Common Mistakes:
- Calling Add() inside goroutine (race condition)
- Forgetting defer wg.Done()
- Passing WaitGroup by value instead of pointer
- Reusing WaitGroup without proper reset
- Not handling panics (Done() won't be called)

🎯 Next Steps:
- Learn about race conditions and data safety
- Understand when to use channels vs WaitGroups
- Master complex coordination patterns
*/
