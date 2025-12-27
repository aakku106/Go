# Code Review: datastructures/list/list.go

## Overall Assessment

**File Purpose**: List data structure implementation  
**Rating**: N/A (Empty File)  
**Status**: Not implemented yet

---

## File Contents

```go
package list
```

**The file only contains a package declaration - no code yet!**

---

## What a List Should Implement

A **linked list** (or dynamic array list) is a fundamental data structure. Here's what you should implement:

---

## 📚 List Types

### 1. Singly Linked List

```
[Data|Next] → [Data|Next] → [Data|Next] → nil
```

**Operations**: Insert, Delete, Search, Traverse

### 2. Doubly Linked List

```
nil ← [Prev|Data|Next] ↔ [Prev|Data|Next] ↔ [Prev|Data|Next] → nil
```

**Operations**: Same as singly, but can traverse backwards

### 3. ArrayList (Dynamic Array)

```
[0][1][2][3][4][ ][ ][ ]
```

**Operations**: Random access, dynamic resizing

---

## 🔧 Complete Implementations

```go
package list

import (
    "errors"
    "fmt"
)

// IMPLEMENTATION 1: Singly Linked List

type Node struct {
    value any
    next  *Node
}

type LinkedList struct {
    head *Node
    tail *Node
    size int
}

func NewLinkedList() *LinkedList {
    return &LinkedList{}
}

// Insert at end
func (l *LinkedList) Append(value any) {
    newNode := &Node{value: value}

    if l.head == nil {
        l.head = newNode
        l.tail = newNode
    } else {
        l.tail.next = newNode
        l.tail = newNode
    }
    l.size++
}

// Insert at beginning
func (l *LinkedList) Prepend(value any) {
    newNode := &Node{value: value, next: l.head}
    l.head = newNode

    if l.tail == nil {
        l.tail = newNode
    }
    l.size++
}

// Insert at index
func (l *LinkedList) InsertAt(index int, value any) error {
    if index < 0 || index > l.size {
        return errors.New("index out of bounds")
    }

    if index == 0 {
        l.Prepend(value)
        return nil
    }

    if index == l.size {
        l.Append(value)
        return nil
    }

    // Find node before insertion point
    prev := l.nodeAt(index - 1)
    newNode := &Node{value: value, next: prev.next}
    prev.next = newNode
    l.size++
    return nil
}

// Delete at index
func (l *LinkedList) DeleteAt(index int) (any, error) {
    if index < 0 || index >= l.size {
        return nil, errors.New("index out of bounds")
    }

    var value any

    if index == 0 {
        value = l.head.value
        l.head = l.head.next
        if l.head == nil {
            l.tail = nil
        }
    } else {
        prev := l.nodeAt(index - 1)
        value = prev.next.value
        prev.next = prev.next.next

        if prev.next == nil {
            l.tail = prev
        }
    }

    l.size--
    return value, nil
}

// Get value at index
func (l *LinkedList) Get(index int) (any, error) {
    if index < 0 || index >= l.size {
        return nil, errors.New("index out of bounds")
    }

    node := l.nodeAt(index)
    return node.value, nil
}

// Helper: get node at index
func (l *LinkedList) nodeAt(index int) *Node {
    current := l.head
    for i := 0; i < index; i++ {
        current = current.next
    }
    return current
}

// Search for value
func (l *LinkedList) Contains(value any) bool {
    current := l.head
    for current != nil {
        if current.value == value {
            return true
        }
        current = current.next
    }
    return false
}

// Size
func (l *LinkedList) Size() int {
    return l.size
}

// IsEmpty
func (l *LinkedList) IsEmpty() bool {
    return l.size == 0
}

// Clear all elements
func (l *LinkedList) Clear() {
    l.head = nil
    l.tail = nil
    l.size = 0
}

// ToSlice converts to slice
func (l *LinkedList) ToSlice() []any {
    result := make([]any, 0, l.size)
    current := l.head
    for current != nil {
        result = append(result, current.value)
        current = current.next
    }
    return result
}

// String representation
func (l *LinkedList) String() string {
    return fmt.Sprintf("%v", l.ToSlice())
}

// IMPLEMENTATION 2: Doubly Linked List

type DNode struct {
    value any
    prev  *DNode
    next  *DNode
}

type DoublyLinkedList struct {
    head *DNode
    tail *DNode
    size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
    return &DoublyLinkedList{}
}

func (l *DoublyLinkedList) Append(value any) {
    newNode := &DNode{value: value}

    if l.head == nil {
        l.head = newNode
        l.tail = newNode
    } else {
        newNode.prev = l.tail
        l.tail.next = newNode
        l.tail = newNode
    }
    l.size++
}

func (l *DoublyLinkedList) Prepend(value any) {
    newNode := &DNode{value: value, next: l.head}

    if l.head != nil {
        l.head.prev = newNode
    } else {
        l.tail = newNode
    }

    l.head = newNode
    l.size++
}

func (l *DoublyLinkedList) DeleteAt(index int) (any, error) {
    if index < 0 || index >= l.size {
        return nil, errors.New("index out of bounds")
    }

    node := l.nodeAt(index)
    value := node.value

    if node.prev != nil {
        node.prev.next = node.next
    } else {
        l.head = node.next
    }

    if node.next != nil {
        node.next.prev = node.prev
    } else {
        l.tail = node.prev
    }

    l.size--
    return value, nil
}

func (l *DoublyLinkedList) nodeAt(index int) *DNode {
    var current *DNode

    // Optimize: start from head or tail
    if index < l.size/2 {
        current = l.head
        for i := 0; i < index; i++ {
            current = current.next
        }
    } else {
        current = l.tail
        for i := l.size - 1; i > index; i-- {
            current = current.prev
        }
    }

    return current
}

// IMPLEMENTATION 3: ArrayList (Dynamic Array)

type ArrayList struct {
    items []any
}

func NewArrayList() *ArrayList {
    return &ArrayList{
        items: make([]any, 0),
    }
}

func (l *ArrayList) Append(value any) {
    l.items = append(l.items, value)
}

func (l *ArrayList) InsertAt(index int, value any) error {
    if index < 0 || index > len(l.items) {
        return errors.New("index out of bounds")
    }

    // Expand slice
    l.items = append(l.items, nil)

    // Shift elements right
    copy(l.items[index+1:], l.items[index:])

    // Insert value
    l.items[index] = value
    return nil
}

func (l *ArrayList) DeleteAt(index int) (any, error) {
    if index < 0 || index >= len(l.items) {
        return nil, errors.New("index out of bounds")
    }

    value := l.items[index]

    // Shift elements left
    copy(l.items[index:], l.items[index+1:])

    // Truncate
    l.items = l.items[:len(l.items)-1]

    return value, nil
}

func (l *ArrayList) Get(index int) (any, error) {
    if index < 0 || index >= len(l.items) {
        return nil, errors.New("index out of bounds")
    }
    return l.items[index], nil
}

func (l *ArrayList) Set(index int, value any) error {
    if index < 0 || index >= len(l.items) {
        return errors.New("index out of bounds")
    }
    l.items[index] = value
    return nil
}

func (l *ArrayList) Size() int {
    return len(l.items)
}

func (l *ArrayList) IsEmpty() bool {
    return len(l.items) == 0
}

// IMPLEMENTATION 4: Generic List (Go 1.18+)

type List[T any] struct {
    head *ListNode[T]
    tail *ListNode[T]
    size int
}

type ListNode[T any] struct {
    value T
    next  *ListNode[T]
}

func NewList[T any]() *List[T] {
    return &List[T]{}
}

func (l *List[T]) Append(value T) {
    newNode := &ListNode[T]{value: value}

    if l.head == nil {
        l.head = newNode
        l.tail = newNode
    } else {
        l.tail.next = newNode
        l.tail = newNode
    }
    l.size++
}

func (l *List[T]) Prepend(value T) {
    newNode := &ListNode[T]{value: value, next: l.head}
    l.head = newNode

    if l.tail == nil {
        l.tail = newNode
    }
    l.size++
}

func (l *List[T]) Get(index int) (T, error) {
    var zero T
    if index < 0 || index >= l.size {
        return zero, errors.New("index out of bounds")
    }

    current := l.head
    for i := 0; i < index; i++ {
        current = current.next
    }
    return current.value, nil
}

func (l *List[T]) Size() int {
    return l.size
}

func (l *List[T]) IsEmpty() bool {
    return l.size == 0
}

func (l *List[T]) ToSlice() []T {
    result := make([]T, 0, l.size)
    current := l.head
    for current != nil {
        result = append(result, current.value)
        current = current.next
    }
    return result
}
```

---

## Usage Examples

```go
package main

import (
    "fmt"
    "yourproject/list"
)

func main() {
    demonstrateLinkedList()
    demonstrateDoublyLinkedList()
    demonstrateArrayList()
    demonstrateGenericList()
}

func demonstrateLinkedList() {
    fmt.Println("=== Singly Linked List ===")

    l := list.NewLinkedList()

    l.Append(1)
    l.Append(2)
    l.Append(3)
    fmt.Println("After appends:", l)  // [1, 2, 3]

    l.Prepend(0)
    fmt.Println("After prepend:", l)  // [0, 1, 2, 3]

    l.InsertAt(2, 99)
    fmt.Println("After insert at 2:", l)  // [0, 1, 99, 2, 3]

    val, _ := l.DeleteAt(2)
    fmt.Printf("Deleted: %v, List: %v\n", val, l)  // 99, [0, 1, 2, 3]

    fmt.Printf("Size: %d\n", l.Size())
    fmt.Printf("Contains 2? %v\n\n", l.Contains(2))
}

func demonstrateDoublyLinkedList() {
    fmt.Println("=== Doubly Linked List ===")

    l := list.NewDoublyLinkedList()

    l.Append(10)
    l.Append(20)
    l.Append(30)

    l.Prepend(5)
    fmt.Printf("List: %v\n", l.ToSlice())  // [5, 10, 20, 30]

    // Can traverse backwards efficiently
    l.DeleteAt(2)
    fmt.Printf("After delete: %v\n\n", l.ToSlice())
}

func demonstrateArrayList() {
    fmt.Println("=== ArrayList ===")

    l := list.NewArrayList()

    l.Append(100)
    l.Append(200)
    l.Append(300)

    l.InsertAt(1, 150)
    fmt.Printf("After insert: %v\n", l.items)  // [100, 150, 200, 300]

    val, _ := l.Get(2)
    fmt.Printf("Value at index 2: %v\n\n", val)
}

func demonstrateGenericList() {
    fmt.Println("=== Generic List ===")

    // Type-safe integer list
    intList := list.NewList[int]()
    intList.Append(1)
    intList.Append(2)
    intList.Append(3)
    fmt.Printf("Int list: %v\n", intList.ToSlice())

    // Type-safe string list
    strList := list.NewList[string]()
    strList.Append("hello")
    strList.Append("world")
    fmt.Printf("String list: %v\n", strList.ToSlice())
}
```

---

## Comparison Table

| Operation | Linked List | Doubly Linked     | ArrayList      |
| --------- | ----------- | ----------------- | -------------- |
| Append    | O(1)        | O(1)              | O(1) amortized |
| Prepend   | O(1)        | O(1)              | O(n)           |
| Insert    | O(n)        | O(n)              | O(n)           |
| Delete    | O(n)        | O(n)              | O(n)           |
| Get       | O(n)        | O(n/2)            | O(1)           |
| Search    | O(n)        | O(n)              | O(n)           |
| Space     | O(n)        | O(n) [2 pointers] | O(n)           |

**Choose Based On**:

- **Random Access**: ArrayList
- **Frequent Insert/Delete at ends**: LinkedList
- **Bidirectional Traversal**: DoublyLinkedList

---

## Test Suite

```go
package list

import "testing"

func TestLinkedListAppend(t *testing.T) {
    l := NewLinkedList()

    l.Append(1)
    l.Append(2)
    l.Append(3)

    if l.Size() != 3 {
        t.Errorf("Expected size 3, got %d", l.Size())
    }

    expected := []any{1, 2, 3}
    for i, v := range expected {
        val, _ := l.Get(i)
        if val != v {
            t.Errorf("Index %d: expected %v, got %v", i, v, val)
        }
    }
}

func TestLinkedListPrepend(t *testing.T) {
    l := NewLinkedList()

    l.Prepend(3)
    l.Prepend(2)
    l.Prepend(1)

    expected := []any{1, 2, 3}
    slice := l.ToSlice()

    for i, v := range expected {
        if slice[i] != v {
            t.Errorf("Index %d: expected %v, got %v", i, v, slice[i])
        }
    }
}

func TestLinkedListInsertAt(t *testing.T) {
    l := NewLinkedList()
    l.Append(1)
    l.Append(3)

    l.InsertAt(1, 2)

    expected := []any{1, 2, 3}
    slice := l.ToSlice()

    for i, v := range expected {
        if slice[i] != v {
            t.Errorf("Index %d: expected %v, got %v", i, v, slice[i])
        }
    }
}

// Add more tests...
```

---

## Key Takeaways

1. **Linked List**: Good for insert/delete at ends
2. **Doubly Linked**: Can traverse both directions
3. **ArrayList**: Best for random access
4. **Generics**: Use for type safety (Go 1.18+)
5. **Choose wisely**: Based on operation patterns

---

## Next Steps for list.go

1. ✅ Choose list type (Singly Linked recommended to start)
2. ✅ Implement Node struct
3. ✅ Implement List struct with head/tail
4. ✅ Add Append(), Prepend() methods
5. ✅ Add Get(), InsertAt(), DeleteAt()
6. ✅ Add Size(), IsEmpty(), Contains()
7. ✅ Write comprehensive tests
8. ✅ Consider migrating to generics

**Start with singly linked list - it's the foundation for understanding all list structures!** 🔗
