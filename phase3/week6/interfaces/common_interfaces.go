package main

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
	"time"
)

// ===== fmt.Stringer INTERFACE =====

// fmt.Stringer is one of the most common interfaces in Go
// type Stringer interface {
//     String() string
// }

type PersonStringer struct {
	Name string
	Age  int
	City string
}

func (p PersonStringer) String() string {
	return fmt.Sprintf("%s (%d years old) from %s", p.Name, p.Age, p.City)
}

type BankAccount struct {
	AccountNumber string
	Balance       float64
	Owner         string
}

func (ba BankAccount) String() string {
	// Mask account number for security
	masked := strings.Repeat("*", len(ba.AccountNumber)-4) + ba.AccountNumber[len(ba.AccountNumber)-4:]
	return fmt.Sprintf("Account %s - Balance: $%.2f (Owner: %s)", masked, ba.Balance, ba.Owner)
}

func demonstrateStringer() {
	fmt.Println("=== fmt.Stringer Interface ===")

	person := PersonStringer{Name: "Alice", Age: 30, City: "New York"}
	account := BankAccount{
		AccountNumber: "1234567890",
		Balance:       1500.75,
		Owner:         "Alice Johnson",
	}

	// String() method called automatically by fmt package
	fmt.Printf("Person: %s\n", person)
	fmt.Printf("Account: %s\n", account)

	// Also works with Printf verbs
	fmt.Printf("Person details: %v\n", person)
	fmt.Printf("Account info: %v\n", account)

	// Slice of Stringers
	items := []fmt.Stringer{person, account}
	fmt.Println("\nAll items:")
	for i, item := range items {
		fmt.Printf("  %d. %s\n", i+1, item)
	}
}

// ===== io.Reader INTERFACE =====

// io.Reader is fundamental for reading data
// type Reader interface {
//     Read([]byte) (n int, err error)
// }

type StringReader struct {
	data     string
	position int
}

func NewStringReader(data string) *StringReader {
	return &StringReader{data: data, position: 0}
}

func (sr *StringReader) Read(p []byte) (n int, err error) {
	if sr.position >= len(sr.data) {
		return 0, io.EOF
	}

	n = copy(p, sr.data[sr.position:])
	sr.position += n
	return n, nil
}

// UpperCaseReader wraps another reader and converts to uppercase
type UpperCaseReader struct {
	reader io.Reader
}

func NewUpperCaseReader(reader io.Reader) *UpperCaseReader {
	return &UpperCaseReader{reader: reader}
}

func (ucr *UpperCaseReader) Read(p []byte) (n int, err error) {
	n, err = ucr.reader.Read(p)
	for i := 0; i < n; i++ {
		if p[i] >= 'a' && p[i] <= 'z' {
			p[i] = p[i] - 'a' + 'A'
		}
	}
	return n, err
}

func demonstrateReader() {
	fmt.Println("\n=== io.Reader Interface ===")

	// Original reader
	original := NewStringReader("hello world from go!")

	// Wrapped reader
	upperReader := NewUpperCaseReader(original)

	// Read all data
	data, err := io.ReadAll(upperReader)
	if err != nil {
		fmt.Printf("Error reading: %v\n", err)
		return
	}

	fmt.Printf("Original: hello world from go!\n")
	fmt.Printf("Uppercase: %s\n", string(data))

	// Working with standard library readers
	fmt.Println("\nUsing strings.NewReader:")
	standardReader := strings.NewReader("this is from strings.NewReader")
	standardData, _ := io.ReadAll(standardReader)
	fmt.Printf("Read: %s\n", string(standardData))

	// Chain multiple readers
	fmt.Println("\nChaining readers:")
	chainedReader := NewUpperCaseReader(strings.NewReader("chained reader example"))
	chainedData, _ := io.ReadAll(chainedReader)
	fmt.Printf("Chained result: %s\n", string(chainedData))
}

// ===== io.Writer INTERFACE =====

// io.Writer is fundamental for writing data
// type Writer interface {
//     Write([]byte) (n int, err error)
// }

type LogWriter struct {
	prefix string
}

func NewLogWriter(prefix string) *LogWriter {
	return &LogWriter{prefix: prefix}
}

func (lw *LogWriter) Write(p []byte) (n int, err error) {
	timestamp := time.Now().Format("15:04:05")
	message := fmt.Sprintf("[%s] %s: %s", timestamp, lw.prefix, string(p))
	fmt.Print(message)
	return len(p), nil
}

type MemoryWriter struct {
	buffer []byte
}

func NewMemoryWriter() *MemoryWriter {
	return &MemoryWriter{buffer: make([]byte, 0)}
}

func (mw *MemoryWriter) Write(p []byte) (n int, err error) {
	mw.buffer = append(mw.buffer, p...)
	return len(p), nil
}

func (mw *MemoryWriter) String() string {
	return string(mw.buffer)
}

func (mw *MemoryWriter) Reset() {
	mw.buffer = mw.buffer[:0]
}

func demonstrateWriter() {
	fmt.Println("\n=== io.Writer Interface ===")

	// Log writer
	logger := NewLogWriter("APP")
	logger.Write([]byte("Application started\n"))
	logger.Write([]byte("Processing request\n"))

	// Memory writer
	memory := NewMemoryWriter()
	memory.Write([]byte("Hello "))
	memory.Write([]byte("World "))
	memory.Write([]byte("from memory!\n"))

	fmt.Printf("Memory buffer contains: %s", memory.String())

	// Using with fmt.Fprintf
	fmt.Println("\nUsing with fmt.Fprintf:")
	fmt.Fprintf(logger, "User %s logged in\n", "alice")
	fmt.Fprintf(memory, "Current time: %s\n", time.Now().Format("15:04:05"))
	fmt.Printf("Memory after fprintf: %s", memory.String())
}

// ===== io.ReadWriter INTERFACE =====

type InMemoryFile struct {
	content  []byte
	position int
}

func NewInMemoryFile() *InMemoryFile {
	return &InMemoryFile{
		content:  make([]byte, 0),
		position: 0,
	}
}

func (imf *InMemoryFile) Read(p []byte) (n int, err error) {
	if imf.position >= len(imf.content) {
		return 0, io.EOF
	}

	n = copy(p, imf.content[imf.position:])
	imf.position += n
	return n, nil
}

func (imf *InMemoryFile) Write(p []byte) (n int, err error) {
	imf.content = append(imf.content, p...)
	return len(p), nil
}

func (imf *InMemoryFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		imf.position = int(offset)
	case io.SeekCurrent:
		imf.position += int(offset)
	case io.SeekEnd:
		imf.position = len(imf.content) + int(offset)
	}

	if imf.position < 0 {
		imf.position = 0
	}
	if imf.position > len(imf.content) {
		imf.position = len(imf.content)
	}

	return int64(imf.position), nil
}

func demonstrateReadWriter() {
	fmt.Println("\n=== io.ReadWriter Interface ===")

	file := NewInMemoryFile()

	// Write some data
	file.Write([]byte("Line 1: Hello World\n"))
	file.Write([]byte("Line 2: Go Programming\n"))
	file.Write([]byte("Line 3: Interfaces Rock!\n"))

	// Reset position to beginning
	file.Seek(0, io.SeekStart)

	// Read all data
	fmt.Println("Reading from in-memory file:")
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Print(string(data))
	}

	// Append more data
	file.Write([]byte("Line 4: Appended data\n"))

	// Read from current position
	fmt.Println("\nReading new data:")
	file.Seek(0, io.SeekStart)
	newData, _ := io.ReadAll(file)
	fmt.Print(string(newData))
}

// ===== sort.Interface =====

// sort.Interface enables custom sorting
// type Interface interface {
//     Len() int
//     Less(i, j int) bool
//     Swap(i, j int)
// }

type Student struct {
	Name  string
	Grade float64
	Age   int
}

type ByGrade []Student

func (bg ByGrade) Len() int           { return len(bg) }
func (bg ByGrade) Less(i, j int) bool { return bg[i].Grade > bg[j].Grade } // Descending
func (bg ByGrade) Swap(i, j int)      { bg[i], bg[j] = bg[j], bg[i] }

type ByAge []Student

func (ba ByAge) Len() int           { return len(ba) }
func (ba ByAge) Less(i, j int) bool { return ba[i].Age < ba[j].Age } // Ascending
func (ba ByAge) Swap(i, j int)      { ba[i], ba[j] = ba[j], ba[i] }

type ByName []Student

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Less(i, j int) bool { return bn[i].Name < bn[j].Name }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }

func demonstrateSortInterface() {
	fmt.Println("\n=== sort.Interface ===")

	students := []Student{
		{Name: "Alice", Grade: 85.5, Age: 20},
		{Name: "Bob", Grade: 92.0, Age: 19},
		{Name: "Charlie", Grade: 78.5, Age: 21},
		{Name: "Diana", Grade: 96.0, Age: 18},
	}

	fmt.Println("Original students:")
	printStudents(students)

	// Sort by grade (descending)
	sort.Sort(ByGrade(students))
	fmt.Println("\nSorted by grade (highest first):")
	printStudents(students)

	// Sort by age (ascending)
	sort.Sort(ByAge(students))
	fmt.Println("\nSorted by age (youngest first):")
	printStudents(students)

	// Sort by name (alphabetical)
	sort.Sort(ByName(students))
	fmt.Println("\nSorted by name (alphabetical):")
	printStudents(students)
}

func printStudents(students []Student) {
	for _, student := range students {
		fmt.Printf("  %s: Grade=%.1f, Age=%d\n", student.Name, student.Grade, student.Age)
	}
}

// ===== error INTERFACE =====

// error is a built-in interface
// type error interface {
//     Error() string
// }

type ValidationErrorCustom struct {
	Field   string
	Value   interface{}
	Message string
}

func (ve ValidationErrorCustom) Error() string {
	return fmt.Sprintf("validation failed for field '%s' with value '%v': %s",
		ve.Field, ve.Value, ve.Message)
}

type NetworkErrorCustom struct {
	Operation string
	URL       string
	Code      int
	Timestamp time.Time
}

func (ne NetworkErrorCustom) Error() string {
	return fmt.Sprintf("network error during %s to %s: HTTP %d at %s",
		ne.Operation, ne.URL, ne.Code, ne.Timestamp.Format("15:04:05"))
}

// Implementing multiple interfaces
func (ne NetworkErrorCustom) String() string {
	return fmt.Sprintf("NetworkError{Op: %s, URL: %s, Code: %d}",
		ne.Operation, ne.URL, ne.Code)
}

func demonstrateErrorInterface() {
	fmt.Println("\n=== error Interface ===")

	// Custom validation error
	valErr := ValidationErrorCustom{
		Field:   "email",
		Value:   "invalid-email",
		Message: "must contain @ symbol",
	}

	// Custom network error
	netErr := NetworkErrorCustom{
		Operation: "GET",
		URL:       "https://api.example.com/users",
		Code:      404,
		Timestamp: time.Now(),
	}

	// Using as error interface
	errors := []error{valErr, netErr}

	fmt.Println("Handling errors:")
	for i, err := range errors {
		fmt.Printf("Error %d: %v\n", i+1, err)
	}

	// Type assertions to get specific behavior
	fmt.Println("\nDetailed error analysis:")
	for _, err := range errors {
		switch e := err.(type) {
		case ValidationErrorCustom:
			fmt.Printf("Validation issue in field: %s\n", e.Field)
		case NetworkErrorCustom:
			fmt.Printf("Network problem with status: %d\n", e.Code)
			// NetworkError also implements Stringer
			fmt.Printf("  Details: %s\n", e.String())
		default:
			fmt.Printf("Unknown error type: %T\n", e)
		}
	}
}

// ===== COMBINING INTERFACES =====

type Logger interface {
	Log(message string)
}

type FileSystemLogger interface {
	Logger
	io.Writer
	SetLogLevel(level string)
}

type ConsoleFileLogger struct {
	logLevel string
	buffer   *bytes.Buffer
}

func NewConsoleFileLogger() *ConsoleFileLogger {
	return &ConsoleFileLogger{
		logLevel: "INFO",
		buffer:   bytes.NewBuffer(nil),
	}
}

func (cfl *ConsoleFileLogger) Log(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] %s: %s\n", timestamp, cfl.logLevel, message)

	// Write to console
	fmt.Print(logEntry)

	// Write to buffer (simulating file)
	cfl.buffer.WriteString(logEntry)
}

func (cfl *ConsoleFileLogger) Write(p []byte) (n int, err error) {
	message := strings.TrimSpace(string(p))
	cfl.Log(message)
	return len(p), nil
}

func (cfl *ConsoleFileLogger) SetLogLevel(level string) {
	cfl.logLevel = level
}

func (cfl *ConsoleFileLogger) GetBufferedLogs() string {
	return cfl.buffer.String()
}

func demonstrateInterfaceCombination() {
	fmt.Println("\n=== Combining Interfaces ===")

	logger := NewConsoleFileLogger()

	// Use as Logger interface
	var log Logger = logger
	log.Log("Application started")

	// Use as io.Writer interface
	var writer io.Writer = logger
	fmt.Fprintf(writer, "User %s logged in", "alice")

	// Use as FileSystemLogger interface (composed)
	var fsLogger FileSystemLogger = logger
	fsLogger.SetLogLevel("ERROR")
	fsLogger.Log("Critical system error")

	// Show buffered content
	fmt.Println("\nBuffered logs:")
	fmt.Print(logger.GetBufferedLogs())
}

// ===== INTERFACE COMPOSITION PATTERNS =====

type Processor interface {
	Process(data string) (string, error)
}

type Validator interface {
	Validate(data string) error
}

type ProcessorValidator interface {
	Processor
	Validator
}

type EmailProcessor struct {
	domain string
}

func NewEmailProcessor(domain string) *EmailProcessor {
	return &EmailProcessor{domain: domain}
}

func (ep *EmailProcessor) Validate(email string) error {
	if !strings.Contains(email, "@") {
		return fmt.Errorf("email must contain @ symbol")
	}
	if !strings.HasSuffix(email, ep.domain) {
		return fmt.Errorf("email must be from domain %s", ep.domain)
	}
	return nil
}

func (ep *EmailProcessor) Process(email string) (string, error) {
	if err := ep.Validate(email); err != nil {
		return "", fmt.Errorf("processing failed: %w", err)
	}

	// Normalize email
	normalized := strings.ToLower(strings.TrimSpace(email))
	return normalized, nil
}

func demonstrateInterfaceCompositionPatterns() {
	fmt.Println("\n=== Interface Composition Patterns ===")

	processor := NewEmailProcessor("@company.com")

	emails := []string{
		"Alice@company.com",
		"bob@other.com",
		"invalid-email",
		"Charlie@COMPANY.COM",
	}

	fmt.Println("Processing emails:")
	for _, email := range emails {
		// Use as composed interface
		var pv ProcessorValidator = processor

		// Validate first
		if err := pv.Validate(email); err != nil {
			fmt.Printf("❌ %s: %v\n", email, err)
			continue
		}

		// Then process
		result, err := pv.Process(email)
		if err != nil {
			fmt.Printf("❌ %s: %v\n", email, err)
		} else {
			fmt.Printf("✅ %s → %s\n", email, result)
		}
	}
}

// ===== MAIN DEMO FUNCTION =====

func runCommonInterfacesDemo() {
	fmt.Println("📚 Go Common Standard Library Interfaces Tutorial")
	fmt.Println("=================================================")

	demonstrateStringer()
	demonstrateReader()
	demonstrateWriter()
	demonstrateReadWriter()
	demonstrateSortInterface()
	demonstrateErrorInterface()
	demonstrateInterfaceCombination()
	demonstrateInterfaceCompositionPatterns()

	fmt.Println("\n✅ Common interfaces concepts covered!")
	fmt.Println("\n🎯 Key Points:")
	fmt.Println("- fmt.Stringer customizes string representation")
	fmt.Println("- io.Reader/Writer enable flexible data handling")
	fmt.Println("- sort.Interface allows custom sorting logic")
	fmt.Println("- error interface enables rich error information")
	fmt.Println("- Interface composition creates powerful abstractions")
	fmt.Println("- Standard interfaces integrate seamlessly with Go's ecosystem")
	fmt.Println("- Implementing standard interfaces makes your types 'Go-idiomatic'")
	fmt.Println("- Small, focused interfaces are easier to implement and test")
}
