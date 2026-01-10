Week 5 of my 8-week Go challenge (Jan 4-10, 2026).

⚠️ REGRESSION WEEK

WHAT I BUILT

This week I explored HTTP servers and built my first linked list from scratch:
• HTTP server basics - net/http, ServeMux, Server struct
• JSON APIs - encoding/decoding, validation, RWMutex
• HTML templating - text/template package
• Singly linked list - InsertAtBeginning, InsertAtLast, PrintList
• 10 files total (7 HTTP + 3 linked list)

Best work: try3_POST/main.go (8.5/10). Production-quality JSON API with proper error handling, RWMutex for thread-safe map access, validation, and correct HTTP status codes. Shows I CAN write good code when focused.

But here's the problem: only 1 of 7 HTTP files has error handling. The other 6 ignore errors completely.

KEY DISCOVERIES

HTTP patterns learned:
• Reading stdlib source code - Found how `ListenAndServe()` wraps `Server` struct by reading `/net/http/server.go`
• Go 1.22+ routing - `"POST /cat"`, `"GET /cat/{id}"` with method constraints
• RWMutex for concurrent maps - Correct read/write lock usage
• JSON validation - Proper error responses with status codes

Linked list mechanics:
• InsertAtBeginning - Textbook correct O(1) implementation
• InsertAtLast - Correct algorithm, but returns WRONG node (bug!)
• Pointer manipulation - Understanding head pointer management

THE PROBLEM

I regressed on patterns I already knew:

Week 4: All files had error handling (9.0/10)
Week 5: Only 1 of 7 HTTP files has error handling (6.8/10)

Week 4: Queue tests had comprehensive assertions
Week 5: Linked list tests have ZERO assertions - just print output

This isn't a knowledge gap. This is speed over quality. I explored new topics (HTTP, linked lists) but didn't maintain the discipline from Week 4.

SELF-REVIEW RESULTS

AI code reviews (Claude Sonnet 4.5):

Week 5 Rating: 6.8/10 (C+) ⚠️
Progression: 7.0 → 8.0 → 7.7 → 9.0 → 6.8

Dropped 2.2 points from Week 4. Biggest regression in 5 weeks.

Strengths: Reading stdlib source (professional practice), production-quality JSON API (try3_POST), RWMutex correct, first linked list from scratch

Critical Issues:
• Error handling lost (6 of 7 HTTP files ignore errors)
• Test assertions lost (linked list tests don't verify anything)
• InsertAtLast bug (returns last node instead of head)
• Interface design flaw (signatures don't match implementation)

The review called this "inconsistent quality" - I can write production code (try3_POST) but don't apply the same discipline everywhere.

PROGRESS TRACKING

Week 1: Data structures (7.0/10)
Week 2: OOP concepts (8.0/10)
Week 3: Concurrency basics (7.7/10)
Week 4: Patterns & Testing (9.0/10) ⭐ PEAK
Week 5: HTTP & Linked Lists (6.8/10) ⚠️ REGRESSION

Week 4 was my best week. Week 5 shows what happens when you prioritize exploration over discipline.

WHAT I'M FIXING

Week 6 priorities (before learning anything new):

1. Add error handling to all 6 HTTP files
2. Fix linked list bugs (InsertAtLast return value)
3. Add test assertions to all linked list tests
4. Fix test function names (lowercase 't' → uppercase 'T')

No new topics until I restore Week 4 quality standards.

CODE & REVIEWS

GitHub:
https://github.com/aakku106/Go
Week 5 Code: /0.0015_HTTP_Starts_Here/, /datastructures/list/
AI Reviews: /review/week5/

The review is brutally honest: "You went backwards. Week 4 showed you know how to write proper tests. Why didn't you apply that knowledge here?"

Good question. Week 6 answer: discipline over speed.

Learning that regression is part of the journey. The key is recognizing it and fixing it.

#golang #programming #http #datastructures #softwaredevelopment #learning #CodingChallenge #LearningInPublic #Regression #QualityOverSpeed
