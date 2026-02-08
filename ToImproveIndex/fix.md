# 8-Week Go Challenge Improvement Index

_8 weeks. 96 files. 5.8/10 average. Internship secured._

## What The Data Shows

**Quality Progression:**  
Week 1-4: 7.0 → 8.0 → 7.7 → 9.0 (Learning phase)  
Week 5-6: 6.8 → 4.9 (Regression)  
Week 7-8: 5.6 → 5.8 (Slight recovery)

**Best Work:**  
forSelect.go (Week 4: 10/10), stack.go (Week 8: 8/10)

**Gap Between Best and Worst:**  
Can write production-ready code (8/10). Also commits non-compiling code (2/10). 6-point quality gap exists within same week.

---

## Core Reality

Knowledge exists. Consistency doesn't.

Evidence:

- Week 4: Perfect WaitGroup implementation + systematic debugging (forSelect.go: 10/10)
- Week 6: Zero assertions in tests, broken error handling in 3 consecutive files
- Week 8: Thread-safe stack with sync.Mutex (8/10) + test file with zero assertions (3/10)

You know how to write good code. You don't consistently apply it.

---

## Recurring Patterns (8 Weeks)

### 1. Code Doesn't Compile Before Commit

**Frequency:** 4 of 8 weeks  
**Files Affected:** Week 6 main4.go (TLS), Week 8 Docker client, others

**Example:**

```go
server.ListenAndServeTLS()  // Missing required cert/key parameters
```

**Impact:** Wastes time, shows lack of `go build` before commit.

**Fix:** Run `go build` before every commit. No exceptions.

---

### 2. Tests Without Assertions

**Frequency:** 5 of 8 weeks  
**Files Affected:** Week 3, Week 5, Week 6, Week 7, Week 8

**Pattern:**

```go
func TestStack(t *testing.T) {
    stack.Push(1)
    stack.Pop()
    // No assertions - just calls functions
}
```

**Reality:** You know how to write assertions (Week 4: 9 assertions, Week 8 stack_test: 13 assertions). You don't always do it.

**Fix:** Every test needs at least one assertion. Delete tests that only print.

---

### 3. Error Handling Regression

**Frequency:** Appears, disappears, reappears across weeks  
**Files Affected:** Week 5-6 HTTP files

**Pattern:**

```go
if request.Method != http.MethodGet {
    http.Error(writer, "Not allowed", 405)
}
// Code continues executing - BUG
```

**Reality:** Week 3-4 had proper error handling. Week 5-6 lost it. Week 8 recovered partially.

**Fix:** `http.Error()` doesn't stop execution. Always add `return` after.

---

### 4. Spelling Errors Propagate Into Code

**Frequency:** Every single week (8/8)  
**Impact:** Type names, filenames, exported APIs

**Examples:**

- `ProrityQueue` (missing 'i') - Week 1-6
- `SingelyLinkList` (missing 'l') - Week 5-6
- `skack/` folder (missing 't') - Week 1-2
- `pacakages/` (missing 'g') - Week 1-2

**Reality:** Comments having typos is one thing. Type names affect every file that imports your code.

**Fix:** Run spell checker before commit. Rename types/folders with typos.

---

### 5. Copy-Paste Without Understanding

**Frequency:** 3 of 8 weeks  
**Evidence:** Same bug appears in 3 consecutive files (Week 6)

**Pattern:** Week 6 main1.go, main2.go, main3.go all have identical `http.Error()` bug.

**Reality:** Copied broken code without testing it.

**Fix:** If copying code, understand what it does. Run it.

---

### 6. Hardcoded Credentials/Magic Numbers

**Frequency:** 4 of 8 weeks  
**Files:** Week 5 linked list, Week 6 stack (uint16 limit), Week 8 PostgreSQL

**Examples:**

- PostgreSQL password in source code
- Stack capped at 65,535 elements (uint16) for no reason

**Fix:** Use environment variables. Remove arbitrary limits.

---

### 7. Inconsistent Exploration vs Implementation

**Best Weeks (Deep Exploration):**

- Week 4: forSelect.go (313 lines, systematic debugging, 10/10)
- Week 8: defer.go (101 lines, hypothesis testing, 7.5/10)

**Worst Weeks (Surface-level, Rushed):**

- Week 6: Copy-pasted HTTP handlers without understanding
- Week 7: Explored game dev + JSON but ignored all Week 6 feedback

**Pattern:** Quality improves when exploring deeply. Drops when chasing new topics.

**Reality:** You learn best by diving deep, not breadth-first exploration.

---

## What Actually Works (Proven)

These patterns produced your best work:

### 1. Hypothesis-Driven Testing (forSelect.go Week 4)

You created test case "weee" to expose the bug:

```go
// Discovered: `for i, value := range <-ch` iterates over string, not channel
// Designed "weee" (4 chars) + slice size 4 to prove it
```

This is professional debugging methodology. Do this more.

### 2. Reading Standard Library (Week 5 Eg2)

You read `/net/http/server.go` to understand how `ListenAndServe` works internally.

Result: 8/10 file with deep understanding.

Do this more. Read stdlib when learning new packages.

### 3. Thread-Safe Patterns (Week 8 stack.go)

```go
func (s *Stack) Push(value any) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    // ...
}
```

Perfect sync.Mutex usage, production-ready code (8/10).

You know concurrency. Apply it everywhere data is shared.

### 4. Test Assertions (Week 4, Week 8 stack_test)

Week 4: 9 assertions in queue tests  
Week 8: 13 assertions in stack tests

When you write assertions, quality jumps to 7-10/10.  
When you skip assertions, quality drops to 3-6/10.

Pattern is clear. Write assertions.

---

## Priority Actions (Based on 8-Week Data)

### Critical (Prevent Regressions)

1. **Run `go build` before every commit**  
   Prevents non-compiling code. No excuses.

2. **Every test needs assertions**  
   Delete tests that only print. You know how to write assertions (proven Week 4, 8).

3. **Run spell checker before commit**  
   Typos in comments are fine. Typos in type names affect everyone who uses your code.

4. **Add `return` after `http.Error()`**  
   You know this (Week 3-4). Don't regress (Week 5-6).

### High Priority (Consistency)

1. **Fix existing non-compiling code**  
   Week 6 main4.go, Week 8 Docker client. Delete or fix.

2. **Remove hardcoded credentials**  
   Week 8 PostgreSQL password in source code is security risk.

3. **Rename types with typos**  
   `ProrityQueue` → `PriorityQueue`  
   `SingelyLinkList` → `SinglyLinkedList`

4. **Fix filename typos**  
   `movments.go` → `movements.go`  
   `unMarsal.go` → `unmarshal.go`

### Medium Priority (Code Quality)

1. **Complete or delete partial implementations**  
   Week 5: InsertAfter(), InsertBefore(), InsertAt() are stubs.  
   Either implement or delete.

2. **Add edge case tests**  
    Your tests cover happy paths. Add: empty input, nil pointers, overflow.

3. **Replace PrintList() tests with assertions**  
    Week 5 linked list tests. Convert to actual test logic.

### Low Priority (Learning)

1. **Study these gaps** (identified across reviews):
    - Context usage for cancellation
    - Middleware patterns
    - Table-driven tests (you partially use them)
    - `go test -cover` for coverage analysis

---

## What Not To Do

Based on 8-week regression patterns:

❌ **Don't explore new topics while ignoring existing issues**  
Week 7: Started game dev + JSON. Ignored 8 of 8 Week 6 issues.

❌ **Don't copy-paste without testing**  
Week 6: Same bug in 3 files proves code wasn't run.

❌ **Don't commit without `go build`**  
4 weeks had non-compiling code. This is preventable.

❌ **Don't write tests without assertions**  
You've done this 5 times. You know better (proven Week 4, 8).

❌ **Don't use testing framework for non-testing tasks**  
Week 6 main_test.go: Using test for 600s auto-kill timer.

---

## The Gap

**Can write:** Production-ready, thread-safe code with proper patterns (Week 4: 10/10, Week 8: 8/10)  
**Actually writes:** Mix of excellent (8/10) and broken (2/10) in same week

**The issue isn't knowledge. It's discipline.**

Fix: Apply Week 4 methodology to everything. Don't commit without running code.
