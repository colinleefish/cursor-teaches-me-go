package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Exercise 1: Basic If/Else Statements
func exerciseIfElse() {
	fmt.Println("=== Exercise 1: If/Else Statements ===")

	// TODO: Create a variable 'age' with value 25
	// YOUR CODE HERE

	// TODO: Check if age is >= 18 and print "Adult" or "Minor"
	// Hint: No parentheses needed around condition
	// YOUR CODE HERE

	// TODO: Create a score variable with value 85
	// YOUR CODE HERE

	// TODO: Use if/else if/else to print grade based on score:
	// 90-100: "A", 80-89: "B", 70-79: "C", 60-69: "D", below 60: "F"
	// YOUR CODE HERE

	// TODO: Use if with initialization: if x := getValue(); x > 50 { ... }
	// Create a function getValue() that returns a random number 1-100
	// YOUR CODE HERE

	fmt.Printf("Random value check result\n")
}

// TODO: Implement getValue() function for the above exercise
// Hint: Use rand.Intn(100) + 1
// YOUR CODE HERE

// Exercise 2: For Loop Patterns
func exerciseForLoops() {
	fmt.Println("\n=== Exercise 2: For Loop Patterns ===")

	// TODO: Pattern 1 - Traditional C-style for loop
	// Print numbers 1 to 5
	fmt.Println("C-style loop (1-5):")
	// YOUR CODE HERE

	// TODO: Pattern 2 - While-style loop
	// Print numbers that double each time: 1, 2, 4, 8, 16 (stop before 32)
	fmt.Println("While-style loop (powers of 2):")
	// YOUR CODE HERE

	// TODO: Pattern 3 - Infinite loop with break
	// Print numbers 1-3, then break
	fmt.Println("Infinite loop with break:")
	// YOUR CODE HERE

	// TODO: Pattern 4 - Range loop over slice
	// Create a slice of strings and print index and value
	colors := []string{"red", "green", "blue"}
	fmt.Println("Range loop over slice:")
	_ = colors // Remove this line and use colors in your loop
	// YOUR CODE HERE

	// TODO: Pattern 5 - Range loop over map
	// Create a map of student names to grades and print them
	fmt.Println("Range loop over map:")
	// YOUR CODE HERE

	// TODO: Pattern 6 - Range loop with continue
	// Skip printing even numbers from 1-10
	fmt.Println("Loop with continue (odd numbers only):")
	// YOUR CODE HERE
}

// Exercise 3: Switch Statements
func exerciseSwitchStatements() {
	fmt.Println("\n=== Exercise 3: Switch Statements ===")

	// TODO: Basic switch - check day of week
	day := "Monday"
	fmt.Printf("Day: %s - ", day)
	// Switch on day and print if it's "Weekday" or "Weekend"
	// YOUR CODE HERE

	// TODO: Switch with multiple values
	month := "December"
	fmt.Printf("Month: %s - ", month)
	// Switch to print season: Dec/Jan/Feb = "Winter", etc.
	// YOUR CODE HERE

	// TODO: Switch without expression (replaces if/else chain)
	score := 85
	fmt.Printf("Score: %d - ", score)
	// Use switch without expression to determine grade
	// YOUR CODE HERE

	// TODO: Type switch (advanced)
	// Create an interface{} variable and use type switch
	var value interface{} = 42
	fmt.Printf("Value: %v - Type: ", value)
	// YOUR CODE HERE
}

// Exercise 4: Defer Statements
func exerciseDefer() {
	fmt.Println("\n=== Exercise 4: Defer Statements ===")

	// TODO: Basic defer - prints happen in reverse order
	fmt.Println("Defer execution order:")
	// YOUR CODE HERE - Add 3 defer statements that print "First", "Second", "Third"

	// TODO: Defer with variables - capture current value vs reference
	fmt.Println("Defer with variables:")
	// YOUR CODE HERE

	// TODO: Defer for resource cleanup simulation
	fmt.Println("Resource cleanup simulation:")
	// YOUR CODE HERE - simulate opening/closing a file or connection

	fmt.Println("Function ending...")
}

// TODO: Create a function to simulate file operations with defer
// YOUR CODE HERE

// Exercise 5: Real-World Scenarios
func exerciseRealWorld() {
	fmt.Println("\n=== Exercise 5: Real-World Scenarios ===")

	// TODO: Number guessing game
	// Generate random number 1-10, let user guess (simulate with fixed guess)
	target := rand.Intn(10) + 1
	guess := 7 // Simulate user guess

	fmt.Printf("Target: %d, Guess: %d - ", target, guess)
	// YOUR CODE HERE - Use if/else to check if correct, too high, or too low

	// TODO: Grade calculator with validation
	// Calculate average and assign letter grade
	scores := []int{85, 92, 78, 88, 95}
	fmt.Printf("Scores: %v - ", scores)
	// YOUR CODE HERE - Calculate average using loop, then assign grade using switch

	// TODO: FizzBuzz (classic programming problem)
	// Print numbers 1-15, but replace multiples of 3 with "Fizz",
	// multiples of 5 with "Buzz", and multiples of both with "FizzBuzz"
	fmt.Println("FizzBuzz (1-15):")
	// YOUR CODE HERE

	// TODO: Input validation pattern
	// Validate an email address (simple check for @ symbol)
	emails := []string{"user@example.com", "invalid-email", "test@test.org"}
	fmt.Println("Email validation:")
	_ = emails // Remove this line and use emails in your loop
	// YOUR CODE HERE - Use loop and if/else to validate each email
}

// Exercise 6: Combining Control Flow
func exerciseCombined() {
	fmt.Println("\n=== Exercise 6: Combined Control Flow ===")

	// TODO: Menu system simulation
	// Create a simple menu system using switch and loops
	fmt.Println("Menu System:")
	options := []string{"View Profile", "Edit Settings", "Logout", "Exit"}
	_ = options // Remove this line and use options in your implementation
	// YOUR CODE HERE - Display menu options and simulate user selection

	// TODO: Data processing pipeline
	// Process a slice of numbers: filter evens, square them, sum the result
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Input numbers: %v\n", numbers)

	// YOUR CODE HERE - Use loops and conditions to process the data

	// TODO: Error simulation with goto (just to show it exists)
	// Create a retry mechanism using goto (though this is not recommended)
	fmt.Println("Goto example (not recommended):")
	// YOUR CODE HERE
}

// Main function to run all exercises
func main() {
	fmt.Println("🚀 Go Control Flow Practice")
	fmt.Println("==========================")

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	exerciseIfElse()
	exerciseForLoops()
	exerciseSwitchStatements()
	exerciseDefer()
	exerciseRealWorld()
	exerciseCombined()

	fmt.Println("\n✅ Control flow exercises completed!")
	fmt.Println("\n💡 Key Takeaways:")
	fmt.Println("- Go's for loop is incredibly flexible")
	fmt.Println("- Switch statements don't fallthrough by default")
	fmt.Println("- Defer is perfect for cleanup and runs in LIFO order")
	fmt.Println("- No parentheses needed around conditions")
	fmt.Println("- Control flow can be combined for powerful patterns")
}
