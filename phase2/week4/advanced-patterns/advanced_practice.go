package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Exercise 1: Recursive Functions
func exerciseRecursion() {
	fmt.Println("=== Exercise 1: Recursive Functions ===")

	// TODO: Implement factorial(n int) (int, error)
	// Handle negative numbers with error, use recursion
	// YOUR CODE HERE

	// Test factorial
	fmt.Println("Factorial tests:")
	factorialTests := []int{0, 1, 5, -1, 10}
	for _, n := range factorialTests {
		result, err := factorial(n)
		if err != nil {
			fmt.Printf("factorial(%d) -> Error: %v\n", n, err)
		} else {
			fmt.Printf("factorial(%d) -> %d\n", n, result)
		}
	}

	// TODO: Implement fibonacci(n int) (int, error)
	// Classic fibonacci with proper error handling
	// YOUR CODE HERE

	fmt.Println("\nFibonacci tests:")
	fibTests := []int{0, 1, 2, 5, 8, -1, 15}
	for _, n := range fibTests {
		result, err := fibonacci(n)
		if err != nil {
			fmt.Printf("fibonacci(%d) -> Error: %v\n", n, err)
		} else {
			fmt.Printf("fibonacci(%d) -> %d\n", n, result)
		}
	}

	// TODO: Implement gcd(a, b int) int (Greatest Common Divisor)
	// Use Euclidean algorithm with recursion
	// YOUR CODE HERE

	fmt.Println("\nGCD tests:")
	gcdTests := [][2]int{{48, 18}, {100, 25}, {17, 13}, {0, 5}}
	for _, test := range gcdTests {
		result := gcd(test[0], test[1])
		fmt.Printf("gcd(%d, %d) -> %d\n", test[0], test[1], result)
	}

	// TODO: Implement sumDigits(n int) int
	// Sum all digits in a number using recursion
	// YOUR CODE HERE

	fmt.Println("\nSum digits tests:")
	digitTests := []int{123, 9876, 0, 1, 999}
	for _, n := range digitTests {
		result := sumDigits(n)
		fmt.Printf("sumDigits(%d) -> %d\n", n, result)
	}
}

// TODO: Implement factorial function
// YOUR CODE HERE
func factorial(n int) (int, error) {
	// YOUR CODE HERE
	return 0, errors.New("not implemented")
}

// TODO: Implement fibonacci function
// YOUR CODE HERE
func fibonacci(n int) (int, error) {
	// YOUR CODE HERE
	return 0, errors.New("not implemented")
}

// TODO: Implement gcd function
// YOUR CODE HERE
func gcd(a, b int) int {
	// YOUR CODE HERE
	return 0
}

// TODO: Implement sumDigits function
// YOUR CODE HERE
func sumDigits(n int) int {
	// YOUR CODE HERE
	return 0
}

// Exercise 2: Higher-Order Functions
func exerciseHigherOrder() {
	fmt.Println("\n=== Exercise 2: Higher-Order Functions ===")

	// TODO: Implement mapInt(slice []int, fn func(int) int) []int
	// Apply function to each element of slice
	// YOUR CODE HERE

	// TODO: Implement filterInt(slice []int, predicate func(int) bool) []int
	// Filter elements that satisfy the predicate
	// YOUR CODE HERE

	// TODO: Implement reduceInt(slice []int, initial int, fn func(int, int) int) int
	// Reduce slice to single value using function
	// YOUR CODE HERE

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Original numbers: %v\n", numbers)

	// Test map
	squared := mapInt(numbers, func(x int) int { return x * x })
	fmt.Printf("Squared: %v\n", squared)

	doubled := mapInt(numbers, func(x int) int { return x * 2 })
	fmt.Printf("Doubled: %v\n", doubled)

	// Test filter
	evens := filterInt(numbers, func(x int) bool { return x%2 == 0 })
	fmt.Printf("Evens: %v\n", evens)

	greaterThan5 := filterInt(numbers, func(x int) bool { return x > 5 })
	fmt.Printf("Greater than 5: %v\n", greaterThan5)

	// Test reduce
	sum := reduceInt(numbers, 0, func(acc, x int) int { return acc + x })
	fmt.Printf("Sum: %d\n", sum)

	product := reduceInt(numbers[:4], 1, func(acc, x int) int { return acc * x })
	fmt.Printf("Product of first 4: %d\n", product)

	// TODO: Implement findInt(slice []int, predicate func(int) bool) (int, bool)
	// Find first element that satisfies predicate
	// YOUR CODE HERE

	fmt.Println("\nFind tests:")
	firstEven, found := findInt(numbers, func(x int) bool { return x%2 == 0 })
	if found {
		fmt.Printf("First even number: %d\n", firstEven)
	}

	firstGreaterThan8, found := findInt(numbers, func(x int) bool { return x > 8 })
	if found {
		fmt.Printf("First number > 8: %d\n", firstGreaterThan8)
	}
}

// TODO: Implement mapInt function
// YOUR CODE HERE
func mapInt(slice []int, fn func(int) int) []int {
	// YOUR CODE HERE
	return nil
}

// TODO: Implement filterInt function
// YOUR CODE HERE
func filterInt(slice []int, predicate func(int) bool) []int {
	// YOUR CODE HERE
	return nil
}

// TODO: Implement reduceInt function
// YOUR CODE HERE
func reduceInt(slice []int, initial int, fn func(int, int) int) int {
	// YOUR CODE HERE
	return 0
}

// TODO: Implement findInt function
// YOUR CODE HERE
func findInt(slice []int, predicate func(int) bool) (int, bool) {
	// YOUR CODE HERE
	return 0, false
}

// Exercise 3: Function Composition
func exerciseFunctionComposition() {
	fmt.Println("\n=== Exercise 3: Function Composition ===")

	// TODO: Define function type Transform for int -> int transformations
	// YOUR CODE HERE

	// TODO: Implement compose function that combines two Transform functions
	// compose(f, g Transform) Transform where result(x) = f(g(x))
	// YOUR CODE HERE

	// TODO: Create a chain method for Transform type
	// Allow chaining multiple transformations: t1.Chain(t2).Chain(t3)
	// YOUR CODE HERE

	// Basic transformations
	addOne := Transform(func(x int) int { return x + 1 })
	double := Transform(func(x int) int { return x * 2 })
	square := Transform(func(x int) int { return x * x })

	// Test composition
	fmt.Println("Function composition tests:")

	// Compose two functions
	addOneThenDouble := compose(double, addOne)
	result1 := addOneThenDouble(5) // (5+1)*2 = 12
	fmt.Printf("addOneThenDouble(5) = %d\n", result1)

	doubleThenSquare := compose(square, double)
	result2 := doubleThenSquare(3) // (3*2)^2 = 36
	fmt.Printf("doubleThenSquare(3) = %d\n", result2)

	// Test chaining
	fmt.Println("\nFunction chaining tests:")
	pipeline := addOne.Chain(double).Chain(square)
	result3 := pipeline(2) // ((2+1)*2)^2 = 36
	fmt.Printf("addOne.Chain(double).Chain(square)(2) = %d\n", result3)

	// TODO: Implement pipe function for multiple transformations
	// pipe(transforms ...Transform) Transform
	// YOUR CODE HERE

	fmt.Println("\nPipeline tests:")
	multistep := pipe(addOne, double, square)
	result4 := multistep(2)
	fmt.Printf("pipe(addOne, double, square)(2) = %d\n", result4)
}

// TODO: Define Transform type
// YOUR CODE HERE
type Transform func(int) int

// TODO: Implement compose function
// YOUR CODE HERE
func compose(f, g Transform) Transform {
	// YOUR CODE HERE
	return func(x int) int { return 0 }
}

// TODO: Implement Chain method for Transform
// YOUR CODE HERE
func (t Transform) Chain(next Transform) Transform {
	// YOUR CODE HERE
	return func(x int) int { return 0 }
}

// TODO: Implement pipe function
// YOUR CODE HERE
func pipe(transforms ...Transform) Transform {
	// YOUR CODE HERE
	return func(x int) int { return 0 }
}

// Exercise 4: Memoization
func exerciseMemoization() {
	fmt.Println("\n=== Exercise 4: Memoization ===")

	// TODO: Implement memoize function for int -> int functions
	// memoize(fn func(int) int) func(int) int
	// YOUR CODE HERE

	// Create expensive fibonacci function
	var expensiveFib func(int) int
	expensiveFib = func(n int) int {
		fmt.Printf("Computing fibonacci(%d)\n", n)
		if n <= 1 {
			return n
		}
		return expensiveFib(n-1) + expensiveFib(n-2)
	}

	// Test without memoization
	fmt.Println("Without memoization:")
	start := time.Now()
	result1 := expensiveFib(10)
	duration1 := time.Since(start)
	fmt.Printf("fibonacci(10) = %d (took %v)\n\n", result1, duration1)

	// Test with memoization
	fmt.Println("With memoization:")
	memoizedFib := memoize(expensiveFib)

	start = time.Now()
	result2 := memoizedFib(10)
	duration2 := time.Since(start)
	fmt.Printf("First call: fibonacci(10) = %d (took %v)\n", result2, duration2)

	start = time.Now()
	result3 := memoizedFib(10)
	duration3 := time.Since(start)
	fmt.Printf("Second call: fibonacci(10) = %d (took %v)\n\n", result3, duration3)

	// TODO: Implement memoizeWithTTL for cache expiration
	// memoizeWithTTL(fn func(int) int, ttl time.Duration) func(int) int
	// YOUR CODE HERE

	fmt.Println("Memoization with TTL:")
	ttlMemoizedFib := memoizeWithTTL(expensiveFib, 2*time.Second)

	fmt.Printf("Call 1: %d\n", ttlMemoizedFib(8))
	fmt.Printf("Call 2 (cached): %d\n", ttlMemoizedFib(8))

	fmt.Println("Waiting for cache to expire...")
	time.Sleep(3 * time.Second)
	fmt.Printf("Call 3 (expired): %d\n", ttlMemoizedFib(8))
}

// TODO: Implement memoize function
// YOUR CODE HERE
func memoize(fn func(int) int) func(int) int {
	// YOUR CODE HERE
	return fn
}

// TODO: Implement memoizeWithTTL function
// YOUR CODE HERE
func memoizeWithTTL(fn func(int) int, ttl time.Duration) func(int) int {
	// YOUR CODE HERE
	_ = ttl // Remove this line when implementing
	return fn
}

// Exercise 5: Decorator Patterns
func exerciseDecorators() {
	fmt.Println("\n=== Exercise 5: Decorator Patterns ===")

	// TODO: Define DecoratedFunc type for functions that can be decorated
	// YOUR CODE HERE

	// TODO: Implement timing decorator
	// withTiming(fn DecoratedFunc) DecoratedFunc
	// YOUR CODE HERE

	// TODO: Implement logging decorator
	// withLogging(name string, fn DecoratedFunc) DecoratedFunc
	// YOUR CODE HERE

	// TODO: Implement retry decorator
	// withRetry(maxAttempts int, fn DecoratedFunc) DecoratedFunc
	// YOUR CODE HERE

	// Base function to decorate
	slowFunction := DecoratedFunc(func() error {
		time.Sleep(100 * time.Millisecond)
		return nil
	})

	// Test timing decorator
	fmt.Println("Testing timing decorator:")
	timedFunction := withTiming(slowFunction)
	timedFunction()

	// Test logging decorator
	fmt.Println("\nTesting logging decorator:")
	loggedFunction := withLogging("SlowOperation", slowFunction)
	loggedFunction()

	// Test combining decorators
	fmt.Println("\nTesting combined decorators:")
	decoratedFunction := withTiming(withLogging("CombinedOperation", slowFunction))
	decoratedFunction()

	// TODO: Implement error-prone function for retry testing
	attemptCount := 0
	flakyFunction := DecoratedFunc(func() error {
		attemptCount++
		if attemptCount < 3 {
			return errors.New("temporary failure")
		}
		attemptCount = 0 // Reset for next test
		return nil
	})

	fmt.Println("\nTesting retry decorator:")
	retryFunction := withRetry(5, flakyFunction)
	err := retryFunction()
	if err != nil {
		fmt.Printf("Function failed: %v\n", err)
	} else {
		fmt.Println("Function succeeded after retries")
	}
}

// TODO: Define DecoratedFunc type
// YOUR CODE HERE
type DecoratedFunc func() error

// TODO: Implement withTiming decorator
// YOUR CODE HERE
func withTiming(fn DecoratedFunc) DecoratedFunc {
	// YOUR CODE HERE
	return fn
}

// TODO: Implement withLogging decorator
// YOUR CODE HERE
func withLogging(name string, fn DecoratedFunc) DecoratedFunc {
	// YOUR CODE HERE
	_ = name // Remove this line when implementing
	return fn
}

// TODO: Implement withRetry decorator
// YOUR CODE HERE
func withRetry(maxAttempts int, fn DecoratedFunc) DecoratedFunc {
	// YOUR CODE HERE
	_ = maxAttempts // Remove this line when implementing
	return fn
}

// Exercise 6: Data Processing Pipelines
func exerciseDataPipelines() {
	fmt.Println("\n=== Exercise 6: Data Processing Pipelines ===")

	// TODO: Define Pipeline type for data processing
	// type Pipeline[T any] func([]T) []T (if using generics)
	// For now, use IntPipeline for simplicity
	// YOUR CODE HERE

	// TODO: Implement Then method for chaining pipelines
	// YOUR CODE HERE

	// TODO: Create basic pipeline operations
	// filterEvens, mapDouble, filterGreaterThan(threshold int) IntPipeline
	// YOUR CODE HERE

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Printf("Original data: %v\n", data)

	// Create pipeline operations
	// YOUR CODE HERE - implement these in the pipeline section

	// Build pipeline
	pipeline := filterEvens.
		Then(mapDouble).
		Then(filterGreaterThan(10))

	result := pipeline(data)
	fmt.Printf("Pipeline result: %v\n", result)

	// TODO: Implement parallel pipeline processing
	// parallelMap(data []int, fn func(int) int, workers int) []int
	// YOUR CODE HERE

	fmt.Println("\nParallel processing test:")
	expensiveOperation := func(x int) int {
		time.Sleep(10 * time.Millisecond) // Simulate work
		return x * x
	}

	smallData := []int{1, 2, 3, 4, 5}

	start := time.Now()
	sequentialResult := mapInt(smallData, expensiveOperation)
	sequentialTime := time.Since(start)

	start = time.Now()
	parallelResult := parallelMap(smallData, expensiveOperation, 3)
	parallelTime := time.Since(start)

	fmt.Printf("Sequential: %v (took %v)\n", sequentialResult, sequentialTime)
	fmt.Printf("Parallel: %v (took %v)\n", parallelResult, parallelTime)
}

// TODO: Define IntPipeline type
// YOUR CODE HERE
type IntPipeline func([]int) []int

// TODO: Implement Then method for IntPipeline
// YOUR CODE HERE
func (p IntPipeline) Then(next IntPipeline) IntPipeline {
	// YOUR CODE HERE
	return func(input []int) []int { return nil }
}

// TODO: Implement pipeline operations
// YOUR CODE HERE
var filterEvens = IntPipeline(func(input []int) []int {
	// YOUR CODE HERE
	return nil
})

var mapDouble = IntPipeline(func(input []int) []int {
	// YOUR CODE HERE
	return nil
})

func filterGreaterThan(threshold int) IntPipeline {
	// YOUR CODE HERE
	_ = threshold // Remove when implementing
	return func(input []int) []int { return nil }
}

// TODO: Implement parallelMap function
// YOUR CODE HERE
func parallelMap(data []int, fn func(int) int, workers int) []int {
	// YOUR CODE HERE
	_ = workers             // Remove when implementing
	return mapInt(data, fn) // Fallback to sequential for now
}

// Exercise 7: Tree Operations with Functions
func exerciseTreeOperations() {
	fmt.Println("\n=== Exercise 7: Tree Operations ===")

	// TODO: Define TreeNode struct
	// YOUR CODE HERE

	// TODO: Implement tree traversal methods
	// InOrder(visit func(int)), PreOrder(visit func(int)), PostOrder(visit func(int))
	// YOUR CODE HERE

	// TODO: Implement tree transformation functions
	// Map(fn func(int) int) *TreeNode, Filter(predicate func(int) bool) *TreeNode
	// YOUR CODE HERE

	// Create sample tree
	//       5
	//      / \
	//     3   8
	//    / \ / \
	//   2  4 7  9
	root := &TreeNode{
		Value: 5,
		Left: &TreeNode{
			Value: 3,
			Left:  &TreeNode{Value: 2},
			Right: &TreeNode{Value: 4},
		},
		Right: &TreeNode{
			Value: 8,
			Left:  &TreeNode{Value: 7},
			Right: &TreeNode{Value: 9},
		},
	}

	// Test traversals
	fmt.Println("Tree traversals:")

	fmt.Print("InOrder: ")
	root.InOrder(func(val int) { fmt.Printf("%d ", val) })
	fmt.Println()

	fmt.Print("PreOrder: ")
	root.PreOrder(func(val int) { fmt.Printf("%d ", val) })
	fmt.Println()

	fmt.Print("PostOrder: ")
	root.PostOrder(func(val int) { fmt.Printf("%d ", val) })
	fmt.Println()

	// Test tree transformations
	fmt.Println("\nTree transformations:")

	doubledTree := root.Map(func(x int) int { return x * 2 })
	fmt.Print("Doubled tree (InOrder): ")
	doubledTree.InOrder(func(val int) { fmt.Printf("%d ", val) })
	fmt.Println()

	evenTree := root.Filter(func(x int) bool { return x%2 == 0 })
	fmt.Print("Even values only (InOrder): ")
	if evenTree != nil {
		evenTree.InOrder(func(val int) { fmt.Printf("%d ", val) })
	}
	fmt.Println()
}

// TODO: Define TreeNode struct
// YOUR CODE HERE
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// TODO: Implement TreeNode methods
// YOUR CODE HERE
func (t *TreeNode) InOrder(visit func(int)) {
	// YOUR CODE HERE
}

func (t *TreeNode) PreOrder(visit func(int)) {
	// YOUR CODE HERE
}

func (t *TreeNode) PostOrder(visit func(int)) {
	// YOUR CODE HERE
}

func (t *TreeNode) Map(fn func(int) int) *TreeNode {
	// YOUR CODE HERE
	_ = fn // Remove when implementing
	return nil
}

func (t *TreeNode) Filter(predicate func(int) bool) *TreeNode {
	// YOUR CODE HERE
	_ = predicate // Remove when implementing
	return nil
}

// Exercise 8: Real-World Application
func exerciseRealWorld() {
	fmt.Println("\n=== Exercise 8: Real-World Application ===")

	// TODO: Build a text processing pipeline
	// Create functions for: splitWords, filterEmpty, toLowerCase, removeDuplicates
	// YOUR CODE HERE

	text := "Hello World! This is a Test. Hello Go programming. Go is AWESOME!"
	fmt.Printf("Original text: %s\n", text)

	// TODO: Create text processing pipeline
	textPipeline := func(text string) []string {
		// YOUR CODE HERE
		_ = strings.Fields // Remove when implementing
		return []string{}  // Replace with actual implementation
	}

	words := textPipeline(text)
	fmt.Printf("Processed words: %v\n", words)

	// TODO: Implement word frequency counter using higher-order functions
	// countWords(words []string) map[string]int
	// YOUR CODE HERE

	freq := countWords(words)
	fmt.Printf("Word frequencies: %v\n", freq)

	// TODO: Implement caching decorator for expensive operations
	// Create a function that simulates expensive API calls and cache results
	// YOUR CODE HERE

	fmt.Println("\nAPI call simulation with caching:")
	apiCall := func(id string) (string, error) {
		time.Sleep(100 * time.Millisecond) // Simulate network delay
		return fmt.Sprintf("data-%s", id), nil
	}

	cachedAPICall := cacheResults(apiCall)

	// Test caching
	ids := []string{"user1", "user2", "user1", "user3", "user1"}
	for _, id := range ids {
		start := time.Now()
		result, err := cachedAPICall(id)
		duration := time.Since(start)
		if err != nil {
			fmt.Printf("Error for %s: %v\n", id, err)
		} else {
			fmt.Printf("ID %s: %s (took %v)\n", id, result, duration)
		}
	}
}

// TODO: Implement text processing functions
// YOUR CODE HERE

// TODO: Implement countWords function
// YOUR CODE HERE
func countWords(words []string) map[string]int {
	// YOUR CODE HERE
	return nil
}

// TODO: Implement cacheResults function
// YOUR CODE HERE
func cacheResults(fn func(string) (string, error)) func(string) (string, error) {
	// YOUR CODE HERE
	return fn
}

// Main function to run all exercises
func main() {
	fmt.Println("🚀 Go Advanced Function Patterns Practice")
	fmt.Println("=========================================")

	exerciseRecursion()
	exerciseHigherOrder()
	exerciseFunctionComposition()
	exerciseMemoization()
	exerciseDecorators()
	exerciseDataPipelines()
	exerciseTreeOperations()
	exerciseRealWorld()

	fmt.Println("\n✅ Advanced pattern exercises completed!")
	fmt.Println("\n💡 Key Takeaways:")
	fmt.Println("- Recursion with proper base cases and error handling")
	fmt.Println("- Higher-order functions enable flexible, reusable code")
	fmt.Println("- Function composition builds complex operations from simple ones")
	fmt.Println("- Memoization can dramatically improve performance")
	fmt.Println("- Decorators add cross-cutting concerns without modifying core logic")
	fmt.Println("- Pipelines enable clean, readable data transformations")
	fmt.Println("- Functional patterns make code more predictable and testable")
	fmt.Println("- Go's explicit nature applies to advanced patterns too")
}
