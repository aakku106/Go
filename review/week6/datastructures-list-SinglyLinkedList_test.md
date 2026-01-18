# Code Review: datastructures/list/SinglyLinkedList_test.go

**File**: `datastructures/list/SinglyLinkedList_test.go`  
**Topic**: Singly Linked List Tests with Assertions  
**Rating**: 7.5/10  
**Reviewed**: January 18, 2026

---

## Summary

Good test coverage with **assertions added** (improvement from Week 5). Tests verify InsertAtBeginning, InsertAtLast, InsertAfter, and InsertBefore. **Critical bug**: Line 11 has `T.Fatal` (capital T) instead of `t.Fatal` - this test won't compile. Otherwise solid work.

---

## What You Did

### 1. Tests with Assertions (Week 5 Improvement!)

```go
func TestInsertAtBeginning(t *testing.T) {
    list := NewSinglyLinkedList(106)
    list = list.InsertAtBeginning("aww").(*SingelyLinkList)

    if list.data != "aww" {
        t.Errorf("Expected head data to be 'aww', got %v", list.data)
    }
}
```

**Week 5 had NO assertions.** Week 6 has comprehensive checks. Excellent improvement.

### 2. Test Order Verification

```go
// Verify order: cat -> aww -> 106
if list.data != "cat" {
    t.Errorf("Expected head data to be 'cat', got %v", list.data)
}
if list.next.data != "aww" {
    t.Errorf("Expected second node data to be 'aww', got %v", list.next.data)
}
if list.next.next.data != 106 {
    t.Errorf("Expected third node data to be 106, got %v", list.next.next.data)
}
```

Verifies entire list structure. Good thoroughness.

### 3. Error Case Testing

```go
// Test out of bounds index - should return error
result, err = list.InsertAfter("out of bounds", 100)
if err == nil {
    t.Error("Expected error for out of bounds index, got nil")
}
if result != nil {
    t.Error("Expected nil result when error occurs")
}
```

Tests error conditions. This is proper testing.

---

## Critical Issue

### Line 11: Capital T Typo

```go
func testNewSinglyLinkedList(t *testing.T) {  // lowercase "test" - won't run!
    head := NewSinglyLinkedList(123)

    if head == nil {
        T.Fatal("NewSinglyLinkedList() returned nil, expected a node")  // ❌ Capital T!
    }
```

**TWO bugs here**:

1. Function name `testNewSinglyLinkedList` (lowercase "test") - won't run with `go test`
2. Line 11: `T.Fatal` should be `t.Fatal` (lowercase t)

**This test won't compile.** Run `go test` and you'll see error.

**Fix**:

```go
func TestNewSinglyLinkedList(t *testing.T) {  // Capital T in "Test"
    head := NewSinglyLinkedList(123)

    if head == nil {
        t.Fatal("NewSinglyLinkedList() returned nil, expected a node")  // lowercase t
    }
```

---

## What's Good

### 1. Week 5 Regression FIXED

**Week 5**: Tests had ZERO assertions (just printed output)  
**Week 6**: All tests have proper assertions

**This shows you learned from Week 4's feedback.**

### 2. Comprehensive Test Cases

```go
TestInsertAtBeginning  // Tests head insertion
TestInsertAtLast       // Tests tail insertion
TestInsertAfter        // Tests middle insertion + error case
TestInsertBefore       // Tests before insertion + index 0 edge case
```

Good coverage of all insert methods.

### 3. Edge Case Testing

```go
// TestInsertBefore with index 0
result, err = list.InsertBefore(106, 1)
```

Tests edge cases like inserting at beginning via InsertBefore.

### 4. Error Verification

```go
if err != nil {
    t.Fatalf("Unexpected error: %v", err)
}
```

Fails fast if unexpected error occurs. Good practice.

---

## Issues

### Major: T.Fatal Won't Compile

Line 11: `T.Fatal` should be `t.Fatal`

### Minor: testNewSinglyLinkedList Won't Run

```go
func testNewSinglyLinkedList(t *testing.T) {  // lowercase "test"
```

Should be:

```go
func TestNewSinglyLinkedList(t *testing.T) {  // Capital "Test"
```

### Minor: TestTEmp is Debug Code

```go
func TestTEmp(t *testing.T) {
    list := NewSinglyLinkedList(106)
    var result LinkList
    result, _ = list.InsertBefore("wee", 0)
    result, _ = result.InsertBefore("awww", 0)
    result, _ = result.InsertBefore("awww", 10)  // This will error
    result.PrintList()
}
```

This looks like debugging/experimentation code. Should be removed or renamed.

**Issues**:

- No assertions (just prints)
- Ignores errors with `_`
- Line 3 will error (index 10 out of bounds)

---

## What Could Be Better

### 1. Table-Driven Tests

```go
func TestInsertAtBeginning(t *testing.T) {
    tests := []struct {
        name     string
        initial  any
        inserts  []any
        expected []any
    }{
        {"single insert", 106, []any{"aww"}, []any{"aww", 106}},
        {"multiple inserts", 106, []any{"aww", "cat"}, []any{"cat", "aww", 106}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test logic here
        })
    }
}
```

This is Go's idiomatic testing pattern.

### 2. Helper Function for List Verification

```go
func verifyList(t *testing.T, list *SingelyLinkList, expected []any) {
    current := list
    for i, want := range expected {
        if current == nil {
            t.Fatalf("List ended early at index %d", i)
        }
        if current.data != want {
            t.Errorf("Index %d: got %v, want %v", i, current.data, want)
        }
        current = current.next
    }
    if current != nil {
        t.Error("List has more elements than expected")
    }
}
```

Reduces repetitive verification code.

---

## Progress from Week 5

**Week 5 Critical Issues**:

1. Tests had NO assertions ❌
2. Two test functions had lowercase 't' (wouldn't run) ❌

**Week 6 Status**:

1. Tests now have assertions ✅
2. **Still has uppercase T typo** ❌
3. **Still has lowercase test function** ❌

**Partial improvement.** Assertions added (major win), but typos remain.

---

## Recommendations

### Critical

**1. Fix Line 11**:

```go
// Current (won't compile):
T.Fatal("...")

// Fix:
t.Fatal("...")
```

**2. Fix Function Name**:

```go
// Current (won't run):
func testNewSinglyLinkedList(t *testing.T) {

// Fix:
func TestNewSinglyLinkedList(t *testing.T) {
```

**3. Run Tests**:

```bash
cd datastructures/list
go test -v
```

This will show the compilation error.

### Major

**4. Remove or Fix TestTEmp**:

Either add assertions:

```go
func TestInsertBeforeEdgeCases(t *testing.T) {
    list := NewSinglyLinkedList(106)

    // Test index 0
    result, err := list.InsertBefore("wee", 0)
    if err != nil {
        t.Fatal(err)
    }

    // Test out of bounds
    _, err = result.InsertBefore("awww", 100)
    if err == nil {
        t.Error("Expected error for out of bounds index")
    }
}
```

Or remove it.

---

## Rating Justification

**7.5/10**

**Good**:

- Assertions added from Week 5 (+3)
- Comprehensive test coverage (+2)
- Error case testing (+1.5)
- Edge case testing (+1)
- Clear test organization (+1)

**Bad**:

- T.Fatal typo (won't compile) (-1)
- testNewSinglyLinkedList won't run (-0.5)
- TestTEmp is debug code (-0.5)

**Huge improvement from Week 5.** Assertions show you learned from feedback. But typos prevent tests from running. **Run `go test` to catch these errors.**

---

## What You're Learning

1. Test assertions (t.Error, t.Errorf, t.Fatal)
2. Error case testing
3. Edge case identification
4. Test organization
5. Nil checking
6. List structure verification

**You're writing real tests now.** Assertions verify behavior instead of just printing. Fix the typos and these tests will be solid.
