# Week 6 Progress: HTTP Request Deep Dive

Week 6 (Jan 11-17): Strong recovery from Week 5's regression.

**Rating**: 8.42/10 (⬆️ +1.62 from Week 5's 6.8)

## Key Achievement

Week 5's critical error handling issue **completely fixed**. All 5 new files have proper error handling:

```go
if err := server.ListenAndServe(); err != nil {
    log.Println("Error: ", err)
}
```

## What I Explored

Systematically explored http.Request internals:

**Headers** (main1.go - 8.5/10):

- Discovered Header is `map[string][]string`
- Compared browser vs curl headers
- Learned header values are arrays

**Body** (main2.go - 9/10):

- Discovered request.Body is io.ReadCloser, not string
- Used io.ReadAll to read properly
- Documented dangerous pattern (missing return after error)

**URL** (main3.go - 9.5/10):

- Discovered URL.Query() returns url.Values (`map[string][]string`)
- Learned RawQuery vs Query() difference
- Explored Path, Scheme, Host, RawQuery fields
- **Looked up type definitions before using** (professional approach)

**TLS** (main4.go - 8/10):

- Used request.TLS to detect HTTPS
- Explained why TLS is nil in development
- Showed ListenAndServeTLS for HTTPS

**Auto-Shutdown** (main_test.go - 7.5/10):

- Created battery-saving test with 10-min timeout
- Creative problem-solving, wrong tool

## Learning Methodology

Developed professional API exploration pattern:

1. Print value
2. Discover type
3. Read type definition
4. Explore all fields
5. Document findings

**This is how professionals learn new technologies.**

## Progress vs Week 5

✅ Error handling: 5/10 → 8/10 (major improvement)
✅ HTTP development: 7/10 → 8.5/10 (deep exploration)
✅ Conceptual understanding: Maintained 8.5/10

⚠️ try4 issues from Week 5 still not fixed
⚠️ Spelling errors across all files
⚠️ Inconsistent exploration depth (main4 shallow)

## What's Next

**Week 7 Critical**:

- Fix try4 template error handling
- Run spell-check on all files
- Explore HTTP client side (making requests)

**Learning Goals**:

- Context package (timeouts, cancellation)
- Proper test patterns (assertions, table-driven)
- Middleware patterns

## Week 6 in Numbers

- **Files**: 6 total (1 modified, 5 new)
- **Lines**: ~450 lines
- **Commits**: 32 in main repo
- **Average Rating**: 8.42/10
- **Best File**: main3.go (9.5/10) - URL exploration
- **Improvement**: +1.62 points from Week 5

**Trajectory**: Back on upward trend. Week 5 was anomaly.

---

_"You're building a mental model of how HTTP works internally. This is deep learning, not surface-level copying."_

#golang #learning #http #webdevelopment #systematiclearning
