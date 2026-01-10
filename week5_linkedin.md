Week 5 of my 8-week Go challenge (Jan 4-10, 2026).

⚠️ REGRESSION WEEK

WHAT I BUILT

HTTP servers and my first linked list from scratch:
• HTTP servers (net/http, ServeMux, JSON APIs, templating)
• Singly linked list (InsertAtBeginning, InsertAtLast, PrintList)
• 10 files total (7 HTTP + 3 linked list)

Best work: try3_POST/main.go (8.5/10). Production-quality JSON API with error handling, RWMutex for thread-safe maps, and validation. Shows I CAN write good code.

But only 1 of 7 HTTP files has error handling. The other 6 ignore errors completely.

THE PROBLEM

Week 4: All files had error handling (9.0/10)
Week 5: Only 1 of 7 HTTP files has error handling (6.8/10)

Week 4: Queue tests had comprehensive assertions
Week 5: Linked list tests have ZERO assertions - just print output

This isn't a knowledge gap. This is speed over quality. I explored new topics but didn't maintain discipline from Week 4.

SELF-REVIEW

AI reviews (Claude Sonnet 4.5):

Week 5: 6.8/10 ⚠️
Progression: 7.0 → 8.0 → 7.7 → 9.0 → 6.8

Dropped 2.2 points. Biggest regression in 5 weeks.

Strengths: Read stdlib source (/net/http/server.go), production JSON API, RWMutex correct, first linked list from scratch

Issues: Error handling lost (6 of 7 files), test assertions lost, InsertAtLast bug, interface design flaw

The review: "You went backwards. Week 4 showed you know how to write proper tests. Why didn't you apply that knowledge here?"

WHAT I'M FIXING

Week 6 priorities (before learning anything new):

1. Add error handling to all 6 HTTP files
2. Fix linked list bugs
3. Add test assertions to all tests
4. Fix test function names (lowercase 't' → uppercase 'T')

No new topics until I restore Week 4 standards.

CODE & REVIEWS

GitHub: https://github.com/aakku106/Go
Week 5: /0.0015_HTTP_Starts_Here/, /datastructures/list/
Reviews: /review/week5/

Week 6 answer: discipline over speed.

Learning that regression is part of the journey. The key is recognizing it and fixing it.

#golang #programming #http #datastructures #learning #LearningInPublic
