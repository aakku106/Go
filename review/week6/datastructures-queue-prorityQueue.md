# Code Review: datastructures/queue/prorityQueue.go

**File**: `datastructures/queue/prorityQueue.go`  
**Category**: Priority Queue Implementation  
**Lines**: 101  
**Rating**: 5/10

---

## Overview

Implements priority queue with 5 priority levels (0 = highest priority). Uses array of slices with separate front/rear pointers per priority. Includes memory reclamation similar to linearQueue. Code works correctly but has filename typo, unreachable code after panic, and complex nested function in Dequeue.

---

## Strengths

1. **Priority Ordering** - Correctly dequeues highest priority (0) first
2. **Memory Reclamation** - Reuses slice space per priority level
3. **Bounds Validation** - Panics if priority >= MAX (prevents array overflow)
4. **Debug Logging** - Tracks state across all priority levels
5. **isEmpty() Logic** - Checks all priority queues, resets when empty
6. **Working Implementation** - Priority ordering verified by tests

---

## Issues

### Critical

**1. Filename Typo**

```
prorityQueue.go  // WRONG: "prority" missing 'i'
```

Should be `priorityQueue.go`. Matches type name typo from queue.go.

**2. Type Name Typo**

```go
type ProrityQueue struct {  // WRONG: "Prority" missing 'i'
```

Should be `PriorityQueue`. This propagates to:

- All method receivers
- Test file name
- All test code
- Documentation

**3. Unreachable Code After Panic**

```go
if priority >= MAX {
    log.Panic("Cant have prority greater than: ", MAX-1, ...)
    return  // UNREACHABLE - panic stops execution
}
```

`log.Panic` calls `panic()` which stops execution. The `return` never executes. Dead code.

**Fix**:

```go
if priority >= MAX {
    panic(fmt.Sprintf("priority %d exceeds max %d", priority, MAX-1))
}
// No return needed
```

Or better, return error instead of panicking:

```go
func (q *ProrityQueue) Enqueue(value any, priority uint8) error {
    if priority >= MAX {
        return fmt.Errorf("priority %d exceeds max %d", priority, MAX-1)
    }
    // ...
    return nil
}
```

**4. Complex Nested Function in Dequeue**

Lines 37-65: Dequeue wraps logic in anonymous function:

```go
value := func() (val any) {
    if Debug {
        fmt.Println("______Entered func______")
    }
    for i := range MAX {
        // 15 lines of logic
    }
    return
}()
```

Why? No clear benefit. Makes code harder to read and debug. Should be:

```go
func (q *ProrityQueue) Dequeue() (any, bool) {
    if q.isEmpty() {
        return nil, false
    }

    // Find first non-empty priority queue
    for i := range MAX {
        if q.rear[i] != q.front[i] {
            val := q.queue[i][q.front[i]]
            q.front[i]++
            if Debug {
                fmt.Printf("Dequeued: %v from priority %d\n", val, i)
            }
            return val, true
        }
    }

    return nil, false  // All empty
}
```

Simpler, clearer, same logic.

### Major

**1. Excessive Debug Output**

Debug prints:

- Function entry/exit
- Each loop iteration
- Before/after state
- Queue state dumps

For priority queue with 100 items across 5 priorities, this prints hundreds of lines.

**2. Typos in Comments**

- "prority" → "priority" (appears multiple times)
- Comments inherit typo from type name

**3. Memory Reclamation Duplicated**

Same logic as linearQueue (lines 11-23, 25-30). Should extract to shared function:

```go
func reclaimOrAppend(slice []any, front, rear *uint, value any) []any {
    // Shared logic for both LinearQueue and PriorityQueue
}
```

DRY principle violated.

**4. No Concurrency Safety**

Same issue as linearQueue - no mutex protection. Race conditions possible.

**5. isEmpty() Side Effects (Again)**

```go
func (q *ProrityQueue) isEmpty() bool {
    for i := range uint(MAX) {
        if q.rear[i] != q.front[i] {
            return false
        }
    }
    // SIDE EFFECT: Resets all front/rear pointers
    for i := range MAX {
        q.front[i] = 0
        q.rear[i] = 0
    }
    return true
}
```

Method named "isEmpty" mutates state. Same issue as linearQueue.

**6. Length() Calculates Every Time**

```go
func (q ProrityQueue) Length() uint {
    var total uint
    for i := range uint(MAX) {
        total += q.rear[i] - q.front[i]  // Recalculates every call
    }
    return total
}
```

Could cache and update on Enqueue/Dequeue for O(1) instead of O(MAX).

### Minor

**1. Value Receiver (Length)**

```go
func (q ProrityQueue) Length() uint {  // Value receiver
```

Copies entire struct (including 5 slices). Should be pointer:

```go
func (q *ProrityQueue) Length() uint {
```

**2. Debug Formatting**

```go
fmt.Println("______Entered func______")
fmt.Println("______EXIT func______")
```

Underscores don't add value. Standard logging:

```go
fmt.Println("Dequeue: entering")
fmt.Println("Dequeue: exiting")
```

**3. Inconsistent Loop Ranges**

```go
for i := range MAX {  // Using const
for i := range uint(MAX) {  // Converting const
```

Pick one style.

---

## Suggested Improvements

1. **Rename file/type** - `prorityQueue` → `priorityQueue` throughout
2. **Remove unreachable return** - After log.Panic()
3. **Simplify Dequeue** - Remove nested anonymous function
4. **Return errors** - Instead of panicking on invalid priority
5. **Extract shared logic** - DRY violation with linearQueue
6. **Remove isEmpty side effects** - Separate reset() method
7. **Cache length** - Update counter instead of recalculating
8. **Add mutex** - Concurrency safety
9. **Fix debug verbosity** - Reduce output
10. **Pointer receivers** - All methods

---

## Priority Queue Logic Analysis

Priority ordering works correctly:

```
Enqueue(106, priority=0)  // Highest
Enqueue(69, priority=0)
Enqueue(2, priority=1)
Enqueue(1, priority=4)   // Lowest

Dequeue order:
1. 106 (priority 0, first in)
2. 69 (priority 0, second in)
3. 2 (priority 1)
4. 1 (priority 4)
```

Correct FIFO within each priority level. ✅

---

## What You Learned

✅ Priority queue data structure  
✅ Multi-level queue management  
✅ Priority-based dequeuing  
✅ Bounds validation  
❌ Clean function design (nested anonymous function)  
❌ Error handling over panics  
❌ Code reuse (DRY principle)

---

## Testing

Comprehensive tests exist in `prorityQueue_test.go`:

- Priority ordering ✅
- Multiple items per priority ✅
- isEmpty behavior ✅
- Length tracking ✅

Test coverage better than linearQueue tests.

---

## Final Verdict

**5/10** - Working priority queue with correct priority ordering and memory reclamation, but dragged down by filename/type typos, unreachable code after panic, and overly complex Dequeue with nested anonymous function.

**Best part**: Priority ordering logic is correct  
**Worst part**: Complex nested function makes simple logic hard to follow

**Comparison**:

- linearQueue: 5.5/10 (confusing logic)
- prorityQueue: 5/10 (complex implementation)

Both work but need simplification.

**Fix priority**: Rename file/type, simplify Dequeue, return errors instead of panicking.
