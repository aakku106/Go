# Week 7 Code Review Summary (January 18-24, 2026)

**Review Period**: January 18-24, 2026  
**Overall Rating**: 5.6/10  
**Status**: Mixed progress - new topics explored but old issues ignored

---

## Overview

Week 7 splits into two distinct tracks:

**Main Repository (5 files, 5.2/10 average)**:

- 0.0016: Ebiten 2D game with WASD movement (3.5/10 avg)
- 0.0017_PlayingWith_JSON: JSON marshaling/unmarshaling education (7.0/10 avg with tests)

**Datastructures Repository (4 files, 6.1/10 average)**:

- Implemented InsertAt method (stubbed in Week 6)
- Fixed critical test discovery bug
- Added new test file (poorly)
- Expanded documentation (+93 lines)
- **Ignored 7 of 8 Week 6 issues**

Main repo continues pattern of exploring new topics (game dev, JSON) instead of fixing fundamentals. Datastructures shows selective improvement: fixed one critical issue but ignored all other Week 6 feedback.

---

## File Breakdown

### Main Repository

#### 0.0016 - Ebiten Game (3.5/10 average)

**1. main.go** - 3/10

- **Purpose**: Ebiten 2D game with WASD player movement
- **Critical Issues**:
  - Global variable `aakku` (mutable global state)
  - Claims "Raylib" but uses Ebiten (wrong library name)
  - Unused Enemy struct (dead code)
  - Zero error handling
- **What Works**: Game runs, basic Ebiten integration
- **Verdict**: Functional but poorly structured, wrong library claim

**2. movments.go** - 4/10

- **Purpose**: Movement methods for Position type
- **Critical Issues**:
  - Filename typo: "movments" missing 'e'
  - No boundary checking (infinite movement off-screen)
  - Magic number 2 not explained
- **What Works**: Methods work syntactically
- **Verdict**: Basic functionality, missing essential features

---

#### 0.0017_PlayingWith_JSON (7.0/10 average) ⭐ BEST MAIN REPO FOLDER

**3. main.go** - 7/10 ⭐ BEST MAIN REPO FILE

- **Purpose**: Educational JSON marshaling examples (call1-call5)
- **Critical Issues**: Functions never called (no entry point)
- **Major Issues**: Inconsistent struct naming, variable reassignments
- **What Works**:
  - Progressive pedagogical structure
  - Security awareness (password handling)
  - Accurate comments
  - Proper error handling
- **Verdict**: Best educational code in main repo, but no execution

**4. unMarsal.go** - 5.5/10

- **Purpose**: JSON unmarshaling from jsonplaceholder API
- **Critical Issues**:
  - Filename typo: "unMarsal" missing 'h'
  - Struct doesn't match API (expects age/country/province but API returns username/phone/company)
  - Functions never called
- **Major Issues**: Poor function names (um, um2)
- **Verdict**: Educational intent clear, execution flawed

**5. main_test.go** - 8.5/10 ⭐ BEST WEEK 7 FILE

- **Purpose**: Comprehensive JSON marshal/unmarshal tests
- **Major Issues**:
  - String comparison of JSON (brittle)
  - Test uses idealized JSON instead of real API structure
- **What Works**:
  - 9 test functions, table-driven tests
  - Edge case coverage
  - Round-trip validation
  - Proper assertions with t.Error/t.Fatal
  - 265 lines comprehensive testing
- **Verdict**: Best testing in Week 7, most complete file

---

### Datastructures Repository

**6. list/SingelyLinkList.go** - 6.5/10 (MODIFIED)

- **Purpose**: Singly linked list implementation
- **Week 7 Changes**:
  - ✅ InsertAt implemented (was stubbed in Week 6)
  - ✅ Added bounds checking to InsertAt
- **Week 6 Issues NOT Fixed** (7 issues ignored):
  - ❌ Filename typo: "SingelyLinkList" (still missing 'l')
  - ❌ PrintList value receiver (should be pointer)
  - ❌ Labeled break not removed
  - ❌ Error grammar: "to index which greater" (still wrong)
  - ❌ Debug typo: "Insearting" (still present)
  - ❌ Inconsistent returns (error vs nil)
  - ❌ No comprehensive tests for InsertAt
- **Verdict**: Fixed 1 of 8 Week 6 issues, ignored all other feedback

**7. list/SinglyLinkedListtest.go** - 7/10 (RENAMED)

- **Purpose**: Linked list tests
- **Week 7 Changes**:
  - ✅ Fixed test discovery bug (testNewSinglyLinkedList → TestNewSinglyLinkedList)
  - ✅ Fixed undefined variable (T → t)
  - ✅ Removed type assertions
  - ❌ Broke naming convention (removed underscore: SinglyLinkedList_test.go → SinglyLinkedListtest.go)
- **Major Issues**: Filename breaks Go convention (should be \*\_test.go)
- **Verdict**: Fixed critical bugs but created new convention violation

**8. list/SingallyLinkedListtest2_test.go** - 3/10 💀 WORST DATASTRUCTURES FILE (NEW)

- **Purpose**: Tests for InsertAt method
- **Critical Issues**:
  - Double filename typo: "Singally" + "test2" (no underscore)
  - Zero assertions (not a real test)
  - Just prints output (manual inspection)
- **Major Issues**:
  - Variable named 'cat'
  - Inconsistent error handling
  - Should be in SinglyLinkedListtest.go, not separate file
- **Verdict**: Not a test - it's debug print code

**9. doc/SingallyLinkedList.md** - 8/10 ⭐ BEST DATASTRUCTURES FILE (MODIFIED)

- **Purpose**: Technical documentation explaining concrete types vs interfaces
- **Week 7 Changes**:
  - Completely rewritten (210 → 303 lines, +44%)
  - Added trade-off analysis
  - Added comparison table
  - Added use case guidance
  - Removed Week 6 interface internals content
- **Critical Issues**:
  - Filename typo: "Singally" (still present)
  - InsertAt example comment wrong (says "replaces" but code inserts)
- **Major Issues**: Lost Week 6 interface internals explanation
- **Verdict**: Excellent documentation, addresses why concrete types used

---

## Rating Distribution

| Rating | Files | Percentage                      |
| ------ | ----- | ------------------------------- |
| 8-10   | 2     | 22% (main_test.go, doc.md)      |
| 6-7.9  | 3     | 33% (main.go JSON, tests, impl) |
| 4-5.9  | 2     | 22% (movments, unMarsal)        |
| 0-3.9  | 2     | 22% (main.go game, test2)       |

**Average**: 5.6/10

**Distribution**:

- Main repo: 5.2/10 (range: 3-8.5)
- Datastructures: 6.1/10 (range: 3-8)

---

## Best and Worst

### ⭐ Best Files

**1. main_test.go (8.5/10)** - BEST OVERALL

- Comprehensive testing with 9 test functions
- Table-driven tests, edge cases, round-trip validation
- Proper assertions throughout
- Best testing in entire Week 7

**2. doc/SingallyLinkedList.md (8/10)** - BEST DOCUMENTATION

- Complete rewrite explaining concrete vs interface trade-offs
- Comparison table, use case guidance
- 303 lines of quality technical writing
- Addresses real codebase behavior

**3. 0.0017/main.go (7/10)** - BEST MAIN REPO CODE

- Progressive pedagogical structure (call1-call5)
- Security awareness, proper error handling
- Clear educational intent

### 💀 Worst Files

**1. 0.0016/main.go (3/10)** - WORST OVERALL

- Global variable `aakku`
- Wrong library claim (says Raylib, uses Ebiten)
- Unused Enemy struct
- No error handling, no boundary checking

**2. list/SingallyLinkedListtest2_test.go (3/10)** - WORST TEST

- Double filename typo
- Zero assertions (not a real test)
- Just prints output for manual inspection
- Variable named 'cat'

**3. 0.0016/movments.go (4/10)** - WORST IMPLEMENTATION

- Filename typo
- No boundary checking (infinite off-screen movement)
- Magic numbers, no return values

---

## Statistics

### Code Volume

- **Main repo**: 5 files, ~500 lines total
  - 0.0016: 2 files, ~120 lines
  - 0.0017: 3 files + 1 test, ~380 lines
- **Datastructures**: 4 files
  - Implementation: 156 lines (unchanged)
  - Test files: 2 files, ~90 lines total
  - Documentation: 303 lines (+93 from Week 6)

### Commits

- **Main repo**: 18+ commits
- **Datastructures**: 4 commits (105c629, a02b4a5, 6231fd3, 75250b4)

### Testing

- **Main repo**: 1 test file (main_test.go, 9 tests, 8.5/10) ⭐
- **Datastructures**: 2 test files
  - SinglyLinkedListtest.go: 7/10 (real tests, fixed bugs)
  - SingallyLinkedListtest2_test.go: 3/10 (no assertions)

### Issues Fixed from Week 6

- ✅ InsertAt implemented (was stubbed)
- ✅ Test discovery bug fixed (capitalization)
- ✅ Undefined variable T fixed
- ❌ Filename typos unchanged (7 instances)
- ❌ PrintList value receiver not fixed
- ❌ Labeled break not removed
- ❌ Error grammar not fixed

**Fix Rate**: 3 of 10 issues (30%)

---

## Week 7 Patterns

### What Improved

1. **Testing Quality** - main_test.go shows comprehensive test writing (8.5/10)
2. **Documentation** - doc.md complete rewrite with trade-off analysis (8/10)
3. **Critical Fixes** - InsertAt implementation, test discovery bug fixed
4. **New Topics** - Game development (Ebiten), JSON marshaling/unmarshaling

### What Regressed

1. **Ignored Feedback** - 7 of 8 Week 6 issues not addressed
2. **New Typos** - Created new typo in test2 file ("Singally")
3. **Non-Tests** - test2 has zero assertions
4. **Convention Breaking** - Removed underscore from test filename

### Persistent Issues

1. **Filename Typos** - Week 6 typos still present: movments, unMarsal, Singelly/Singally (8 instances total)
2. **No Entry Points** - 0.0017 functions never called
3. **Struct Mismatches** - unMarsal struct doesn't match API
4. **Global Variables** - aakku global in game
5. **Missing Boundaries** - No movement limits in game

---

## Topic Exploration

### New Topics Week 7

1. **Game Development** - Ebiten 2D library, game loop, input handling
2. **JSON Operations** - Marshal/unmarshal, struct tags, error handling
3. **API Integration** - jsonplaceholder.typicode.com
4. **Trade-off Analysis** - Concrete types vs interfaces (documentation)

### Repeated Topics

1. Linked lists (continuing from Week 6)
2. Testing (improving from Week 6)
3. Documentation (expanding from Week 6)

---

## Comparison to Week 6

| Metric              | Week 6 | Week 7 | Change  |
| ------------------- | ------ | ------ | ------- |
| Overall Rating      | 4.9/10 | 5.6/10 | +0.7 ⬆️ |
| Main Repo Avg       | N/A    | 5.2/10 | New     |
| Datastructures Avg  | 4.9/10 | 6.1/10 | +1.2 ⬆️ |
| Files Reviewed      | 4      | 9      | +5      |
| Best File Rating    | 6.5/10 | 8.5/10 | +2.0 ⬆️ |
| Worst File Rating   | 3/10   | 3/10   | Same    |
| Week 6 Issues Fixed | N/A    | 3/10   | 30%     |
| Test Quality (best) | 5/10   | 8.5/10 | +3.5 ⬆️ |
| Documentation Lines | 210    | 303    | +93 ⬆️  |

**Improvement**: +0.7 overall, +1.2 datastructures, +3.5 best test quality  
**Stagnation**: Worst file still 3/10, 70% of Week 6 issues ignored

---

## Final Verdict

**5.6/10** - Slight improvement from Week 6 (4.9/10) but uneven progress.

**What Works**:

- Testing dramatically improved (main_test.go 8.5/10 vs Week 6 best 5/10)
- Documentation quality high (doc.md 8/10)
- Fixed critical InsertAt stub and test discovery bug
- Explored new topics (game dev, JSON) successfully

**What Doesn't Work**:

- Ignored 70% of Week 6 feedback
- Created new issues while fixing old (test2 typos, convention breaking)
- Main repo still has no execution (0.0017 functions never called)
- Persistent filename typos across 8 instances
- Worst files unchanged (game 3/10, test2 3/10)

**Main Problems**:

1. Selective improvement (fixes one issue, ignores rest)
2. Focus on new topics instead of fixing fundamentals
3. Typos persist across weeks
4. Breaking conventions while fixing bugs
5. Testing improved but still incomplete

**Recommendation**: Week 8 should focus on:

1. Fix all persisting filename typos (8 instances)
2. Address remaining Week 6 issues (7 items)
3. Add entry points to 0.0017 code
4. Fix test2 to have real assertions
5. Add boundary checking to game movement
6. Restore test filename convention
7. Fix struct to match real API in unMarsal.go

**Progress Assessment**: Moving forward but leaving too many issues behind. Quality of best files (8-8.5/10) shows capability. Quality of worst files (3/10) shows lack of fundamentals. Need to consolidate improvements instead of constantly exploring new topics.
