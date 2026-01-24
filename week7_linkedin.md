Week 7 of my 8-week Go challenge (Jan 18-24, 2026).

IMPROVEMENT BUT SELECTIVE

WHAT I BUILT

Main repo: Ebiten 2D game, JSON marshaling/unmarshaling with tests (5 files)
Datastructures: InsertAt implementation, documentation expanded 210→303 lines (4 files)

Best work: main_test.go (8.5/10) - 9 test functions with table-driven tests and round-trip validation. Best testing in any week.

THE REALITY CHECK

Week 7: 5.6/10 (Main: 5.2/10, Datastructures: 6.1/10)
Progression: 7.0 → 8.0 → 7.7 → 9.0 → 6.8 → 4.9 → 5.6

First improvement after three weeks of decline.

THE PROBLEM

Selective improvement: Fixed 3 of 10 Week 6 issues (30% fix rate).

Fixed:
• InsertAt implementation
• Test discovery bug
• Undefined variable

Ignored (7 issues):
• Filename typos: movments, unMarsal, SingelyLinkList
• PrintList value receiver
• Error grammar
• Debug typo

Created 7 new issues:
• Test file with zero assertions
• Broke naming convention
• No entry points (functions never called)
• Global variable in game

Net result: Fixed 3, created 7. -4 issues.

SELF-REVIEW

AI review (Claude Sonnet 4.5): Week 7: 5.6/10

"Quality ceiling raised (+2 best file: 8.5/10) but floor unchanged (worst: 3/10). Fixes functional issues, ignores quality issues. Creates as many problems as fixes."

Pattern: Can write 8.5/10 code (tests, docs) but doesn't apply same care everywhere. Gap between best (8.5/10) and worst (3/10) is 5.5 points. Not a knowledge problem - attention to detail problem.

WEEK 8 PLAN

Consolidation week. No new topics until issues fixed:

1. Fix all 8 filename typos
2. Add entry points (functions never called)
3. Fix test2 or delete it
4. Address remaining Week 6 issues

Goal: Close the gap between best and worst work.

CODE & REVIEWS

GitHub: https://github.com/aakku106/Go
Week 7: /0.0016/, /0.0017_PlayingWith_JSON/, /datastructures/
Reviews: /review/week7/

#golang #programming #gamedev #json #datastructures #learning #LearningInPublic #CodeQuality
