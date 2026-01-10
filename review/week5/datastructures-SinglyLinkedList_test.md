# Code Review: datastructures/list/SinglyLinkedList_test.go

**File**: `datastructures/list/SinglyLinkedList_test.go`  
**Topic**: Singly Linked List Tests  
**Rating**: 4/10  
**Reviewed**: January 11, 2026

---

## Summary

Test file with 3 functions, but only 1 runs automatically. Missing test assertions - tests only print output without verification.

---

## Critical Issue: Test Naming

```go
func testCreate(t *testing.T) {           // ❌ lowercase 't'
func testInsertAtBeginning(t *testing.T) {  // ❌ lowercase 't'
func TestInsertAtLast(t *testing.T) {      // ✓ uppercase 'T'
```

**Only `TestInsertAtLast` runs with `go test`.** The other two are ignored.

**You JUST fixed this in queue tests.** Why did you make the same mistake again?

---

## What You Did

### Test 1: testCreate (doesn't run)

```go
func testCreate(t *testing.T) {
    head := Create(123)
    head.PrintList()
}
```

Creates a node with value 123, prints it.

### Test 2: testInsertAtBeginning (doesn't run)

```go
func testInsertAtBeginning(t *testing.T) {
    list := Create(106)
    list.PrintList()
    list = list.InsertAtBeginning("aww")
    list.PrintList()
    list = list.InsertAtBeginning("cat")
    list.PrintList()
}
```

Tests inserting at beginning with type mixing (int, string).

### Test 3: TestInsertAtLast (runs)

```go
func TestInsertAtLast(t *testing.T) {
    list := Create("Car")
    list.PrintList()
    list = list.InsertAtLast("weee")
    list.PrintList()
    list = list.InsertAtLast("Awwww")
    list.PrintList()
    list = list.InsertAtBeginning(106)
    list.PrintList()
}
```

Tests insert at last and beginning with mixed types.

---

## Major Issues

### Issue 1: No Assertions

**Every test just prints, none verify correctness.**

```go
func testCreate(t *testing.T) {
    head := Create(123)
    head.PrintList()  // Prints, but doesn't test anything!
}
```

**What you should do**:

```go
func TestCreate(t *testing.T) {
    head := Create(123)
    if head == nil {
        t.Fatal("Create returned nil")
    }
    if head.data != 123 {
        t.Fatalf("Expected data 123, got %v", head.data)
    }
    if head.next != nil {
        t.Fatal("New node should have nil next")
    }
}
```

### Issue 2: PrintList in Tests

```go
list.PrintList()  // Manual verification?
```

**Tests should be automated.** If you have to read the output to verify correctness, it's not a real test.

### Issue 3: Commented Placeholders

```go
//func TestInsertAfter(t *testing.T) {
//
//}
```

**Delete these.** Commented placeholders are noise.

---

## What's Good

### 1. TestInsertAtLast Has Correct Name

```go
func TestInsertAtLast(t *testing.T) {  // ✓ Will run with go test
```

One out of three functions has the correct name.

### 2. Mixed Type Testing

```go
list := Create("Car")
// ...
list = list.InsertAtBeginning(106)  // int after strings
```

Good testing practice - verifies `any` type works with different types.

---

## What This Test File Actually Tests

**Nothing.** It prints output. A test that doesn't assert is not a test.

Run `go test` in this directory:

```bash
$ go test
ok      list    0.002s
```

It passes, but it didn't verify **anything** except "the code didn't crash."

---

## Recommendations

### Critical: Fix Test Names

```go
func TestCreate(t *testing.T) {           // ✓ Uppercase T
func TestInsertAtBeginning(t *testing.T) { // ✓ Uppercase T
func TestInsertAtLast(t *testing.T) {     // ✓ Already correct
```

### Critical: Add Assertions

```go
func TestInsertAtBeginning(t *testing.T) {
    list := Create(106)

    // Verify initial state
    if list.data != 106 {
        t.Fatalf("Expected 106, got %v", list.data)
    }

    // Insert at beginning
    list = list.InsertAtBeginning("aww")

    // Verify new head
    if list.data != "aww" {
        t.Fatalf("Expected 'aww' at head, got %v", list.data)
    }

    // Verify old head is second
    if list.next == nil {
        t.Fatal("Second node should not be nil")
    }
    if list.next.data != 106 {
        t.Fatalf("Expected 106 in second node, got %v", list.next.data)
    }
}
```

### Major: Remove PrintList from Tests

```go
// Wrong:
func TestCreate(t *testing.T) {
    head := Create(123)
    head.PrintList()  // Manual verification
}

// Right:
func TestCreate(t *testing.T) {
    head := Create(123)
    if head.data != 123 {
        t.Fatalf("Expected 123, got %v", head.data)
    }
}
```

### Minor: Delete Commented Code

Remove all commented placeholder functions.

---

## Rating Justification

**4/10**

**Good**:

- One test has correct name (+1)
- Tests multiple data types (+1)
- Tests exist (+1)
- Covers main functions (+1)

**Bad**:

- 2 of 3 tests don't run (-2)
- Zero assertions (-2)
- PrintList instead of verification (-1)
- Commented placeholders (-1)

**These aren't really tests.** They're manual verification scripts.

---

## Comparison to Week 4

**Week 4 queue tests**: 9.5/10

```go
func TestLinearQueue(t *testing.T) {
    q := LinearQueue{}
    q.Enqueue(106)
    if q.LengthOfQueue() != 1 {  // ✓ Assertion
        t.Fatal("Expected length 1, got", q.LengthOfQueue())
    }
}
```

**Week 5 linked list tests**: 4/10

```go
func testCreate(t *testing.T) {  // ✗ Wrong name
    head := Create(123)
    head.PrintList()  // ✗ No assertion
}
```

**You went backwards.** Week 4 showed you know how to write proper tests. Why didn't you apply that knowledge here?

---

## What You're Learning

1. Table-driven testing (not yet)
2. Test assertions (not yet)
3. Test naming (inconsistent)
4. Test structure (basic)

**You wrote tests, but they don't test anything.** Add assertions and fix naming - both are issues you've already solved in other files.
