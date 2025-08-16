// Level 1.3: Channels - Communication Between Goroutines
// This file teaches basic channel operations for goroutine communication

package main

import (
	"fmt"
	// "time" // TODO: Uncomment when implementing timing demos
)

// TUTOR: Channels are Go's way of letting goroutines communicate safely.
// Think of channels as typed pipes - you put data in one end, get it out the other.
// Unlike shared memory, channels prevent race conditions by design.
// The key insight: "Don't communicate by sharing memory; share memory by communicating."
// Channels turn concurrent programming from error-prone to elegant and safe.
// TODO: Demonstrate basic channel creation and operations
func demonstrateBasicChannels() {
	fmt.Println("=== Basic Channel Operations ===")

	// TODO: Create a channel with make(chan type)
	// TODO: Send a value with ch <- value
	// TODO: Receive a value with value := <-ch
	// TODO: Show that this blocks until both sender and receiver are ready
}

// TUTOR: Channels are blocking by default - this is crucial for synchronization.
// Send blocks until another goroutine receives; receive blocks until another sends.
// This blocking behavior creates natural synchronization points in your program.
// Blocking ensures that data handoff is synchronized and safe.
// Understanding blocking behavior is key to designing correct concurrent programs.
// TODO: Demonstrate channel blocking behavior
func demonstrateChannelBlocking() {
	fmt.Println("\n=== Channel Blocking Behavior ===")

	// TODO: Show how sending blocks without a receiver
	// TODO: Show how receiving blocks without a sender
	// TODO: Demonstrate synchronization using channel blocking
	// TODO: Use goroutines to make send/receive operations work together
}

// TUTOR: Channels enable safe data passing between goroutines without race conditions.
// Each value sent through a channel is received by exactly one receiver.
// This prevents the data races you'd get with shared variables.
// Channel communication is the foundation of safe concurrent programming in Go.
// Data flows through channels like water through pipes - clean and predictable.
// TODO: Demonstrate safe data passing through channels
func demonstrateDataPassing() {
	fmt.Println("\n=== Safe Data Passing ===")

	// TODO: Pass different data types through channels
	// TODO: Show how each value is received by exactly one receiver
	// TODO: Demonstrate that there are no race conditions
	// TODO: Compare with shared variable approaches (conceptually)
}

// TUTOR: Channel closing signals "no more values" to receivers.
// Close with close(ch) - only the sender should close channels.
// Receiving from a closed channel returns the zero value and false.
// Closing enables clean termination patterns in concurrent programs.
// Proper closing prevents goroutine leaks and enables elegant shutdown.
// TODO: Demonstrate channel closing semantics
func demonstrateChannelClosing() {
	fmt.Println("\n=== Channel Closing ===")

	// TODO: Create a channel and send several values
	// TODO: Close the channel with close(ch)
	// TODO: Show receiving with value, ok := <-ch syntax
	// TODO: Demonstrate that closed channels return zero value and false
}

// TUTOR: Range loops over channels receive all values until the channel is closed.
// The range automatically handles the receive and checks for channel closure.
// This is the most elegant way to consume all values from a channel.
// Range loops eliminate the need for manual closure checking in most cases.
// This pattern is essential for producer-consumer relationships.
// TODO: Demonstrate range loops over channels
func demonstrateChannelRange() {
	fmt.Println("\n=== Channel Range Loops ===")

	// TODO: Create a channel and send multiple values in a goroutine
	// TODO: Close the channel when done sending
	// TODO: Use for value := range ch to receive all values
	// TODO: Show that range exits automatically when channel closes
}

// TUTOR: Channels coordinate goroutines naturally through their blocking behavior.
// Send/receive operations create synchronization points without explicit locks.
// This coordination is more intuitive than mutex-based approaches.
// Channel coordination scales well and composes nicely.
// Understanding coordination through channels is fundamental to Go's concurrency model.
// TODO: Demonstrate goroutine coordination using channels
func demonstrateChannelCoordination() {
	fmt.Println("\n=== Channel Coordination ===")

	// TODO: Use a channel to coordinate when work is complete
	// TODO: Show how channel operations naturally synchronize goroutines
	// TODO: Demonstrate coordination without WaitGroups
	// TODO: Compare channel coordination vs WaitGroup coordination
}

// TUTOR: Channels have types - you can only send/receive values of the channel's type.
// Type safety prevents many concurrent programming errors at compile time.
// Different types need different channels - Go won't convert automatically.
// Channel types make concurrent programs more reliable and easier to understand.
// Type safety in channels extends Go's type safety to concurrent programming.
// TODO: Demonstrate channel type safety
func demonstrateChannelTypes() {
	fmt.Println("\n=== Channel Type Safety ===")

	// TODO: Create channels of different types (int, string, struct)
	// TODO: Show that types must match exactly
	// TODO: Demonstrate compile-time type checking with channels
	// TODO: Use struct types to pass complex data through channels
}

// TUTOR: Nil channels have special behavior - operations on them block forever.
// This might seem like a bug, but it's actually a feature for advanced patterns.
// For now, just remember: always initialize channels with make().
// Understanding nil channel behavior prevents mysterious blocking bugs.
// Proper channel initialization is essential for working concurrent programs.
// TODO: Demonstrate nil channel behavior and proper initialization
func demonstrateNilChannels() {
	fmt.Println("\n=== Nil Channel Behavior ===")

	// TODO: Show that var ch chan int creates a nil channel
	// TODO: Demonstrate that nil channels block forever (with timeout)
	// TODO: Show proper initialization with make(chan int)
	// TODO: Explain why nil channel behavior is actually useful (briefly)
}

// Helper function to simulate work and send results
func produceData(ch chan<- string, count int) {
	// TODO: Implement a simple producer that sends data and closes channel
}

// Helper function to consume data from channel
func consumeData(ch <-chan string) {
	// TODO: Implement a consumer using range loop
}

func main() {
	fmt.Println("📡 Welcome to Channels - Goroutine Communication! 📡")
	fmt.Println("Learn safe communication before advanced patterns")

	// TODO: Implement each demonstration function
	// Focus on understanding communication before coordination patterns

	demonstrateBasicChannels()
	// demonstrateChannelBlocking()
	// demonstrateDataPassing()
	// demonstrateChannelClosing()
	// demonstrateChannelRange()
	// demonstrateChannelCoordination()
	// demonstrateChannelTypes()
	// demonstrateNilChannels()

	fmt.Println("\n🎉 Congratulations! You can communicate through channels!")
	fmt.Println("Next: Learn choice and control with select.go")
}

/*
📡 Channel Foundation Concepts:

1. **Purpose**: Safe communication between goroutines
2. **Creation**: make(chan Type) creates a typed channel
3. **Operations**: ch <- value (send), value := <-ch (receive)
4. **Blocking**: Send/receive block until counterpart is ready
5. **Closing**: close(ch) signals no more values
6. **Range**: for value := range ch receives until closed

🎯 Essential Patterns:
```go
// Basic communication
ch := make(chan string)
go func() {
    ch <- "Hello"    // Send
}()
msg := <-ch         // Receive

// Producer-Consumer
go func() {
    defer close(ch)  // Always close when done sending
    for i := 0; i < 5; i++ {
        ch <- fmt.Sprintf("Item %d", i)
    }
}()
for item := range ch {  // Receive all items
    fmt.Println(item)
}
```

🚨 Common Beginner Mistakes:
- Forgetting that channels block (causes deadlocks)
- Not closing channels (range loops never exit)
- Closing channels from receiver side (should be sender)
- Using nil channels accidentally (blocks forever)
- Trying to send/receive on closed channels

🔗 What's Next:
Channels enable communication, but what if you need to handle multiple channels?
Next, learn select statements for advanced channel control!
*/
