# Week 5 Code Review Index

Week 5 code reviews covering HTTP server development.

**Review Period**: January 4-10, 2026  
**Focus**: net/http package, ServeMux routing, JSON APIs, HTML templating

---

## Quick Summary

**Overall Rating: 7.2/10**  
**Total Files Reviewed: 7**  
**Best File: try3_POST/main.go (8.5/10)**  
**Worst File: stripedDown/main.go (3/10)**  
**Topics: HTTP servers, routing, JSON encoding, templates, POST/GET**  
**Critical Issue: Error handling present in only 1 of 7 files**

---

## Week 5 Highlights

### Notable Achievements

1. **try3_POST/main.go** - First HTTP file with proper error handling, concurrency safety (RWMutex), JSON encoding/decoding, and validation
2. **Template Usage** - Learned text/template package for HTML rendering
3. **HTTP Method Routing** - Used Go 1.22+ syntax ("POST /cat", "GET /cat/{id}")
4. **Progressive Learning** - Systematic progression from basic server to JSON API

### Major Issues

1. **Error Handling Regression** - Only 1 of 7 files has error handling (try3_POST)
2. **Spelling Errors** - Persistent across all files despite Week 4 feedback
3. **Race Conditions** - Global template variable in try4
4. **No Testing** - Zero test files for HTTP handlers

---

## File Reviews

### HTTP Server Basics

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
**Topics:** Explicit Server struct, abstraction understanding  
**Key Issues:** Server.Addr bug (says :8080, listens on :80), no error handling, 6 spelling errors  
**Strengths:** Discovered standard library abstraction, read source code, explained wrapper pattern

Strong conceptual understanding marred by critical Addr bug. Shows active learning by reading stdlib source.

---

#### [0.0015_HTTP_Starts_Here/try2_mux/Eg2/eg2.2/main.go](0.0015-try2-Eg2-eg2.2-main.md)

**Rating: 6.5/10**  
**Topics:** Default port behavior (port 80)  
**Key Issues:** Port 80 requires sudo (undocumented), no error handling, PUT/POST confusion, no logging  
**Strengths:** Minimal code demonstrating concept

Shows default port 80 behavior but doesn't document sudo requirement or handle permission errors. Feels rushed.

---

### JSON API

#### [0.0015_HTTP_Starts_Here/try3_POST/main.go](0.0015-try3-POST-main.md) 🌟

**Rating: 8.5/10** - **BEST FILE**  
**Topics:** POST/GET, JSON encoding/decoding, sync.RWMutex, path parameters  
**Key Issues:** ID generation bug (uses len(map)), Content-Type typo ("app/json"), 4 spelling errors  
**Strengths:** Error handling throughout, RWMutex for concurrency, proper status codes, validation

First Week 5 file with comprehensive error handling and production patterns. Shows significant improvement. Best work of the week.

---

#### [0.0015_HTTP_Starts_Here/try3_POST/stripedDown/main.go](0.0015-try3-POST-stripedDown-main.md)

**Rating: 3/10** - **WORST FILE**  
**Topics:** Anti-pattern demonstration  
**Key Issues:** Every error ignored, no concurrency safety (race conditions), no validation, misleading comments  
**Strengths:** None (intentionally broken)

Stripped version with all error handling removed. This code will crash under load. Should be deleted or clearly marked as anti-pattern example.

---

### HTML Templating

#### [0.0015_HTTP_Starts_Here/try4/main.go](0.0015-try4-main.md)

**Rating: 7/10**  
**Topics:** HTML templates, text/template package  
**Key Issues:** Race condition (global template variable), template.Must panics, parsing on every request, 8 spelling errors  
**Strengths:** Template usage, HTML separation, hot reload awareness

Demonstrates templates but has critical race condition from global variable shared across concurrent requests. Parses templates on every request (performance issue).

---

## Rating Summary

| File Category | Count | Avg Rating | Best       | Status               |
| ------------- | ----- | ---------- | ---------- | -------------------- |
| Basic HTTP    | 2     | 7.25/10    | 7.5/10     | Functional           |
| Server Struct | 2     | 7.25/10    | 8/10       | Learning             |
| JSON API      | 2     | 5.75/10    | **8.5/10** | One good, one broken |
| Templates     | 1     | 7/10       | 7/10       | Race condition       |
| **Overall**   | **7** | **7.2/10** | **8.5/10** | **Passing**          |

---

## Week Progression

| Week | Focus            | Rating     | Improvement |
| ---- | ---------------- | ---------- | ----------- |
| 1    | Basics           | 7.0/10     | Baseline    |
| 2    | OOP              | 8.0/10     | +1.0        |
| 3    | Concurrency      | 7.7/10     | -0.3        |
| 4    | Patterns & Tests | **9.0/10** | **+1.3**    |
| 5    | **HTTP Servers** | **7.2/10** | **-1.8**    |

**Significant drop from Week 4.** Error handling regression is the primary cause.

---

## Technical Highlights

### 1. Error Handling in try3_POST

First Week 5 file with proper error handling:

```go
err := json.NewDecoder(r.Body).Decode(&cat)
if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
}
```

All errors checked, proper status codes used, validation implemented.

### 2. Concurrency Safety (RWMutex)

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

Correct use of RWMutex for concurrent access to shared map.

### 3. HTTP Method Routing (Go 1.22+)

```go
mux.HandleFunc("POST /cat", createCat)
mux.HandleFunc("GET /cat/{id}", getCat)

// In handler:
id := r.PathValue("id")
```

Modern Go 1.22+ routing syntax with path parameters.

### 4. Template Parsing

```go
tem = template.Must(template.ParseFiles("./htmlFiles/cat.html"))
tem.Execute(w, nil)
```

Basic template usage (though has race condition from global variable).

---

## Critical Issues

### 1. Error Handling Regression

**Week 3-4**: Error handling patterns learned (comma-ok, error wrapping, http.Error)  
**Week 5**: Only 1 of 7 files has error handling

**Files without error handling:**

- try1/main.go
- try2_mux/main.go
- Eg2/main.go
- eg2.2/main.go
- stripedDown/main.go
- try4/main.go

**You learned this in Week 3. Why aren't you applying it?**

### 2. Spelling Errors Persist

**Week 4 Review Said**: "Enable spell-check in your editor"  
**Week 5 Reality**: Spelling errors in every file

Common errors:

- "Initilizating" (appears in 5+ files)
- "simpally" (appears in 4+ files)
- "lest" instead of "let's"
- "handeler" instead of "handler"

**This is not a knowledge gap. This is not proofreading.**

### 3. Race Conditions

**try4/main.go**:

```go
var tem *template.Template  // Global

func handleCat(w http.ResponseWriter, r *http.Request) {
    tem = template.Must(...)  // Overwrites global
}

func handleDog(w http.ResponseWriter, r *http.Request) {
    tem = template.Must(...)  // Overwrites global
}
```

Concurrent requests will corrupt template state.

### 4. No Tests

**Week 4**: Added test assertions to all data structures  
**Week 5**: Zero HTTP handler tests

You should have:

- TestHandleRoot
- TestCreateCat
- TestGetCat (found)
- TestGetCat (not found)
- TestInvalidJSON

---

## Areas for Improvement

### 1. Restore Error Handling Discipline

Every HTTP handler should:

```go
if err := server.ListenAndServe(); err != nil {
    log.Fatal(err)
}
```

Not optional. Required.

### 2. Fix Spelling

Enable spell-check. The same errors appear in every file.

### 3. Add HTTP Tests

Learn httptest package:

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

### 4. Parse Templates Once

```go
// Wrong (current):
func handler(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles("file.html"))  // Every request
    t.Execute(w, nil)
}

// Right:
var tpl = template.Must(template.ParseFiles("file.html"))  // Once

func handler(w http.ResponseWriter, r *http.Request) {
    tpl.Execute(w, nil)
}
```

### 5. Learn Middleware

Logging, auth, CORS all done via middleware:

```go
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path)
        next(w, r)
    }
}

mux.HandleFunc("/cat", loggingMiddleware(handleCat))
```

---

## Comparison to Week 4

| Aspect         | Week 4               | Week 5                                |
| -------------- | -------------------- | ------------------------------------- |
| Error Handling | Comprehensive        | Mostly absent                         |
| Testing        | Added assertions     | No tests                              |
| Spelling       | Reduced              | Same level                            |
| Concurrency    | WaitGroup mastery    | RWMutex (good) + race condition (bad) |
| Best File      | forSelect.go (10/10) | try3_POST (8.5/10)                    |
| Consistency    | High                 | Low                                   |

**Week 4 showed systematic improvement. Week 5 shows inconsistent application of learned principles.**

---

## Final Verdict

**7.2/10 (C+)** - Learned HTTP server basics, routing, JSON APIs, and templates. try3_POST shows you can write production-quality code when focused. However, most files lack error handling, spelling errors persist, and Week 4's discipline was not maintained. One excellent file (try3_POST) brings up the average, but the other files show rushed work.

**This is passing but disappointing after Week 4's strong performance.**

---

## Recommendations for Week 6

### Must Fix

1. **Error Handling**: Every file, every error, every time
2. **Spelling**: Enable spell-check immediately
3. **Testing**: Write httptest tests for all handlers

### Should Learn

1. Middleware pattern
2. Context for request scoping
3. Database integration (replace in-memory maps)
4. Proper template structure (parse once, execute many)

### Could Explore

1. WebSockets
2. gRPC
3. GraphQL
4. Authentication/Authorization

---

## Summary

**Strengths**: JSON API with error handling (try3_POST), RWMutex concurrency, template usage  
**Critical Issues**: Error handling regression (6 of 7 files), spelling errors (all files), no tests  
**Grade**: C+ (Learned HTTP basics, missing production discipline)
