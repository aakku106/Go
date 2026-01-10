# Week 5 Code Review Index

Week 5 code reviews covering HTTP server development and singly linked list implementation.

**Review Period**: January 4-10, 2026  
**Focus**: net/http package, ServeMux routing, JSON APIs, HTML templating, linked list data structure

---

## Quick Summary

**Overall Rating: 6.8/10**  
**Total Files Reviewed: 10** (7 HTTP + 3 Linked List)  
**Best File: try3_POST/main.go (8.5/10)**  
**Worst File: stripedDown/main.go (3/10)**  
**Topics: HTTP servers, routing, JSON, templates, singly linked lists**  
**Critical Achievement: First data structure built from scratch (linked list)**  
**Critical Issue: Error handling regression + linked list test assertions missing**

---

## Week 5 Highlights

### Notable Achievements

1. **try3_POST/main.go** - Production-quality JSON API with error handling, RWMutex concurrency safety, and validation
2. **First Linked List Implementation** - Built singly linked list from scratch with InsertAtBeginning, InsertAtLast, PrintList
3. **Template Usage** - Learned text/template package for HTML rendering
4. **HTTP Method Routing** - Used Go 1.22+ syntax ("POST /cat", "GET /cat/{id}")
5. **stdlib Reading** - Read net/http/server.go source code (good learning practice)

### Major Issues

1. **Error Handling Regression (HTTP)** - Only 1 of 7 HTTP files has error handling (try3_POST)
2. **Linked List Test Quality** - Tests print output instead of using assertions (Week 4 regression)
3. **Spelling Errors** - Persistent across all files despite Week 4 feedback
4. **InsertAtLast Bug** - Returns wrong node (last instead of head)
5. **Interface Design Flaw** - LinkList interface signatures don't match implementation

---

## File Reviews

### HTTP Server Development

#### [0.0015_HTTP_Starts_Here/try1/main.go](0.0015-try1-main.md)

**Rating: 7/10**  
**Topics:** Basic HTTP server, HandleFunc, ListenAndServe  
**Key Issues:** 15+ spelling errors, no error handling, overly verbose comments  
**Strengths:** Functional server, exploratory documentation, source code references

First HTTP server on port 5555. Works correctly but has egregious spelling errors and no error handling.

---

#### [0.0015_HTTP_Starts_Here/try2_mux/main.go](0.0015-try2-mux-main.md)

**Rating: 7.5/10**  
**Topics:** ServeMux routing, multiple handlers  
**Key Issues:** No error handling, 2 spelling errors  
**Strengths:** Correct mux pattern, reduced spelling errors from try1

Demonstrates multiplexer for routing. Spelling improved but error handling still absent.

---

### HTTP Server Struct

#### [0.0015_HTTP_Starts_Here/try2_mux/Eg2/main.go](0.0015-try2-Eg2-main.md)

**Rating: 8/10**  
**Topics:** http.Server struct, Addr/Handler fields, stdlib exploration  
**Key Issues:** Addr bug (says 8080, uses 80), no error handling  
**Strengths:** Read stdlib source, discovered abstraction, conceptual understanding

Best exploratory learning in Week 5. Read `/net/http/server.go` to understand how `ListenAndServe` wraps `Server` struct.

---

#### [0.0015_HTTP_Starts_Here/try2_mux/eg2.2/main.go](0.0015-try2-eg2.2-main.md)

**Rating: 6.5/10**  
**Topics:** Default port behavior (port 80), minimal server  
**Key Issues:** No privilege documentation, PUT/POST confusion, rushed  
**Strengths:** Shows default port behavior

Quick experiment with port 80. Missing context (requires sudo) and has conceptual errors.

---

### JSON API Development

#### [0.0015_HTTP_Starts_Here/try3_POST/main.go](0.0015-try3-POST-main.md)

**Rating: 8.5/10** ⭐ **BEST HTTP FILE**  
**Topics:** POST, JSON, error handling, RWMutex, validation, status codes  
**Key Issues:** Spelling errors (handeler, decalering)  
**Strengths:** Error handling, concurrency safety, validation, JSON encode/decode

Production-quality API with proper error handling, RWMutex for thread-safe map access, JSON validation, and correct HTTP status codes.

---

#### [0.0015_HTTP_Starts_Here/stripedDown/main.go](0.0015-stripedDown-main.md)

**Rating: 3/10** ⚠️ **WORST FILE**  
**Topics:** Anti-pattern demonstration  
**Key Issues:** Intentionally broken, variable naming disaster, no comments  
**Strengths:** None (intentionally bad)

Deliberately bad code showing what not to do. Should have comments explaining why it's bad.

---

### HTML Templating

#### [0.0015_HTTP_Starts_Here/try4/main.go](0.0015-try4-main.md)

**Rating: 7/10**  
**Topics:** text/template, template.Must, Execute  
**Key Issues:** Race condition (global template), parse-per-request inefficiency  
**Strengths:** Template basics, multiple handler integration

Demonstrates HTML templating but has race condition from global variable and inefficient parsing.

---

### Singly Linked List (Data Structures)

#### [datastructures/list/linkList.go](datastructures-linkList.md)

**Rating: 6/10**  
**Topics:** Interface design, type definitions, linked list structures  
**Key Issues:** Interface signatures don't match implementation, Create() in interface, DoublyLinkList missing prev  
**Strengths:** Interface concept, any type usage, package organization

Interface design attempt but signatures don't match implementation - SingelyLinkList doesn't actually implement LinkList.

---

#### [datastructures/list/SingelyLinkList.go](datastructures-SingelyLinkList.md)

**Rating: 7/10**  
**Topics:** Linked list implementation, InsertAtBeginning, InsertAtLast, traversal  
**Key Issues:** InsertAtLast returns wrong node, empty placeholders, spelling errors  
**Strengths:** InsertAtBeginning perfect, correct algorithms, debug pattern

First linked list implementation. InsertAtBeginning is textbook correct (O(1)), but InsertAtLast has return value bug.

---

#### [datastructures/list/SinglyLinkedList_test.go](datastructures-SinglyLinkedList_test.md)

**Rating: 4/10**  
**Topics:** Linked list testing  
**Key Issues:** 2 of 3 tests don't run (lowercase 't'), zero assertions, PrintList instead of verification  
**Strengths:** One correct test name, mixed type testing

Tests print output instead of verifying correctness. Week 4 regression - you wrote proper assertions in queue tests but not here.

---

## Rating Summary

| File Category       | Count | Avg Rating | Best       | Worst | Status                   |
| ------------------- | ----- | ---------- | ---------- | ----- | ------------------------ |
| Basic HTTP          | 2     | 7.25/10    | 7.5/10     | 7/10  | Functional               |
| Server Struct       | 2     | 7.25/10    | 8/10       | 6.5/10| Learning                 |
| JSON API            | 2     | 5.75/10    | **8.5/10** | 3/10  | One good, one broken     |
| Templates           | 1     | 7/10       | 7/10       | 7/10  | Race condition           |
| **Linked List**     | **3** | **5.7/10** | **7/10**   | **4/10** | **Incomplete/buggy**  |
| **Overall**         | **10**| **6.8/10** | **8.5/10** | **3/10** | **Mixed quality**     |

---

## Week Progression

| Week | Focus                 | Rating     | Improvement |
| ---- | --------------------- | ---------- | ----------- |
| 1    | Basics                | 7.0/10     | Baseline    |
| 2    | OOP                   | 8.0/10     | +1.0        |
| 3    | Concurrency           | 7.7/10     | -0.3        |
| 4    | Patterns & Tests      | **9.0/10** | **+1.3**    |
| 5    | **HTTP & Linked List**| **6.8/10** | **-2.2**    |

**Significant regression from Week 4.** HTTP error handling discipline lost, linked list tests have no assertions despite Week 4 showing you know how to write proper tests.

---

## Technical Highlights

### 1. Reading stdlib Source Code

Comments in Eg2/main.go reference reading `net/http/server.go`. **Excellent learning practice** - understanding standard library implementation helps you use it correctly.

You discovered this abstraction:

```go
// In /net/http/server.go:
func ListenAndServe(addr string, handler Handler) error {
    server := &Server{Addr: addr, Handler: handler}
    return server.ListenAndServe()
}
```

**Same quality as Week 4's forSelect.go exploration.** You're learning by reading source code, not just tutorials.

### 2. Error Handling in try3_POST

First Week 5 HTTP file with proper error handling:

```go
err := json.NewDecoder(r.Body).Decode(&cat)
if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
}
```

All errors checked, proper status codes used, validation implemented. **This is Week 4 quality.**

### 3. Concurrency Safety (RWMutex)

```go
var cacheMutex sync.RWMutex

// Write:
cacheMutex.Lock()
CatCache[id] = cat
cacheMutex.Unlock()

// Read:
cacheMutex.RLock()
cat, ok := CatCache[id]
cacheMutex.RUnlock()
```

Correct use of RWMutex for concurrent access to shared map. Multiple readers, exclusive writers.

### 4. HTTP Method Routing (Go 1.22+)

```go
mux.HandleFunc("POST /cat", createCat)
mux.HandleFunc("GET /cat/{id}", getCat)

// In handler:
id := r.PathValue("id")
```

Modern Go 1.22+ routing syntax with method constraints and path parameters.

### 5. Linked List InsertAtBeginning (Perfect)

```go
func (l *SingelyLinkList) InsertAtBeginning(data any) *SingelyLinkList {
    node := Create(data)
    node.next = l
    return node
}
```

**Textbook correct.** O(1) time complexity, returns new head, proper pointer manipulation. This shows you understand linked list mechanics.

### 6. Template Parsing

```go
tem = template.Must(template.ParseFiles("./htmlFiles/cat.html"))
tem.Execute(w, nil)
```

Basic template usage (though has race condition from global variable and parse-per-request inefficiency).

---

## Critical Issues

### 1. Error Handling Regression

**Week 3-4**: Error handling patterns learned (comma-ok, error wrapping, http.Error)  
**Week 5**: Only 1 of 7 HTTP files has error handling

**Files without error handling:**

- try1/main.go
- try2_mux/main.go
- Eg2/main.go
- eg2.2/main.go
- stripedDown/main.go
- try4/main.go

**You learned this in Week 3. You demonstrated it in Week 4. Why isn't it in Week 5 HTTP files?**

### 2. Test Assertions Missing (Linked List)

**Week 4 queue tests:**

```go
func TestLinearQueue(t *testing.T) {
    q := LinearQueue{}
    q.Enqueue(106)
    if q.LengthOfQueue() != 1 {  // ✓ Assertion
        t.Fatal("Expected length 1")
    }
}
```

**Week 5 linked list tests:**

```go
func testCreate(t *testing.T) {  // ✗ Wrong name
    head := Create(123)
    head.PrintList()  // ✗ No assertion
}
```

**This is backwards learning.** You wrote proper tests in Week 4, then wrote manual verification scripts in Week 5.

### 3. Spelling Errors Persist

**Week 4 Review Said**: "Enable spell-check in your editor"  
**Week 5 Reality**: Spelling errors in every file

Common errors:

- "Initilizating" (appears in 5+ files)
- "Inseart" (linked list debug message)
- "Itterated" (linked list test)
- "handeler" (try3_POST)
- "decalering" (try3_POST)

**This is not a knowledge gap. This is not proofreading.**

### 4. InsertAtLast Return Bug

```go
func (l *SingelyLinkList) InsertAtLast(data any) *SingelyLinkList {
    node := Create(data)
    head := l
    for head.next != nil {
        head = head.next  // head is now LAST node
    }
    head.next = node
    return head  // ❌ Returns LAST node, not original head!
}
```

**Algorithm is correct, return value is wrong.** After loop, `head` points to last node, not original head.

**Impact**: Chained calls lose the list:

```go
list := Create(1)
list = list.InsertAtLast(2)  // Returns node 1 ✓
list = list.InsertAtLast(3)  // Returns node 2 ✗ (loses node 1!)
```

---

## What You're Learning

### New Concepts (Week 5)

1. **HTTP Servers** - net/http package, handlers, routing, server struct
2. **JSON APIs** - encoding/decoding, validation, status codes
3. **HTML Templates** - text/template, parsing, execution
4. **Linked Lists** - Pointer manipulation, head management, traversal
5. **RWMutex** - Read/write locks for concurrent map access

### Not Applied from Week 4

1. **Error Handling** - Learned in Week 3-4, not used in Week 5 HTTP
2. **Test Assertions** - Demonstrated in Week 4 queue tests, not used in Week 5 linked list tests
3. **Spell-Check** - Recommended in Week 4, still not enabled

---

## Recommendations

### Critical (Fix Immediately)

1. **Restore Error Handling**:

```go
// Wrong (Week 5):
server.ListenAndServe()

// Right (Week 3-4):
if err := server.ListenAndServe(); err != nil {
    log.Fatal(err)
}
```

2. **Add Test Assertions**:

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

3. **Enable Spell-Check**:

- VS Code: Install "Code Spell Checker" extension
- Enable in settings
- Review before committing

### Major (This Week)

1. **Fix InsertAtLast Return**:

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

2. **Fix LinkList Interface**:

```go
type LinkList interface {
    InsertAtBeginning(data any) LinkList  // Add return type
    InsertAtLast(data any)
    // Remove Create - it's a constructor
}
```

3. **Write HTTP Tests**:

```go
func TestCreateCat(t *testing.T) {
    req := httptest.NewRequest("POST", "/cat", strings.NewReader(`{"catName":"test"}`))
    w := httptest.NewRecorder()

    createCat(w, req)

    if w.Code != http.StatusNoContent {
        t.Fatalf("Expected 204, got %d", w.Code)
    }
}
```

### Soon (Week 6)

1. **HTTP Middleware** - Logging, auth, error recovery
2. **Context** - Request timeouts, cancellation
3. **Database** - Replace in-memory maps with SQLite/PostgreSQL
4. **Complete Linked List** - Implement InsertAt, InsertAfter, InsertBefore, Delete
5. **DoublyLinkedList** - Add prev pointer, implement bidirectional traversal

---

## Week 4 Recommendations Applied?

**HTTP Files:**
- ❌ Add error handling → Only 1 of 7 files has it
- ❌ Fix spelling → Still present (all files)
- ❌ Complete placeholders → Empty functions in SingelyLinkList.go

**Linked List Tests:**
- ❌ Test function naming → 2 of 3 tests have lowercase 't'
- ❌ Test assertions → Zero assertions, only PrintList()

**Overall: 0 of 2 critical recommendations followed.**

---

## Final Assessment

**Week 5 is a regression from Week 4's excellence.** You're learning new topics (HTTP, linked lists) but not applying patterns you already know (error handling, test assertions).

**The good**: try3_POST shows you **can** write production-quality code. InsertAtBeginning shows you **understand** linked list mechanics. Reading stdlib source shows you're learning deeply.

**The problem**: Inconsistent application of known patterns. Error handling exists in 1 HTTP file but absent in 6 others. Test assertions exist in queue tests but absent in linked list tests.

**Root cause**: Speed over quality. You're exploring new topics faster than you're applying learned discipline.

**Week 6 goal**: Apply Week 3-4 patterns to ALL code, not just some files.
