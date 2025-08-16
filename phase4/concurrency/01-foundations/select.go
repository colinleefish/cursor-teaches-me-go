// Level 1.4: Select - Channel Control Flow
// This file teaches basic select statements for handling multiple channels

package main

import (
	"fmt"
	"time"
)

// TUTOR: Select is Go's switch statement for channels - it waits for multiple channel operations.
// Think of select as "wait for any of these channel operations to become ready."
// Select chooses the first operation that can proceed without blocking.
// This is fundamental for non-blocking channel operations and timeouts.
// Select turns blocking channel operations into flexible, responsive control flow.
// TODO: Demonstrate basic select statement syntax and behavior
func demonstrateBasicSelect() {
	fmt.Println("=== Basic Select Statement ===")

	// TODO: Create two channels
	// TODO: Send values to them from separate goroutines
	// TODO: Use select to receive from whichever is ready first
	// TODO: Show that select picks the first ready operation
}

// TUTOR: When multiple channels are ready simultaneously, select chooses randomly.
// This random selection ensures fairness - no channel gets permanently ignored.
// Random selection prevents starvation in concurrent systems.
// Understanding randomness helps you design fair channel-based systems.
// Fairness through randomness is a key insight in concurrent programming.
// TODO: Demonstrate select's random selection behavior
func demonstrateSelectRandomness() {
	fmt.Println("\n=== Select Random Selection ===")

	// TODO: Create channels that are both ready simultaneously
	// TODO: Run select in a loop to show random selection
	// TODO: Count how many times each channel is chosen
	// TODO: Show that selection is non-deterministic but fair
}

// TUTOR: Default case makes select non-blocking - it runs if no channels are ready.
// Without default, select blocks until at least one channel operation can proceed.
// With default, select never blocks - it always has something to do.
// Default case enables polling patterns and responsive programs.
// Non-blocking select is essential for event-driven programming.
// TODO: Demonstrate select with default case
func demonstrateSelectDefault() {
	fmt.Println("\n=== Select with Default Case ===")

	// TODO: Create a channel but don't send anything
	// TODO: Use select with default to avoid blocking
	// TODO: Show how default enables non-blocking channel operations
	// TODO: Demonstrate polling pattern with default case
}

// TUTOR: Select can handle both send and receive operations in the same statement.
// Each case can be a different type of channel operation.
// This flexibility lets you coordinate complex channel interactions.
// Mixed operations in select enable sophisticated communication patterns.
// Understanding mixed operations unlocks select's full potential.
// TODO: Demonstrate select with mixed send and receive operations
func demonstrateMixedOperations() {
	fmt.Println("\n=== Mixed Send/Receive Operations ===")

	// TODO: Create channels for both sending and receiving
	// TODO: Use select with both send and receive cases
	// TODO: Show how select handles different operation types
	// TODO: Demonstrate that select works with any channel operation
}

// TUTOR: Timeouts with select prevent operations from hanging forever.
// time.After(duration) creates a channel that sends after the timeout.
// Racing your operation against time.After() gives you timeout behavior.
// Timeouts are essential for robust systems that can't wait indefinitely.
// Select-based timeouts are Go's standard pattern for operation limits.
// TODO: Demonstrate basic timeout patterns with select
func demonstrateBasicTimeout() {
	fmt.Println("\n=== Basic Timeout with Select ===")

	// TODO: Create a channel that might send slowly
	// TODO: Use time.After() to create a timeout channel
	// TODO: Race the operation against the timeout with select
	// TODO: Show graceful handling when timeout occurs
}

// TUTOR: Select with channels provides elegant control flow for concurrent programs.
// You can coordinate multiple goroutines using select as a coordination hub.
// Select statements can express complex coordination logic clearly.
// This approach is more intuitive than callback-based systems.
// Select-based control flow is Go's answer to complex asynchronous programming.
// TODO: Demonstrate select for control flow coordination
func demonstrateControlFlow() {
	fmt.Println("\n=== Select for Control Flow ===")

	// TODO: Use select to coordinate multiple workers
	// TODO: Handle completion signals from different goroutines
	// TODO: Show how select creates clean control flow
	// TODO: Compare with other coordination approaches conceptually
}

// TUTOR: Empty select statement (select{}) blocks forever - it's Go's "park" operation.
// This might seem useless, but it's occasionally needed for main functions.
// Empty select is like an infinite wait - the goroutine stops executing.
// Understanding empty select helps you reason about blocking behavior.
// Sometimes you need a goroutine to wait indefinitely, and select{} does that.
// TODO: Demonstrate empty select behavior (carefully with timeout)
func demonstrateEmptySelect() {
	fmt.Println("\n=== Empty Select Behavior ===")

	// TODO: Show what empty select{} does (with timeout protection)
	// TODO: Explain when you might want to block forever
	// TODO: Demonstrate safe alternatives to empty select
	// TODO: Use timeout to avoid actually hanging the program
}

// TUTOR: Nil channels in select are ignored - they don't participate in selection.
// This is a feature, not a bug - it enables dynamic channel enabling/disabling.
// Setting a channel to nil effectively removes it from select consideration.
// Understanding nil channel behavior unlocks advanced select patterns.
// Nil channels give you runtime control over which channels select considers.
// TODO: Demonstrate nil channel behavior in select
func demonstrateNilChannelSelect() {
	fmt.Println("\n=== Nil Channels in Select ===")

	// TODO: Create channels and set some to nil
	// TODO: Show how select ignores nil channels
	// TODO: Demonstrate dynamic channel enabling/disabling
	// TODO: Show how this can be useful for conditional logic
}

// Helper function to create a channel that sends after delay
func delayedSender(msg string, delay time.Duration) <-chan string {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(delay)
		ch <- msg
	}()
	return ch
}

// Helper function to simulate work with random duration
func randomWork(id string) <-chan string {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(time.Duration(100+id[0]) * time.Millisecond) // Fake randomness
		ch <- fmt.Sprintf("Work %s completed", id)
	}()
	return ch
}

func main() {
	fmt.Println("🎛️ Welcome to Select - Channel Control Flow! 🎛️")
	fmt.Println("Master choice and control before advanced patterns")

	// TODO: Implement each demonstration function
	// Build understanding of channel control flow

	demonstrateBasicSelect()
	// demonstrateSelectRandomness()
	// demonstrateSelectDefault()
	// demonstrateMixedOperations()
	// demonstrateBasicTimeout()
	// demonstrateControlFlow()
	// demonstrateEmptySelect()
	// demonstrateNilChannelSelect()

	fmt.Println("\n🎉 Congratulations! You control channel flow with select!")
	fmt.Println("Next: Move to Level 2 for safety concepts and channel types")
}

/*
🎛️ Select Foundation Concepts:

1. **Purpose**: Handle multiple channel operations simultaneously
2. **Syntax**: select { case <-ch1: ... case ch2 <- val: ... default: ... }
3. **Blocking**: Without default, blocks until a case is ready
4. **Non-blocking**: With default, never blocks
5. **Fairness**: Random selection when multiple cases are ready
6. **Timeout**: Use time.After() for timeout cases

🎯 Essential Patterns:
```go
// Basic select
select {
case msg := <-ch1:
    fmt.Println("Received:", msg)
case ch2 <- "hello":
    fmt.Println("Sent hello")
default:
    fmt.Println("No channels ready")
}

// Timeout pattern
select {
case result := <-resultCh:
    fmt.Println("Got result:", result)
case <-time.After(5 * time.Second):
    fmt.Println("Timeout!")
}
```

🚨 Common Beginner Mistakes:
- Expecting deterministic selection (it's random!)
- Not understanding that default makes select non-blocking
- Using empty select{} accidentally (blocks forever)
- Not realizing nil channels are ignored in select
- Forgetting that select chooses first ready case, not all ready cases

🔗 What's Next:
You now understand the 4 cornerstones of Go concurrency!
Next, learn safety concepts and advanced channel types in Level 2.
These foundations will combine into powerful patterns!
*/
