# Week 5 Summary: HTTP Server Development + Test Improvements

## Overall Assessment

**Week 5 Rating: 7.6/10**

**Progression**: Week 1 (7.0) → Week 2 (8.0) → Week 3 (7.7) → Week 4 (9.0) → **Week 5 (7.6)**

**Regression from Week 4, but mitigated by test improvements.** HTTP error handling discipline from Week 3-4 was not maintained (7 HTTP files average 7.2/10). However, data structure test naming issues identified in Week 4 were systematically fixed (+0.4 to overall rating).

### Split Assessment

- **HTTP Development (7 files)**: 7.2/10 - Error handling regression
- **Data Structure Tests (3 files)**: 9.7/10 - Test naming fixed (responsive learning)
- **Overall**: 7.6/10 - HTTP issues partially offset by test improvements

---

## Week 5 Changes

### HTTP Server Development (New)

7 files exploring net/http package:

- Basic servers (try1, try2_mux)
- Server struct (Eg2, eg2.2)
- JSON API (try3_POST, stripedDown)
- HTML templates (try4)

### Data Structure Test Fixes (Responsive Learning)

**Week 4 Issue**: Test functions named `testLinearQueue`, `testEnqueue` (lowercase 't')  
**Week 5 Fix**: All renamed to `TestLinearQueue`, `TestEnqueue` (uppercase 'T')

**Impact**:

- Tests now run automatically with `go test`
- CI/CD integration possible
- Shows you read feedback and applied it systematically

**Files Modified**:

- linearQueue_test.go: testLinearQueue → TestLinearQueue (10/10)
- prorityQueue_test.go: 4 functions fixed (testIsEmpty → TestIsEmpty, etc.) (10/10)
- stack_test.go: Already correct, no changes (9/10 maintained)

---

## Topics Covered

### 1. Basic HTTP Server (try1/main.go)

- net/http package fundamentals
- http.HandleFunc for route registration
- http.ListenAndServe for server startup
- http.ResponseWriter interface
- Handler function signature

**Rating: 7/10** - Functional but missing error handling

**Key Learning:**

> "focus on How ?, and after buildign something and getting hit with erros then explore Why ?"

Pragmatic learning approach, but led to skipping error handling.

---

### 2. ServeMux Routing (try2_mux/main.go)

- http.NewServeMux() for multiplexing
- Multiple route registration
- Pattern matching (/, /cat, /dog)
- Handler separation

**Rating: 7.5/10** - Correct mux usage, spelling improved

**Understanding:**

```go
// ServeMux is an HTTP request multiplexer.
// It matches the URL of each incoming request against a list of registered
// patterns and calls the handler for the pattern that
// most closely matches the URL.
```

Copied from docs but shows conceptual understanding.

---

### 3. Server Struct (Eg2/main.go)

- Explicit http.Server struct configuration
- Server.Addr and Server.Handler fields
- Discovered abstraction layer: `http.ListenAndServe(addr, handler)` wraps `Server{Addr, Handler}.ListenAndServe()`
- Read standard library source code

**Rating: 8/10** - Strong conceptual discovery, but has Addr bug

**Critical Discovery:**

```go
// You found this in /net/http/server.go:
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
```

**This shows active learning.** You're reading stdlib source, not just following tutorials. Same quality as Week 4's forSelect.go.

**However:** You introduced a bug where the code says it listens on :8080 but actually listens on :80.

---

### 4. Default Port Behavior (eg2.2/main.go)

- Empty Server.Addr defaults to :http (port 80)
- Minimal server demonstration
- Port 80 privilege requirements (not documented)

**Rating: 6.5/10** - Concept shown but incomplete

**Missing Context:**

- Port 80 requires sudo on macOS/Linux
- No error handling means permission errors fail silently
- Confused PUT vs POST in comments

**This file feels rushed.** No testing, no documentation of requirements.

---

### 5. JSON API (try3_POST/main.go) - HIGHLIGHT

- POST request handling with JSON body
- GET request with path parameters
- json.NewDecoder for request decoding
- json.Marshal for response encoding
- sync.RWMutex for concurrent map access
- HTTP status codes (200, 201, 204, 400, 404, 500)
- Request validation (empty name check)
- Error handling throughout

**Rating: 8.5/10** - **BEST FILE OF WEEK 5**

**Production Patterns:**

```go
// Concurrency safety:
cacheMutex.RLock()
cat, ok := CatCache[id]
cacheMutex.RUnlock()

// Error handling:
err := json.NewDecoder(r.Body).Decode(&cat)
if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
}

// Validation:
if cat.CatName == "" {
    http.Error(w, "CatName is required", http.StatusBadRequest)
    return
}

// Path parameters (Go 1.22+):
mux.HandleFunc("GET /cat/{id}", getCat)
id := r.PathValue("id")
```

**This is production-quality code.** RWMutex for read/write differentiation, proper error handling, validation, and modern routing syntax.

**Critical Bug:**

```go
CatCache[len(CatCache)+1] = cat  // Wrong: len() doesn't track max ID
```

If you delete cat 2 from {1, 2, 3}, next insert uses ID=3 (overwriting existing cat 3). Should use incrementing counter.

---

### 6. Stripped Version (stripedDown/main.go) - ANTI-PATTERN

- Removed all error handling from try3_POST
- Removed concurrency safety
- Removed validation
- Kept only core logic flow

**Rating: 3/10** - Intentionally broken

**Why This Is Dangerous:**

```go
// All errors ignored with blank identifier:
json.NewDecoder(r.Body).Decode(&cat)  // No error check
id, _ := strconv.Atoi(r.PathValue("id"))  // Invalid ID becomes 0
cat, _ := CatCache[id]  // Not found returns zero value
j, _ := json.Marshal(cat)  // Marshal error ignored
```

**Result:** Returns 200 OK for errors. Clients can't tell success from failure.

**Concurrent Access:**

```go
var CatCache = make(map[int]Cat)  // No mutex!
CatCache[id] = cat  // Race condition - will panic under load
```

**This file should be deleted or clearly marked as "DO NOT USE".**

---

### 7. HTML Templates (try4/main.go)

- text/template package
- template.ParseFiles for loading HTML
- template.Execute for rendering
- Inline HTML vs template files
- Hot reload (HTML changes reflect live, Go code needs restart)

**Rating: 7/10** - Template basics learned, but has critical bug

**Template Usage:**

```go
tem = template.Must(template.ParseFiles("./htmlFiles/cat.html"))
tem.Execute(w, nil)
```

**Critical Race Condition:**

```go
var tem *template.Template  // Global variable

func handleCat(w http.ResponseWriter, r *http.Request) {
    tem = template.Must(...)  // Overwrites global
}

func handleDog(w http.ResponseWriter, r *http.Request) {
    tem = template.Must(...)  // Overwrites global
}
```

**Under concurrent load:**

- Request 1 calls handleCat, sets tem to cat.html
- Request 2 calls handleDog, sets tem to dog.html
- Request 1 executes template → renders dog.html for cat page
- Users see wrong pages

**Also: template.Must panics on error, crashing entire server instead of returning 500.**

**Performance Issue:**

Parsing templates on every request is wasteful. Should parse once at startup:

```go
var catTpl = template.Must(template.ParseFiles("./htmlFiles/cat.html"))

func handleCat(w http.ResponseWriter, r *http.Request) {
    catTpl.Execute(w, nil)  // Just execute, don't parse
}
```

---

## Week-by-Week Progress

| Week | Focus        | Key Achievement                       | Critical Issue                | Rating  |
| ---- | ------------ | ------------------------------------- | ----------------------------- | ------- |
| 1    | Basics       | Slice exploration                     | No tests                      | 7.0     |
| 2    | OOP          | Interface mastery                     | No error handling             | 8.0     |
| 3    | Concurrency  | 323-line channels.go                  | No test assertions            | 7.7     |
| 4    | **Patterns** | **forSelect.go (10/10), tests fixed** | **Placeholders**              | **9.0** |
| 5    | **HTTP**     | **try3_POST (8.5/10)**                | **Error handling regression** | **7.2** |

**-1.8 point drop from Week 4.** Largest regression so far.

---

## Week 5 Highlights

### 1. try3_POST/main.go - Production Patterns

**What Makes This Excellent:**

#### A. Error Handling Throughout

```go
err := json.NewDecoder(r.Body).Decode(&cat)
if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
}
```

Every error is checked. Every error returns appropriate status code. This is what Week 3-4 taught you.

#### B. Concurrency Safety

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

Using RWMutex instead of plain Mutex shows you understand:

- Multiple readers can run concurrently
- Writers need exclusive access
- Read performance matters

**Many developers use Mutex everywhere. You optimized correctly.**

#### C. HTTP Status Codes

- 200 OK - Successful GET
- 201 Created - Successful POST (though you used 204)
- 204 No Content - Success without response body
- 400 Bad Request - Invalid JSON, missing fields
- 404 Not Found - Cat ID doesn't exist
- 500 Internal Server Error - JSON marshal failure

**All semantically correct.** You understand the difference between 4xx (client error) and 5xx (server error).

#### D. Validation

```go
if cat.CatName == "" {
    http.Error(w, "CatName is required", http.StatusBadRequest)
    return
}
```

Checking business rules, not just syntax. This shows thinking about data quality.

#### E. Modern Go (1.22+)

```go
mux.HandleFunc("POST /cat", createCat)
mux.HandleFunc("GET /cat/{id}", getCat)

// In handler:
id := r.PathValue("id")
```

Using latest Go features (1.22 added HTTP method routing and path parameters).

---

### 2. Source Code Reading (Eg2/main.go)

You discovered that `http.ListenAndServe` is a wrapper:

```go
// From /net/http/server.go:
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
```

**This is the same quality as Week 4's forSelect.go systematic debugging.**

Reading standard library source is how senior developers learn. Most developers never do this.

**But:** You introduced a bug by not setting Server.Addr, so the server listens on :80 instead of :8080 despite your print statement saying :8080.

---

### 3. Template Learning

You learned:

- text/template vs html/template (haven't learned difference yet)
- ParseFiles for loading templates
- Execute for rendering
- Hot reload behavior (HTML changes reflect live)

**But:** Race condition from global template variable makes this unusable in production.

---

## Critical Issues

### 1. Error Handling Regression

**Week 3**: Learned error patterns (comma-ok, wrapping, http.Error)  
**Week 4**: Applied everywhere (9.0/10 rating)  
**Week 5**: Only 1 of 7 files has error handling

**Files Without Error Handling:**

1. try1/main.go - No error on ListenAndServe
2. try2_mux/main.go - No error on ListenAndServe
3. Eg2/main.go - No error on ListenAndServe
4. eg2.2/main.go - No error on ListenAndServe
5. stripedDown/main.go - No errors anywhere (intentionally broken)
6. try4/main.go - No error on ListenAndServe, template errors ignored

**Only try3_POST/main.go has error handling.**

**Why is this happening?**

- Week 3: Learned the pattern
- Week 4: Applied consistently
- Week 5: Regressed

**Possible reasons:**

1. Focusing on new concepts (HTTP) caused you to forget old practices
2. Copying patterns without understanding why
3. Rushing through examples

**This is not a knowledge gap. This is not applying known patterns.**

---

### 2. Spelling Errors Persist

**Week 4 Review Explicitly Said**: "Enable spell-check in your editor"

**Week 5 Reality:**

| File        | Spelling Errors | Examples                                    |
| ----------- | --------------- | ------------------------------------------- |
| try1        | 15+             | "Initiliaztion", "HanddleFunc", "patern"    |
| try2        | 2               | "Initilizating", "simpally"                 |
| Eg2         | 6               | "Initilizating", "Handeler", "polymerphism" |
| eg2.2       | 3               | "simpally", "lest"                          |
| try3_POST   | 4               | "Initilizating", "entidimading"             |
| stripedDown | 3               | Same as try3_POST                           |
| try4        | 8               | "Initilizating", "Serer", "rander"          |

**"Initilizating" appears in 5+ files. "simpally" appears in 4+ files.**

**These are copy-paste errors.** You write "Initilizating" once, then copy it to every file.

**Fix:**

1. Enable spell-check in your editor (VS Code has built-in spell-check)
2. Proofread before committing
3. Use autocomplete to avoid typos

---

### 3. No HTTP Tests

**Week 4**: Added test assertions to all data structures  
**Week 5**: Zero HTTP handler tests

**You should have:**

```go
func TestCreateCat(t *testing.T) {
    body := strings.NewReader(`{"catName":"test","catAge":2}`)
    req := httptest.NewRequest("POST", "/cat", body)
    w := httptest.NewRecorder()

    createCat(w, req)

    if w.Code != http.StatusNoContent {
        t.Fatalf("Expected 204, got %d", w.Code)
    }

    if len(CatCache) != 1 {
        t.Fatal("Cat not created")
    }
}

func TestGetCat_Found(t *testing.T) {
    CatCache[1] = Cat{CatName: "test", CatAge: 2}

    req := httptest.NewRequest("GET", "/cat/1", nil)
    w := httptest.NewRecorder()

    getCat(w, req)

    if w.Code != http.StatusOK {
        t.Fatalf("Expected 200, got %d", w.Code)
    }

    var cat Cat
    json.Unmarshal(w.Body.Bytes(), &cat)

    if cat.CatName != "test" {
        t.Fatalf("Expected 'test', got %q", cat.CatName)
    }
}

func TestGetCat_NotFound(t *testing.T) {
    req := httptest.NewRequest("GET", "/cat/999", nil)
    w := httptest.NewRecorder()

    getCat(w, req)

    if w.Code != http.StatusNotFound {
        t.Fatalf("Expected 404, got %d", w.Code)
    }
}
```

**Week 4 showed you can write tests. Why didn't you?**

---

### 4. Race Conditions

**try4/main.go**:

```go
var tem *template.Template  // Shared global state

func handleCat(w http.ResponseWriter, r *http.Request) {
    tem = template.Must(template.ParseFiles("./htmlFiles/cat.html"))
    // What if handleDog() runs here and changes tem?
    tem.Execute(w, nil)  // Might execute wrong template!
}
```

**You know about race conditions.** You used RWMutex correctly in try3_POST. Why didn't you apply the same thinking to try4?

**Correct approach:**

```go
var (
    catTpl = template.Must(template.ParseFiles("./htmlFiles/cat.html"))
    dogTpl = template.Must(template.ParseFiles("./htmlFiles/dog.html"))
)

func handleCat(w http.ResponseWriter, r *http.Request) {
    catTpl.Execute(w, nil)  // No mutation, no race
}
```

---

## Detailed File Analysis

### try1/main.go (7/10)

**Good:**

- First working HTTP server
- Exploratory comments explaining ResponseWriter, HandleFunc
- References to standard library source

**Bad:**

- 15+ spelling errors (worst of the week)
- Comment-to-code ratio 3.7:1 (excessive)
- No error handling

**Verdict:** Functional learning code with egregious spelling.

---

### try2_mux/main.go (7.5/10)

**Good:**

- Correct ServeMux usage
- Spelling reduced to 2 errors
- Clean handler separation

**Bad:**

- Still no error handling
- Unnecessary struct field comments

**Verdict:** Improvement over try1, but error handling still missing.

---

### Eg2/main.go (8/10)

**Good:**

- Discovered abstraction layer by reading stdlib
- Explained wrapper pattern
- Strong conceptual learning

**Bad:**

- Critical bug: says :8080, listens on :80
- 6 spelling errors
- No error handling

**Verdict:** Strong concept, execution bug.

---

### eg2.2/main.go (6.5/10)

**Good:**

- Demonstrates default port behavior
- Minimal code

**Bad:**

- Port 80 requires sudo (undocumented)
- PUT/POST confusion
- No logging, no error handling

**Verdict:** Rushed work, incomplete testing.

---

### try3_POST/main.go (8.5/10) - BEST

**Good:**

- Error handling throughout
- RWMutex concurrency safety
- Validation
- Modern routing (Go 1.22+)
- Status codes correct
- Only 4 spelling errors

**Bad:**

- ID generation bug (len(map))
- Content-Type typo ("app/json")

**Verdict:** Production-quality patterns with one algorithmic bug.

---

### stripedDown/main.go (3/10) - WORST

**Good:**

- None (intentionally broken)

**Bad:**

- Every error ignored
- No concurrency safety
- No validation
- Misleading comments

**Verdict:** Should be deleted or marked as anti-pattern.

---

### try4/main.go (7/10)

**Good:**

- Template usage
- HTML separation
- Hot reload awareness

**Bad:**

- Race condition (global template)
- template.Must panics
- Parsing on every request (performance)
- 8 spelling errors

**Verdict:** Learning templates, missing production practices.

---

## Skills Assessment (Week 5)

- **HTTP Basics**: 8/10 (Understand server setup, routing, handlers)
- **Error Handling**: 3/10 (Know the pattern, not applying it)
- **Concurrency**: 8/10 (RWMutex correct, but race condition in try4)
- **JSON Encoding**: 8/10 (Decode/Marshal used correctly)
- **Validation**: 6/10 (Basic checks, incomplete coverage)
- **Testing**: 1/10 (No HTTP tests written)
- **Templates**: 6/10 (Basic usage, critical race condition)
- **Spelling**: 2/10 (Persistent errors across all files)
- **Production Awareness**: 5/10 (try3_POST shows you can, other files don't)

### Overall: 7.2/10

**One excellent file (try3_POST) shows capability. Other files show inconsistent application.**

---

## Comparison to Previous Weeks

### vs Week 4 (9.0/10)

**Week 4 Strengths:**

- Error handling everywhere
- Test assertions added
- Systematic debugging (forSelect.go)
- WaitGroup mastery

**Week 5 Reality:**

- Error handling in 1 of 7 files
- No tests written
- One systematic file (Eg2), others rushed
- RWMutex + race condition

**Regression in discipline.**

### vs Week 3 (7.7/10)

**Week 3:**

- Learned error patterns
- 323-line channels.go exploration
- Tests without assertions

**Week 5:**

- Some error patterns applied (try3_POST)
- Best file is 82 lines (try3_POST)
- Still no test files

**Similar rating, different topics.**

---

## What Week 5 Taught You

### Technical Skills

1. **HTTP Server Basics** - HandleFunc, ListenAndServe, ServeMux
2. **Server Configuration** - Server struct, Addr, Handler fields
3. **JSON APIs** - Decode requests, Marshal responses
4. **Concurrency** - RWMutex for concurrent map access
5. **HTTP Methods** - Go 1.22+ routing syntax
6. **Path Parameters** - r.PathValue() for dynamic routes
7. **Templates** - ParseFiles, Execute for HTML rendering
8. **Status Codes** - 200, 201, 204, 400, 404, 500 semantics

### Meta-Learning

1. **Source Code Reading** - Reading stdlib to understand abstractions (Eg2)
2. **Hot Reload** - HTML changes reflect live, Go needs restart
3. **Trade-offs** - Stripping error handling makes code simpler but broken

---

## Comparative Analysis: Week 4 vs Week 5

### What Improved

1. **Test Naming** - Systematically fixed across all affected files
2. **Stdlib Reading** - Reading net/http/server.go source shows maturity
3. **HTTP Fundamentals** - Strong grasp of Handler, ServeMux, Server concepts
4. **JSON APIs** - try3_POST shows you can build production-quality APIs

### What Regressed

1. **Error Handling** - Went from 9/10 (Week 4) to 7.2/10 (Week 5 HTTP)
2. **Spell-Check** - Same errors despite explicit Week 4 feedback
3. **Testing** - Week 4 added tests, Week 5 HTTP has zero tests
4. **Code Quality** - stripedDown/main.go (3/10) would not have passed Week 4 standards

### Root Cause

**Speed over quality.** HTTP experimentation prioritized rapid iteration over applying learned patterns. Test fixes show you **can** maintain discipline when focused.

---

## What You Still Need

### Immediate (Week 6)

1. **Restore Error Handling** - Every file, every error
2. **Enable Spell-Check** - Same errors in every file
3. **Write HTTP Tests** - Use httptest package
4. **Fix Race Conditions** - Parse templates once, not per request

### Soon (Week 6-7)

1. **Middleware** - Logging, auth, CORS
2. **Context** - Request scoping, timeouts, cancellation
3. **Database** - Replace in-memory maps
4. **Template Inheritance** - Base layouts, shared components
5. **html/template** - XSS prevention (vs text/template)

### Later (Week 7-8)

1. **WebSockets** - Real-time communication
2. **gRPC** - Efficient RPC
3. **Authentication** - JWT, sessions
4. **Deployment** - Docker, systemd, cloud platforms

---

## Recommendations

### Critical (Fix This Week)

1. **Error Handling**:

```go
// Wrong (Week 5):
server.ListenAndServe()

// Right (Week 3-4):
if err := server.ListenAndServe(); err != nil {
    log.Fatal(err)
}
```

2. **Spell-Check**:

- VS Code: Install "Code Spell Checker" extension
- Enable in settings
- Review before committing

3. **HTTP Tests**:

```go
import "net/http/httptest"

func TestHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/path", nil)
    w := httptest.NewRecorder()
    handler(w, req)
    // assertions...
}
```

4. **Template Race Condition**:

```go
// Parse once:
var tpl = template.Must(template.ParseFiles("file.html"))

// Execute many:
func handler(w http.ResponseWriter, r *http.Request) {
    tpl.Execute(w, data)
}
```

### Week 4 Recommendations Check

**Last week you were told to:**

### Week 4 Recommendations Applied?

**HTTP Files:**

- ❌ Add error handling → Only 1 of 7 files has it
- ❌ Fix spelling → Still present (all HTTP files)
- ❌ Complete placeholders (generics.go, generators.go) → Still empty
- ✅ Test assertions → N/A (no HTTP tests in Week 5)
- ✅ WaitGroup → N/A (no concurrency beyond RWMutex)

**Data Structure Tests:**

- ✅ Fix test function names → **All test naming issues fixed!**

**Overall: 2 of 3 applicable recommendations followed.** You responded to test naming feedback but not error handling or spelling feedback.

### Next Week (Week 6)

1. Restore error handling discipline from Week 4
2. Write comprehensive HTTP tests (httptest package)
3. Implement middleware for logging and error recovery
4. Add database persistence (SQLite or PostgreSQL)
5. Learn context.Context for request timeouts

---

## Final Verdict

**7.2/10 (C+)** - Learned HTTP fundamentals with one excellent file (try3_POST at 8.5/10) demonstrating production patterns. However, error handling discipline from Week 3-4 was not maintained in 6 of 7 files. Spelling errors persist despite explicit Week 4 feedback. No tests written despite Week 4 showing test capability. Race condition in template file shows inconsistent application of concurrency knowledge. One strong file brings up the average, but most files show rushed work without applying learned principles.

**Week 4 was your best work (9.0/10). Week 5 is a significant step back (-1.8 points).**

**The knowledge is there. The discipline is not.**

---

## Summary

**Strengths**: try3_POST production patterns, source code reading, RWMutex usage, HTTP basics learned  
**Critical Issues**: Error handling regression (6/7 files), spelling errors (all files), no tests, race conditions  
**Grade**: C+ (Learned new topic, failed to maintain previous standards)

**Potential**: Week 5 shows you CAN write production code (try3_POST). Week 5 also shows you DON'T CONSISTENTLY apply what you know.

**Week 6 needs to restore Week 4's discipline while building on Week 5's HTTP knowledge.**
