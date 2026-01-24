# Code Review: datastructures/list/SingallyLinkedListtest2_test.go (NEW)

**File**: `datastructures/list/SingallyLinkedListtest2_test.go`  
**Category**: InsertAt Tests  
**Lines**: 18  
**Rating**: 3/10

---

## Overview

New test file created Week 7 for InsertAt method. Single test function TestInsertAt with minimal coverage. Tests basic insertion and out-of-bounds error. No assertions, only prints output. Filename has double typo: "Singally" (wrong) + "test2" (non-standard numbering).

---

## Strengths

1. **Tests InsertAt** - Only file testing newly implemented InsertAt method
2. **Error Case** - Tests out of bounds index (106)
3. **Multiple Insertions** - Creates list, inserts at end, inserts at middle

---

## Issues

### Critical

**1. Double Filename Typo**

```
SingallyLinkedListtest2_test.go
```

Three problems:

- "Singally" → "Singly" (missing 'i' after 'S')
- "test2" → Should follow Go convention
- Missing underscore before first "test"

Should be:

```
SinglyLinkedList_insertAt_test.go  // Descriptive
```

Or:

```
SinglyLinkedList_test2.go  // If numbering needed
```

Current name has typo different from other files (`SingelyLinkList.go` vs `SingallyLinkedListtest2`).

**2. Zero Assertions**

```go
func TestInsertAt(t *testing.T) {
 list := NewSinglyLinkedList("1st")
 list = list.InsertAtLast("3rd")
 if cat, err := list.InsertAt("2nd", 1); err == nil {
  cat.PrintList()  // Just prints, doesn't verify
 }
 list.PrintList()  // Prints again
 list, err := list.InsertAt(106, 106)
 fmt.Println(err)  // Prints error, doesn't assert
 list.PrintList()  // Prints final state
}
```

Test **prints output** but **doesn't verify** anything. No `t.Error`, `t.Fatal`, `t.Errorf`. This is not a test, it's a manual inspection tool.

**Should verify**:

```go
func TestInsertAt(t *testing.T) {
 list := NewSinglyLinkedList("1st")
 list = list.InsertAtLast("3rd")

 // Test valid insertion
 result, err := list.InsertAt("2nd", 1)
 if err != nil {
  t.Fatalf("Unexpected error: %v", err)
 }

 // Verify order: "1st" -> "2nd" -> "3rd"
 if result.data != "1st" {
  t.Errorf("Expected first node '1st', got %v", result.data)
 }
 if result.next.data != "2nd" {
  t.Errorf("Expected second node '2nd', got %v", result.next.data)
 }
 if result.next.next.data != "3rd" {
  t.Errorf("Expected third node '3rd', got %v", result.next.next.data)
 }

 // Test out of bounds
 _, err = result.InsertAt(106, 106)
 if err == nil {
  t.Error("Expected error for out of bounds index, got nil")
 }
}
```

**3. Variable Named 'cat'**

```go
if cat, err := list.InsertAt("2nd", 1); err == nil {
 cat.PrintList()
}
```

Why is the returned list called `cat`? Meaningless variable name. Should be:

```go
if updated, err := list.InsertAt("2nd", 1); err == nil {
 updated.PrintList()
}
```

Or:

```go
result, err := list.InsertAt("2nd", 1)
if err != nil {
 t.Fatal(err)
}
result.PrintList()
```

**4. Inconsistent Error Handling**

```go
if cat, err := list.InsertAt("2nd", 1); err == nil {  // Checks err == nil
 cat.PrintList()
}
// ...
list, err := list.InsertAt(106, 106)  // Doesn't check err at all
fmt.Println(err)
```

First call checks `err == nil`, second call ignores error and just prints. Should consistently check errors:

```go
result, err := list.InsertAt("2nd", 1)
if err != nil {
 t.Fatalf("InsertAt failed: %v", err)
}

result, err = result.InsertAt(106, 106)
if err == nil {
 t.Error("Expected error for index 106, got nil")
}
```

### Major

**1. Unused Package Import**

```go
import (
 "fmt"
 "testing"
)
```

Uses `fmt.Println` which is wrong for tests. Should use `t.Log` for test output:

```go
t.Logf("Error: %v", err)
```

Or remove prints entirely and use assertions.

**2. No Edge Cases**

Missing tests:

- InsertAt(0) - insert at beginning
- InsertAt on single-element list
- InsertAt at end of list vs InsertAtLast
- InsertAt with negative index (wraps to large uint)
- InsertAt maintaining list integrity

**3. No Test Documentation**

```go
func TestInsertAt(t *testing.T) {
 // No comment explaining what this tests
```

Should document:

```go
// TestInsertAt verifies InsertAt inserts nodes at specified index
// and returns error for out-of-bounds indices
func TestInsertAt(t *testing.T) {
```

**4. Test Name Too Generic**

Single function `TestInsertAt` tests multiple scenarios. Should split:

```go
func TestInsertAt_ValidIndex(t *testing.T) { ... }
func TestInsertAt_OutOfBounds(t *testing.T) { ... }
func TestInsertAt_EmptyList(t *testing.T) { ... }
func TestInsertAt_SingleElement(t *testing.T) { ... }
```

Or use subtests:

```go
func TestInsertAt(t *testing.T) {
 t.Run("valid index", func(t *testing.T) { ... })
 t.Run("out of bounds", func(t *testing.T) { ... })
}
```

### Minor

**1. Variable Shadowing**

```go
list := NewSinglyLinkedList("1st")
list = list.InsertAtLast("3rd")
// ...
list, err := list.InsertAt(106, 106)  // Shadows err from earlier scope
```

First `err` from line 7 is shadowed by second `err` on line 14. Not a bug but confusing.

**2. Magic Numbers**

```go
list.InsertAt("2nd", 1)  // Why 1?
list.InsertAt(106, 106)  // Why 106 for both data and index?
```

Should use named constants or document why:

```go
const (
 middleIndex = 1
 invalidIndex = 106
)
```

**3. Inconsistent Data Types**

Uses strings ("1st", "2nd", "3rd") then switches to int (106). Mixing types is fine for `any` but inconsistent test data makes it harder to read.

---

## What You Learned

1. **Testing New Methods** - Writing tests for InsertAt after implementation
2. **Error Testing** - Checking out of bounds behavior
3. **Test File Organization** - Creating separate test file for specific functionality

---

## Testing Coverage

This file tests:

- ✅ InsertAt with valid index
- ✅ InsertAt with out-of-bounds index
- ❌ InsertAt(0)
- ❌ InsertAt on single element
- ❌ InsertAt at list end
- ❌ Assertions (no actual verification)
- ❌ Edge cases
- ❌ List integrity after insertion

---

## Final Verdict

**3/10** - Minimal test with zero assertions. File name has double typo, variable named 'cat', no verification logic. This is not a real test - it prints output for manual inspection. Test exists but doesn't actually test anything automatically.

**Main problems**: Filename double typo, zero assertions (just prints), variable named 'cat', no test verification.

**What works**: Calls InsertAt with valid and invalid indices, demonstrates usage.

**Recommended fixes**:

1. Rename to SinglyLinkedList_insertAt_test.go or SinglyLinkedList_test.go
2. Add actual assertions with t.Error/t.Fatal
3. Rename 'cat' variable to 'result' or 'updated'
4. Split into multiple test functions or subtests
5. Test InsertAt(0), single element, end of list
6. Remove fmt.Println, use assertions instead
7. Add test documentation comment
8. Fix typo in filename (Singally → Singly)
