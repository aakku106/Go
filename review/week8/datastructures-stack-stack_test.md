# Code Review: datastructures/stack/stack_test.go

**Rating: 7.5/10**

## Overview

Stack unit tests with real assertions. Two test functions covering Push and Pop operations with multiple assertions per test.

## What This Code Does

TestPush: Creates stack, pushes 6 values, verifies length is 6.
TestPop: Creates stack, pushes 106, pops and verifies value, pushes 5 more, pops one, verifies value and remaining length.

45 lines with clear error messages.

## Strengths

1. **Real assertions** - Uses t.Fatal() properly, not just running code
2. **Clear error messages** - "failed: the length of stack shall be 6, buts it's: X"
3. **Mixed types** - Tests any type by pushing ints and strings
4. **Multiple scenarios** - Push then pop, multiple operations
5. **Value verification** - Checks both ok boolean and actual value from Pop
6. **Tests edge case** - Pop on empty stack (line 22 calls Pop on newly pushed stack)
7. **State verification** - Checks length after operations

## Issues

### Critical

None - tests work and verify behavior.

### Major

1. **Typo: "buts"** - Line 13, should be "but"
2. **Typo: "aspected"** - Line 23, should be "expected"
3. **Direct stack field access** - Line 13 uses `len(newStack.stack)` instead of `newStack.LengthOfStack()`
4. **Test ignores return value** - Line 34 `value, _` discards ok boolean without checking
5. **No empty stack Pop test** - Should test Pop on empty stack (expected: nil, false)
6. **No test for LengthOfStack panic** - Should verify panic when exceeding 65535 elements

### Minor

1. **Inconsistent spacing** - Blank line after TestPush closing brace (line 16) but not after TestPop
2. **Magic numbers** - 106 used multiple times with no explanation
3. **Hardcoded strings** - "wee", "weeee", "cat", "weeeeeeeeeeee" with no clear pattern
4. **Test naming** - TestPush/TestPop test multiple scenarios, could split into subtests
5. **No table-driven tests** - Multiple push/pop scenarios could use table-driven approach
6. **Direct field access** - Line 13 and 36 access `newStack.stack` directly (unexported field), should use LengthOfStack()
7. **Comment style** - No comments explaining test scenarios

## What You Learned

- Unit test structure in Go
- t.Fatal() for assertions
- Testing with multiple data types (any)
- Verifying state after operations
- (value, ok) pattern testing

Improvement from Week 8 other tests:

- defer_test.go: 0 assertions (3/10)
- files_test.go: 2 assertions (6.5/10)
- stack_test.go: 5 assertions (7.5/10)

## Testing

TestPush tests scenario:

- Push 6 items
- Verify length == 6
  Coverage: Push operation with multiple types

TestPop tests scenarios:

- Push 106, Pop, verify value and ok
- Push 5 items, Pop 1, verify value
- Verify length after operations
  Coverage: Pop operation, return values, state changes

Missing tests:

- Pop on empty stack
- Push then Pop multiple times (drain stack)
- LengthOfStack() panic scenario
- Peek operation (if it existed)
- Type-specific values (pushing nil, pushing Stack as value, etc.)

## Final Verdict

**Best test file in Week 8.** This has 5 real assertions checking actual behavior, clear error messages, and tests with mixed types. Significantly better than defer_test.go (0 assertions) and files_test.go (2 weak assertions).

**Issues:**

- Typos: "buts", "aspected"
- Direct field access breaking encapsulation (line 13, 36)
- Ignoring ok return value (line 34)
- Missing edge cases (empty stack Pop)

**Why 7.5/10 not higher:** Tests verify happy path but miss error cases. Direct access to `newStack.stack` instead of using `LengthOfStack()` method breaks encapsulation. Should test edge cases like:

```go
func TestPopEmpty(t *testing.T) {
    s := Stack{}
    val, ok := s.Pop()
    if ok {
        t.Fatal("Pop on empty stack should return false")
    }
    if val != nil {
        t.Fatal("Pop on empty stack should return nil")
    }
}
```

For Week 8 final week, this shows the strongest testing discipline. Has clear assertions, tests multiple scenarios, provides good error messages. With edge case tests and fixing encapsulation breaks, could be 8.5-9/10.

---

**Previous issues from other weeks:** Week 7 test2: 0 assertions. Week 8 defer_test: 0 assertions. Week 8 stack_test: 5 assertions. Clear improvement trend.
