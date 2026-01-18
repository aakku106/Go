# Code Review: datastructures/queue/linearQueue.go

**File**: `datastructures/queue/linearQueue.go`  
**Category**: Linear Queue Implementation  
**Lines**: 52  
**Rating**: 5.5/10

---

## Overview

Implements FIFO linear queue with memory reclamation optimization. Queue reuses slice space before appending new elements. Code works correctly in single-threaded use but has confusing control flow, no concurrency safety, and logic flaw in Enqueue.

---

## Strengths

1. **Memory Optimization** - Reuses slice space instead of always appending
2. **isEmpty() Side Effect** - Resets front/rear to reclaim memory
3. **Peek() Method** - Non-destructive inspection
4. **Debug Logging** - Tracks internal state (len, cap, front, rear)
5. **Bool Return Values** - Dequeue/Peek return success indicator
6. **Working FIFO** - Correctly maintains queue order

---

## Issues

### Critical

**1. Logic Flaw in Enqueue (Line 6-8)**

```go
func (q *LinearQueue) Enqueue(value any) {
    if len(q.queue) == 0 {
        q.queue = append(q.queue, value)  // Line 6
    }
    if q.isEmpty() || len(q.queue) != int(q.LengthOfQueue()) && q.front == 0 {  // Line 8
        // ...
    }
```

**Flow**:

1. If queue empty, append value (line 6)
2. Then check `isEmpty()` on line 8
3. But we just added an element, so queue is NOT empty
4. Line 8 condition won't match for newly created queue

**This works by accident** because line 8's second condition (`q.front == 0`) catches it. But logic is confusing and shows unclear thinking.

**Fix**:

```go
func (q *LinearQueue) Enqueue(value any) {
    // First enqueue - initialize slice
    if len(q.queue) == 0 {
        q.queue = append(q.queue, value)
        q.rear++
        return
    }

    // Memory reclamation - reuse freed space
    if q.front > 0 && q.rear < uint(len(q.queue)) {
        q.queue[q.rear] = value
        q.rear++
        return
    }

    // Normal enqueue - grow slice
    q.queue = append(q.queue, value)
    q.rear++
}
```

**2. No Capacity Limit**

Queue can grow indefinitely:

```go
q.queue = append(q.queue, value)  // Unbounded growth
```

No maximum size. Memory leak risk if producer faster than consumer.

Should have:

```go
const MaxQueueSize = 1000

func (q *LinearQueue) Enqueue(value any) error {
    if q.LengthOfQueue() >= MaxQueueSize {
        return fmt.Errorf("queue full")
    }
    // ...
}
```

**3. No Concurrency Safety**

No mutex protection. Concurrent access causes race conditions:

```go
// Goroutine 1: Enqueue
q.rear++

// Goroutine 2: Dequeue (simultaneously)
q.front++

// DATA RACE
```

Should have:

```go
type LinearQueue struct {
    sync.RWMutex
    queue []any
    front uint
    rear  uint
}

func (q *LinearQueue) Enqueue(value any) {
    q.Lock()
    defer q.Unlock()
    // ...
}
```

### Major

**1. Confusing Control Flow**

Lines 8-13:

```go
if q.isEmpty() || len(q.queue) != int(q.LengthOfQueue()) && q.front == 0 {
    q.queue[q.rear] = value
    q.rear++
    if Debug { /* ... */ }
    return
}
```

Condition is: `isEmpty() OR (len != length AND front == 0)`

What does this actually check? Hard to understand. Better to split into clear cases:

- First enqueue
- Memory reclamation
- Normal append

**2. isEmpty() Has Side Effect**

```go
func (q *LinearQueue) isEmpty() bool {
    if q.front == q.rear {
        q.front, q.rear = 0, 0  // MUTATES STATE
        return true
    }
    return false
}
```

Method named `isEmpty()` shouldn't mutate state. Violates principle of least surprise.

**Better**: Separate concerns:

```go
func (q *LinearQueue) isEmpty() bool {
    return q.front == q.rear
}

func (q *LinearQueue) reset() {
    q.front, q.rear = 0, 0
}
```

**3. LengthOfQueue Can Underflow (Theoretically)**

```go
func (q LinearQueue) LengthOfQueue() uint {
    return q.rear - q.front  // If rear < front, underflows
}
```

With uint, `0 - 1 = 18446744073709551615` (underflow). Current code prevents this because isEmpty() resets both to 0, but it's fragile.

**4. Verbose Debug Output**

```go
fmt.Println("length: ", len(q.queue), " cap: ", cap(q.queue), "queue:", q, " LengthOfQueue: ", q.LengthOfQueue())
```

Prints on every Enqueue/Dequeue. For 1000 operations, 1000 debug lines. Should be configurable or use proper logging levels.

### Minor

**1. Enqueue Returns Nothing**

```go
func (q *LinearQueue) Enqueue(value any) {  // No return value
```

Caller can't detect failures (though none currently possible). Better to return error for future capacity limits.

**2. Method Name LengthOfQueue**

Verbose. Standard Go would be `Len()` or `Size()`:

```go
func (q LinearQueue) Len() uint {
    return q.rear - q.front
}
```

**3. Value Receiver (LengthOfQueue)**

```go
func (q LinearQueue) LengthOfQueue() uint {  // Value receiver
```

Copies struct unnecessarily. Should be:

```go
func (q *LinearQueue) LengthOfQueue() uint {  // Pointer receiver
```

---

## Suggested Improvements

1. **Clarify Enqueue logic** - Split into clear cases (first, reclaim, append)
2. **Add capacity limit** - Prevent unbounded growth
3. **Add concurrency safety** - Use sync.RWMutex
4. **Remove isEmpty() side effect** - Separate mutation into reset()
5. **Simplify debug** - Reduce verbosity or use levels
6. **Return errors** - Enqueue should return error
7. **Rename Length** - `LengthOfQueue()` → `Len()`
8. **Pointer receivers** - All methods should use pointer receiver

---

## Memory Reclamation Analysis

The memory reclamation strategy is clever:

```
After several Enqueue/Dequeue:
queue: [_ _ _ 99 88 77]
        ^     ^
      front  rear

Instead of appending 66, reuse index 0-2:
queue: [66 _ _ 99 88 77]
            ^     ^
          rear  front (would wrap)
```

But implementation is confusing. Clearer approach:

```go
// Check if we can reuse freed front slots
hasSpace := q.rear < uint(len(q.queue))
hasFreedSlots := q.front > 0

if hasSpace && hasFreedSlots {
    // Compact: shift elements to front
    copy(q.queue, q.queue[q.front:q.rear])
    q.rear = q.rear - q.front
    q.front = 0
}
```

---

## What You Learned

✅ Queue FIFO mechanics  
✅ Memory optimization strategies  
✅ Slice capacity management  
❌ Clear control flow design  
❌ Concurrency safety  
❌ Side effect-free methods

---

## Testing

Tests exist in `linearQueue_test.go`:

- Basic Enqueue/Dequeue ✅
- Length tracking ✅
- Empty queue behavior ✅
- Memory reclamation ❌ (not explicitly tested)
- Concurrency ❌ (not tested)
- Capacity limits ❌ (none exist)

---

## Final Verdict

**5.5/10** - Working FIFO queue with clever memory optimization but confusing implementation. Logic flaw in Enqueue (works by accident), no concurrency safety, unbounded growth, and isEmpty() with side effects reduce score. Code functions correctly for single-threaded use cases.

**Best part**: Memory reclamation concept  
**Worst part**: Confusing control flow and logic that works by accident

**Fix priority**: Clarify Enqueue logic before building on this foundation.
