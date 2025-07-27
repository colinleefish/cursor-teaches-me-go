package main

import (
	"fmt"
)

// ===== EXERCISE 1: BASIC INTERFACE IMPLEMENTATION =====

// TODO: Define a Vehicle interface with the following methods:
// - Start() string
// - Stop() string
// - GetSpeed() int
// YOUR CODE HERE

// TODO: Implement the Vehicle interface for:
// 1. Car struct with fields: Brand, Model, Speed
// 2. Bicycle struct with fields: Type, Speed
// 3. Airplane struct with fields: Model, Speed, Altitude
// Each should have appropriate Start(), Stop(), and GetSpeed() implementations
// YOUR CODE HERE

func exerciseBasicInterface() {
	fmt.Println("=== Exercise 1: Basic Interface Implementation ===")

	// TODO: Create instances of Car, Bicycle, and Airplane
	// YOUR CODE HERE

	// TODO: Store them in a slice of Vehicle interfaces
	// YOUR CODE HERE

	// TODO: Loop through and call Start(), Stop(), and GetSpeed() on each
	// YOUR CODE HERE
}

// ===== EXERCISE 2: INTERFACE COMPOSITION =====

// TODO: Define these interfaces:
// - Flyer interface with Fly() string method
// - Swimmer interface with Swim() string method
// - AmphibiousVehicle interface that combines Flyer and Swimmer
// YOUR CODE HERE

// TODO: Implement a SeaPlane struct that satisfies AmphibiousVehicle
// It should have Name, Speed, and Altitude fields
// YOUR CODE HERE

func exerciseInterfaceComposition() {
	fmt.Println("\n=== Exercise 2: Interface Composition ===")

	// TODO: Create a SeaPlane instance
	// YOUR CODE HERE

	// TODO: Use it as different interface types (Flyer, Swimmer, AmphibiousVehicle)
	// YOUR CODE HERE

	// TODO: Create functions that accept each interface type and test them
	// YOUR CODE HERE
}

// ===== EXERCISE 3: IMPLEMENTING STANDARD INTERFACES =====

// TODO: Create a Product struct with:
// - ID (int)
// - Name (string)
// - Price (float64)
// - Category (string)
// YOUR CODE HERE

// TODO: Implement fmt.Stringer for Product
// Format: "Product: [Name] ($[Price]) - Category: [Category]"
// YOUR CODE HERE

// TODO: Create a ProductList type ([]Product) and implement sort.Interface
// Sort by Price in ascending order
// YOUR CODE HERE

func exerciseStandardInterfaces() {
	fmt.Println("\n=== Exercise 3: Standard Interfaces ===")

	// TODO: Create a slice of products with sample data
	// YOUR CODE HERE

	// TODO: Print products before sorting (using Stringer)
	// YOUR CODE HERE

	// TODO: Sort products using sort.Sort()
	// YOUR CODE HERE

	// TODO: Print products after sorting
	// YOUR CODE HERE
}

// ===== EXERCISE 4: CUSTOM READER/WRITER =====

// TODO: Implement a ROT13Reader that wraps an io.Reader
// It should apply ROT13 encoding to the text as it's read
// ROT13: A→N, B→O, C→P, ..., N→A, O→B, P→C, etc.
// YOUR CODE HERE

// TODO: Implement a CountingWriter that wraps an io.Writer
// It should count the number of bytes written and provide a Count() method
// YOUR CODE HERE

func exerciseCustomReaderWriter() {
	fmt.Println("\n=== Exercise 4: Custom Reader/Writer ===")

	// TODO: Test ROT13Reader with a string input
	// Example: "Hello World" should become "Uryyb Jbeyq"
	// YOUR CODE HERE

	// TODO: Test CountingWriter by writing some data and checking count
	// YOUR CODE HERE

	// TODO: Chain them together: write ROT13 encoded data and count bytes
	// YOUR CODE HERE
}

// ===== EXERCISE 5: ERROR INTERFACE =====

// TODO: Create custom error types:
// 1. ValidationError with Field and Reason
// 2. NetworkError with URL, StatusCode, and Message
// 3. TimeoutError with Operation and Duration
// All should implement the error interface
// YOUR CODE HERE

// TODO: Create a function processRequest(url string, timeout time.Duration) error
// It should return different error types based on conditions:
// - ValidationError if url is empty
// - TimeoutError if timeout < 1 second
// - NetworkError if url doesn't start with "http"
// - nil if everything is ok
// YOUR CODE HERE

func exerciseErrorInterface() {
	fmt.Println("\n=== Exercise 5: Error Interface ===")

	// TODO: Test processRequest with different inputs to trigger each error type
	// YOUR CODE HERE

	// TODO: Use type assertions or type switches to handle each error type differently
	// YOUR CODE HERE
}

// ===== EXERCISE 6: POLYMORPHISM PATTERN =====

// TODO: Design a payment system with:
// - PaymentProcessor interface with Process(amount float64) error method
// - CreditCardProcessor struct with CardNumber field
// - PayPalProcessor struct with Email field
// - BankTransferProcessor struct with AccountNumber field
// YOUR CODE HERE

// TODO: Create a PaymentService struct that uses a PaymentProcessor
// It should have a SetProcessor method to change processors dynamically
// YOUR CODE HERE

func exercisePolymorphism() {
	fmt.Println("\n=== Exercise 6: Polymorphism Pattern ===")

	// TODO: Create different payment processors
	// YOUR CODE HERE

	// TODO: Create a payment service and test with different processors
	// YOUR CODE HERE

	// TODO: Show that the same PaymentService can use different processors
	// YOUR CODE HERE
}

// ===== EXERCISE 7: INTERFACE DESIGN PATTERNS =====

// TODO: Implement the Observer pattern:
// - Observer interface with Update(message string) method
// - Subject interface with Attach/Detach/Notify methods
// - NewsPublisher struct implementing Subject
// - Subscriber struct implementing Observer with Name field
// YOUR CODE HERE

func exerciseObserverPattern() {
	fmt.Println("\n=== Exercise 7: Observer Pattern ===")

	// TODO: Create a news publisher
	// YOUR CODE HERE

	// TODO: Create several subscribers and attach them
	// YOUR CODE HERE

	// TODO: Publish some news and show all subscribers receive it
	// YOUR CODE HERE

	// TODO: Detach one subscriber and publish again
	// YOUR CODE HERE
}

// ===== EXERCISE 8: ADVANCED INTERFACE USAGE =====

// TODO: Create a generic data transformer system:
// - Transformer interface with Transform(data interface{}) (interface{}, error)
// - StringUpperTransformer that converts strings to uppercase
// - NumberDoubler that doubles numeric values (int, float64)
// - SliceReverser that reverses any slice
// Use type assertions and type switches
// YOUR CODE HERE

// TODO: Create a Pipeline struct that chains multiple transformers
// It should have Add(transformer Transformer) and Process(data interface{}) methods
// YOUR CODE HERE

func exerciseAdvancedInterfaces() {
	fmt.Println("\n=== Exercise 8: Advanced Interface Usage ===")

	// TODO: Create individual transformers
	// YOUR CODE HERE

	// TODO: Create a pipeline and add transformers
	// YOUR CODE HERE

	// TODO: Test the pipeline with different data types
	// YOUR CODE HERE
}

// ===== EXERCISE 9: INTERFACE BEST PRACTICES =====

// TODO: Design a logging system following interface best practices:
// - Small, focused interfaces (Logger, Formatter, Writer)
// - Interface composition for complex behaviors
// - Accept interfaces, return structs principle
// YOUR CODE HERE

func exerciseBestPractices() {
	fmt.Println("\n=== Exercise 9: Interface Best Practices ===")

	// TODO: Demonstrate small, focused interfaces
	// YOUR CODE HERE

	// TODO: Show interface composition
	// YOUR CODE HERE

	// TODO: Apply "accept interfaces, return structs" principle
	// YOUR CODE HERE
}

// ===== EXERCISE 10: REAL-WORLD SCENARIO =====

// TODO: Build a simple HTTP-like system:
// - Handler interface with Handle(request Request) Response
// - Request struct with Method, URL, Body
// - Response struct with StatusCode, Body
// - Different handlers: StaticFileHandler, APIHandler, NotFoundHandler
// - Router that selects appropriate handler based on URL pattern
// YOUR CODE HERE

func exerciseRealWorldScenario() {
	fmt.Println("\n=== Exercise 10: Real-World Scenario ===")

	// TODO: Create different handlers
	// YOUR CODE HERE

	// TODO: Create a router and register handlers
	// YOUR CODE HERE

	// TODO: Simulate handling different requests
	// YOUR CODE HERE
}

// ===== HELPER FUNCTIONS =====

// TODO: Implement helper functions as needed for your exercises
// YOUR CODE HERE

// ===== MAIN FUNCTION =====

func main() {
	fmt.Println("🔌 Week 6: Interfaces & Polymorphism Practice")
	fmt.Println("==============================================")

	exerciseBasicInterface()
	exerciseInterfaceComposition()
	exerciseStandardInterfaces()
	exerciseCustomReaderWriter()
	exerciseErrorInterface()
	exercisePolymorphism()
	exerciseObserverPattern()
	exerciseAdvancedInterfaces()
	exerciseBestPractices()
	exerciseRealWorldScenario()

	fmt.Println("\n✅ Week 6 exercises completed!")
	fmt.Println("\n💡 Key Learnings:")
	fmt.Println("- Interfaces enable polymorphism and flexible design")
	fmt.Println("- Implicit implementation makes interfaces natural in Go")
	fmt.Println("- Standard library interfaces integrate seamlessly")
	fmt.Println("- Small, focused interfaces are easier to implement")
	fmt.Println("- Interface composition builds complex behaviors")
	fmt.Println("- Type assertions and switches handle runtime types")
	fmt.Println("- Design patterns become natural with interfaces")
	fmt.Println("- 'Accept interfaces, return structs' enables flexibility")
}
