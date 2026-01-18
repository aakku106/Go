# Week 6 Code Review Index

Week 6 code reviews covering HTTP request internals exploration (try5 package).

**Review Period**: January 11-17, 2026  
**Focus**: HTTP request properties (Header, Method, URL, Body, TLS)

---

## Quick Summary

**Overall Rating: 3.8/10**  
**Total Files Reviewed: 5** (try5 package only, try4 unchanged)  
**Best File: try5/main3.go (6.5/10)**  
**Worst File: try5/main_test.go (1/10)**  
**Topics: HTTP headers, request body, URL parsing, TLS detection (failed)**  
**Critical Issue: Broken error handling in 3/5 files + non-functional TLS code**  
**Pattern: All files print to server console, none send responses to clients**

---

## Week 6 Highlights

### Notable Achievements

1. **URL Internals** - main3.go has excellent explanation of Query() parsing, RawQuery, Path importance
2. **Header Analysis** - main1.go documents curl vs Safari vs Firefox header differences
3. **Body Reading** - main2.go demonstrates io.ReadAll correctly (despite broken error handling)
4. **Security Awareness** - main3.go warns against passing sensitive data in query strings

### Critical Failures

1. **Broken Error Handling (3 files)** - main1.go, main2.go, main3.go all call `http.Error()` without `return`, code continues execution
2. **Non-Functional TLS Code** - main4.go calls `ListenAndServeTLS()` with no parameters, cannot compile/run
3. **Testing Framework Abuse** - main_test.go misuses testing for 600s auto-kill timer instead of actual testing
4. **No Client Responses** - All 5 files print to server console but send nothing (or garbage) to HTTP clients
5. **Persistent Spelling Errors** - "Initilize" appears in ALL 5 files despite Week 5 feedback

### Week 5 Regression

- **Error Handling**: Week 5's try3_POST had proper error handling (8.5/10). Week 6 completely regressed.
- **Testing**: Week 5 had basic tests. Week 6's test file is a 1/10 anti-pattern.

---

## File Reviews

### HTTP Request Properties

#### [0.0015_HTTP_Starts_Here/try5/main1.go](0.0015-try5-main1.md)

**Rating: 5/10**  
**Topics:** Header, Method, URL basic properties  
**Key Issues:** Broken error handling (no return after http.Error), duplicate error check, no client responses  
**Strengths:** Good header analysis (curl vs browsers), educational comments

Explores request.Header showing how different clients send different headers. Broken error handling in handleContact - calls `http.Error()` but continues executing. Also has duplicate if condition checking same thing twice.

---

#### [0.0015_HTTP_Starts_Here/try5/main2.go](0.0015-try5-main2.md)

**Rating: 4/10**  
**Topics:** Request body reading with io.ReadAll  
**Key Issues:** Broken error handling (again), misleading comment about body behavior, documents bug as feature  
**Strengths:** Correct io.ReadAll usage, error handling on ReadAll

Demonstrates body reading but has same error handling bug as main1.go. Worse: comment **acknowledges** the bug exists "for testing" instead of fixing it. Misleading explanation of why body prints nothing.

---

#### [0.0015_HTTP_Starts_Here/try5/main3.go](0.0015-try5-main3.md)

**Rating: 6.5/10** ⭐ **BEST FILE THIS WEEK**  
**Topics:** URL.Query(), RawQuery, Path, User, Scheme, Host  
**Key Issues:** Nil pointer crash risk (User.Username), 100% code duplication (handleContact3), no client responses  
**Strengths:** Excellent Query() explanation, practical use cases, security advice, URL.Path routing importance

Best technical content this week. Explains URL parsing thoroughly with good examples (query params, multiple values, security). Has nil pointer crash risk on `request.URL.User.Username()` and handleContact3 duplicates all of handleRoot3.

---

#### [0.0015_HTTP_Starts_Here/try5/main4.go](0.0015-try5-main4.md)

**Rating: 2/10** ⭐ **WORST CODE THIS WEEK**  
**Topics:** TLS detection (failed attempt)  
**Key Issues:** Calls ListenAndServeTLS() with 0 params (requires 2), code doesn't compile, non-functional  
**Strengths:** Correct TLS nil check, notes TLS is transport layer

**Non-functional code.** Calls `server.ListenAndServeTLS()` without required certificate files. Cannot compile. All output labels say "TLS version:" but print different properties. This file is **completely broken**.

---

#### [0.0015_HTTP_Starts_Here/try5/main_test.go](0.0015-try5-main_test.md)

**Rating: 1/10** ⭐ **WORST TEST FILE EVER**  
**Topics:** Testing framework abuse  
**Key Issues:** Zero test logic, misuses testing for 600s timeout, calls broken main4 code, teaches bad practices  
**Strengths:** Honest about being "useless"

Abuses Go testing framework as auto-kill timer for battery saving. Has zero assertions, zero verifications, blocks until test timeout. Self-described as "useless and un-profectional" [sic]. Should be deleted and replaced with proper context.WithTimeout.

---

### Unchanged Files

#### [0.0015_HTTP_Starts_Here/try4/main.go](0.0015-try4-main.md)

**Modified**: +2 lines (comment only)  
**Functional Changes**: None  
**Week 6 Impact**: Zero

Only added `// NEXT: ../try5/main.go` comment. No code changes. See Week 5 review for analysis.

---

## Week 6 Statistics

**Lines of Code**: ~325 (excluding try4)  
**Spelling Errors**: 35+ across all files  
**"Initilize" Count**: 5 (one per file)  
**Functional Files**: 3/5 (main1, main2, main3 work despite bugs)  
**Broken Files**: 2/5 (main4 doesn't compile, main_test abuses framework)  
**Files with Error Handling Bugs**: 3/5 (main1, main2, main3)  
**Files Sending Client Responses**: 0/5

---

## Critical Patterns

### Broken Error Handling (Repeated 3 Times)

```go
if condition {
    http.Error(writer, "message", statusCode)
}
// NO RETURN - BUG CONTINUES IN 3 FILES
```

This pattern appears in main1.go, main2.go, main3.go. **You must add `return` after `http.Error()`.**

### No Client Responses

All 5 files print to server console but send nothing meaningful to clients. They're logging scripts, not HTTP services.

### Persistent Typos

"Initilize" → "Initialize" appears in EVERY file despite Week 5 review noting this.

---

## Week Comparison

| Metric           | Week 5    | Week 6    | Change   |
| ---------------- | --------- | --------- | -------- |
| Avg Rating       | 6.8/10    | 3.8/10    | -3.0 ⬇️  |
| Best File        | 8.5/10    | 6.5/10    | -2.0 ⬇️  |
| Error Handling   | 1/7 files | 0/5 files | -100% ⬇️ |
| Functional Tests | Basic     | Abused    | -100% ⬇️ |
| Spelling Errors  | Improving | Same      | 0%       |

**Week 6 is significantly worse than Week 5.**

---

## What Worked

✅ URL parsing explanation (main3.go)  
✅ Header analysis across clients (main1.go)  
✅ Body reading mechanics (main2.go)  
✅ Security awareness (query strings)  
✅ Educational comments (when accurate)

---

## What Failed

❌ Error handling (3 files broken)  
❌ TLS implementation (doesn't compile)  
❌ Testing (framework abuse)  
❌ Spell-checking (35+ errors)  
❌ Client responses (none functional)  
❌ Code duplication (handleContact3)  
❌ Following Week 5 feedback

---

## Immediate Actions Required

1. **Fix error handling** - Add `return` after EVERY `http.Error()` call (3 files)
2. **Fix or delete main4.go** - Either add certificates or remove broken code
3. **Delete main_test.go** - Replace with proper timeout using `context.WithTimeout`
4. **Enable spell-check** - "Initilize" in 5 consecutive files is unacceptable
5. **Add client responses** - These are HTTP servers, they should respond to clients
6. **Remove handleContact3** - It's 100% duplicate of handleRoot3

---

## Final Verdict

**Week 6: 3.8/10 - Major Regression**

HTTP request exploration with good educational content but **critically broken implementation**. Three files have identical error handling bugs, one file doesn't compile, and the test file abuses the testing framework. This is worse than Week 5 in every measurable way.

**The good**: You're exploring HTTP internals and documenting your learning.  
**The bad**: You're not applying feedback from previous weeks.  
**The ugly**: You documented bugs as features instead of fixing them.

**Fix the error handling immediately before moving forward.**
