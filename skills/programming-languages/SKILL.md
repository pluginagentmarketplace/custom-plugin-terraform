---
name: Programming Languages & Fundamentals
description: Master multiple programming languages, from Python's simplicity to Rust's safety. Learn syntax, idioms, and best practices across Python, JavaScript, Go, Java, and Rust. Covers fundamentals, OOP, functional programming, memory management, and language-specific optimization patterns.
---

# 📚 Programming Languages & Fundamentals

**Master syntax, idioms, and best practices across modern programming languages.**

Learn to write idiomatic code in Python, JavaScript, Go, Java, and Rust. Understand when to use each language and how to leverage its unique strengths.

## Quick Start: Your First Programs

### Python
```python
# Modern Python with type hints
def greet(name: str, age: int) -> str:
    """Greet a person with their age."""
    return f"Hello {name}! You are {age} years old."

# List comprehension and unpacking
numbers = [1, 2, 3, 4, 5]
squared = [n**2 for n in numbers if n % 2 == 0]
print(squared)  # [4, 16]

# Context managers for resource handling
with open('file.txt', 'r') as f:
    content = f.read()

# Async/await for concurrent operations
import asyncio

async def fetch_data(url):
    # Simulate API call
    await asyncio.sleep(1)
    return f"Data from {url}"

asyncio.run(fetch_data("https://api.example.com"))
```

### JavaScript
```javascript
// Modern ES6+ JavaScript
const greet = (name, age) => `Hello ${name}! You are ${age} years old.`;

// Array methods and destructuring
const numbers = [1, 2, 3, 4, 5];
const [first, second, ...rest] = numbers;
const squared = numbers
  .filter(n => n % 2 === 0)
  .map(n => n ** 2);

// Promises and async/await
async function fetchData(url) {
  try {
    const response = await fetch(url);
    return await response.json();
  } catch (error) {
    console.error('Fetch failed:', error);
  }
}

// Closures and higher-order functions
const createMultiplier = factor => number => number * factor;
const double = createMultiplier(2);
console.log(double(5));  // 10
```

### Go
```go
package main

import (
    "fmt"
    "log"
)

// Simple function with error handling
func greet(name string, age int) string {
    return fmt.Sprintf("Hello %s! You are %d years old.", name, age)
}

// Concurrency with goroutines
func fetchData(url string, ch chan string) {
    // Simulate API call
    result := "Data from " + url
    ch <- result
}

func main() {
    // Goroutines and channels
    ch := make(chan string)
    go fetchData("https://api.example.com", ch)
    data := <-ch
    fmt.Println(data)

    // Interfaces and composition
    var reader interface{ Read(p []byte) (n int, err error) }
    _ = reader
}
```

### Java
```java
public class GreetingService {
    // Type-safe generics
    public static <T> void printArray(T[] array) {
        for (T element : array) {
            System.out.println(element);
        }
    }

    // Functional interfaces and lambdas
    public String greet(String name, int age) {
        return String.format("Hello %s! You are %d years old.", name, age);
    }

    // Stream API for functional programming
    public static void main(String[] args) {
        var numbers = java.util.Arrays.asList(1, 2, 3, 4, 5);
        var squared = numbers.stream()
            .filter(n -> n % 2 == 0)
            .map(n -> n * n)
            .collect(java.util.stream.Collectors.toList());
        System.out.println(squared);
    }
}
```

### Rust
```rust
// Memory safety without garbage collection
fn greet(name: &str, age: u32) -> String {
    format!("Hello {}! You are {} years old.", name, age)
}

// Ownership and borrowing
fn main() {
    let mut numbers = vec![1, 2, 3, 4, 5];
    let first = &numbers[0];
    println!("First: {}", first);

    // Iterator pattern
    let squared: Vec<i32> = numbers.iter()
        .filter(|n| *n % 2 == 0)
        .map(|n| n * n)
        .collect();
    println!("{:?}", squared);

    // Pattern matching
    match first {
        1 => println!("It's one"),
        n if n > 0 => println!("It's positive: {}", n),
        _ => println!("Other"),
    }
}

// Async/await in Rust
async fn fetch_data(url: &str) -> String {
    format!("Data from {}", url)
}
```

## Language Comparison Matrix

| Aspect | Python | JavaScript | Go | Java | Rust |
|--------|--------|------------|----|----|------|
| **Typing** | Dynamic (optional hints) | Dynamic | Static | Static | Static |
| **Compilation** | Interpreted | Interpreted/JIT | Compiled | Compiled (JVM) | Compiled |
| **Memory Management** | Garbage Collected | Garbage Collected | Garbage Collected | Garbage Collected | Manual (Safe) |
| **Concurrency** | Threading/Async | Async/Promises | Goroutines | Threads | Async/Threads |
| **Learning Curve** | Gentle | Moderate | Easy | Steep | Very Steep |
| **Performance** | Medium | Medium-High | Very High | High | Very High |
| **Use Cases** | Data science, scripting | Web, frontend | Systems, servers | Enterprise, Android | Systems, performance |

## Core Concepts Mastery

### 1. Type Systems
- **Strongly Typed** (Java, Go, Rust): Catch errors at compile-time
- **Dynamically Typed** (Python, JavaScript): Flexibility at cost of runtime errors
- **Gradually Typed**: Python type hints, TypeScript provide optional typing
- **Type Inference**: Rust, Go infer types without explicit declaration

### 2. Object-Oriented Programming
```python
# Python OOP with duck typing
class Animal:
    def __init__(self, name):
        self.name = name

    def speak(self):
        pass  # Subclasses implement

class Dog(Animal):
    def speak(self):
        return f"{self.name} barks"

# Interfaces without explicit inheritance
class Bird:
    def speak(self):
        return "chirp"
```

### 3. Functional Programming
```javascript
// Pure functions and immutability
const add = (a, b) => a + b;
const multiply = (a, b) => a * b;

// Function composition
const compose = (f, g) => x => f(g(x));
const addThenMultiply = compose(
    x => multiply(x, 2),
    x => add(x, 3)
);

console.log(addThenMultiply(5));  // (5 + 3) * 2 = 16
```

### 4. Memory Management
- **Garbage Collection** (Python, JavaScript, Java): Automatic, some overhead
- **Manual Safety** (Rust): Ownership system prevents memory errors at compile-time
- **Reference Counting** (Swift, Objective-C): Track object references

### 5. Concurrency Models
- **Threads** (Java): OS-level, heavyweight, complex synchronization
- **Goroutines** (Go): Lightweight, thousands easily, channels for communication
- **Async/Await** (JavaScript, Python, Rust): Single-threaded with event loop
- **Actors** (Erlang, Akka): Message passing between isolated entities

## Detailed Learning Path

### Phase 1: Foundations (Weeks 1-4)
**Choose Your Primary Language**

#### Python Path
- Variables, data types, operators
- Control flow (if/else, loops)
- Functions and scope
- Lists, dictionaries, sets
- String manipulation
- **Projects**:
  - Calculator with operator support
  - Grade calculator with averages
  - Simple to-do list manager
  - Palindrome checker
  - Number guessing game

#### JavaScript Path
- Variables (var/let/const)
- Data types and coercion
- Functions and closures
- DOM manipulation (browser only)
- Array and object methods
- **Projects**:
  - Interactive calculator
  - Todo list with add/delete
  - Color palette generator
  - Quiz application
  - Weather display app

#### Go Path
- Packages and imports
- Variables and constants
- Functions and multiple returns
- Strings and formatting
- Basic file I/O
- **Projects**:
  - Command-line calculator
  - File word counter
  - Simple HTTP server
  - Configuration parser
  - Log analyzer

#### Java Path
- Classes and objects
- Constructors and methods
- Packages and imports
- Arrays and collections
- Exception handling
- **Projects**:
  - Bank account class
  - Student grade tracker
  - Library management system
  - Simple game (tic-tac-toe)
  - Task scheduler

#### Rust Path
- Variable binding
- Ownership and borrowing
- Pattern matching
- Error handling (Result/Option)
- Structs and implementations
- **Projects**:
  - Number guessing game
  - Temperature converter
  - File search tool
  - Config file parser
  - Simple CLI application

### Phase 2: Intermediate (Weeks 5-12)
**Master Language Idioms**

#### All Languages
- Object-Oriented Programming
- Functional programming concepts
- Error handling patterns
- Collections deep dive
- Working with external libraries/packages
- Testing frameworks
- **Common Projects**:
  - Web scraper
  - REST API client
  - Data processing pipeline
  - Task scheduler with persistence
  - Blog or wiki application

#### Python Specifics
- Decorators and metaclasses
- Context managers
- Generators and iterators
- Async/await
- Package management (pip, poetry)
- **Advanced Patterns**:
  - Dependency injection
  - Factory pattern
  - Observer pattern

#### JavaScript Specifics
- Prototypes and inheritance
- Async patterns (callbacks, promises, async/await)
- Module systems (CommonJS, ESM)
- Event-driven architecture
- Testing with Jest/Vitest
- **Advanced Patterns**:
  - Higher-order functions
  - Currying and partial application
  - Memoization

#### Go Specifics
- Interfaces and type embedding
- Goroutines and channels
- Context package
- Error handling best practices
- Build and deployment
- **Advanced Patterns**:
  - Pipeline pattern
  - Worker pool pattern
  - Fan-out/Fan-in

#### Java Specifics
- Generics and type bounds
- Stream API
- Lambda expressions
- Reflection and annotations
- Design patterns (Gang of Four)
- **Advanced Patterns**:
  - Dependency injection (Spring)
  - Proxy and decorator
  - Strategy pattern

#### Rust Specifics
- Trait system
- Lifetimes and borrowing rules
- Macro system
- Error propagation with ?
- Module system
- **Advanced Patterns**:
  - Builder pattern
  - Type-state pattern
  - Zero-cost abstractions

### Phase 3: Advanced (Weeks 13-24)
**Build Production-Grade Applications**

#### All Languages
- Performance profiling and optimization
- Debugging and troubleshooting
- Security best practices
- Testing strategies (unit, integration, e2e)
- Documentation and code quality
- Deployment and DevOps
- **Capstone Projects**:
  - Full-stack application
  - Distributed system component
  - High-performance service
  - Complex data processing pipeline
  - Real-time application

#### Language-Specific Advanced Topics

**Python**:
- Async concurrency patterns
- C extensions and ctypes
- Memory profiling
- Advanced metaclasses
- Machine learning frameworks (NumPy, Pandas, TensorFlow)

**JavaScript**:
- Advanced async patterns
- WebWorkers and threading
- Performance optimization
- Build tools (Webpack, Vite)
- Full-stack frameworks (Next.js, Remix)

**Go**:
- Advanced concurrency patterns
- Network programming
- Building CLI tools
- Microservices architecture
- Container orchestration

**Java**:
- Enterprise frameworks (Spring, Jakarta EE)
- Concurrency utilities
- JVM internals and tuning
- Design patterns at scale
- Distributed systems

**Rust**:
- Systems programming
- Unsafe code and FFI
- Advanced trait design
- Procedural macros
- Embedded systems

## 20+ Production Projects

### Beginner Projects (1-2 weeks each)
1. **Command-line Todo App** - Basic CRUD operations with file persistence
2. **Budget Tracker** - Track income/expenses with categories
3. **Quiz Game** - Multiple choice questions with scoring
4. **Password Generator** - Customizable password creation
5. **Weather CLI** - Fetch and display weather data
6. **Unit Converter** - Convert between different units
7. **Markdown to HTML Converter** - Parse and render markdown
8. **Dice Roller Simulator** - D&D-style dice rolling with statistics

### Intermediate Projects (2-4 weeks each)
9. **REST API Backend** - User management with authentication
10. **Blog Platform** - Posts, comments, user accounts
11. **Chat Application** - Real-time messaging system
12. **E-commerce API** - Products, shopping cart, orders
13. **Real-time Dashboard** - Live data visualization
14. **File Encryption Tool** - Encrypt/decrypt with multiple algorithms
15. **Web Scraper** - Extract and store data from websites
16. **Collaborative Editor** - Real-time collaborative editing
17. **Job Queue System** - Background job processing
18. **Data Analytics Pipeline** - ETL and analysis

### Advanced Projects (4-8 weeks each)
19. **Microservices Architecture** - Multiple services with event bus
20. **Real-time Trading Platform** - Market data, order execution
21. **Distributed Cache** - Redis-like in-memory store
22. **GraphQL Server** - Complex schema with resolvers
23. **Container Orchestration Platform** - Simplified Kubernetes
24. **AI Chatbot** - NLP and conversation system
25. **Video Streaming Service** - Adaptive bitrate, streaming
26. **Search Engine** - Indexing and ranking system

## Best Practices by Language

### Python
- Use virtual environments (venv, poetry)
- Follow PEP 8 style guide
- Use type hints for clarity
- Document with docstrings
- Test with pytest
- Use dataclasses for data structures

### JavaScript
- Use const/let, never var
- Follow ESLint rules
- Use TypeScript for large projects
- Async/await over callbacks
- Test with Jest/Vitest
- Use modern frameworks (React, Vue, Svelte)

### Go
- Keep interfaces small (1-3 methods)
- Use defer for resource cleanup
- Handle errors explicitly
- Use goroutines for concurrency
- Test with built-in testing
- Follow effective Go conventions

### Java
- Use design patterns appropriately
- Leverage Java 17+ features
- Use dependency injection
- Write comprehensive tests
- Document public APIs
- Use build tools (Maven/Gradle)

### Rust
- Embrace the compiler's safety
- Use idiomatic patterns
- Write zero-copy code
- Test with cargo test
- Use appropriate lifetimes
- Leverage macros for metaprogramming

## 🏆 Success Milestones

- [ ] Master one primary language at intermediate level (3 months)
- [ ] Build 3 real projects in your primary language
- [ ] Learn basics of 2 additional languages
- [ ] Write idiomatic code (not "C code in Python")
- [ ] Understand language-specific design patterns
- [ ] Achieve 90%+ test coverage in projects
- [ ] Optimize code based on profiling data
- [ ] Contribute to open-source projects

---

**Ready to master programming languages? Start with your favorite language and build 3 projects!**
