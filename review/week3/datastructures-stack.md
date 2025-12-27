# Code Review: datastructures/stack/stack.go

## Overall Assessment

**File Purpose**: Generic stack data structure implementation  
**Rating**: 8/10  
**Status**: Clean generic implementation with proper methods!

---

## Detailed Review

### ✅ STRENGTHS

1. **Generic Type**: Uses `any` for type flexibility
2. **Clean Interface**: Push, Pop, LengthOfStack methods
3. **Proper Encapsulation**: Stack struct with internal slice
4. **Idiomatic Go**: Simple, straightforward design
5. **Package Structure**: Separate from main learning code

**PROFESSIONAL APPROACH!** 👍

### ⚠️ **ISSUES TO FIX**

#### Issue #1: Method Naming

```go
func (s *Stack) LengthOfStack() int {
    return len(s.items)
}
```

**Not idiomatic!** In Go, prefer short, simple names:

```go
// ❌ Verbose:
func (s *Stack) LengthOfStack() int

// ✅ Idiomatic:
func (s *Stack) Len() int
func (s *Stack) Size() int
```

**Rule**: Method names don't need to repeat type name. We already know it's a stack from `s.`

#### Issue #2: Pop() Doesn't Return Error

```go
func (s *Stack) Pop() any {
    if len(s.items) == 0 {
        return nil
    }
    // ...
}
```

**Problem**: Returning `nil` on empty is ambiguous!

- What if `nil` was pushed onto stack?
- How to distinguish empty vs nil value?

**Better Design**:

```go
// Option 1: Return error
func (s *Stack) Pop() (any, error) {
    if len(s.items) == 0 {
        return nil, errors.New("stack is empty")
    }
    index := len(s.items) - 1
    value := s.items[index]
    s.items = s.items[:index]
    return value, nil
}

// Option 2: Boolean indicator
func (s *Stack) Pop() (any, bool) {
    if len(s.items) == 0 {
        return nil, false
    }
    index := len(s.items) - 1
    value := s.items[index]
    s.items = s.items[:index]
    return value, true
}
```

#### Issue #3: Missing Peek() Method

**Common operation**: Look at top without removing.

```go
func (s *Stack) Peek() (any, bool) {
    if len(s.items) == 0 {
        return nil, false
    }
    return s.items[len(s.items)-1], true
}
```

#### Issue #4: Missing IsEmpty() Method

```go
func (s *Stack) IsEmpty() bool {
    return len(s.items) == 0
}
```

Clearer than checking `s.LengthOfStack() == 0`.

#### Issue #5: No Capacity Limit

**Real-world stacks** often have max capacity:

```go
type Stack struct {
    items    []any
    capacity int  // 0 = unlimited
}

func (s *Stack) Push(item any) error {
    if s.capacity > 0 && len(s.items) >= s.capacity {
        return errors.New("stack overflow")
    }
    s.items = append(s.items, item)
    return nil
}
```

#### Issue #6: Not Using Go Generics (1.18+)

**Since Go 1.18**, you can use actual generics instead of `any`:

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    var zero T
    if len(s.items) == 0 {
        return zero, false
    }
    index := len(s.items) - 1
    value := s.items[index]
    s.items = s.items[:index]
    return value, true
}

// Usage:
intStack := Stack[int]{}
intStack.Push(42)

stringStack := Stack[string]{}
stringStack.Push("hello")
```

**Benefits**:

- Type safety at compile time
- No type assertions needed
- Better performance (no interface overhead)

### 📚 Stack Deep Dive

#### What is a Stack?

**LIFO**: Last In, First Out

```
Push(1) → [1]
Push(2) → [1, 2]
Push(3) → [1, 2, 3]
Pop()   → 3, [1, 2]
Pop()   → 2, [1]
Pop()   → 1, []
```

#### Common Operations

| Operation | Description     | Time Complexity |
| --------- | --------------- | --------------- |
| Push(x)   | Add to top      | O(1) amortized  |
| Pop()     | Remove from top | O(1)            |
| Peek()    | View top        | O(1)            |
| IsEmpty() | Check if empty  | O(1)            |
| Size()    | Get count       | O(1)            |

#### Real-World Uses

1. **Function Call Stack**: Runtime execution
2. **Undo/Redo**: Editor operations
3. **Expression Evaluation**: `3 + 4 * 5`
4. **Backtracking**: Maze solving, parsing
5. **Browser History**: Back button

### 🔧 Complete Stack Implementation

```go
package stack

import (
    "errors"
    "fmt"
)

// Generic stack using Go 1.18+ generics
type Stack[T any] struct {
    items    []T
    capacity int // 0 = unlimited
}

// NewStack creates a new stack with optional capacity limit
func NewStack[T any](capacity int) *Stack[T] {
    return &Stack[T]{
        items:    make([]T, 0),
        capacity: capacity,
    }
}

// Push adds item to top of stack
func (s *Stack[T]) Push(item T) error {
    if s.capacity > 0 && len(s.items) >= s.capacity {
        return errors.New("stack overflow: capacity reached")
    }
    s.items = append(s.items, item)
    return nil
}

// Pop removes and returns top item
func (s *Stack[T]) Pop() (T, error) {
    var zero T
    if s.IsEmpty() {
        return zero, errors.New("stack underflow: empty stack")
    }
    index := len(s.items) - 1
    value := s.items[index]
    s.items = s.items[:index]
    return value, nil
}

// Peek returns top item without removing
func (s *Stack[T]) Peek() (T, error) {
    var zero T
    if s.IsEmpty() {
        return zero, errors.New("empty stack")
    }
    return s.items[len(s.items)-1], nil
}

// IsEmpty checks if stack is empty
func (s *Stack[T]) IsEmpty() bool {
    return len(s.items) == 0
}

// Len returns number of items
func (s *Stack[T]) Len() int {
    return len(s.items)
}

// Cap returns capacity (0 = unlimited)
func (s *Stack[T]) Cap() int {
    return s.capacity
}

// Clear removes all items
func (s *Stack[T]) Clear() {
    s.items = s.items[:0]
}

// String implements fmt.Stringer
func (s *Stack[T]) String() string {
    return fmt.Sprintf("Stack%v", s.items)
}

// YOUR ORIGINAL IMPLEMENTATION (with fixes):

type StackAny struct {
    items []any
}

func NewStackAny() *StackAny {
    return &StackAny{
        items: make([]any, 0),
    }
}

func (s *StackAny) Push(item any) {
    s.items = append(s.items, item)
}

func (s *StackAny) Pop() (any, bool) {
    if len(s.items) == 0 {
        return nil, false  // ✅ Return false to indicate empty
    }
    index := len(s.items) - 1
    value := s.items[index]
    s.items = s.items[:index]
    return value, true
}

func (s *StackAny) Peek() (any, bool) {
    if len(s.items) == 0 {
        return nil, false
    }
    return s.items[len(s.items)-1], true
}

func (s *StackAny) Len() int {  // ✅ Renamed from LengthOfStack
    return len(s.items)
}

func (s *StackAny) IsEmpty() bool {
    return len(s.items) == 0
}
```

### 🎯 Usage Examples

```go
package main

import (
    "fmt"
    "yourproject/stack"
)

func main() {
    demonstrateGenericStack()
    demonstrateStackOperations()
    demonstrateRealWorldUse()
}

func demonstrateGenericStack() {
    fmt.Println("=== Generic Stack ===")

    // Type-safe integer stack
    intStack := stack.NewStack[int](0)
    intStack.Push(1)
    intStack.Push(2)
    intStack.Push(3)

    val, _ := intStack.Pop()
    fmt.Printf("Popped: %d\n", val)  // 3

    // Type-safe string stack
    strStack := stack.NewStack[string](0)
    strStack.Push("hello")
    strStack.Push("world")

    val2, _ := strStack.Pop()
    fmt.Printf("Popped: %s\n\n", val2)  // "world"
}

func demonstrateStackOperations() {
    fmt.Println("=== Stack Operations ===")

    s := stack.NewStack[int](0)

    // Push
    s.Push(10)
    s.Push(20)
    s.Push(30)
    fmt.Printf("After pushes: len=%d\n", s.Len())

    // Peek
    top, _ := s.Peek()
    fmt.Printf("Top (peek): %d\n", top)
    fmt.Printf("Len after peek: %d\n", s.Len())

    // Pop
    val, _ := s.Pop()
    fmt.Printf("Popped: %d\n", val)
    fmt.Printf("Len after pop: %d\n", s.Len())

    // IsEmpty
    fmt.Printf("Is empty? %v\n\n", s.IsEmpty())
}

func demonstrateRealWorldUse() {
    fmt.Println("=== Real-World: Undo/Redo ===")

    type Action struct {
        Type string
        Data string
    }

    undoStack := stack.NewStack[Action](10)

    // Perform actions
    undoStack.Push(Action{"INSERT", "Hello"})
    undoStack.Push(Action{"INSERT", " World"})
    undoStack.Push(Action{"DELETE", "World"})

    // Undo last action
    action, err := undoStack.Pop()
    if err == nil {
        fmt.Printf("Undo: %s %s\n", action.Type, action.Data)
    }
}

/*
STACK IMPLEMENTATION PATTERNS:

1. SLICE-BASED (most common in Go):
   items []T
   - Push: append
   - Pop: items[len-1], reslice
   - Simple, efficient

2. LINKED-LIST BASED:
   type Node struct {
       value T
       next  *Node
   }
   - Push: prepend node
   - Pop: remove head
   - No reallocation

3. ARRAY-BASED (fixed size):
   items [N]T
   top   int
   - Push: items[top++] = value
   - Pop: return items[--top]
   - No allocation, cache-friendly

GO BEST PRACTICES:

✅ Use generics (Go 1.18+) for type safety
✅ Return (value, error) or (value, bool)
✅ Provide IsEmpty() method
✅ Provide Peek() for non-destructive read
✅ Use short method names (Len not LengthOfStack)
❌ Don't return nil for empty (ambiguous)
❌ Don't panic on Pop from empty
❌ Don't use interface{} if generics available

THREAD SAFETY:

Your stack is NOT thread-safe. For concurrent use:

type SafeStack[T any] struct {
    items []T
    mu    sync.Mutex
}

func (s *SafeStack[T]) Push(item T) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.items = append(s.items, item)
}
*/
```

---

## Rating Breakdown

| Category       | Rating | Comments                    |
| -------------- | ------ | --------------------------- |
| Design         | 8/10   | Clean, simple interface     |
| Type Safety    | 6/10   | Uses any, not generics      |
| Error Handling | 6/10   | Pop returns nil (ambiguous) |
| Completeness   | 7/10   | Missing Peek, IsEmpty       |
| Code Quality   | 9/10   | Clean, readable             |

**Overall: 8/10**

---

## Key Takeaways

1. ✅ **Stack = LIFO** (Last In, First Out)
2. ⚠️ **Use generics** instead of `any` (Go 1.18+)
3. ⚠️ **Return errors** from Pop, don't return nil
4. ✅ **Add Peek()** for non-destructive read
5. ✅ **Short method names** (`Len`, not `LengthOfStack`)
6. ✅ **Provide IsEmpty()** for clarity

---

## Next Steps

1. ✅ Migrate to generics: `Stack[T any]`
2. ✅ Fix Pop to return `(T, error)` or `(T, bool)`
3. ✅ Add Peek() method
4. ✅ Add IsEmpty() method
5. ✅ Rename LengthOfStack() to Len()
6. ✅ Add Clear() method
7. ✅ Consider thread-safe version with mutex

**Good first data structure! Clean and simple. Upgrade to generics for type safety!** 📚
