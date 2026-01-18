# Week 6 Summary

**Period**: January 11-17, 2026  
**Rating**: 3.8/10  
**Status**: Major Regression from Week 5

---

## Overview

Week 6 explored HTTP request internals through the try5 package, focusing on Header, Method, URL, Body, and TLS properties. While educational content quality improved (especially URL parsing explanation), implementation quality **collapsed**. Three files have identical error handling bugs, one file is non-functional, and the test file abuses the testing framework.

**This is the worst-performing week so far.**

---

## Files Created/Modified

| File              | Status   | Lines | Rating | Key Issue                             |
| ----------------- | -------- | ----- | ------ | ------------------------------------- |
| try5/main1.go     | New      | 128   | 5/10   | Broken error handling, no return      |
| try5/main2.go     | New      | 67    | 4/10   | Broken error handling + documents bug |
| try5/main3.go     | New      | 102   | 6.5/10 | Nil pointer risk, code duplication    |
| try5/main4.go     | New      | 46    | 2/10   | Non-functional, doesn't compile       |
| try5/main_test.go | New      | 13    | 1/10   | Abuses testing framework              |
| try4/main.go      | Modified | +2    | N/A    | Comment only, no code change          |

**Average Rating**: 3.7/10 (excluding try4)

---

## Critical Issues

### 1. Broken Error Handling (3 Files)

**Files Affected**: main1.go, main2.go, main3.go

```go
// WRONG - appears in 3 files
if request.Method != http.MethodGet {
    http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
}
// Code continues - BUG
fmt.Println("This still executes!")
```

`http.Error()` sends error response to client but **doesn't stop function execution**. Must add `return`:

```go
// CORRECT
if request.Method != http.MethodGet {
    http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
    return  // REQUIRED
}
```

**Impact**: Server sends error response, then continues processing and sends more data, creating malformed HTTP responses.

**Why This Is Critical**: This is the **third consecutive file** with the same bug. It shows:

1. Not understanding http.Error behavior
2. Copy-pasting broken code
3. Not testing code (would catch this immediately)

### 2. Non-Functional TLS Code

**File**: main4.go

```go
// BROKEN - doesn't compile
if err := server.ListenAndServeTLS(); err != nil {
    log.Println("Error: ", err)
}
```

`ListenAndServeTLS()` requires two parameters: `certFile` and `keyFile`. This code has zero parameters and **cannot compile**. Even if it compiled, server would fail without actual certificate files.

**Impact**: Entire file is non-functional. Cannot demonstrate TLS.

### 3. Testing Framework Abuse

**File**: main_test.go

```go
func TestDummy(t *testing.T) {
    InitilizeServer4()  // Blocks for 600s
}
```

- Zero assertions
- Zero test logic
- Blocks until Go test timeout (600s)
- Calls broken main4.go code

**Purpose** (per comment): Battery-saving auto-kill timer for TMUX sessions.

**Why This Is Wrong**: Using testing framework for process management. Should use `context.WithTimeout` or signal handling instead.

---

## Pattern Analysis

### The "Initilize" Epidemic

**"Initilize" appears in ALL 5 files.**

Week 5 review specifically called out this typo. Week 6 has it in **every single file**. This shows:

- Not reading feedback
- Not using spell-check
- Copy-pasting without reviewing

### No Client Responses

**0 of 5 files send meaningful responses to HTTP clients.**

All files print to server console but send nothing (or garbage) to clients:

- main1.go: Prints to console only
- main2.go: Sends "weeeeeeeeeeeeeeeeeeeeaaaaaaaaaaaaaaaaaa%"
- main3.go: Prints to console only
- main4.go: Doesn't run
- main_test.go: Not an HTTP handler

These are **logging scripts**, not HTTP services.

### Documented Bugs

main2.go line 63:

```go
// Althow we had if condition to filter out only GET, but since we dont have return in that if block
// whatever below will strill run and response (Could be dengerious in producton, this is just to show/test)
```

**This comment acknowledges the bug exists, then justifies it as "just to show/test".**

Never document bugs as features. If you know it's wrong, **fix it**.

---

## What Worked

### URL Parsing Deep Dive (main3.go - 6.5/10)

**Best technical content this week.** Excellent explanations:

- How Query() splits on `&` and `=`
- RawQuery use cases (proxies, verification, logging)
- URL.Path as "spine of HTTP" (routing)
- Security: don't pass sensitive data in query strings
- Multiple values for same key handled correctly

### Header Analysis (main1.go - 5/10)

Good documentation of header differences:

- curl: minimal (`Accept: */*`)
- Safari: optimized for battery/speed
- Firefox: verbose for cross-platform compatibility

Shows understanding of client behavior.

### Body Reading (main2.go - 4/10)

Correctly demonstrates:

- `io.ReadAll(request.Body)` usage
- Body is `io.ReadCloser`
- Stream consumption after reading

Implementation broken but concept demonstrated.

---

## What Failed

### Error Handling Collapse

**Week 5**: try3_POST (8.5/10) had proper error handling with `return`  
**Week 6**: 3/5 files have broken error handling, 0/5 have proper handling

This is a **100% regression**.

### Testing Quality

**Week 5**: Basic but functional tests  
**Week 6**: Test file abuses framework for process management (1/10)

### Code Quality

**Week 5**: Average rating 6.8/10  
**Week 6**: Average rating 3.8/10  
**Decline**: -3.0 points (-44%)

### Following Feedback

Week 5 review identified:

- "Initilize" typo (appears in 5/5 Week 6 files)
- Error handling issues (worse in Week 6)
- Spelling errors (35+ in Week 6)

**Zero improvement.**

---

## Statistics

| Metric                      | Value                            |
| --------------------------- | -------------------------------- |
| Total Files                 | 5 new, 1 modified (comment only) |
| Lines of Code               | ~325                             |
| Spelling Errors             | 35+                              |
| "Initilize" Count           | 5                                |
| Broken Error Handling       | 3 files                          |
| Non-Functional Code         | 2 files                          |
| Proper Tests                | 0                                |
| Client-Facing HTTP Services | 0                                |

---

## Learning Outcomes

### Concepts Understood ✅

- HTTP headers vary by client
- `map[string][]string` for Header and Query
- Body is `io.ReadCloser`, must use `io.ReadAll()`
- URL.Query() parsing (`&` and `=` splitting)
- URL.Path for routing
- TLS is transport layer (conceptually)
- Security: no sensitive data in query strings

### Implementation Failures ❌

- Error handling (http.Error requires return)
- TLS server setup (requires certificates)
- Testing framework purpose
- Client response patterns
- Code duplication
- Spell-checking

---

## Comparison to Week 5

| Category            | Week 5     | Week 6      | Change        |
| ------------------- | ---------- | ----------- | ------------- |
| **Average Rating**  | 6.8/10     | 3.8/10      | **-3.0**      |
| **Best File**       | 8.5/10     | 6.5/10      | **-2.0**      |
| **Worst File**      | 3/10       | 1/10        | **-2.0**      |
| **Error Handling**  | Some files | Broken in 3 | **Worse**     |
| **Testing**         | Basic      | Abused      | **Worse**     |
| **Spell Errors**    | Improving  | Same        | **No change** |
| **Functional Code** | Most files | 3/5 files   | **Worse**     |

**Every metric declined.**

---

## Root Cause Analysis

### Why Did Week 6 Fail?

1. **Not reading previous feedback** - "Initilize" typo repeated 5 times after being flagged in Week 5
2. **Copy-paste without review** - Same error handling bug in 3 files
3. **Not testing code** - None of these files were run or tested properly
4. **Documenting bugs instead of fixing** - main2.go comment justifies broken error handling
5. **Incomplete research** - Started TLS file without understanding certificate requirements
6. **Clever hacks over proper solutions** - Using test timeout for battery management

### What Changed From Week 5?

Week 5 had try3_POST (8.5/10) with:

- Proper error handling with `return`
- RWMutex for concurrency
- JSON validation
- Status codes

**Week 6 forgot everything Week 5 did right.**

---

## Recommendations

### Immediate (Before Week 7)

1. **Fix error handling** - Add `return` after every `http.Error()` in main1.go, main2.go, main3.go
2. **Fix or delete main4.go** - Either generate certificates or remove broken code
3. **Delete main_test.go** - Use `context.WithTimeout` in main functions instead
4. **Enable spell-check** - "Initilize" × 5 is unacceptable
5. **Test your code** - Run it before committing

### Short-term (Week 7-8)

1. **Read Week 5 feedback** - Review what worked (try3_POST) and apply it
2. **Set up spell-checker** - VS Code, vim, whatever editor you use
3. **Write real tests** - Use httptest.NewRecorder
4. **Add client responses** - HTTP servers should respond to clients
5. **Code review before commit** - Read your own code

### Long-term

1. **Learn error patterns** - Understand functions that require manual return
2. **Study testing** - Read https://go.dev/doc/tutorial/add-a-test
3. **TLS fundamentals** - Understand certificate requirements before coding
4. **Stop documenting bugs** - If you know it's wrong, fix it
5. **Apply feedback** - Previous reviews are worth reading

---

## Final Verdict

**Week 6: 3.8/10 - Major Regression**

Good educational exploration of HTTP request internals, **terrible implementation**. Three files have identical error handling bugs showing copy-paste without understanding. One file doesn't compile. Test file abuses framework. Zero improvement from Week 5 feedback.

**This week demonstrates**:

- ✅ Curiosity about HTTP internals
- ✅ Ability to read documentation
- ✅ Writing explanatory comments
- ❌ Following feedback
- ❌ Testing code before committing
- ❌ Understanding error control flow
- ❌ Spell-checking
- ❌ Code review discipline

**You can explore and learn concepts, but you're not consolidating that knowledge into working code.**

---

## Action Items for Week 7

**Must Do**:

1. Fix all `http.Error` calls - add `return`
2. Delete or fix main4.go
3. Delete main_test.go
4. Enable spell-check
5. Read Week 5 review

**Should Do**:

1. Write actual tests
2. Add client responses
3. Remove code duplication
4. Apply Week 5 feedback

**Could Do**:

1. Refactor into `package main` with proper main()
2. Add proper shutdown handling
3. Learn certificate generation for HTTPS

**Bottom Line**: Week 6 is a step backward. Fix the error handling before moving to Week 7.
