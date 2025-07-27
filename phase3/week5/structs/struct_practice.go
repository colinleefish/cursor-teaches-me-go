package main

import (
	"fmt"
)

// ===== EXERCISE 1: BASIC STRUCT DEFINITION =====

// TODO: Define a Book struct with the following fields:
// - Title (string)
// - Author (string)
// - Pages (int)
// - PublishedYear (int)
// - InStock (bool)
// YOUR CODE HERE

func exerciseBasicStruct() {
	fmt.Println("=== Exercise 1: Basic Struct Definition ===")

	// TODO: Create a Book instance with sample data
	// YOUR CODE HERE

	// TODO: Print the book information
	// YOUR CODE HERE

	// TODO: Modify some fields and print again
	// YOUR CODE HERE
}

// ===== EXERCISE 2: METHODS AND RECEIVERS =====

// TODO: Add methods to your Book struct:
// 1. String() string - returns formatted book information
// 2. IsClassic() bool - returns true if published before 1980
// 3. UpdateStock(inStock bool) - updates the InStock field (use pointer receiver)
// 4. PageCategory() string - returns "Short" (<200), "Medium" (200-500), "Long" (>500)
// YOUR CODE HERE

func exerciseMethodsAndReceivers() {
	fmt.Println("\n=== Exercise 2: Methods and Receivers ===")

	// TODO: Create a book and test all your methods
	// YOUR CODE HERE

	// TODO: Test the UpdateStock method (should modify the original)
	// YOUR CODE HERE
}

// ===== EXERCISE 3: STRUCT EMBEDDING =====

// TODO: Define a MediaItem struct with common fields:
// - Title (string)
// - Creator (string)
// - Year (int)
// YOUR CODE HERE

// TODO: Define a Movie struct that embeds MediaItem and adds:
// - Duration (int) // in minutes
// - Genre (string)
// YOUR CODE HERE

// TODO: Define a Song struct that embeds MediaItem and adds:
// - Duration (int) // in seconds
// - Album (string)
// YOUR CODE HERE

// TODO: Add methods:
// - (MediaItem) Info() string - returns basic info
// - (Movie) Play() string - returns "Playing movie: [title]"
// - (Song) Play() string - returns "Playing song: [title]"
// YOUR CODE HERE

func exerciseStructEmbedding() {
	fmt.Println("\n=== Exercise 3: Struct Embedding ===")

	// TODO: Create instances of Movie and Song
	// YOUR CODE HERE

	// TODO: Demonstrate accessing embedded fields directly
	// YOUR CODE HERE

	// TODO: Demonstrate calling embedded methods
	// YOUR CODE HERE

	// TODO: Demonstrate method shadowing if you implemented Play() differently
	// YOUR CODE HERE
}

// ===== EXERCISE 4: STRUCT TAGS AND JSON =====

// TODO: Define a Product struct with JSON tags:
// - ID (int) -> "id"
// - Name (string) -> "product_name"
// - Price (float64) -> "price" with string conversion
// - Category (string) -> "category"
// - InStock (bool) -> "in_stock"
// - Description (string) -> "description" with omitempty
// - InternalNotes (string) -> excluded from JSON (use "-")
// YOUR CODE HERE

func exerciseStructTags() {
	fmt.Println("\n=== Exercise 4: Struct Tags and JSON ===")

	// TODO: Create a Product instance with sample data
	// YOUR CODE HERE

	// TODO: Marshal to JSON and print
	// YOUR CODE HERE

	// TODO: Unmarshal from this JSON string:
	jsonData := `{"id":2,"product_name":"Laptop","price":"999.99","category":"Electronics","in_stock":true}`
	_ = jsonData // Remove this line when implementing
	// YOUR CODE HERE

	// TODO: Print the unmarshaled product
	// YOUR CODE HERE
}

// ===== EXERCISE 5: COMPOSITION PATTERN =====

// TODO: Define interfaces:
// - Writer interface with Write(content string) error method
// - Reader interface with Read() (string, error) method
// YOUR CODE HERE

// TODO: Define a FileManager struct that embeds both Writer and Reader interfaces
// and adds a Filename field
// YOUR CODE HERE

// TODO: Implement concrete types:
// - ConsoleWriter that prints to console
// - FileWriter that simulates writing to a file
// - MemoryReader that returns predefined content
// YOUR CODE HERE

func exerciseComposition() {
	fmt.Println("\n=== Exercise 5: Composition Pattern ===")

	// TODO: Create a FileManager with ConsoleWriter and MemoryReader
	// YOUR CODE HERE

	// TODO: Test reading and writing
	// YOUR CODE HERE

	// TODO: Create another FileManager with FileWriter
	// YOUR CODE HERE
}

// ===== EXERCISE 6: ADVANCED STRUCT PATTERNS =====

// TODO: Define a Customer struct with:
// - ID (int)
// - Name (string)
// - Email (string)
// - CreatedAt (time.Time)
// - Address (embedded struct with Street, City, ZipCode)
// - Orders (slice of Order structs)
// YOUR CODE HERE

// TODO: Define an Order struct with:
// - ID (int)
// - Total (float64)
// - Items (slice of strings)
// - OrderDate (time.Time)
// YOUR CODE HERE

// TODO: Add methods:
// - (Customer) AddOrder(order Order) - adds order to customer's orders
// - (Customer) TotalSpent() float64 - calculates total from all orders
// - (Customer) RecentOrders(days int) []Order - returns orders from last N days
// YOUR CODE HERE

func exerciseAdvancedPatterns() {
	fmt.Println("\n=== Exercise 6: Advanced Struct Patterns ===")

	// TODO: Create a customer with embedded address
	// YOUR CODE HERE

	// TODO: Create several orders and add them to the customer
	// YOUR CODE HERE

	// TODO: Calculate total spent
	// YOUR CODE HERE

	// TODO: Get recent orders (use time.Now().AddDate(0, 0, -7) for 7 days ago)
	// YOUR CODE HERE
}

// ===== EXERCISE 7: CONSTRUCTOR FUNCTIONS =====

// TODO: Create constructor functions:
// - NewBook(title, author string, pages, year int) (*Book, error)
//   - Validate that title and author are not empty
//   - Validate that pages > 0 and year > 1400
// - NewCustomer(name, email string) (*Customer, error)
//   - Validate email contains "@"
//   - Initialize empty orders slice and set CreatedAt to now
// YOUR CODE HERE

func exerciseConstructors() {
	fmt.Println("\n=== Exercise 7: Constructor Functions ===")

	// TODO: Test valid book creation
	// YOUR CODE HERE

	// TODO: Test invalid book creation (should return error)
	// YOUR CODE HERE

	// TODO: Test valid customer creation
	// YOUR CODE HERE

	// TODO: Test invalid customer creation (should return error)
	// YOUR CODE HERE
}

// ===== MAIN FUNCTION =====

func main() {
	fmt.Println("🏗️ Week 5: Structs & Methods Practice")
	fmt.Println("=====================================")

	exerciseBasicStruct()
	exerciseMethodsAndReceivers()
	exerciseStructEmbedding()
	exerciseStructTags()
	exerciseComposition()
	exerciseAdvancedPatterns()
	exerciseConstructors()

	fmt.Println("\n✅ Week 5 exercises completed!")
	fmt.Println("\n💡 Key Learnings:")
	fmt.Println("- Structs define custom data types")
	fmt.Println("- Methods attach behavior to types")
	fmt.Println("- Pointer receivers modify originals, value receivers work on copies")
	fmt.Println("- Embedding enables composition over inheritance")
	fmt.Println("- Struct tags provide metadata for serialization")
	fmt.Println("- Constructor functions enable validation and initialization")
	fmt.Println("- Composition patterns create flexible, reusable designs")
}
