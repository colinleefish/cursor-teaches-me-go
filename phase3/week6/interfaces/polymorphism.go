package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// ===== BASIC POLYMORPHISM =====

type Animal interface {
	Speak() string
	Move() string
}

type DogPoly struct {
	Name  string
	Breed string
}

func (d DogPoly) Speak() string {
	return fmt.Sprintf("%s says Woof!", d.Name)
}

func (d DogPoly) Move() string {
	return fmt.Sprintf("%s runs around", d.Name)
}

type BirdPoly struct {
	Name    string
	Species string
}

func (b BirdPoly) Speak() string {
	return fmt.Sprintf("%s chirps melodiously", b.Name)
}

func (b BirdPoly) Move() string {
	return fmt.Sprintf("%s flies gracefully", b.Name)
}

type FishPoly struct {
	Name string
	Type string
}

func (f FishPoly) Speak() string {
	return fmt.Sprintf("%s makes bubbles", f.Name)
}

func (f FishPoly) Move() string {
	return fmt.Sprintf("%s swims smoothly", f.Name)
}

func demonstrateBasicPolymorphism() {
	fmt.Println("=== Basic Polymorphism ===")

	// Different types, same interface
	animals := []Animal{
		DogPoly{Name: "Rex", Breed: "German Shepherd"},
		BirdPoly{Name: "Tweety", Species: "Canary"},
		FishPoly{Name: "Nemo", Type: "Clownfish"},
		DogPoly{Name: "Buddy", Breed: "Golden Retriever"},
	}

	// Polymorphic behavior - same method calls, different implementations
	for i, animal := range animals {
		fmt.Printf("Animal %d:\n", i+1)
		fmt.Printf("  %s\n", animal.Speak())
		fmt.Printf("  %s\n", animal.Move())
		fmt.Printf("  Type: %T\n", animal)
		fmt.Println()
	}

	// Functions can work with any Animal implementation
	makeAnimalShow(animals)
}

func makeAnimalShow(animals []Animal) {
	fmt.Println("🎪 Welcome to the Animal Show!")
	for _, animal := range animals {
		fmt.Printf("- %s\n", animal.Speak())
		fmt.Printf("- %s\n", animal.Move())
	}
}

// ===== STRATEGY PATTERN =====

type SortStrategy interface {
	Sort([]int) []int
	Name() string
}

type BubbleSort struct{}

func (bs BubbleSort) Sort(data []int) []int {
	result := make([]int, len(data))
	copy(result, data)

	n := len(result)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

func (bs BubbleSort) Name() string {
	return "Bubble Sort"
}

type QuickSort struct{}

func (qs QuickSort) Sort(data []int) []int {
	result := make([]int, len(data))
	copy(result, data)
	quickSortHelper(result, 0, len(result)-1)
	return result
}

func (qs QuickSort) Name() string {
	return "Quick Sort"
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSortHelper(arr, low, pi-1)
		quickSortHelper(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

type DataSorter struct {
	strategy SortStrategy
}

func NewDataSorter(strategy SortStrategy) *DataSorter {
	return &DataSorter{strategy: strategy}
}

func (ds *DataSorter) SetStrategy(strategy SortStrategy) {
	ds.strategy = strategy
}

func (ds *DataSorter) SortData(data []int) []int {
	fmt.Printf("Using %s...\n", ds.strategy.Name())
	start := time.Now()
	result := ds.strategy.Sort(data)
	duration := time.Since(start)
	fmt.Printf("%s completed in %v\n", ds.strategy.Name(), duration)
	return result
}

func demonstrateStrategyPattern() {
	fmt.Println("\n=== Strategy Pattern ===")

	data := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Original data: %v\n", data)

	sorter := NewDataSorter(BubbleSort{})

	// Use bubble sort
	result1 := sorter.SortData(data)
	fmt.Printf("Bubble sorted: %v\n", result1)

	// Change strategy to quick sort
	sorter.SetStrategy(QuickSort{})
	result2 := sorter.SortData(data)
	fmt.Printf("Quick sorted: %v\n", result2)
}

// ===== OBSERVER PATTERN =====

type Observer interface {
	Update(message string)
	GetID() string
}

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify(message string)
}

type NewsAgency struct {
	observers []Observer
	news      string
}

func NewNewsAgency() *NewsAgency {
	return &NewsAgency{
		observers: make([]Observer, 0),
	}
}

func (na *NewsAgency) Attach(observer Observer) {
	na.observers = append(na.observers, observer)
	fmt.Printf("Observer %s attached\n", observer.GetID())
}

func (na *NewsAgency) Detach(observer Observer) {
	for i, obs := range na.observers {
		if obs.GetID() == observer.GetID() {
			na.observers = append(na.observers[:i], na.observers[i+1:]...)
			fmt.Printf("Observer %s detached\n", observer.GetID())
			break
		}
	}
}

func (na *NewsAgency) Notify(message string) {
	na.news = message
	fmt.Printf("📰 Broadcasting news: %s\n", message)
	for _, observer := range na.observers {
		observer.Update(message)
	}
}

type NewsChannel struct {
	name string
}

func NewNewsChannel(name string) *NewsChannel {
	return &NewsChannel{name: name}
}

func (nc *NewsChannel) Update(message string) {
	fmt.Printf("📺 %s received: %s\n", nc.name, message)
}

func (nc *NewsChannel) GetID() string {
	return nc.name
}

type NewsPaper struct {
	name string
}

func NewNewsPaper(name string) *NewsPaper {
	return &NewsPaper{name: name}
}

func (np *NewsPaper) Update(message string) {
	fmt.Printf("📰 %s printing: %s\n", np.name, message)
}

func (np *NewsPaper) GetID() string {
	return np.name
}

func demonstrateObserverPattern() {
	fmt.Println("\n=== Observer Pattern ===")

	agency := NewNewsAgency()

	// Create observers
	cnn := NewNewsChannel("CNN")
	bbc := NewNewsChannel("BBC")
	times := NewNewsPaper("New York Times")
	post := NewNewsPaper("Washington Post")

	// Attach observers
	agency.Attach(cnn)
	agency.Attach(bbc)
	agency.Attach(times)
	agency.Attach(post)

	// Broadcast news
	agency.Notify("Go 1.21 Released with New Features!")
	fmt.Println()

	// Detach one observer
	agency.Detach(bbc)
	agency.Notify("Major Breakthrough in AI Technology!")
}

// ===== FACTORY PATTERN =====

type Transport interface {
	Deliver() string
	GetType() string
}

type Truck struct {
	capacity int
}

func (t Truck) Deliver() string {
	return fmt.Sprintf("Delivering by truck (capacity: %d tons)", t.capacity)
}

func (t Truck) GetType() string {
	return "Land Transport"
}

type Ship struct {
	containers int
}

func (s Ship) Deliver() string {
	return fmt.Sprintf("Delivering by ship (%d containers)", s.containers)
}

func (s Ship) GetType() string {
	return "Sea Transport"
}

type Plane struct {
	passengers int
}

func (p Plane) Deliver() string {
	return fmt.Sprintf("Delivering by plane (%d passengers)", p.passengers)
}

func (p Plane) GetType() string {
	return "Air Transport"
}

type TransportFactory interface {
	CreateTransport() Transport
}

type LandTransportFactory struct{}

func (ltf LandTransportFactory) CreateTransport() Transport {
	return Truck{capacity: 10}
}

type SeaTransportFactory struct{}

func (stf SeaTransportFactory) CreateTransport() Transport {
	return Ship{containers: 50}
}

type AirTransportFactory struct{}

func (atf AirTransportFactory) CreateTransport() Transport {
	return Plane{passengers: 200}
}

type LogisticsManager struct {
	factory TransportFactory
}

func NewLogisticsManager(factory TransportFactory) *LogisticsManager {
	return &LogisticsManager{factory: factory}
}

func (lm *LogisticsManager) PlanDelivery() {
	transport := lm.factory.CreateTransport()
	fmt.Printf("🚚 %s\n", transport.Deliver())
	fmt.Printf("   Type: %s\n", transport.GetType())
}

func demonstrateFactoryPattern() {
	fmt.Println("\n=== Factory Pattern ===")

	// Different factories create different transports
	factories := []TransportFactory{
		LandTransportFactory{},
		SeaTransportFactory{},
		AirTransportFactory{},
	}

	for i, factory := range factories {
		fmt.Printf("Delivery Option %d:\n", i+1)
		manager := NewLogisticsManager(factory)
		manager.PlanDelivery()
		fmt.Println()
	}
}

// ===== ADAPTER PATTERN =====

// Legacy system interface
type LegacyPrinter interface {
	PrintOldFormat(data string)
}

type OldPrinter struct {
	name string
}

func (op OldPrinter) PrintOldFormat(data string) {
	fmt.Printf("🖨️ [Legacy %s] %s\n", op.name, strings.ToUpper(data))
}

// Modern interface
type ModernPrinter interface {
	Print(document DocumentAdapter) error
}

type DocumentAdapter struct {
	Title   string
	Content string
	Author  string
}

// Adapter to make legacy printer work with modern interface
type PrinterAdapter struct {
	legacyPrinter LegacyPrinter
}

func NewPrinterAdapter(legacy LegacyPrinter) *PrinterAdapter {
	return &PrinterAdapter{legacyPrinter: legacy}
}

func (pa *PrinterAdapter) Print(doc DocumentAdapter) error {
	// Convert modern document to legacy format
	legacyFormat := fmt.Sprintf("TITLE: %s | AUTHOR: %s | CONTENT: %s",
		doc.Title, doc.Author, doc.Content)

	pa.legacyPrinter.PrintOldFormat(legacyFormat)
	return nil
}

type AdvancedPrinter struct {
	model string
}

func (ap AdvancedPrinter) Print(doc DocumentAdapter) error {
	fmt.Printf("🖨️ [Advanced %s]\n", ap.model)
	fmt.Printf("   Title: %s\n", doc.Title)
	fmt.Printf("   Author: %s\n", doc.Author)
	fmt.Printf("   Content: %s\n", doc.Content)
	return nil
}

func demonstrateAdapterPattern() {
	fmt.Println("\n=== Adapter Pattern ===")

	doc := DocumentAdapter{
		Title:   "Go Programming Guide",
		Content: "Learning interfaces and polymorphism",
		Author:  "Go Developer",
	}

	// Modern printers
	printers := []ModernPrinter{
		AdvancedPrinter{model: "HP LaserJet Pro"},
		NewPrinterAdapter(OldPrinter{name: "Dot Matrix 1985"}),
		AdvancedPrinter{model: "Canon PIXMA"},
		NewPrinterAdapter(OldPrinter{name: "IBM Typewriter"}),
	}

	for i, printer := range printers {
		fmt.Printf("Printer %d:\n", i+1)
		printer.Print(doc)
		fmt.Println()
	}
}

// ===== DECORATOR PATTERN =====

type Coffee interface {
	Cost() float64
	Description() string
}

type SimpleCoffee struct{}

func (sc SimpleCoffee) Cost() float64 {
	return 2.0
}

func (sc SimpleCoffee) Description() string {
	return "Simple coffee"
}

type CoffeeDecorator struct {
	coffee Coffee
}

type MilkDecorator struct {
	CoffeeDecorator
}

func NewMilkDecorator(coffee Coffee) *MilkDecorator {
	return &MilkDecorator{CoffeeDecorator{coffee: coffee}}
}

func (md *MilkDecorator) Cost() float64 {
	return md.coffee.Cost() + 0.5
}

func (md *MilkDecorator) Description() string {
	return md.coffee.Description() + ", milk"
}

type SugarDecorator struct {
	CoffeeDecorator
}

func NewSugarDecorator(coffee Coffee) *SugarDecorator {
	return &SugarDecorator{CoffeeDecorator{coffee: coffee}}
}

func (sd *SugarDecorator) Cost() float64 {
	return sd.coffee.Cost() + 0.2
}

func (sd *SugarDecorator) Description() string {
	return sd.coffee.Description() + ", sugar"
}

type WhipDecorator struct {
	CoffeeDecorator
}

func NewWhipDecorator(coffee Coffee) *WhipDecorator {
	return &WhipDecorator{CoffeeDecorator{coffee: coffee}}
}

func (wd *WhipDecorator) Cost() float64 {
	return wd.coffee.Cost() + 0.7
}

func (wd *WhipDecorator) Description() string {
	return wd.coffee.Description() + ", whipped cream"
}

func demonstrateDecoratorPattern() {
	fmt.Println("\n=== Decorator Pattern ===")

	// Start with simple coffee
	coffee := SimpleCoffee{}
	fmt.Printf("☕ %s - $%.2f\n", coffee.Description(), coffee.Cost())

	// Add milk
	coffeeWithMilk := NewMilkDecorator(coffee)
	fmt.Printf("☕ %s - $%.2f\n", coffeeWithMilk.Description(), coffeeWithMilk.Cost())

	// Add sugar
	coffeeWithMilkAndSugar := NewSugarDecorator(coffeeWithMilk)
	fmt.Printf("☕ %s - $%.2f\n", coffeeWithMilkAndSugar.Description(), coffeeWithMilkAndSugar.Cost())

	// Add whipped cream
	fancyCoffee := NewWhipDecorator(coffeeWithMilkAndSugar)
	fmt.Printf("☕ %s - $%.2f\n", fancyCoffee.Description(), fancyCoffee.Cost())

	fmt.Println("\nCustom combinations:")

	// Different combinations
	combinations := []Coffee{
		NewSugarDecorator(SimpleCoffee{}),
		NewWhipDecorator(NewMilkDecorator(SimpleCoffee{})),
		NewMilkDecorator(NewSugarDecorator(NewWhipDecorator(SimpleCoffee{}))),
	}

	for i, combo := range combinations {
		fmt.Printf("%d. ☕ %s - $%.2f\n", i+1, combo.Description(), combo.Cost())
	}
}

// ===== POLYMORPHIC BEHAVIOR WITH RANDOM SELECTION =====

type GameCharacter interface {
	Attack() string
	Defend() string
	GetHealth() int
	TakeDamage(damage int)
}

type Warrior struct {
	name   string
	health int
	armor  int
}

func NewWarrior(name string) *Warrior {
	return &Warrior{name: name, health: 100, armor: 20}
}

func (w *Warrior) Attack() string {
	damage := 25 + rand.Intn(10)
	return fmt.Sprintf("%s swings sword for %d damage!", w.name, damage)
}

func (w *Warrior) Defend() string {
	return fmt.Sprintf("%s raises shield (armor: %d)", w.name, w.armor)
}

func (w *Warrior) GetHealth() int {
	return w.health
}

func (w *Warrior) TakeDamage(damage int) {
	actualDamage := damage - w.armor/2
	if actualDamage < 0 {
		actualDamage = 0
	}
	w.health -= actualDamage
	if w.health < 0 {
		w.health = 0
	}
}

type Mage struct {
	name   string
	health int
	mana   int
}

func NewMage(name string) *Mage {
	return &Mage{name: name, health: 70, mana: 100}
}

func (m *Mage) Attack() string {
	if m.mana >= 20 {
		m.mana -= 20
		damage := 30 + rand.Intn(15)
		return fmt.Sprintf("%s casts fireball for %d damage! (mana: %d)", m.name, damage, m.mana)
	}
	return fmt.Sprintf("%s has no mana for spells!", m.name)
}

func (m *Mage) Defend() string {
	if m.mana >= 10 {
		m.mana -= 10
		return fmt.Sprintf("%s casts magic shield (mana: %d)", m.name, m.mana)
	}
	return fmt.Sprintf("%s dodges attack", m.name)
}

func (m *Mage) GetHealth() int {
	return m.health
}

func (m *Mage) TakeDamage(damage int) {
	m.health -= damage
	if m.health < 0 {
		m.health = 0
	}
}

func demonstrateGamePolymorphism() {
	fmt.Println("\n=== Game Character Polymorphism ===")

	characters := []GameCharacter{
		NewWarrior("Conan"),
		NewMage("Gandalf"),
		NewWarrior("Aragorn"),
		NewMage("Merlin"),
	}

	fmt.Println("🗡️ Battle Simulation!")

	for round := 1; round <= 3; round++ {
		fmt.Printf("\nRound %d:\n", round)

		for _, char := range characters {
			if char.GetHealth() > 0 {
				action := rand.Intn(2)
				if action == 0 {
					fmt.Printf("  %s\n", char.Attack())
				} else {
					fmt.Printf("  %s\n", char.Defend())
				}
			}
		}

		// Simulate some damage
		for _, char := range characters {
			if char.GetHealth() > 0 {
				damage := rand.Intn(15) + 5
				char.TakeDamage(damage)
			}
		}
	}

	fmt.Println("\nFinal Status:")
	for i, char := range characters {
		health := char.GetHealth()
		status := "💀 KO"
		if health > 0 {
			status = "✅ Alive"
		}
		fmt.Printf("  Character %d: Health=%d %s\n", i+1, health, status)
	}
}

// ===== MAIN DEMO FUNCTION =====

func runPolymorphismDemo() {
	fmt.Println("🔄 Go Polymorphism and Design Patterns Tutorial")
	fmt.Println("===============================================")

	rand.Seed(time.Now().UnixNano())

	demonstrateBasicPolymorphism()
	demonstrateStrategyPattern()
	demonstrateObserverPattern()
	demonstrateFactoryPattern()
	demonstrateAdapterPattern()
	demonstrateDecoratorPattern()
	demonstrateGamePolymorphism()

	fmt.Println("\n✅ Polymorphism and design patterns covered!")
	fmt.Println("\n🎯 Key Points:")
	fmt.Println("- Polymorphism enables different types to be treated uniformly")
	fmt.Println("- Strategy pattern allows algorithm selection at runtime")
	fmt.Println("- Observer pattern enables loose coupling between objects")
	fmt.Println("- Factory pattern abstracts object creation")
	fmt.Println("- Adapter pattern makes incompatible interfaces work together")
	fmt.Println("- Decorator pattern adds behavior without modifying original types")
	fmt.Println("- Interfaces enable flexible and maintainable design patterns")
	fmt.Println("- Go's implicit interface satisfaction makes patterns natural")
}
