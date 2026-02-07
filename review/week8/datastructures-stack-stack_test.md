# Code Review: datastructures/stack/stack_test.go

**Rating: 7/10**

## Overview

Comprehensive stack tests with 13+ assertions. Two test functions: TestStack (all operations) and TestAmount (100 million element performance test). Has critical bug in TestAmount's Peek verification loop.

## What This Code Does

TestStack: Full integration test covering Push, Peek, Pop, Len, and empty stack behavior with 13 assertions.
TestAmount: Performance test pushing/popping 100 million elements and measuring elapsed time.

101 lines with detailed error messages.

## Strengths

1. **13+ assertions** - TestStack thoroughly verifies all operations
2. **Tests all methods** - Push, Pop, Peek, Len all covered
3. **Edge case testing** - Empty stack Pop returns (nil, false) correctly
4. **Clear error messages** - Each assertion explains expected vs actual
5. **Mixed type testing** - Uses both string and int (tests any type)
6. **Value and ok testing** - Verifies both return values from Pop/Peek
7. **Performance testing** - TestAmount with 100M elements shows scalability
8. **Sequence testing** - Push → Peek → Pop → verify order (LIFO)
9. **Empty verification** - Tests all methods return correct values on empty stack
10. **Uses constructor** - Calls NewStack() properly with size parameter

## Issues

### Critical

1. **TestAmount Peek loop never executes** - Line 68: `for i := 1; i > iterate; i++` condition is backwards. Should be `i < iterate`. This loop will never run since 1 is not greater than 100,000,000

### Major

1. **Typo: "somem"** - Line 32, should be "some"
2. **Typo: "itterate"** - Line 62 variable name, should be "iterate" (appears correctly elsewhere)
3. **Typo: "insted"** - Line 75, should be "instead"
4. **Uses log.Println in test** - Line 86 should use `t.Logf()` instead of `log.Println()`

### Minor

1. **TestAmount doesn't actually verify Peek** - Due to bug, the verification loop is skipped
2. **No subtests** - Could use t.Run() for better organization
3. **Magic number 106** - No explanation for why this value
4. **Verbose error messages** - Some repeat information ("1st element in 0th index or element in top")
5. **No table-driven tests** - Multiple scenarios could use struct-based test cases
6. **Anonymous functions in TestAmount** - Line 64, 67, 83 use needless func(){...}() wrappers
7. **TestAmount iteration count** - 100M is huge, could make tests slow, no skip for short mode

## What You Learned

- **Comprehensive testing** - Single test function with multiple scenarios
- **Testing (value, ok) patterns** - Verify both return values
- **Edge case coverage** - Empty stack behavior
- **Performance testing** - Large dataset handling (100M elements)
- **time.Since()** - Measuring execution time
- **Testing order** - LIFO verification through sequence

What you practiced:

- t.Fatal() for test assertions
- Mixed type testing with any
- Constructor usage in tests
- State verification after operations

What needs improvement:

- Loop condition logic (line 68 bug)
- Spell checking (somem, itterate, insted)
- Using t.Logf() instead of log.Println()

## Testing

**TestStack coverage:**

- Push operation: 2 pushes, verify lengths
- Peek operation: 2 peeks, verify values
- Pop operation: 3 pops (including empty stack)
- Len operation: Verified after each state change
- Empty stack behavior: All methods tested on empty stack
- Mixed types: string and int
- LIFO order: Verified through Pop sequence

Total assertions: 13

**TestAmount coverage:**

- Push 100 million elements: ✓ Works
- Peek verification during hold: ✗ **Bug - loop never runs**
- Pop all elements: ✓ Works
- Verify empty after operations: ✓ Works (3 checks)
- Time measurement: ✓ Logged

**Missing tests:**

- Concurrent access (thread safety)
- NewStack with various sizes (0, negative, huge)
- Capacity testing
- Multiple goroutines pushing/popping simultaneously

## Final Verdict

**Strong test coverage with excellent assertions but contains critical bug in TestAmount.** TestStack is exceptional - 13 assertions covering all operations, edge cases, and LIFO behavior. Clear error messages, mixed types, proper (value, ok) verification.

**Critical issue:** Line 68 bug means the Peek verification loop in TestAmount never executes:

```go
for i := 1; i > iterate; i++ {  // BUG: 1 > 100000000 is false
```

Should be:

```go
for i := 1; i < iterate; i++ {  // Correct: 1 < 100000000 is true
```

This means TestAmount pushes 100M elements, skips verification, pops all, checks empty. The performance test works but doesn't verify data integrity during the hold phase.

**Why 7/10:**

✓ TestStack is comprehensive (would be 9/10 alone)  
✓ 13+ real assertions checking behavior  
✓ Tests edge cases (empty stack)  
✓ Clear error messages  
✗ Critical logic bug in TestAmount loop  
✗ 3 typos  
✗ Uses log.Println instead of t.Logf

**Comparison to Week 8 main repo tests:**

- defer_test.go: 0 assertions (3/10)
- files_test.go: 2 assertions (6.5/10)
- **stack_test.go: 13 assertions (7/10)** ← Best testing in Week 8

Even with the bug, this shows massive improvement in testing discipline. Fix line 68 and it would rate 8.5/10.

---

**Bug fix needed:**

```go
// Line 68 - change from:
for i := 1; i > iterate; i++ {
// To:
for i := 0; i < iterate; i++ {
    if val, ok := stack.Peek(); !ok {
        t.Fatal("there shall exist values in stack")
    } else {
        if val != iterate-1-i {  // Top element decreases
            t.Fatal("We shall get:", iterate-1-i, "but instead got", val)
        }
    }
}
```
