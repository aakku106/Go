Week 6 of my 8-week Go challenge (Jan 11-17, 2026).

⚠️ CONTINUED REGRESSION

WHAT I BUILT

Two repositories this week:

Main repo - HTTP request internals:
• Headers, body reading, URL parsing, TLS attempt (failed)
• 5 files exploring net/http Request properties

Datastructures repo (NEW):
• Singly linked list, linear queue, priority queue, stack
• 12 files with Go interfaces
• Best work: SingallyLinkedList.md (7.5/10) - excellent interface documentation

THE REALITY CHECK

Main Repository: 3.8/10
Datastructures Repository: 5.6/10
Combined Week 6: 4.9/10

Week progression: 7.0 → 8.0 → 7.7 → 9.0 → 6.8 → **4.9**

Three weeks of decline. From 9.0/10 to 4.9/10.

WHAT WENT WRONG

Copy-pasted the SAME broken error handling in 3 files:

```go
// WRONG
if request.Method != http.MethodGet {
    http.Error(writer, "Method Not Allowed", 405)
}
// Code continues - BUG
```

Missing `return` after `http.Error()`. Three identical bugs.

Other failures:
• TLS code doesn't compile (0 params, needs 2)
• Test file abuses framework (1/10)
• "Initilize" typo in ALL 5 files despite Week 5 feedback
• Datastructures: `SingelyLinkList` typo 8+ times, `ProrityQueue` 7+ times
• stack.go has 12-line comment admitting code is wrong but keeping it

SELF-REVIEW

AI reviews (Claude Sonnet 4.5):

Week 6: 4.9/10 ⚠️

The review: "Week 4 proved you can write 9/10 code. Week 6 shows you're not applying ANY lessons learned."

WHAT THIS MEANS

Main repo: Speed without discipline. 3.8/10.
Datastructures repo: Learning visible but sloppy. 5.6/10.

I CAN learn new concepts (interfaces). I WON'T apply basic discipline (error handling).

WEEK 7 PLAN

1. STOP writing new code
2. Fix all broken error handling (Week 5 + Week 6)
3. Fix all typos
4. Delete self-aware bad code comment

No new exploration until I restore Week 4 standards (9.0/10).

CODE & REVIEWS

GitHub: https://github.com/aakku106/Go
Week 6: /0.0015_HTTP_Starts_Here/try5/, /datastructures/
Reviews: /review/week6/

Three weeks of regression. Week 7 is about breaking that pattern.

#golang #programming #http #datastructures #learning #LearningInPublic #CodeQuality
