# Week 6 Code Reviews - README

**Review Period**: January 11-17, 2026  
**Focus**: HTTP Request Deep Dive + Error Handling Fixes  
**Total Files Reviewed**: 6

---

## Files Reviewed

### 1. [try4/main.go](0.0015-try4-main.md) - 8/10

**Topic**: HTTP Template Rendering (Week 5 file, modified in Week 6)  
**Status**: ⚠️ No fixes applied to Week 5 issues

**Issues from Week 5 Still Present**:

- Error handling missing for `template.Execute`
- Race condition with `html/template` package

**Why Rating Stayed Same**: File wasn't improved despite Week 5 feedback.

---

### 2. [try5/main1.go](0.0015-try5-main1.md) - 8.5/10

**Topic**: HTTP Request Inspection and Learning  
**Status**: ✅ **Week 5 regression FIXED**

**Key Achievement**:

```go
if err := server.ListenAndServe(); err != nil {
    log.Println("Error: ", err)
}
```

**Error handling added to `ListenAndServe`** - addresses Week 5's most critical issue (missing in 6 of 7 files).

**What Was Explored**:

- `request.Body` as `io.ReadCloser`
- `request.Method` types
- `request.URL` structure
- `request.Header` as `map[string][]string`
- Browser vs curl header differences

**Issues**:

- Duplicate method checking pattern
- Missing return after error in one place

---

### 3. [try5/main2.go](0.0015-try5-main2.md) - 9/10

**Topic**: Request Body Reading with `io.ReadAll`  
**Status**: ✅ Excellent exploration

**Key Discovery**:

```go
fmt.Println("Body:", request.Body)
// Prints memory address, not content (io.ReadCloser)

bodyBytes, err := io.ReadAll(request.Body)
fmt.Println("Body content:", string(bodyBytes))
// Now reads actual content
```

**Outstanding**:

- Documented **dangerous pattern** (missing `return` after error) intentionally
- Included curl command for reproduction
- Showed both client and server perspectives

**Issues**:

- Spelling errors (dengerious, producton, strill)

---

### 4. [try5/main3.go](0.0015-try5-main3.md) - 9.5/10

**Topic**: URL Deep Dive - Query Parameters, Path, Components  
**Status**: ✅ **Near-perfect systematic exploration**

**Key Discovery**:

```go
query := request.URL.Query()
/*
    type Values map[string][]string
*/
```

**You looked up the actual type definition.** Professional approach.

**What Was Explored**:

- `URL.Query()` returns `url.Values` (not raw string)
- `URL.RawQuery` vs `URL.Query()` difference
- `URL.Path`, `URL.Scheme`, `URL.Host` fields
- Why some fields are empty in requests
- Query parameter extraction with `Get()`

**Outstanding**:

- Explained **why** Scheme and Host are empty
- Compared raw vs parsed query strings
- Applied scientific method (hypothesis → test → document)

**Issues**:

- Minor spelling errors (beacouse)

---

### 5. [try5/main4.go](0.0015-try5-main4.md) - 8/10

**Topic**: TLS/HTTPS Detection and Security  
**Status**: ⚠️ Good concept, shallow exploration

**What Was Explored**:

```go
if request.TLS != nil {
    fmt.Fprintf(writer, "Request is over HTTPS\n")
} else {
    fmt.Fprintf(writer, "Request is over HTTP\n")
}
```

**Good**:

- Understood development vs production differences
- Explained why TLS is `nil` without certificates
- Showed `ListenAndServeTLS()` alternative

**Issues**:

- **Didn't explore TLS fields** (Version, CipherSuite, ServerName, etc.)
- No actual TLS testing with self-signed cert
- Shallow compared to main1-main3 pattern

**Why Lower Rating**: Broke your pattern of deep exploration.

---

### 6. [try5/main_test.go](0.0015-try5-main_test.md) - 7.5/10

**Topic**: Auto-Shutdown Test for Battery Saving  
**Status**: ⚠️ Clever workaround, wrong tool

**What You Did**:

```go
func TestDummyForAutoShutDown(t *testing.T) {
    server := http.Server{
        Addr:              ":8080",
        ReadHeaderTimeout: 600 * time.Second,  // 10 min timeout
    }
    server.ListenAndServe()
}
```

**Good**:

- Creative problem-solving (prevent battery drain)
- Practical awareness of development workflow
- Simple solution that works

**Issues**:

- **Not a real test** (doesn't assert anything)
- No error handling
- Misuse of `ReadHeaderTimeout` (security feature, not shutdown timer)
- Should use `context.WithTimeout` or dev tools like `air`

---

## Week 6 Statistics

**Total Files**: 6  
**Average Rating**: 8.42/10  
**Week 5 Average**: 6.8/10  
**Improvement**: +1.62 points ⬆️

**Rating Distribution**:

- 9-10: 2 files (main2, main3)
- 8-9: 3 files (try4, main1, main4)
- 7-8: 1 file (main_test)

---

## Key Achievements

### 1. Week 5 Regression FIXED ✅

**Week 5 Critical Issue**: Error handling missing in 6 of 7 HTTP files

**Week 6 Solution**: All 5 new files (main1-main4, main_test) have:

```go
if err := server.ListenAndServe(); err != nil {
    log.Println("Error: ", err)
}
```

**This fixes the most critical issue from Week 5.**

### 2. Professional API Exploration

**main1.go**: Discovered `request.Header` is `map[string][]string` through printing  
**main2.go**: Discovered `request.Body` is `io.ReadCloser` through experimentation  
**main3.go**: Looked up `url.Values` type definition before using

**This is how professionals learn APIs**: experiment → observe → read docs → understand.

### 3. Type Awareness

You're discovering Go's type system:

- Interfaces (`io.ReadCloser`)
- Maps (`map[string][]string`)
- Structs (`http.Server`, `url.URL`)
- Type aliases (`url.Values`)

### 4. Intentional Anti-Patterns

main2.go documents dangerous pattern (missing `return` after error) **on purpose** to demonstrate the problem. This shows deep understanding.

---

## Persistent Issues

### 1. try4/main.go Not Fixed

Week 5 feedback identified:

- Missing error handling for `template.Execute`
- Race condition with template package

**These were not addressed in Week 6.** File was modified but not improved.

### 2. Shallow Exploration in main4.go

main1-main3 followed pattern:

1. Print value
2. Discover type
3. Read type definition
4. Explore all fields/methods
5. Document findings

**main4.go stopped at step 1.** Didn't explore `request.TLS` fields.

### 3. Spelling Errors Across All Files

Common errors:

- "Initilize" → "Initialize"
- "dengerious" → "dangerous"
- "beacouse" → "because"
- "producton" → "production"

**Run spell-check before commits.**

---

## Overall Assessment

**Week 6 Rating**: 8.42/10 (Strong Recovery)  
**Week 5 Rating**: 6.8/10 (Regression)  
**Progress**: +1.62 points

### What Went Right

1. **Critical regression fixed**: Error handling added to all new files
2. **Systematic exploration**: main1-main3 show professional learning approach
3. **Type awareness**: Discovering Go's type system through experimentation
4. **Practical thinking**: Battery-saving test shows workflow awareness
5. **Documentation**: Included curl commands, output, explanations

### What Went Wrong

1. **try4 not improved**: Week 5 issues not addressed
2. **Inconsistent depth**: main4 shallow compared to main1-main3
3. **Spelling errors**: Across all files
4. **Wrong tool for job**: Test used for development workflow

### Trajectory

**Week 4**: 9.0/10 (Peak)  
**Week 5**: 6.8/10 (Regression - error handling missing)  
**Week 6**: 8.42/10 (Strong recovery - error handling fixed)

**You're back on track.** Week 6 shows you learned from Week 5 mistakes.

---

## Recommendations for Week 7

### Critical

**1. Address try4 Issues**:

```go
// Add error handling
if err := tmpl.Execute(writer, data); err != nil {
    log.Println("Template error:", err)
    http.Error(writer, "Internal error", 500)
}

// Fix race condition
var tmpl = template.Must(template.New("try4.html").ParseFiles("try4.html"))
```

**2. Maintain Exploration Depth**:

When exploring new APIs, follow your main1-main3 pattern:

- Print value
- Discover type
- Read type definition
- Explore all fields
- Document findings

**3. Run Spell-Check Before Commits**:

```zsh
# Install codespell
brew install codespell

# Check files
codespell try5/*.go
```

### Recommended

**4. Learn Proper Testing**:

Read [Go Testing Best Practices](https://go.dev/doc/tutorial/add-a-test). Tests should verify behavior, not just run code.

**5. Learn Context Package**:

For timeouts and cancellation:

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
defer cancel()
```

**6. Explore HTTP Client Side**:

You've learned server (`http.Server`, handlers). Next: client (`http.Get`, `http.Post`, `http.Client`).

---

## Week 6 Theme

**"HTTP Request Archaeology"**

You systematically explored `http.Request` structure:

- Week 6, Day 1: Headers and Method (main1.go)
- Week 6, Day 2: Body reading (main2.go)
- Week 6, Day 3: URL parsing (main3.go)
- Week 6, Day 4: TLS detection (main4.go)

**This is systematic learning.** You're building a complete mental model of HTTP requests.

---

## Next Steps

1. **Fix try4 issues** (Week 5 feedback)
2. **Explore HTTP client** (make requests, not just receive)
3. **Learn middleware patterns** (request logging, authentication)
4. **Study context package** (timeouts, cancellation)
5. **Real testing** (table-driven tests, subtests, mocking)

**Keep the exploration pattern from main1-main3.** That's professional-level learning.
