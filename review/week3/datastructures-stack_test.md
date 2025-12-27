# Code Review: datastructures/stack/stack_test.go

## Overall Assessment

**File Purpose**: Unit tests for stack data structure  
**Rating**: 7.5/10  
**Status**: Good test coverage with room for improvement!

---

## Detailed Review

### ✅ STRENGTHS

1. **Table-Driven Tests**: Proper use of test cases!
2. **Multiple Scenarios**: Tests push, pop, empty cases
3. **Proper Structure**: Uses testing package correctly
4. **Edge Cases**: Tests empty stack behavior

**SOLID TESTING FOUNDATION!** 👍

### ⚠️ **ISSUES TO FIX**

#### Issue #1: Hardcoded Expected Values in Loop

```go
for i, v := range s.items {
    if v != i+1 {  // ❌ Assumes values are 1, 2, 3...
        t.Errorf("Expected %d, got %v", i+1, v)
    }
}
```

**Problem**: Test is coupled to specific push values.

**Better**:

```go
// Store expected values
expected := []any{1, 2, 3}

for i, v := range s.items {
    if v != expected[i] {
        t.Errorf("Index %d: expected %v, got %v", i, expected[i], v)
    }
}

// OR use test case expectations:
if got != tt.expected {
    t.Errorf("Expected %v, got %v", tt.expected, got)
}
```

#### Issue #2: Missing Assertions

**TestPop doesn't verify LIFO order!**

```go
func TestPop(t *testing.T) {
    s := &Stack{}
    s.Push(1)
    s.Push(2)
    s.Push(3)

    val := s.Pop()
    // ❌ No assertion that val == 3!
    fmt.Println("Popped item:", val)
}
```

**Fix**:

```go
func TestPop(t *testing.T) {
    s := &Stack{}
    s.Push(1)
    s.Push(2)
    s.Push(3)

    // Should pop in LIFO order
    if val := s.Pop(); val != 3 {
        t.Errorf("Expected 3, got %v", val)
    }
    if val := s.Pop(); val != 2 {
        t.Errorf("Expected 2, got %v", val)
    }
    if val := s.Pop(); val != 1 {
        t.Errorf("Expected 1, got %v", val)
    }
}
```

#### Issue #3: Not Testing Error Cases

```go
// Missing test: Pop from empty stack
// Missing test: Push to full stack (if capacity exists)
// Missing test: Peek on empty stack
```

**Add**:

```go
func TestPopEmpty(t *testing.T) {
    s := &Stack{}
    val := s.Pop()

    // Should return nil or error
    if val != nil {
        t.Errorf("Pop from empty should return nil, got %v", val)
    }
}
```

#### Issue #4: TestTestStack Has Confusing Name

```go
func TestTestStack(t *testing.T)  // ❌ "TestTest" is redundant
```

**Better**:

```go
func TestStackOperations(t *testing.T)
func TestStackWithMultipleValues(t *testing.T)
```

#### Issue #5: Not Testing All Methods

**Missing tests for**:

- `LengthOfStack()` explicitly
- `Peek()` if you add it
- `IsEmpty()` if you add it

### 📚 Comprehensive Test Examples

```go
package stack

import (
    "testing"
)

// Test 1: Basic Push
func TestPush(t *testing.T) {
    s := &Stack{}

    s.Push(1)
    if s.Len() != 1 {
        t.Errorf("Expected length 1, got %d", s.Len())
    }

    s.Push(2)
    s.Push(3)
    if s.Len() != 3 {
        t.Errorf("Expected length 3, got %d", s.Len())
    }
}

// Test 2: Pop maintains LIFO order
func TestPopLIFO(t *testing.T) {
    s := &Stack{}
    s.Push(1)
    s.Push(2)
    s.Push(3)

    // Pop should return 3, 2, 1 (LIFO)
    tests := []any{3, 2, 1}
    for i, expected := range tests {
        got := s.Pop()
        if got != expected {
            t.Errorf("Pop %d: expected %v, got %v", i, expected, got)
        }
    }
}

// Test 3: Pop from empty stack
func TestPopEmpty(t *testing.T) {
    s := &Stack{}

    val := s.Pop()
    if val != nil {
        t.Errorf("Pop from empty stack should return nil, got %v", val)
    }

    // Should still be empty
    if s.Len() != 0 {
        t.Errorf("Length should be 0, got %d", s.Len())
    }
}

// Test 4: Table-driven test for multiple operations
func TestStackOperations(t *testing.T) {
    tests := []struct {
        name      string
        pushes    []any
        pops      int
        wantLen   int
        wantPeek  any
    }{
        {
            name:     "single item",
            pushes:   []any{1},
            pops:     0,
            wantLen:  1,
            wantPeek: 1,
        },
        {
            name:     "multiple items",
            pushes:   []any{1, 2, 3},
            pops:     0,
            wantLen:  3,
            wantPeek: 3,
        },
        {
            name:     "push and pop",
            pushes:   []any{1, 2, 3, 4},
            pops:     2,
            wantLen:  2,
            wantPeek: 2,
        },
        {
            name:     "empty after pops",
            pushes:   []any{1, 2},
            pops:     2,
            wantLen:  0,
            wantPeek: nil,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            s := &Stack{}

            // Push items
            for _, item := range tt.pushes {
                s.Push(item)
            }

            // Pop items
            for i := 0; i < tt.pops; i++ {
                s.Pop()
            }

            // Check length
            if got := s.Len(); got != tt.wantLen {
                t.Errorf("Len() = %d, want %d", got, tt.wantLen)
            }

            // Check peek (if not empty)
            if tt.wantLen > 0 {
                // Assuming you add Peek() method
                // if got := s.Peek(); got != tt.wantPeek {
                //     t.Errorf("Peek() = %v, want %v", got, tt.wantPeek)
                // }
            }
        })
    }
}

// Test 5: Length tracking
func TestLength(t *testing.T) {
    s := &Stack{}

    if s.Len() != 0 {
        t.Errorf("New stack should have length 0, got %d", s.Len())
    }

    s.Push(1)
    if s.Len() != 1 {
        t.Errorf("After 1 push, length should be 1, got %d", s.Len())
    }

    s.Push(2)
    s.Push(3)
    if s.Len() != 3 {
        t.Errorf("After 3 pushes, length should be 3, got %d", s.Len())
    }

    s.Pop()
    if s.Len() != 2 {
        t.Errorf("After 1 pop, length should be 2, got %d", s.Len())
    }
}

// Test 6: Different types (since using any)
func TestDifferentTypes(t *testing.T) {
    s := &Stack{}

    s.Push(42)
    s.Push("hello")
    s.Push(3.14)
    s.Push(true)

    // Pop in reverse order
    if val := s.Pop(); val != true {
        t.Errorf("Expected true, got %v", val)
    }
    if val := s.Pop(); val != 3.14 {
        t.Errorf("Expected 3.14, got %v", val)
    }
    if val := s.Pop(); val != "hello" {
        t.Errorf("Expected 'hello', got %v", val)
    }
    if val := s.Pop(); val != 42 {
        t.Errorf("Expected 42, got %v", val)
    }
}

// Test 7: IsEmpty (if you add this method)
func TestIsEmpty(t *testing.T) {
    s := &Stack{}

    // Assuming you add IsEmpty() method
    /*
    if !s.IsEmpty() {
        t.Error("New stack should be empty")
    }

    s.Push(1)
    if s.IsEmpty() {
        t.Error("Stack with item should not be empty")
    }

    s.Pop()
    if !s.IsEmpty() {
        t.Error("Stack after popping all items should be empty")
    }
    */
}

// Test 8: Peek (if you add this method)
func TestPeek(t *testing.T) {
    s := &Stack{}

    // Assuming you add Peek() method
    /*
    // Peek on empty
    if val := s.Peek(); val != nil {
        t.Errorf("Peek on empty should return nil, got %v", val)
    }

    s.Push(1)
    s.Push(2)

    // Peek doesn't remove
    if val := s.Peek(); val != 2 {
        t.Errorf("Peek should return 2, got %v", val)
    }
    if s.Len() != 2 {
        t.Error("Peek should not change length")
    }

    // Peek again, same value
    if val := s.Peek(); val != 2 {
        t.Errorf("Peek should still return 2, got %v", val)
    }
    */
}

// Test 9: Push-Pop sequence
func TestPushPopSequence(t *testing.T) {
    s := &Stack{}

    // Interleaved pushes and pops
    s.Push(1)
    s.Push(2)

    if val := s.Pop(); val != 2 {
        t.Errorf("Expected 2, got %v", val)
    }

    s.Push(3)
    s.Push(4)

    if val := s.Pop(); val != 4 {
        t.Errorf("Expected 4, got %v", val)
    }
    if val := s.Pop(); val != 3 {
        t.Errorf("Expected 3, got %v", val)
    }
    if val := s.Pop(); val != 1 {
        t.Errorf("Expected 1, got %v", val)
    }
}

// Test 10: Large number of items
func TestLargeStack(t *testing.T) {
    s := &Stack{}

    n := 10000

    // Push n items
    for i := 0; i < n; i++ {
        s.Push(i)
    }

    if s.Len() != n {
        t.Errorf("Expected length %d, got %d", n, s.Len())
    }

    // Pop all items in reverse order
    for i := n - 1; i >= 0; i-- {
        val := s.Pop()
        if val != i {
            t.Errorf("Expected %d, got %v", i, val)
        }
    }

    if s.Len() != 0 {
        t.Errorf("Stack should be empty, length is %d", s.Len())
    }
}

// BENCHMARK TESTS

func BenchmarkPush(b *testing.B) {
    s := &Stack{}
    for i := 0; i < b.N; i++ {
        s.Push(i)
    }
}

func BenchmarkPop(b *testing.B) {
    s := &Stack{}
    for i := 0; i < b.N; i++ {
        s.Push(i)
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        s.Pop()
    }
}

func BenchmarkPushPop(b *testing.B) {
    s := &Stack{}
    for i := 0; i < b.N; i++ {
        s.Push(i)
        s.Pop()
    }
}

/*
TESTING BEST PRACTICES:

1. TEST CATEGORIES:
   - Happy path (normal usage)
   - Edge cases (empty, full, single item)
   - Error cases (invalid operations)
   - Boundary conditions (max size, nil values)
   - Performance (benchmarks)

2. WHAT TO TEST:
   ✅ Push increases length
   ✅ Pop decreases length
   ✅ Pop returns LIFO order
   ✅ Pop from empty returns nil/error
   ✅ Peek doesn't modify stack
   ✅ IsEmpty works correctly
   ✅ Different data types (if using any)

3. ASSERTIONS:
   - Check return values
   - Check state changes (length)
   - Check side effects
   - Check error conditions

4. TABLE-DRIVEN TESTS:
   - Test multiple scenarios efficiently
   - Clear test case descriptions
   - Easy to add new cases

5. BENCHMARKS:
   - Measure Push performance
   - Measure Pop performance
   - Compare with alternative implementations

6. COVERAGE:
   go test -cover
   - Aim for > 80% coverage
   - Cover all public methods
   - Cover error paths

RUN TESTS:
go test -v               # Verbose output
go test -run TestPush   # Run specific test
go test -cover          # With coverage
go test -bench .        # Run benchmarks
go test -race           # Race detection
*/
```

---

## Rating Breakdown

| Category     | Rating | Comments                       |
| ------------ | ------ | ------------------------------ |
| Structure    | 9/10   | Good use of table-driven tests |
| Coverage     | 6/10   | Missing edge cases             |
| Assertions   | 6/10   | TestPop missing assertions     |
| Completeness | 7/10   | Not testing all methods        |
| Code Quality | 8/10   | Clean, readable                |

**Overall: 7.5/10**

---

## Key Takeaways

1. ✅ **Table-driven tests** - you're using them!
2. ⚠️ **Add assertions** in TestPop (verify values!)
3. ⚠️ **Test edge cases** (empty, single item)
4. ✅ **Test error conditions** (pop from empty)
5. ✅ **Run with coverage**: `go test -cover`
6. ✅ **Add benchmarks** for performance testing

---

## Next Steps

1. ✅ Add assertions to TestPop (verify LIFO!)
2. ✅ Test Pop from empty stack
3. ✅ Test Length() method explicitly
4. ✅ Add TestIsEmpty (when you implement it)
5. ✅ Add TestPeek (when you implement it)
6. ✅ Add benchmarks (BenchmarkPush, BenchmarkPop)
7. ✅ Run: `go test -cover` and aim for > 80%

**Good testing foundation! Add more assertions and edge cases to make it comprehensive!** 🧪
