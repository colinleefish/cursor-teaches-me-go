package main

import (
	"fmt"
)

// ===== fmt.Stringer INTERFACE =====

// TODO: Understand fmt.Stringer interface
// The fmt.Stringer interface has one method: String() string
// When you implement this, fmt.Print, fmt.Printf %s, %v will use your String() method

// TODO: Define a PersonStringer struct with Name (string), Age (int), City (string) fields
// YOUR CODE HERE

// TODO: Implement String() method for PersonStringer
// Should return "[Name] ([Age] years old) from [City]"
// YOUR CODE HERE

// TODO: Define a BankAccount struct with AccountNumber (string), Balance (float64), Owner (string)
// YOUR CODE HERE

// TODO: Implement String() method for BankAccount
// Mask account number for security (show only last 4 digits)
// Should return "Account ****[last4] - Balance: $[balance] (Owner: [owner])"
// Hint: Use strings.Repeat("*", count) to create asterisks
// YOUR CODE HERE

func demonstrateStringer() {
	fmt.Println("=== fmt.Stringer Interface ===")

	// TODO: Create a PersonStringer instance with sample data
	// YOUR CODE HERE

	// TODO: Create a BankAccount instance with sample data
	// YOUR CODE HERE

	// TODO: Print both using %s (this will call String() method automatically)
	// YOUR CODE HERE

	// TODO: Print both using %v (this also calls String() method)
	// YOUR CODE HERE

	// TODO: Create a slice of fmt.Stringer containing both person and account
	// Loop through and print each item
	// YOUR CODE HERE
}

// ===== io.Reader INTERFACE =====

// TODO: Understand io.Reader interface
// The io.Reader interface has one method: Read([]byte) (n int, err error)
// It reads data into the provided byte slice and returns bytes read and error
// Return io.EOF when no more data to read

// TODO: Define a StringReader struct with data (string) and position (int) fields
// YOUR CODE HERE

// TODO: Implement NewStringReader constructor function
// Should return *StringReader with given data and position 0
// YOUR CODE HERE

// TODO: Implement Read method for StringReader (use pointer receiver)
// - Return 0, io.EOF if position >= len(data)
// - Use copy(p, sr.data[sr.position:]) to copy data
// - Update position and return bytes copied
// YOUR CODE HERE

// TODO: Define UpperCaseReader struct that wraps another io.Reader
// Should have reader field of type io.Reader
// YOUR CODE HERE

// TODO: Implement NewUpperCaseReader constructor
// YOUR CODE HERE

// TODO: Implement Read method for UpperCaseReader
// - Call wrapped reader's Read method first
// - Convert lowercase letters to uppercase in the byte slice
// - Check if p[i] >= 'a' && p[i] <= 'z', then p[i] = p[i] - 'a' + 'A'
// YOUR CODE HERE

func demonstrateReader() {
	fmt.Println("\n=== io.Reader Interface ===")

	// TODO: Create a StringReader with some sample text
	// YOUR CODE HERE

	// TODO: Wrap it with UpperCaseReader
	// YOUR CODE HERE

	// TODO: Use io.ReadAll to read all data and handle error
	// Print original text and uppercase result
	// YOUR CODE HERE

	// TODO: Demonstrate with standard library strings.NewReader
	// Use io.ReadAll to read the data
	// YOUR CODE HERE

	// TODO: Chain readers - wrap strings.NewReader with UpperCaseReader
	// Read and print the result
	// YOUR CODE HERE
}

// ===== io.Writer INTERFACE =====

// TODO: Understand io.Writer interface
// The io.Writer interface has one method: Write([]byte) (n int, err error)
// It writes data from byte slice and returns bytes written and error

// TODO: Define LogWriter struct with prefix (string) field
// YOUR CODE HERE

// TODO: Implement NewLogWriter constructor
// YOUR CODE HERE

// TODO: Implement Write method for LogWriter
// - Get timestamp with time.Now().Format("15:04:05")
// - Format message as "[timestamp] prefix: content"
// - Print the message and return len(p), nil
// YOUR CODE HERE

// TODO: Define MemoryWriter struct with buffer ([]byte) field
// YOUR CODE HERE

// TODO: Implement NewMemoryWriter constructor
// Initialize with empty byte slice using make([]byte, 0)
// YOUR CODE HERE

// TODO: Implement Write method for MemoryWriter
// Append p to buffer using append(mw.buffer, p...)
// Return len(p), nil
// YOUR CODE HERE

// TODO: Implement String() method for MemoryWriter
// Convert buffer to string and return
// YOUR CODE HERE

// TODO: Implement Reset() method for MemoryWriter
// Reset buffer to empty: mw.buffer = mw.buffer[:0]
// YOUR CODE HERE

func demonstrateWriter() {
	fmt.Println("\n=== io.Writer Interface ===")

	// TODO: Create a LogWriter with prefix "APP"
	// YOUR CODE HERE

	// TODO: Write some log messages using Write method
	// Convert strings to []byte for writing
	// YOUR CODE HERE

	// TODO: Create a MemoryWriter and write some data to it
	// YOUR CODE HERE

	// TODO: Print the memory buffer contents using String() method
	// YOUR CODE HERE

	// TODO: Demonstrate fmt.Fprintf with both writers
	// fmt.Fprintf accepts io.Writer as first argument
	// YOUR CODE HERE
}

// ===== io.ReadWriter INTERFACE =====

// TODO: Define InMemoryFile struct with content ([]byte) and position (int) fields
// This will implement both io.Reader and io.Writer (so it's an io.ReadWriter)
// YOUR CODE HERE

// TODO: Implement NewInMemoryFile constructor
// Initialize with empty content slice and position 0
// YOUR CODE HERE

// TODO: Implement Read method for InMemoryFile (similar to StringReader)
// - Return 0, io.EOF if position >= len(content)
// - Copy data from content[position:] to p
// - Update position and return bytes copied
// YOUR CODE HERE

// TODO: Implement Write method for InMemoryFile
// - Append p to content using append(imf.content, p...)
// - Return len(p), nil
// YOUR CODE HERE

// TODO: Implement Seek method for InMemoryFile
// Parameters: offset (int64), whence (int)
// Handle io.SeekStart, io.SeekCurrent, io.SeekEnd
// Clamp position between 0 and len(content)
// Return final position as int64
// YOUR CODE HERE

func demonstrateReadWriter() {
	fmt.Println("\n=== io.ReadWriter Interface ===")

	// TODO: Create a new InMemoryFile
	// YOUR CODE HERE

	// TODO: Write several lines of data to the file
	// YOUR CODE HERE

	// TODO: Reset position to beginning using Seek(0, io.SeekStart)
	// YOUR CODE HERE

	// TODO: Read all data using io.ReadAll and print it
	// Handle the error properly
	// YOUR CODE HERE

	// TODO: Append more data to the file
	// YOUR CODE HERE

	// TODO: Reset to beginning and read all data again
	// Show that new data was appended
	// YOUR CODE HERE
}

// ===== sort.Interface =====

// sort.Interface enables custom sorting
// type Interface interface {
//     Len() int
//     Less(i, j int) bool
//     Swap(i, j int)
// }

// TODO: Define Student struct with Name (string), Grade (float64), Age (int) fields
// YOUR CODE HERE

// TODO: Define ByGrade type as []Student and implement sort.Interface
// - Len() should return length of slice
// - Less(i, j) should sort by Grade in descending order
// - Swap(i, j) should swap elements i and j
// YOUR CODE HERE

// TODO: Define ByAge type as []Student and implement sort.Interface
// - Len() should return length of slice
// - Less(i, j) should sort by Age in ascending order
// - Swap(i, j) should swap elements i and j
// YOUR CODE HERE

// TODO: Define ByName type as []Student and implement sort.Interface
// - Len() should return length of slice
// - Less(i, j) should sort by Name alphabetically
// - Swap(i, j) should swap elements i and j
// YOUR CODE HERE

func demonstrateSortInterface() {
	fmt.Println("\n=== sort.Interface ===")

	// TODO: Create a slice of students with sample data
	// YOUR CODE HERE

	// TODO: Print original students
	// YOUR CODE HERE

	// TODO: Sort by grade (descending) and print
	// YOUR CODE HERE

	// TODO: Sort by age (ascending) and print
	// YOUR CODE HERE

	// TODO: Sort by name (alphabetical) and print
	// YOUR CODE HERE
}

// Helper function to print students
// TODO: Implement printStudents function
// YOUR CODE HERE

// ===== error INTERFACE =====

// error is a built-in interface
// type error interface {
//     Error() string
// }

// TODO: Define ValidationErrorCustom struct with Field (string), Value (interface{}), Message (string) fields
// YOUR CODE HERE

// TODO: Implement Error() method for ValidationErrorCustom
// Should return "validation failed for field '[Field]' with value '[Value]': [Message]"
// YOUR CODE HERE

// TODO: Define NetworkErrorCustom struct with Operation (string), URL (string), Code (int), Timestamp (time.Time) fields
// YOUR CODE HERE

// TODO: Implement Error() method for NetworkErrorCustom
// Should return "network error during [Operation] to [URL]: HTTP [Code] at [Timestamp]"
// Format timestamp using Format("15:04:05")
// YOUR CODE HERE

// TODO: Implement String() method for NetworkErrorCustom (fmt.Stringer)
// Should return "NetworkError{Op: [Operation], URL: [URL], Code: [Code]}"
// YOUR CODE HERE

func demonstrateErrorInterface() {
	fmt.Println("\n=== error Interface ===")

	// TODO: Create sample validation and network errors
	// YOUR CODE HERE

	// TODO: Create slice of error interface containing both errors
	// YOUR CODE HERE

	// TODO: Range over errors and print each one
	// YOUR CODE HERE

	// TODO: Use type assertion to handle each error type specifically
	// For ValidationErrorCustom: Print field name
	// For NetworkErrorCustom: Print status code and use String() method
	// YOUR CODE HERE
}

// ===== COMBINING INTERFACES =====

// TODO: Define Logger interface with Log(message string) method
// YOUR CODE HERE

// TODO: Define FileSystemLogger interface that combines Logger and io.Writer
// Also add SetLogLevel(level string) method
// YOUR CODE HERE

// TODO: Define ConsoleFileLogger struct with logLevel (string) and buffer (*bytes.Buffer) fields
// YOUR CODE HERE

// TODO: Implement NewConsoleFileLogger constructor
// Initialize with "INFO" level and new buffer
// YOUR CODE HERE

// TODO: Implement Log method for ConsoleFileLogger
// - Format with timestamp, level, message
// - Write to console and buffer
// YOUR CODE HERE

// TODO: Implement Write method for ConsoleFileLogger (io.Writer)
// - Convert bytes to string, trim space
// - Call Log method
// - Return length and nil error
// YOUR CODE HERE

// TODO: Implement SetLogLevel method
// YOUR CODE HERE

// TODO: Implement GetBufferedLogs method to return buffer contents
// YOUR CODE HERE

func demonstrateInterfaceCombination() {
	fmt.Println("\n=== Combining Interfaces ===")

	// TODO: Create new logger
	// YOUR CODE HERE

	// TODO: Use as Logger interface
	// YOUR CODE HERE

	// TODO: Use as io.Writer interface with fmt.Fprintf
	// YOUR CODE HERE

	// TODO: Use as FileSystemLogger interface
	// YOUR CODE HERE

	// TODO: Show buffered logs
	// YOUR CODE HERE
}

// ===== INTERFACE COMPOSITION PATTERNS =====

// TODO: Define Processor interface with Process(data string) (string, error) method
// YOUR CODE HERE

// TODO: Define Validator interface with Validate(data string) error method
// YOUR CODE HERE

// TODO: Define ProcessorValidator interface combining both interfaces
// YOUR CODE HERE

// TODO: Define EmailProcessor struct with domain (string) field
// YOUR CODE HERE

// TODO: Implement NewEmailProcessor constructor
// YOUR CODE HERE

// TODO: Implement Validate method for EmailProcessor
// - Check if email contains @
// - Check if email ends with correct domain
// - Return appropriate errors
// YOUR CODE HERE

// TODO: Implement Process method for EmailProcessor
// - Call Validate first
// - If valid, normalize email (lowercase, trim space)
// - Return result or error
// YOUR CODE HERE

func demonstrateInterfaceCompositionPatterns() {
	fmt.Println("\n=== Interface Composition Patterns ===")

	// TODO: Create email processor with "@company.com" domain
	// YOUR CODE HERE

	// TODO: Create slice of test emails (valid and invalid)
	// YOUR CODE HERE

	// TODO: Process each email:
	// - Use as ProcessorValidator interface
	// - Validate first
	// - If valid, process and show result
	// - Handle errors appropriately
	// YOUR CODE HERE
}

func main() {
	fmt.Println("📚 Go Common Standard Library Interfaces Practice")
	fmt.Println("=================================================")

	demonstrateStringer()
	demonstrateReader()
	demonstrateWriter()
	demonstrateReadWriter()
	// TODO: Uncomment these as you implement them
	// demonstrateSortInterface()
	// demonstrateErrorInterface()
	// demonstrateInterfaceCombination()
	// demonstrateInterfaceCompositionPatterns()

	fmt.Println("\n✅ Common interfaces practice completed!")
	fmt.Println("\n🎯 Learning Goals:")
	fmt.Println("- Implement fmt.Stringer for custom string representation")
	fmt.Println("- Create io.Reader/Writer for flexible data handling")
	fmt.Println("- Use sort.Interface for custom sorting logic")
	fmt.Println("- Implement error interface for rich error information")
	fmt.Println("- Apply interface composition for powerful abstractions")
	fmt.Println("- Make types integrate with Go's standard library")
}
