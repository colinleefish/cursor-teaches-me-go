package main

import (
	"fmt"
	"math"
)

// ===== BASIC INTERFACE DEFINITION =====

// Simple interface with one method
type Speaker interface {
	Speak() string
}

// Types that implement Speaker interface (implicitly)
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("%s says Woof!", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("%s says Meow!", c.Name)
}

type Robot struct {
	Model string
}

func (r Robot) Speak() string {
	return fmt.Sprintf("Robot %s says: BEEP BOOP!", r.Model)
}

func demonstrateBasicInterfaces() {
	fmt.Println("=== Basic Interface Implementation ===")

	// All these types automatically implement Speaker
	animals := []Speaker{
		Dog{Name: "Buddy"},
		Cat{Name: "Whiskers"},
		Robot{Model: "R2D2"},
	}

	for _, speaker := range animals {
		fmt.Println(speaker.Speak())
	}

	// Interface variable can hold any implementing type
	var s Speaker
	s = Dog{Name: "Rex"}
	fmt.Printf("Speaker is now: %s\n", s.Speak())

	s = Cat{Name: "Mittens"}
	fmt.Printf("Speaker is now: %s\n", s.Speak())
}

// ===== INTERFACE WITH MULTIPLE METHODS =====

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func demonstrateMultiMethodInterface() {
	fmt.Println("\n=== Multi-Method Interface ===")

	shapes := []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 3},
		Rectangle{Width: 7, Height: 7},
	}

	totalArea := 0.0
	for i, shape := range shapes {
		area := shape.Area()
		perimeter := shape.Perimeter()
		totalArea += area

		fmt.Printf("Shape %d: Area=%.2f, Perimeter=%.2f\n", i+1, area, perimeter)
	}

	fmt.Printf("Total area of all shapes: %.2f\n", totalArea)
}

// ===== INTERFACE SATISFACTION =====

type Walker interface {
	Walk() string
}

type Runner interface {
	Run() string
}

type Athlete struct {
	Name string
}

func (a Athlete) Walk() string {
	return fmt.Sprintf("%s is walking", a.Name)
}

func (a Athlete) Run() string {
	return fmt.Sprintf("%s is running", a.Name)
}

func demonstrateInterfaceSatisfaction() {
	fmt.Println("\n=== Interface Satisfaction ===")

	athlete := Athlete{Name: "John"}

	// Athlete implements both Walker and Runner
	var walker Walker = athlete
	var runner Runner = athlete

	fmt.Println(walker.Walk())
	fmt.Println(runner.Run())

	// Functions accepting interfaces
	makeWalk(athlete)
	makeRun(athlete)
}

func makeWalk(w Walker) {
	fmt.Printf("Making someone walk: %s\n", w.Walk())
}

func makeRun(r Runner) {
	fmt.Printf("Making someone run: %s\n", r.Run())
}

// ===== EMPTY INTERFACE =====

func demonstrateEmptyInterface() {
	fmt.Println("\n=== Empty Interface (interface{}) ===")

	// interface{} can hold any value
	var anything interface{}

	anything = 42
	fmt.Printf("anything = %v (type: %T)\n", anything, anything)

	anything = "hello world"
	fmt.Printf("anything = %v (type: %T)\n", anything, anything)

	anything = []int{1, 2, 3}
	fmt.Printf("anything = %v (type: %T)\n", anything, anything)

	anything = Dog{Name: "Rover"}
	fmt.Printf("anything = %v (type: %T)\n", anything, anything)

	// Slice of interface{} can hold mixed types
	mixed := []interface{}{
		42,
		"string",
		true,
		3.14,
		Dog{Name: "Mixed"},
	}

	fmt.Println("\nMixed slice:")
	for i, item := range mixed {
		fmt.Printf("  [%d] %v (type: %T)\n", i, item, item)
	}
}

// ===== TYPE ASSERTIONS =====

func demonstrateTypeAssertions() {
	fmt.Println("\n=== Type Assertions ===")

	var data interface{} = "Hello, World!"

	// Safe type assertion with ok idiom
	if str, ok := data.(string); ok {
		fmt.Printf("data is a string: %q\n", str)
	} else {
		fmt.Println("data is not a string")
	}

	// Unsafe type assertion (can panic)
	str := data.(string)
	fmt.Printf("Unsafe assertion: %q\n", str)

	// This would panic:
	// num := data.(int) // panic: interface conversion

	// Testing multiple types
	testTypeAssertion(42)
	testTypeAssertion("text")
	testTypeAssertion(3.14)
	testTypeAssertion(Dog{Name: "Asserter"})
}

func testTypeAssertion(data interface{}) {
	if str, ok := data.(string); ok {
		fmt.Printf("Got string: %q\n", str)
	} else if num, ok := data.(int); ok {
		fmt.Printf("Got int: %d\n", num)
	} else if f, ok := data.(float64); ok {
		fmt.Printf("Got float64: %.2f\n", f)
	} else {
		fmt.Printf("Got unknown type: %T with value %v\n", data, data)
	}
}

// ===== TYPE SWITCHES =====

func demonstrateTypeSwitches() {
	fmt.Println("\n=== Type Switches ===")

	values := []interface{}{
		42,
		"hello",
		3.14,
		true,
		Dog{Name: "Switcher"},
		[]int{1, 2, 3},
		nil,
	}

	for i, v := range values {
		fmt.Printf("Value %d: ", i)
		describeType(v)
	}
}

func describeType(data interface{}) {
	switch v := data.(type) {
	case nil:
		fmt.Println("nil value")
	case int:
		fmt.Printf("integer: %d\n", v)
	case string:
		fmt.Printf("string: %q (length: %d)\n", v, len(v))
	case float64:
		fmt.Printf("float: %.2f\n", v)
	case bool:
		fmt.Printf("boolean: %t\n", v)
	case Dog:
		fmt.Printf("dog: %s\n", v.Name)
	case []int:
		fmt.Printf("int slice: %v (length: %d)\n", v, len(v))
	default:
		fmt.Printf("unknown type: %T with value %v\n", v, v)
	}
}

// ===== INTERFACE COMPOSITION =====

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string) error
}

type Closer interface {
	Close() error
}

// Composed interfaces
type ReadWriter interface {
	Reader
	Writer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

type File struct {
	name     string
	content  string
	position int
	closed   bool
}

func (f *File) Read() string {
	if f.closed {
		return ""
	}
	data := f.content[f.position:]
	f.position = len(f.content)
	return data
}

func (f *File) Write(data string) error {
	if f.closed {
		return fmt.Errorf("file %s is closed", f.name)
	}
	f.content += data
	return nil
}

func (f *File) Close() error {
	if f.closed {
		return fmt.Errorf("file %s already closed", f.name)
	}
	f.closed = true
	fmt.Printf("File %s closed\n", f.name)
	return nil
}

func demonstrateInterfaceComposition() {
	fmt.Println("\n=== Interface Composition ===")

	file := &File{name: "test.txt", content: "Initial content\n"}

	// File implements all three individual interfaces
	var reader Reader = file
	var writer Writer = file
	var closer Closer = file

	// File also implements composed interfaces
	var readWriter ReadWriter = file
	var readWriteCloser ReadWriteCloser = file

	// Using individual interfaces
	fmt.Printf("Read: %s", reader.Read())
	writer.Write("Added content\n")

	// Using composed interfaces
	processReadWriter(readWriter)
	processReadWriteCloser(readWriteCloser)

	closer.Close()
}

func processReadWriter(rw ReadWriter) {
	rw.Write("From ReadWriter\n")
	fmt.Printf("ReadWriter content: %s", rw.Read())
}

func processReadWriteCloser(rwc ReadWriteCloser) {
	rwc.Write("From ReadWriteCloser\n")
	fmt.Printf("ReadWriteCloser content: %s", rwc.Read())
	// Note: We don't close here to avoid closing the file early
}

// ===== INTERFACE VALUES =====

type Printer interface {
	Print() string
}

type DocumentBasic struct {
	Title   string
	Content string
}

func (d DocumentBasic) Print() string {
	return fmt.Sprintf("Document: %s\nContent: %s", d.Title, d.Content)
}

func demonstrateInterfaceValues() {
	fmt.Println("\n=== Interface Values ===")

	var printer Printer

	// Zero value of interface is nil
	fmt.Printf("Zero interface: %v\n", printer)
	fmt.Printf("Is nil: %t\n", printer == nil)

	// Assigning a value
	printer = DocumentBasic{Title: "Go Guide", Content: "Learning Go interfaces"}
	fmt.Printf("Interface with value: %v\n", printer)
	fmt.Printf("Is nil: %t\n", printer == nil)

	// Interface holds both type and value
	fmt.Printf("Type: %T, Value: %v\n", printer, printer)

	// Calling method
	fmt.Println(printer.Print())

	// Setting back to nil
	printer = nil
	fmt.Printf("Back to nil: %v\n", printer)
	fmt.Printf("Is nil: %t\n", printer == nil)

	// This would panic if uncommented:
	// fmt.Println(printer.Print()) // panic: runtime error
}

// ===== INTERFACE BEST PRACTICES =====

func demonstrateBestPractices() {
	fmt.Println("\n=== Interface Best Practices ===")

	fmt.Println("✅ Key Interface Design Principles:")
	fmt.Println("- Keep interfaces small and focused")
	fmt.Println("- Define interfaces where they're used, not implemented")
	fmt.Println("- Use composition to build complex behaviors")
	fmt.Println("- Accept interfaces, return structs")
	fmt.Println("- Favor many small interfaces over few large ones")
	fmt.Println("- Name interfaces by capability (use -er suffix)")

	fmt.Println("\n💡 Best practices demonstrated in other tutorial files!")
}

// ===== MAIN DEMO FUNCTION =====

func runInterfaceBasicsDemo() {
	fmt.Println("🔌 Go Interface Basics Tutorial")
	fmt.Println("===============================")

	demonstrateBasicInterfaces()
	demonstrateMultiMethodInterface()
	demonstrateInterfaceSatisfaction()
	demonstrateEmptyInterface()
	demonstrateTypeAssertions()
	demonstrateTypeSwitches()
	demonstrateInterfaceComposition()
	demonstrateInterfaceValues()
	demonstrateBestPractices()

	fmt.Println("\n✅ Interface basics concepts covered!")
	fmt.Println("\n🎯 Key Points:")
	fmt.Println("- Interfaces are satisfied implicitly (duck typing)")
	fmt.Println("- Any type implementing interface methods satisfies the interface")
	fmt.Println("- Empty interface{} can hold any value")
	fmt.Println("- Use type assertions and type switches for runtime type checking")
	fmt.Println("- Compose interfaces for complex behaviors")
	fmt.Println("- Interface values hold both type and value information")
	fmt.Println("- Keep interfaces small and focused")
	fmt.Println("- Define interfaces where they're used, not implemented")
}
