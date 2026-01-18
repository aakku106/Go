# Code Review: datastructures/stack/stack_test.go

**File**: `datastructures/stack/stack_test.go`  
**Category**: Stack Tests  
**Lines**: 41  
**Rating**: 5/10

---

## Overview

Tests basic stack operations: Push, Pop, and length checking. Covers happy path but directly accesses private `stack` field instead of using public API. Tests are minimal with no edge cases. All tests pass.

---

## Strengths

1. **Multiple Test Functions** - Separates Push/Pop/Length
2. **Type Variety** - Tests with int and string
3. **Failure Messages** - Uses `t.Errorf` with context
4. **Tests Pass** - All functionality verified at basic level
5. **Explicit Values** - Clear test data (10, 20, 30, "geek", "for", "geeks")

---

## Issues

### Critical

**1. Direct Field Access**

```go
func TestStack_Push(t *testing.T) {
    var stack Stack
    stack.Push(10)
    stack.Push(20)
    stack.Push(30)
    if len(stack.stack) != 3 {  // ← Accessing private field
        t.Errorf("Expected stack length 3, got %d", len(stack.stack))
    }
}
```

**Problems**:

- Accesses `stack.stack` (internal field)
- Breaks encapsulation
- Tests implementation, not interface
- Should use `stack.LengthOfStack()` (or better, `stack.Len()`)

**Why This Breaks**:

If implementation changes from slice to linked list:

```go
type Stack struct {
    top  *node  // No longer []any
    size int
}
```

**Tests break** even though public API works fine.

**Should be**:

```go
if stack.LengthOfStack() != 3 {
    t.Errorf("Expected stack length 3, got %d", stack.LengthOfStack())
}
```

Or better:

```go
if stack.Len() != 3 {
    t.Errorf("Expected stack length 3, got %d", stack.Len())
}
```

**2. No Empty Stack Tests**

```go
func TestStack_Pop(t *testing.T) {
    // Never tests popping from empty stack
}
```

**Missing**:

```go
func TestStack_PopEmpty(t *testing.T) {
    var stack Stack
    val, ok := stack.Pop()
    if ok {
        t.Errorf("Expected Pop from empty stack to fail, got (%v, %v)", val, ok)
    }
    if val != nil {
        t.Errorf("Expected nil value, got %v", val)
    }
}
```

**3. No Overflow Tests**

```go
func TestStack_LengthOfStack(t *testing.T) {
    // Never tests the uint16 overflow panic
}
```

Given that `LengthOfStack()` panics at 65,536 elements, should test:

```go
func TestStack_LengthOverflow(t *testing.T) {
    var stack Stack
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("Expected panic for stack > 65535 elements")
        }
    }()

    for i := 0; i <= math.MaxUint16+1; i++ {
        stack.Push(i)
    }
    stack.LengthOfStack()  // Should panic
}
```

(Though the real fix is deleting the uint16 cap entirely.)

### Major

**1. Incomplete Pop Test**

```go
func TestStack_Pop(t *testing.T) {
    var stack Stack
    stack.Push("geek")
    stack.Push("for")
    stack.Push("geeks")

    val, ok := stack.Pop()
    if !ok {
        t.Errorf("Expected successful Pop, got false")
    }
    if val != "geeks" {
        t.Errorf("Expected 'geeks', got %v", val)
    }

    // ← Stops here, doesn't verify LIFO order
}
```

**Should verify complete LIFO**:

```go
// Pop all and verify order
expected := []string{"geeks", "for", "geek"}
for i, want := range expected {
    val, ok := stack.Pop()
    if !ok {
        t.Errorf("Pop %d failed", i)
    }
    if val != want {
        t.Errorf("Pop %d: expected %v, got %v", i, want, val)
    }
}

// Verify empty
if _, ok := stack.Pop(); ok {
    t.Error("Expected empty stack after all pops")
}
```

**2. No Length Assertions After Operations**

```go
func TestStack_Pop(t *testing.T) {
    stack.Push("geek")
    stack.Push("for")
    stack.Push("geeks")
    // ← Should check length == 3

    stack.Pop()
    // ← Should check length == 2
}
```

**3. No Table-Driven Tests**

Tests repeat similar logic:

```go
func TestStack_Push(t *testing.T) {
    var stack Stack
    stack.Push(10)
    stack.Push(20)
    stack.Push(30)
}

func TestStack_LengthOfStack(t *testing.T) {
    var stack Stack
    stack.Push("geek")
    stack.Push("for")
    stack.Push("geeks")
}
```

Could combine:

```go
func TestStack_Operations(t *testing.T) {
    tests := []struct {
        name   string
        values []any
        want   int
    }{
        {"integers", []any{10, 20, 30}, 3},
        {"strings", []any{"geek", "for", "geeks"}, 3},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            var stack Stack
            for _, v := range tt.values {
                stack.Push(v)
            }
            if got := stack.LengthOfStack(); got != tt.want {
                t.Errorf("got %d, want %d", got, tt.want)
            }
        })
    }
}
```

### Minor

**1. Variable Shadowing**

```go
func TestStack_Push(t *testing.T) {
    var stack Stack  // ← Named "stack"
}
```

While legal, naming variable same as type can be confusing. Consider `s` or `stk`:

```go
func TestStack_Push(t *testing.T) {
    var s Stack
}
```

**2. No Benchmarks**

Missing performance tests:

```go
func BenchmarkStack_Push(b *testing.B) {
    var s Stack
    for i := 0; i < b.N; i++ {
        s.Push(i)
    }
}
```

**3. Method Name in Tests**

```go
func TestStack_LengthOfStack(t *testing.T) {
```

Verbose. If method renamed to `Len()`, test becomes:

```go
func TestStack_Len(t *testing.T) {
```

---

## Suggested Improvements

1. **Use public API** - Replace `len(stack.stack)` with `stack.LengthOfStack()`
2. **Test empty stack** - Add Pop from empty test
3. **Test overflow** - Test uint16 panic (or delete cap and test large stacks)
4. **Complete LIFO test** - Pop all elements, verify order
5. **Add length assertions** - Check length after each operation
6. **Table-driven tests** - Reduce duplication
7. **Add edge cases**:
   - Pop until empty
   - Push after Pop
   - Large stack (if cap removed)
8. **Add benchmarks** - Performance testing

---

## Better Test Suite

```go
package stack

import (
    "testing"
)

func TestStack_PushPop(t *testing.T) {
    tests := []struct {
        name   string
        values []any
    }{
        {"integers", []any{10, 20, 30}},
        {"strings", []any{"a", "b", "c"}},
        {"mixed", []any{1, "two", 3.0}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            var s Stack

            // Push all
            for _, v := range tt.values {
                s.Push(v)
            }

            if got := s.Len(); got != len(tt.values) {
                t.Fatalf("After pushes: got len %d, want %d", got, len(tt.values))
            }

            // Pop all (LIFO order)
            for i := len(tt.values) - 1; i >= 0; i-- {
                val, ok := s.Pop()
                if !ok {
                    t.Fatalf("Pop %d failed", i)
                }
                if val != tt.values[i] {
                    t.Errorf("Pop %d: got %v, want %v", i, val, tt.values[i])
                }
                if got := s.Len(); got != i {
                    t.Errorf("After pop %d: got len %d, want %d", i, got, i)
                }
            }

            // Verify empty
            if _, ok := s.Pop(); ok {
                t.Error("Expected empty stack")
            }
        })
    }
}

func TestStack_PopEmpty(t *testing.T) {
    var s Stack
    val, ok := s.Pop()
    if ok {
        t.Errorf("Expected false, got true with value %v", val)
    }
}

func BenchmarkStack_Push(b *testing.B) {
    var s Stack
    for i := 0; i < b.N; i++ {
        s.Push(i)
    }
}
```

---

## What You Learned

✅ Basic test structure  
✅ Multiple test functions  
✅ Type variety testing  
❌ Encapsulation (accessing private fields)  
❌ Edge case testing  
❌ Complete behavior verification

---

## Coverage Analysis

Estimated coverage: **~60%**

**Covered**:

- ✅ Push with values
- ✅ Pop with values (partially)
- ✅ LengthOfStack with values

**Not Covered**:

- ❌ Pop from empty stack
- ❌ Pop until empty
- ❌ uint16 overflow panic
- ❌ Push after Pop
- ❌ Length changes during operations

---

## Comparison to Other Test Files

| File                     | Rating   | Coverage | Edge Cases | Best Aspect              |
| ------------------------ | -------- | -------- | ---------- | ------------------------ |
| SinglyLinkedList_test.go | 4/10     | ~40%     | None       | Attempted multiple tests |
| linearQueue_test.go      | 6/10     | ~60%     | Some       | Size assertions          |
| prorityQueue_test.go     | 7/10     | ~70%     | Some       | Priority ordering        |
| **stack_test.go**        | **5/10** | **~60%** | **None**   | **Type variety**         |

**Stack tests middle of pack**. Better than linked list (broken test names), worse than priority queue (good coverage).

**Common pattern**: All test files access internals directly.

---

## Final Verdict

**5/10** - Basic tests verify Push/Pop work, but directly access private `stack` field breaking encapsulation, miss all edge cases (empty stack, overflow, LIFO order), and don't fully verify behavior.

**Functional**: Tests pass and verify basic operations.  
**Design**: Breaks encapsulation with direct field access.  
**Coverage**: Only ~60%, misses edge cases.

**Main problem**: Tests would break if implementation changed, even if public API stayed same.

**Fix priority**:

1. Replace `len(stack.stack)` with `stack.LengthOfStack()`
2. Add empty stack test
3. Complete LIFO verification in Pop test
4. Add length assertions after each operation

**Compared to stack.go (4/10)**: Tests are better than implementation (implementation has arbitrary cap + rambling comment, tests at least work).
