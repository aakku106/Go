# Week 5 Summary: HTTP Servers + Singly Linked List

## Overall Assessment

**Week 5 Rating: 6.8/10**

**Progression**: Week 1 (7.0) → Week 2 (8.0) → Week 3 (7.7) → Week 4 (9.0) → **Week 5 (6.8)**

**Significant regression from Week 4.** Error handling discipline from Week 3-4 was not maintained in HTTP files (only 1 of 7 has error handling). Linked list tests use PrintList() instead of assertions despite Week 4 demonstrating proper test patterns.

### Split Assessment

- **HTTP Development (7 files)**: 7.1/10 - Learning new concepts but error handling regression
- **Linked List (3 files)**: 5.7/10 - First implementation has bugs, tests lack assertions
- **Overall**: 6.8/10 - New topics explored but prior patterns not applied

---

## Week 5 Focus

### HTTP Server Development (New - 7 Files)

Explored net/http package systematically:

**Basic servers**:

- try1: First HTTP server on port 5555 (7/10)
- try2_mux: ServeMux routing (7.5/10)

**Server struct**:

- Eg2: Read stdlib source, discovered abstraction (8/10)
- eg2.2: Default port 80 behavior (6.5/10)

**JSON APIs**:

- try3_POST: Production-quality JSON API (8.5/10) ⭐ BEST FILE
- stripedDown: Intentionally broken anti-pattern (3/10) ⚠️ WORST FILE

**Templates**:

- try4: HTML templating with race condition (7/10)

### Singly Linked List (New - 3 Files)

First data structure built from scratch:

**linkList.go (6/10)**:

- Interface design attempt
- Type definitions (SingelyLinkList, DoublyLinkList)
- **Issue**: Interface signatures don't match implementation

**SingelyLinkList.go (7/10)**:

- InsertAtBeginning: Perfect O(1) implementation
- InsertAtLast: Correct algorithm, wrong return value (returns last node instead of head)
- PrintList: Correct traversal

**SinglyLinkedList_test.go (4/10)**:

- Only 1 of 3 tests runs (lowercase 't' on 2 functions)
- Zero assertions - tests only print output
- **Week 4 regression** - you demonstrated proper assertions in queue tests

---

## Topics Covered

### 1. HTTP Server Basics (try1/main.go - 7/10)

**Learned**:

- `net/http` package fundamentals
- `http.HandleFunc` for route registration
- `http.ListenAndServe` for server startup
- `http.ResponseWriter` interface
- Handler function signature: `func(w http.ResponseWriter, r *http.Request)`

**Issues**:

- 15+ spelling errors ("Initilizating", "simpally", "lest")
- No error handling on `ListenAndServe()`
- Overly verbose comments

**Key quote**:

> "focus on How ?, and after buildign something and getting hit with erros then explore Why ?"

Pragmatic learning approach, but led to skipping error handling.

---

### 2. ServeMux Routing (try2_mux/main.go - 7.5/10)

**Learned**:

- `http.NewServeMux()` for multiplexing
- Multiple route registration (/, /cat, /dog)
- Pattern matching and handler separation

**Code**:

```go
mux := http.NewServeMux()
mux.HandleFunc("/", handleRoot)
mux.HandleFunc("/cat", handleCat)
mux.HandleFunc("/dog", handleDog)
http.ListenAndServe(":5555", mux)
```

**Improvement**: Spelling errors reduced from 15+ to 2.

**Still missing**: Error handling.

---

### 3. Server Struct (Eg2/main.go - 8/10) ⭐ BEST EXPLORATION

**Learned**:

- Explicit `http.Server` struct configuration
- `Server.Addr` and `Server.Handler` fields
- Discovered abstraction by reading stdlib source

**Critical Discovery**:

```go
// You found this in /net/http/server.go:
func ListenAndServe(addr string, handler Handler) error {
    server := &Server{Addr: addr, Handler: handler}
    return server.ListenAndServe()
}
```

**This shows active learning.** You're reading standard library source code, not just tutorials. **Same quality as Week 4's forSelect.go.**

**Bug Introduced**:

```go
server := http.Server{
    Addr:    ":80",  // Actually listens on :80
    Handler: mux,
}
// Comment says ":8080" but code uses ":80"
```

---

### 4. Default Port 80 (eg2.2/main.go - 6.5/10)

**Learned**:

- Empty `Server.Addr` defaults to `:http` (port 80)
- Minimal server configuration

**Missing**:

- Port 80 requires sudo on macOS/Linux (not documented)
- Confused PUT vs POST in comments
- No error handling means permission errors fail silently

**This file feels rushed.** Quick experiment without proper documentation.

---

### 5. JSON API (try3_POST/main.go - 8.5/10) ⭐ BEST HTTP FILE

**Learned**:

- POST request handling with JSON body
- `json.NewDecoder(r.Body).Decode()` for JSON parsing
- `json.NewEncoder(w).Encode()` for JSON responses
- `sync.RWMutex` for concurrent map access
- `http.Error()` for error responses with status codes
- Go 1.22+ method routing: `"POST /cat"`
- Path parameters: `r.PathValue("id")`

**Code Quality**:

```go
err := json.NewDecoder(r.Body).Decode(&cat)
if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
}

// Validation
if cat.CatName == "" {
    http.Error(w, "catName is required", http.StatusBadRequest)
    return
}

// Thread-safe cache access
cacheMutex.Lock()
CatCache[id] = cat
cacheMutex.Unlock()
```

**This is production-quality code.**

**Issues**:

- Still has spelling errors ("handeler", "decalering")
- But **every error is handled** with proper status codes

**RWMutex Usage**:

```go
// Write (exclusive lock):
cacheMutex.Lock()
CatCache[id] = cat
cacheMutex.Unlock()

// Read (shared lock):
cacheMutex.RLock()
cat, ok := CatCache[id]
cacheMutex.RUnlock()
```

Correct use of read/write locks. Multiple concurrent readers, exclusive writers.

---

### 6. Anti-Pattern Demo (stripedDown/main.go - 3/10) ⚠️ WORST FILE

**What you did**:

- Removed all comments
- Used cryptic variable names (`w`, `r`, `m`, `s`)
- No error handling
- No structure

**Code**:

```go
m := http.NewServeMux()
m.HandleFunc("POST /cat", func(w http.ResponseWriter, r *http.Request) {
    var c Cat
    json.NewDecoder(r.Body).Decode(&c)
    // ... no error handling
})
```

**Why this is 3/10 instead of 0/10**:

You labeled it "stripedDown" which suggests intentional demonstration of bad practices.

**What's missing**:

```go
// THIS FILE DEMONSTRATES ANTI-PATTERNS:
// 1. No error handling
// 2. Cryptic variable names
// 3. No code organization
// DO NOT USE THIS PATTERN IN PRODUCTION
```

Without comments explaining **why** it's bad, it just looks like bad code.

---

### 7. HTML Templating (try4/main.go - 7/10)

**Learned**:

- `text/template` package
- `template.ParseFiles()` for loading templates
- `template.Must()` for panic-on-error
- `template.Execute(w, data)` for rendering

**Code**:

```go
tem = template.Must(template.ParseFiles("./htmlFiles/cat.html"))
tem.Execute(w, nil)
```

**Issues**:

1. **Race Condition**:

```go
var tem *template.Template  // Global variable

func handler(w http.ResponseWriter, r *http.Request) {
    tem = template.Must(template.ParseFiles("./htmlFiles/cat.html"))  // Concurrent writes!
}
```

Multiple requests will write to `tem` simultaneously.

2. **Inefficiency**:

Parsing template on **every request** instead of once at startup.

**Correct Pattern**:

```go
var tem = template.Must(template.ParseFiles("./htmlFiles/cat.html"))  // Parse once

func handler(w http.ResponseWriter, r *http.Request) {
    tem.Execute(w, nil)  // Only execute, don't parse
}
```

---

### 8. Singly Linked List - Interface (linkList.go - 6/10)

**Learned**:

- Interface design in Go
- Type definitions for data structures
- `any` type for generic data storage

**Code**:

```go
type LinkList interface {
    Create(data any)
    InsertAtBeginning(data any)
    InsertAtLast(data any)
    InsertAfter(data any, index uint)
    InsertBefore(data any, index uint)
    InsertAt(data any, index uint)
}

type SingelyLinkList struct {
    data any
    next *SingelyLinkList
}
```

**Critical Flaw**:

```go
// Interface says:
type LinkList interface {
    InsertAtBeginning(data any)  // No return value
}

// Implementation has:
func (l *SingelyLinkList) InsertAtBeginning(data any) *SingelyLinkList {
    return node  // Returns new head
}
```

**Signatures don't match.** `SingelyLinkList` does NOT implement `LinkList` interface because return types differ.

**Test**:

```go
var list LinkList = &SingelyLinkList{}  // Compile error!
```

**Design Issues**:

1. `Create()` in interface makes no sense - constructors aren't instance methods
2. `DoublyLinkList` only has `next`, missing `prev` pointer
3. Spelling: "SingelyLinkList" vs "SinglyLinkedList_test.go"

---

### 9. Singly Linked List - Implementation (SingelyLinkList.go - 7/10)

**Learned**:

- Pointer manipulation
- Linked list traversal
- Head pointer management
- O(1) vs O(n) operations

**InsertAtBeginning - Perfect Implementation**:

```go
func (l *SingelyLinkList) InsertAtBeginning(data any) *SingelyLinkList {
    node := Create(data)
    node.next = l
    return node
}
```

**Textbook correct.** O(1) time complexity, returns new head. This shows you understand linked list mechanics.

**InsertAtLast - Correct Algorithm, Wrong Return**:

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

**Bug**: After loop, `head` points to last node. You return last node instead of original head.

**Impact**:

```go
list := Create(1)
list = list.InsertAtLast(2)  // Returns node 1 ✓
list = list.InsertAtLast(3)  // Returns node 2 ✗ (loses node 1!)
list.PrintList()  // Only prints 2→3
```

**Fix**:

```go
func (l *SingelyLinkList) InsertAtLast(data any) *SingelyLinkList {
    node := Create(data)
    current := l  // Use different variable name
    for current.next != nil {
        current = current.next
    }
    current.next = node
    return l  // Return original head
}
```

**PrintList - Correct Traversal**:

```go
func (l SingelyLinkList) PrintList() {
    fmt.Println("LinkList:")
    for {
        fmt.Println(" ↓↪ Data =", l.data)
        if l.next == nil {
            return
        }
        l = *l.next  // Dereference because value receiver
    }
}
```

Works correctly, but value receiver is inefficient (copies entire node).

**Issues**:

1. Named break `Loop:` unnecessary for single loop
2. Empty placeholder functions (`InsertAfter`, `InsertBefore`, `InsertAt`)
3. Spelling: "Inseart", "Itterated"
4. No nil checks - panics if `l` is nil

---

### 10. Linked List Tests (SinglyLinkedList_test.go - 4/10) ⚠️ WEEK 4 REGRESSION

**What you wrote**:

```go
func testCreate(t *testing.T) {           // ❌ lowercase 't'
    head := Create(123)
    head.PrintList()  // ❌ No assertion
}

func testInsertAtBeginning(t *testing.T) {  // ❌ lowercase 't'
    list := Create(106)
    list.PrintList()  // ❌ No assertion
}

func TestInsertAtLast(t *testing.T) {      // ✓ uppercase 'T'
    list := Create("Car")
    list.PrintList()  // ❌ No assertion
}
```

**Problems**:

1. **Only 1 of 3 tests runs** - `testCreate` and `testInsertAtBeginning` have lowercase 't'
2. **Zero assertions** - tests print output, don't verify correctness
3. **Manual verification** - you have to read output to check if it works

**Comparison to Week 4**:

**Week 4 queue tests (9.5/10)**:

```go
func TestLinearQueue(t *testing.T) {
    q := LinearQueue{}
    q.Enqueue(106)
    if q.LengthOfQueue() != 1 {  // ✓ Assertion
        t.Fatal("Expected length 1, got", q.LengthOfQueue())
    }
}
```

**Week 5 linked list tests (4/10)**:

```go
func testCreate(t *testing.T) {  // ✗ Wrong name
    head := Create(123)
    head.PrintList()  // ✗ No assertion
}
```

**You went backwards.** Week 4 showed you know how to write proper tests with assertions. Week 5 has manual verification scripts.

**What proper assertions look like**:

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

func TestInsertAtBeginning(t *testing.T) {
    list := Create(106)
    list = list.InsertAtBeginning("aww")

    // Verify new head
    if list.data != "aww" {
        t.Fatalf("Expected 'aww' at head, got %v", list.data)
    }

    // Verify old head is second
    if list.next == nil || list.next.data != 106 {
        t.Fatal("Second node should be 106")
    }
}
```

---

## What You're Learning

### New Concepts (Week 5)

1. **HTTP Servers**:

   - Handler functions
   - ServeMux routing
   - Server struct configuration
   - Method-based routing (Go 1.22+)

2. **JSON APIs**:

   - JSON encoding/decoding
   - Request body parsing
   - HTTP status codes
   - Input validation

3. **Concurrency**:

   - RWMutex for shared map access
   - Read vs write locks
   - Race conditions in template parsing

4. **Templates**:

   - text/template package
   - ParseFiles and Execute
   - template.Must for error handling

5. **Linked Lists**:
   - Node structure with pointers
   - Head pointer management
   - Traversal algorithms
   - O(1) vs O(n) operations
   - InsertAtBeginning vs InsertAtLast

### Patterns NOT Applied from Week 4

1. **Error Handling**: Learned Week 3-4, only in 1 of 7 HTTP files
2. **Test Assertions**: Demonstrated in Week 4 queue tests, not in Week 5 linked list tests
3. **Test Naming**: Fixed in queue tests Week 4, broken again in linked list tests Week 5

---

## Critical Issues

### 1. Error Handling Regression (HTTP)

**Week 3-4**: Every error checked, proper error wrapping  
**Week 5**: 6 of 7 HTTP files ignore errors

**Files without error handling**:

```go
// try1/main.go:
http.ListenAndServe(":5555", nil)  // Ignores error

// try2_mux/main.go:
http.ListenAndServe(":5555", mux)  // Ignores error

// Eg2/main.go:
server.ListenAndServe()  // Ignores error
```

**Correct pattern (you already know this)**:

```go
if err := server.ListenAndServe(); err != nil {
    log.Fatal(err)
}
```

### 2. Test Assertions Missing (Linked List)

**Week 4 pattern**:

```go
if q.LengthOfQueue() != 1 {
    t.Fatal("Expected length 1")
}
```

**Week 5 reality**:

```go
head.PrintList()  // Hope it looks right?
```

**This is not testing.** This is manual verification.

### 3. InsertAtLast Return Bug

**Current code**:

```go
return head  // Returns LAST node after loop
```

**Chained calls lose the list**:

```go
list = list.InsertAtLast(2).InsertAtLast(3)  // Loses head!
```

**Fix**:

```go
return l  // Return original head before loop
```

### 4. Spelling Errors Everywhere

**Despite Week 4 feedback**, spelling errors in every file:

- HTTP: "Initilizating", "handeler", "decalering"
- Linked List: "Inseart", "Itterated"
- Tests: No spell-check

**Solution**: Enable Code Spell Checker extension in VS Code.

---

## Comparative Analysis: Week 4 vs Week 5

### What Improved

1. **stdlib Reading** - Reading net/http/server.go source shows maturity
2. **JSON Handling** - try3_POST is production-quality
3. **First Data Structure** - Built linked list from scratch (incomplete but foundational)
4. **RWMutex** - Correct concurrent map access

### What Regressed

1. **Error Handling** - 9/10 (Week 4) → 7.1/10 (Week 5 HTTP)
2. **Test Quality** - 9.5/10 queue tests (Week 4) → 4/10 linked list tests (Week 5)
3. **Code Discipline** - Week 4 had consistent patterns, Week 5 is inconsistent
4. **Spell-Check** - Still not enabled despite feedback

### Root Cause Analysis

**Week 4**: Focused on quality and patterns. Every file had error handling, tests had assertions.

**Week 5**: Focused on speed and exploration. Tried many new topics but didn't maintain discipline.

**Example**:

- try3_POST (8.5/10): You **can** write quality code - error handling, concurrency, validation
- try1 (7/10): No error handling despite knowing how
- SinglyLinkedList_test (4/10): No assertions despite Week 4 demonstrating proper testing

**Conclusion**: Quality is inconsistent because you're prioritizing speed over applying learned patterns.

---

## What You Still Need

### Immediate (Week 6)

1. **Restore Error Handling** - Every error, every file
2. **Add Test Assertions** - Linked list tests need proper verification
3. **Fix InsertAtLast Bug** - Return original head, not last node
4. **Enable Spell-Check** - Install Code Spell Checker extension

### Soon (Week 6-7)

1. **Complete Linked List** - Implement InsertAt, InsertAfter, InsertBefore, Delete, Search
2. **HTTP Middleware** - Logging, error recovery, CORS
3. **Context** - Request timeouts, cancellation
4. **HTTP Testing** - Use httptest package
5. **Database** - Replace in-memory maps with SQLite

### Later (Week 7-8)

1. **DoublyLinkedList** - Add prev pointer, bidirectional traversal
2. **WebSockets** - Real-time communication
3. **gRPC** - Efficient RPC protocol
4. **Authentication** - JWT or sessions
5. **Advanced Data Structures** - Trees, graphs, hash tables

---

## Recommendations

### Critical (Fix This Week)

**1. Error Handling**:

```go
// Wrong (Week 5):
server.ListenAndServe()

// Right (Week 3-4):
if err := server.ListenAndServe(); err != nil {
    log.Fatal(err)
}
```

Apply to ALL 6 HTTP files missing it.

**2. Test Assertions**:

```go
// Wrong (Week 5):
func testCreate(t *testing.T) {
    head := Create(123)
    head.PrintList()
}

// Right (Week 4 pattern):
func TestCreate(t *testing.T) {
    head := Create(123)
    if head == nil {
        t.Fatal("Create returned nil")
    }
    if head.data != 123 {
        t.Fatalf("Expected 123, got %v", head.data)
    }
}
```

Fix all 3 test functions.

**3. InsertAtLast Return**:

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

### Major (This Week)

**4. HTTP Tests**:

```go
func TestCreateCat(t *testing.T) {
    req := httptest.NewRequest("POST", "/cat",
        strings.NewReader(`{"catName":"test"}`))
    w := httptest.NewRecorder()

    createCat(w, req)

    if w.Code != http.StatusNoContent {
        t.Fatalf("Expected 204, got %d", w.Code)
    }
}
```

**5. Fix Template Race**:

```go
// Parse once at startup:
var tem = template.Must(template.ParseFiles("./htmlFiles/cat.html"))

// Execute in handler:
func handler(w http.ResponseWriter, r *http.Request) {
    tem.Execute(w, data)  // Don't parse here
}
```

**6. Fix LinkList Interface**:

```go
type LinkList interface {
    InsertAtBeginning(data any) LinkList  // Add return type
    InsertAtLast(data any)
    Length() int
    PrintList()
    // Remove Create - it's a constructor
}
```

### Minor

**7. Enable Spell-Check**:

VS Code → Extensions → Install "Code Spell Checker"

**8. Add Nil Checks**:

```go
func (l *SingelyLinkList) InsertAtLast(data any) *SingelyLinkList {
    if l == nil {
        return Create(data)
    }
    // ... rest of code
}
```

---

## Week 4 Recommendations Applied?

**Error Handling**:

- ❌ HTTP: Only 1 of 7 files
- ✓ Linked List: N/A (no I/O operations)

**Spelling**:

- ❌ Still present in all files

**Test Naming**:

- ✓ Queue tests: Fixed in Week 4 (minor cleanup Week 5)
- ❌ Linked list tests: 2 of 3 have lowercase 't'

**Test Assertions**:

- ✓ Queue tests: Already had assertions
- ❌ Linked list tests: Zero assertions

**Complete Placeholders**:

- ❌ SingelyLinkList.go: 3 empty functions

**Overall: 1 of 5 critical recommendations followed (queue test names).**

---

## Next Week (Week 6)

1. **Add error handling to all HTTP files** - Apply Week 3-4 patterns
2. **Fix linked list bugs** - InsertAtLast return, test naming, assertions
3. **Write HTTP tests** - Use httptest package
4. **Implement middleware** - Logging, recovery, CORS
5. **Add database** - SQLite for persistence instead of in-memory maps

---

## Final Thoughts

**Week 5 shows you're learning new topics (HTTP, linked lists) but not applying patterns you already know.**

**Evidence you CAN write quality code**:

- try3_POST: Production-quality JSON API
- InsertAtBeginning: Perfect linked list implementation
- Week 4 queue tests: Proper assertions

**Evidence of inconsistency**:

- 6 HTTP files without error handling
- Linked list tests without assertions
- Spelling errors despite feedback

**The gap is not knowledge - it's discipline.**

You know how to handle errors (Week 3-4).  
You know how to write test assertions (Week 4).  
You know how to read docs and stdlib source (Week 4-5).

**Week 6 goal**: Apply what you know **consistently**, not just in some files.
