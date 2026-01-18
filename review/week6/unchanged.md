# Week 6 Unchanged Files

**Review Period**: January 11-17, 2026

---

## Files Not Modified in Week 6

The following files in `0.0015_HTTP_Starts_Here` were **not changed** during Week 6 (January 11-17, 2026):

### Unchanged Folders

**try1/**:

- Purpose: First HTTP server attempt
- Status: Not modified
- Last Week: Week 5 (created and reviewed)

**try2_mux/**:

- Purpose: HTTP multiplexer (mux) exploration
- Status: Not modified
- Last Week: Week 5 (created and reviewed)

**try3_POST/**:

- Purpose: POST request handling
- Status: Not modified
- Last Week: Week 5 (created and reviewed)

---

## Files Modified in Week 6

**try4/**:

- Purpose: Template rendering
- Status: **Modified** in Week 6 (but Week 5 issues not fixed)
- File: `try4/main.go`
- Review: See [0.0015-try4-main.md](0.0015-try4-main.md)

**try5/** (NEW in Week 6):

- Purpose: HTTP request deep dive
- Status: **Created** in Week 6
- Files: `main1.go`, `main2.go`, `main3.go`, `main4.go`, `main_test.go`
- Reviews: See individual review files

---

## Why These Files Were Unchanged

Week 6 focused on:

1. Fixing Week 5's critical error handling regression
2. Deep exploration of `http.Request` structure (try5/ folder)
3. Learning request internals (headers, body, URL, TLS)

try1, try2_mux, and try3_POST were earlier experiments from Week 5 that were **complete for their learning purpose**. No need to modify them.

---

## Note

Week 6 work was **entirely in the main Go repository**. The datastructures repository had **0 commits** during Week 6.

Git confirmed only these folders changed:

- `try4/` (modified)
- `try5/` (created)

All other folders in `0.0015_HTTP_Starts_Here` remained unchanged.
