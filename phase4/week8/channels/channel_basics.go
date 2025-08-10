// Week 8: Channel Basics
// This file demonstrates the fundamentals of channels in Go

package main

import (
	"fmt"
	"sync"
	"time"
)

// TODO: Demonstrate basic channel creation and usage
func demonstrateBasicChannels() {
	fmt.Println("=== Basic Channel Operations ===")

	// TODO: Create an unbuffered channel
	// Use make(chan type) syntax
	ch := make(chan int)

	// TODO: Send and receive in separate goroutines
	// Show that unbuffered channels block until both sender and receiver are ready
	go func() {
		ch <- 1
	}()
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println(<-ch)
	}()

	// TODO: Demonstrate the synchronization property
	// Show how channels coordinate goroutines
	time.Sleep(1 * time.Second)

	fmt.Println("Basic channel operations completed!")
}

// TODO: Demonstrate channel blocking behavior
func demonstrateChannelBlocking() {
	fmt.Println("\n=== Channel Blocking Behavior ===")

	// TODO: Show blocking send on unbuffered channel
	// Create a channel and try to send without a receiver
	// Use a goroutine to avoid deadlock and show the blocking
	ch := make(chan int)
	// go func() {
	// 	fmt.Println("Sending 6 in 1 second")
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("Sending 6")
	// 	ch <- 6
	// 	fmt.Println("Sent 6")
	// }()

	// time.Sleep(3 * time.Second)

	// TODO: Show blocking receive on empty channel
	// Try to receive from channel with no sender
	// go func() {
	// 	fmt.Println("Receiving from channel")
	// 	fmt.Println(<-ch)
	// 	fmt.Println("Received from channel")
	// }()

	// time.Sleep(3 * time.Second)

	// TODO: Show synchronous nature of unbuffered channels
	// Demonstrate that send doesn't complete until receive happens

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Sending 6")
		ch <- 6
		fmt.Println("Sent 6")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Receiving from channel")
		fmt.Println(<-ch)
		fmt.Println("Received from channel")
	}()

	wg.Wait()

	fmt.Println("Channel blocking demonstration completed!")
}

// TODO: Demonstrate channel directions (send-only, receive-only)
func demonstrateChannelDirections() {
	fmt.Println("\n=== Channel Directions ===")

	// TODO: Create functions with directional channel parameters

	// Send-only channel parameter

	ch := make(chan string)
	sender := func(ch chan<- string, message string) {
		ch <- message
	}

	// Receive-only channel parameter
	// receiver := func(ch <-chan string) string {
	//     // TODO: Receive message from channel
	//     // Show that you can only receive, not send
	//     return ""
	// }
	receiver := func(ch <-chan string) string {
		return <-ch
	}

	// TODO: Create a bidirectional channel
	// Pass it to sender and receiver functions
	// Show how Go automatically converts to directional channels

	go sender(ch, "Hello")

	fmt.Println(receiver(ch))

	fmt.Println("Channel directions demonstration completed!")
}

// TODO: Demonstrate channel closing
func demonstrateChannelClosing() {
	fmt.Println("\n=== Channel Closing ===")

	// TODO: Show how to close a channel
	// Demonstrate that closing is done by sender, not receiver
	wg := sync.WaitGroup{}

	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// ch <- 1
		close(ch)
	}()

	// TODO: Show receiving from closed channel
	// Demonstrate that receives from closed channel return zero value
	// Show the two-value receive to check if channel is closed

	time.Sleep(1 * time.Second)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch)
	}()

	// TODO: Show what happens when sending to closed channel
	// This should panic - demonstrate with recover

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: %v\n", r)
			}
		}()
		ch <- 1
	}()

	// TODO: Show multiple closes cause panic
	// Use recover to catch the panic

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: %v\n", r)
			}
		}()
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: %v\n", r)
			}
		}()
		close(ch)
		close(ch)
	}()

	wg.Wait()

	fmt.Println("Channel closing demonstration completed!")
}

// TODO: Demonstrate range over channels
func demonstrateChannelRange() {
	fmt.Println("\n=== Range Over Channels ===")

	// TODO: Create a producer that sends multiple values
	// producer := func(ch chan<- int) {
	//     // TODO: Send numbers 1-5 to channel
	//     // TODO: Close the channel when done
	// }

	// TODO: Create a consumer that uses range
	// consumer := func(ch <-chan int) {
	//     // TODO: Use range to receive all values
	//     // Show that range automatically exits when channel is closed
	// }

	// TODO: Connect producer and consumer
	// Show the clean pattern for producer-consumer with range

	fmt.Println("Channel range demonstration completed!")
}

// TODO: Demonstrate nil channels
func demonstrateNilChannels() {
	fmt.Println("\n=== Nil Channels ===")

	// TODO: Show nil channel behavior
	// var nilCh chan int

	// TODO: Show that sending to nil channel blocks forever
	// Use select with default to avoid blocking

	// TODO: Show that receiving from nil channel blocks forever
	// Use select with default to avoid blocking

	// TODO: Show practical use of nil channels
	// Use nil channels to disable cases in select statements

	fmt.Println("Nil channels demonstration completed!")
}

// TODO: Demonstrate channel as first-class values
func demonstrateChannelsAsValues() {
	fmt.Println("\n=== Channels as First-Class Values ===")

	// TODO: Show channels can be stored in variables
	// TODO: Show channels can be passed as function parameters
	// TODO: Show channels can be returned from functions
	// TODO: Show channels can be stored in slices/maps

	// TODO: Create a channel factory function
	// createIntChannel := func(bufferSize int) chan int {
	//     // TODO: Return a new channel with given buffer size
	//     return nil
	// }

	// TODO: Create a channel registry
	// channelRegistry := make(map[string]chan string)

	// TODO: Store and retrieve channels from registry
	// Show how channels can be managed dynamically

	fmt.Println("Channels as values demonstration completed!")
}

// TODO: Demonstrate channel communication patterns
func demonstrateCommunicationPatterns() {
	fmt.Println("\n=== Communication Patterns ===")

	// TODO: Pattern 1: Request-Response
	requestResponse := func() {
		// TODO: Create request and response channels
		// TODO: Send request and wait for response
		// Show bidirectional communication
	}

	// TODO: Pattern 2: Worker notification
	workerNotification := func() {
		// TODO: Use channel to signal worker completion
		// Show how channels can coordinate work
	}

	// TODO: Pattern 3: Broadcast using channel closing
	broadcast := func() {
		// TODO: Use channel closing to signal multiple goroutines
		// Show how closing can be used for broadcast
	}

	// TODO: Execute each pattern
	requestResponse()
	workerNotification()
	broadcast()

	fmt.Println("Communication patterns demonstration completed!")
}

// TODO: Demonstrate channel vs other synchronization
func demonstrateChannelVsOther() {
	fmt.Println("\n=== Channels vs Other Synchronization ===")

	// TODO: Compare channel-based solution vs mutex-based solution
	// for a simple counter problem

	// TODO: Channel-based counter
	// channelCounter := func(n int) int {
	//     // TODO: Use channel to coordinate counter increments
	//     // Show how channels can replace mutexes
	//     return 0
	// }

	// TODO: Show timing comparison
	// Measure performance of channel vs mutex approach

	fmt.Println("Channel comparison demonstration completed!")
}

// TODO: Demonstrate common channel mistakes
func demonstrateCommonMistakes() {
	fmt.Println("\n=== Common Channel Mistakes ===")

	// TODO: Mistake 1: Forgetting to close channel in range
	// mistake1 := func() {
	//     // TODO: Show what happens when you forget to close
	//     // Use timeout to avoid infinite blocking
	// }

	// TODO: Mistake 2: Sending on closed channel
	// mistake2 := func() {
	//     // TODO: Show panic when sending to closed channel
	//     // Use recover to handle the panic
	// }

	// TODO: Mistake 3: Closing channel multiple times
	// mistake3 := func() {
	//     // TODO: Show panic when closing already closed channel
	//     // Use recover to handle the panic
	// }

	// TODO: Mistake 4: Wrong direction expectations
	// mistake4 := func() {
	//     // TODO: Show compilation errors with wrong channel directions
	//     // Comment out the incorrect code with explanations
	// }

	// TODO: Execute each mistake demonstration safely
	fmt.Println("Demonstrating mistakes (safely)...")

	fmt.Println("Common mistakes demonstration completed!")
}

// TODO: Demonstrate channel performance characteristics
func demonstrateChannelPerformance() {
	fmt.Println("\n=== Channel Performance ===")

	// TODO: Measure channel operation performance
	// Compare send/receive speeds with different patterns

	// TODO: Test 1: Simple send/receive
	simpleTest := func() {
		// TODO: Measure time for many send/receive operations
	}

	// TODO: Test 2: Goroutine creation overhead vs channel reuse
	goroutineTest := func() {
		// TODO: Compare creating new goroutines vs reusing with channels
	}

	// TODO: Test 3: Channel vs direct function call
	callTest := func() {
		// TODO: Compare channel communication vs direct function calls
	}

	// TODO: Run performance tests
	simpleTest()
	goroutineTest()
	callTest()

	fmt.Println("Channel performance demonstration completed!")
}

// Helper function to simulate work
func doWork(id int, duration time.Duration) {
	fmt.Printf("Worker %d: Starting work for %v\n", id, duration)
	time.Sleep(duration)
	fmt.Printf("Worker %d: Work completed\n", id)
}

// Helper function to safely demonstrate panics
func safelyDemonstratePanic(description string, fn func()) {
	fmt.Printf("--- %s ---\n", description)
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Caught panic: %v\n", r)
		}
	}()
	fn()
}

func main() {
	fmt.Println("📡 Welcome to Channel Basics! 📡")
	fmt.Println("This file teaches you the fundamentals of Go channels")

	// TODO: Implement each demonstration function
	// Start with basic operations and progress to advanced concepts

	// demonstrateBasicChannels()
	// demonstrateChannelBlocking()
	// demonstrateChannelDirections()
	demonstrateChannelClosing()
	// demonstrateChannelRange()
	// demonstrateNilChannels()
	// demonstrateChannelsAsValues()
	// demonstrateCommunicationPatterns()
	// demonstrateChannelVsOther()
	// demonstrateCommonMistakes()
	// demonstrateChannelPerformance()

	fmt.Println("\n🎉 Congratulations! You've learned channel basics!")
	fmt.Println("Next: Learn about buffered channels in buffered_channels.go")
}

/*
🔍 Key Concepts to Remember:

1. **Channel Creation**: make(chan Type) for unbuffered, make(chan Type, size) for buffered
2. **Send/Receive**: ch <- value to send, value := <-ch to receive
3. **Blocking**: Unbuffered channels block until both sender and receiver are ready
4. **Directions**: chan<- for send-only, <-chan for receive-only
5. **Closing**: close(ch) to close, value, ok := <-ch to check if closed
6. **Range**: for value := range ch automatically handles channel closing
7. **Nil Channels**: Sending/receiving on nil channels blocks forever

📋 Channel Patterns:
```go
// Basic send/receive
ch := make(chan int)
go func() { ch <- 42 }()
value := <-ch

// Producer-consumer with range
go func() {
    defer close(ch)
    for i := 0; i < 10; i++ {
        ch <- i
    }
}()
for value := range ch {
    process(value)
}

// Check if closed
value, ok := <-ch
if !ok {
    // Channel is closed
}
```

🚨 Common Mistakes:
- Deadlock on unbuffered channels (no receiver/sender)
- Forgetting to close channels when using range
- Sending on closed channels (causes panic)
- Closing channels multiple times (causes panic)
- Wrong channel direction expectations

🎯 Next Steps:
- Learn buffered channels for performance
- Master select statements for multiplexing
- Understand advanced channel patterns
- Practice with real-world scenarios
*/
