# Unchanged Files - Week 6

**Review Period**: January 11-17, 2026

---

## Scope

This document covers **two repositories**:

1. **Main Repository** (`/Users/aakku/Documents/Go`)
2. **Datastructures Repository** (`/Users/aakku/Documents/Go/datastructures`)

---

## Main Repository

### Modified But Functionally Unchanged

#### 0.0015_HTTP_Starts_Here/try4/main.go

**Status**: Modified  
**Change**: +2 lines (comment only)  
**Functional Impact**: None  
**Week 5 Rating**: 7/10

**Diff**:

```diff
+
+// NEXT: ../try5/main.go
```

**Analysis**: Only added navigation comment pointing to try5. No code logic, functionality, or behavior changed. This is a documentation change, not a code change.

**Verdict**: Treated as **UNCHANGED** for Week 6 review. See `review/week5/0.0015-try4-main.md` for full code review.

### Truly Unchanged Files

All other files in main repository (0.0001-0.0014 folders, prior try1-try3 files, etc.) had **zero commits** during January 11-17, 2026.

**Verification**:

```bash
git diff --name-status b1e6aa1..193be90  # Main repo
```

**Result**: Only try4/main.go and try5/\* files appear in Week 6 commits.

---

## Datastructures Repository

### All New Files

**Status**: Entire repository created in Week 6  
**Commit**: 11b4360 "dubug on" (Jan 17, 2026)  
**Files**: 12 (all NEW)

**Verification**:

```bash
cd datastructures
git log --since="2026-01-11" --until="2026-01-18" --oneline
# 11b4360 dubug on

git show --stat 11b4360
# 12 files changed, 914 insertions(+)
```

**All files are NEW**:

- doc/SingallyLinkedList.md (210 lines)
- list/linkList.go (37 lines)
- list/SingelyLinkList.go (93 lines)
- list/SinglyLinkedList_test.go (59 lines)
- queue/queue.go (33 lines)
- queue/linearQueue.go (89 lines)
- queue/linearQueue_test.go (56 lines)
- queue/prorityQueue.go (104 lines)
- queue/prorityQueue_test.go (150 lines)
- stack/stack.go (39 lines)
- stack/stack_test.go (41 lines)
- go.mod (3 lines)

**No files were MODIFIED** - everything is brand new.

---

## Summary

### Main Repository

- **Modified files**: 1 (try4/main.go - comment only)
- **New files**: 5 (try5/main1.go, main2.go, main3.go, main4.go, main_test.go)
- **Functionally unchanged**: All files from 0.0001-0.0014

### Datastructures Repository

- **Modified files**: 0
- **New files**: 12 (entire repository)
- **Functionally unchanged**: N/A (new repository)

### Week 6 Total

- **Modified files**: 1 (comment only)
- **New files**: 17 (5 main + 12 datastructures)
- **Functionally unchanged**: All prior main repo files

**Note**: Git analysis confirms try4 existed in Week 5 (commit b1e6aa1, Jan 10) and was only modified with a comment in Week 6. Datastructures repository is entirely new.
