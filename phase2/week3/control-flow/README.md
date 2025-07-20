# Week 3: Control Flow Structures 🔄

Welcome to Go's control flow! Unlike Python, Go has a simpler but more explicit approach to controlling program execution. Let's master all the ways to control your program's flow.

## 📚 What You'll Learn

- **if/else statements**: Go's condition checking (no parentheses needed!)
- **for loops**: The only loop in Go, but it's incredibly flexible
- **switch statements**: Clean multi-way branching without fallthrough
- **defer statements**: Cleanup and resource management
- **goto/labels**: When you absolutely need them (rarely!)

## 🎯 Learning Objectives

After completing this section, you'll be able to:
- [ ] Write conditional logic with `if/else` statements
- [ ] Use all four patterns of Go's `for` loop
- [ ] Create clean multi-way branches with `switch`
- [ ] Properly manage resources with `defer`
- [ ] Understand when and how to use `goto` (though you probably won't need it)

## 📁 Files in This Section

- `if_else.go` - Conditional statements and scope
- `loops.go` - All patterns of Go's flexible for loop
- `switch.go` - Switch statements and type switches
- `defer.go` - Resource cleanup and defer patterns
- `goto.go` - Labels and goto (for completeness)
- `control_flow_practice.go` - **YOUR PRACTICE FILE** - Fill in the blanks!

## ⚡ Key Differences from Python

### Conditions (No Parentheses!)
```python
# Python
if (x > 0 and y < 10):
    print("valid")

# Go
if x > 0 && y < 10 {
    fmt.Println("valid")
}
```

### Loops (Only `for`, But Flexible)
```python
# Python has while, for
while condition:
    do_something()

for i in range(10):
    print(i)

for item in items:
    process(item)

# Go only has for, but it does everything
for condition {
    doSomething()
}

for i := 0; i < 10; i++ {
    fmt.Println(i)
}

for _, item := range items {
    process(item)
}
```

### Switch (No Fallthrough by Default)
```python
# Python doesn't have switch - uses if/elif
if value == "a":
    handle_a()
elif value == "b" or value == "c":
    handle_bc()
else:
    handle_default()

# Go
switch value {
case "a":
    handleA()
case "b", "c":
    handleBC()
default:
    handleDefault()
}
```

## 🚀 Getting Started

1. Read through the example files to understand the concepts
2. Open `control_flow_practice.go`
3. Fill in the `// YOUR CODE HERE` sections
4. Run with: `go run control_flow_practice.go`
5. Test your understanding with different inputs

## 💡 Pro Tips

1. **No parentheses needed** around conditions in `if` statements
2. **Opening brace must be on same line** as `if`, `for`, `switch`
3. **`defer` executes in LIFO order** (Last In, First Out)
4. **Use `:=` for new variables** in if conditions: `if x := getValue(); x > 0 {}`
5. **Range loops** give you index and value: `for i, v := range slice {}`

## 🧪 Exercises to Complete

Each exercise builds your understanding step by step:

1. **Basic Conditions** - if/else with various operators
2. **Loop Patterns** - All four types of for loops
3. **Switch Logic** - Regular and type switches
4. **Resource Management** - Using defer properly
5. **Real-World Scenarios** - Combining control flow patterns

## 🎯 Success Criteria

You'll know you've mastered this section when you can:
- Write clean conditional logic without parentheses
- Choose the right loop pattern for any situation
- Use switch statements for clean multi-way branching
- Properly clean up resources with defer
- Combine control flow structures effectively

Let's dive in and master Go's control flow! 🐹 