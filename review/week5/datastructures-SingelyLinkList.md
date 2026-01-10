# Code Review: datastructures/list/SingelyLinkList.go

**File**: `datastructures/list/SingelyLinkList.go`  
**Topic**: Singly Linked List Implementation  
**Rating**: 7/10  
**Reviewed**: January 11, 2026

---

## Summary

First working linked list implementation with Create, InsertAtBeginning, InsertAtLast, and PrintList. Core algorithms are correct but has bugs and incomplete functions.

---

## What You Did

### 1. Node Creation

```go
func Create(data any) *SingelyLinkList {
    if Debug {
        fmt.Println("DEBUG: Creating a NODE")
    }
    return &SingelyLinkList{
        data: data,
        next: nil,
    }
}
```

Constructor function returning pointer to new node.

### 2. Insert at Beginning

```go
func (l *SingelyLinkList) InsertAtBeginning(data any) *SingelyLinkList {
    node := Create(data)
    node.next = l
    return node
}
```

**Correct algorithm.** O(1) time complexity, returns new head.

### 3. Insert at Last

```go
func (l *SingelyLinkList) InsertAtLast(data any) *SingelyLinkList {
    node := Create(data)
    head := l
Loop:
    for {
        if head.next != nil {
            head = head.next
        } else {
            break Loop
        }
    }
    head.next = node
    node.next = nil
    return head
}
```

Traverses to last node, appends new node. O(n) time complexity.

### 4. PrintList

```go
func (l SingelyLinkList) PrintList() {
    fmt.Println("LinkList:")
    I := 0
    for {
        if Debug {
            I++
            fmt.Println("Itterated: ", I)
        }
        fmt.Println(" ↓↪ Data =", l.data)
        if l.next == nil {
            return
        }
        l = *l.next
    }
}
```

Traverses and prints all nodes. Works correctly.

---

## Critical Issues

### Issue 1: InsertAtLast Returns Wrong Value

```go
func (l *SingelyLinkList) InsertAtLast(data any) *SingelyLinkList {
    node := Create(data)
    head := l
Loop:
    for {
        if head.next != nil {
            head = head.next  // head is now LAST node
        } else {
            break Loop
        }
    }
    head.next = node
    node.next = nil
    return head  // ❌ Returns LAST node, not original head!
}
```

**Bug**: You return `head` after the loop, but `head` is now pointing to the LAST node, not the original head.

**Correct**:

```go
func (l *SingelyLinkList) InsertAtLast(data any) *SingelyLinkList {
    node := Create(data)
    current := l  // Rename for clarity
    for current.next != nil {
        current = current.next
    }
    current.next = node
    return l  // Return original head
}
```

**Impact**: Your test works because you're only calling it once. Try this:

```go
list := Create(1)
list = list.InsertAtLast(2)  // Returns node 1 ✓
list = list.InsertAtLast(3)  // Returns node 2 ✗ (should return 1)
list.PrintList()  // Only prints 2→3, missing 1!
```

### Issue 2: Named Break is Unnecessary

```go
Loop:
for {
    if head.next != nil {
        head = head.next
    } else {
        break Loop  // Named break is pointless here
    }
}
```

**Simplify**:

```go
for head.next != nil {
    head = head.next
}
```

Named breaks are useful for nested loops, not single loops.

### Issue 3: Placeholder Functions

```go
func (l *SingelyLinkList) InsertAfter()  {}
func (l *SingelyLinkList) InsertBefore() {}
func (l *SingelyLinkList) InsertAt()     {}
```

Empty placeholders. **Week 4 review said: "Complete placeholders or remove them."**

You still have empty placeholders in Week 5.

---

## Spelling Errors

1. **"Inseart"** → "Insert" (in debug message)
2. **"Itterated"** → "Iterated"
3. **Variable "I"** → Should be lowercase `i` (Go convention)

---

## What's Good

### 1. InsertAtBeginning is Perfect

```go
func (l *SingelyLinkList) InsertAtBeginning(data any) *SingelyLinkList {
    node := Create(data)
    node.next = l
    return node
}
```

**Textbook correct.** Creates node, links to old head, returns new head. O(1) time.

### 2. PrintList Algorithm

```go
for {
    fmt.Println(" ↓↪ Data =", l.data)
    if l.next == nil {
        return
    }
    l = *l.next
}
```

Correct traversal pattern for singly linked list. The dereference `*l.next` is necessary since you're using value receiver.

### 3. Debug Pattern

```go
if Debug {
    fmt.Println("DEBUG: Creating a NODE")
}
```

Good debugging practice, toggleable via package variable.

---

## Design Issues

### Issue 1: Value Receiver on PrintList

```go
func (l SingelyLinkList) PrintList() {  // Value receiver
    // ...
    l = *l.next  // Have to dereference
}
```

**Inefficient.** This copies the entire node struct every time you call `PrintList()`.

**Better**:

```go
func (l *SingelyLinkList) PrintList() {  // Pointer receiver
    current := l
    for current != nil {
        fmt.Println(" ↓↪ Data =", current.data)
        current = current.next  // No dereference needed
    }
}
```

### Issue 2: No Error Handling

```go
func (l *SingelyLinkList) InsertAtLast(data any) *SingelyLinkList {
    // What if l is nil?
    head := l
    for head.next != nil {  // PANIC if l is nil!
        head = head.next
    }
}
```

**Missing nil check.** If someone calls `var list *SingelyLinkList; list.InsertAtLast(1)`, your code panics.

---

## Algorithmic Correctness

### InsertAtBeginning: ✅ Correct

- Creates new node
- Links to old head
- Returns new head
- O(1) time

### InsertAtLast: ⚠️ Algorithm correct, return value wrong

- Traverses to last node (correct)
- Appends new node (correct)
- Returns wrong node (bug)
- O(n) time

### PrintList: ✅ Correct

- Traverses all nodes
- Prints data
- Stops at nil
- O(n) time

---

## Recommendations

### Critical: Fix InsertAtLast Return

```go
func (l *SingelyLinkList) InsertAtLast(data any) *SingelyLinkList {
    node := Create(data)
    current := l
    for current.next != nil {
        current = current.next
    }
    current.next = node
    return l  // Return original head, not current
}
```

### Major: Add Nil Checks

```go
func (l *SingelyLinkList) InsertAtLast(data any) *SingelyLinkList {
    if l == nil {
        return Create(data)
    }
    // ... rest of code
}
```

### Major: Use Pointer Receiver on PrintList

```go
func (l *SingelyLinkList) PrintList() {  // Pointer, not value
    current := l
    for current != nil {
        fmt.Println(" ↓↪ Data =", current.data)
        current = current.next
    }
}
```

### Minor: Simplify Loop

```go
// Instead of:
Loop:
for {
    if head.next != nil {
        head = head.next
    } else {
        break Loop
    }
}

// Use:
for head.next != nil {
    head = head.next
}
```

---

## Rating Justification

**7/10**

**Good**:

- InsertAtBeginning is perfect (+2)
- Core algorithms correct (+2)
- PrintList works (+2)
- Debug pattern (+1)

**Bad**:

- InsertAtLast returns wrong value (-2)
- Empty placeholder functions (-1)
- Spelling errors (-1)

**This is solid foundational work.** The algorithms show you understand linked list mechanics. Fix the return value bug and you're at 8/10.

---

## What You're Learning

1. Pointer manipulation in Go
2. Linked list traversal
3. Head pointer management
4. O(1) vs O(n) operations

**You're building data structures from scratch, not using built-in types.** This is the right way to learn.
