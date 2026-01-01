Week 3 of my 8-week Go challenge (Dec 21-27, 2025).

WHAT I BUILT

This week I explored advanced Go concepts:
• Error handling - comma-ok pattern, type assertions, error wrapping
• HTTP client - GET requests with resource management
• Concurrency - goroutines, buffered/unbuffered channels, blocking
• Data structures - stack and queue with tests
• 13 Go files total

Best work: channels.go (323 lines). Systematically explored channel behavior through experimentation - started with simple unbuffered channels, discovered deadlocks, fixed them, then tested 9 different buffered channel patterns (B1-B9) to understand FIFO ordering, capacity limits, and blocking.

Second highlight: Queue with O(1) operations, aggressive memory optimization, and self-healing behavior. Auto-reclaims wasted space when empty, reuses capacity instead of always appending. Most tutorials teach naive O(n) approach - this shows understanding of pointer arithmetic and memory management. See how isEmpty() triggers self-healing in the code.

KEY DISCOVERIES

Found important patterns through trial and error:
• Unbuffered channels block until goroutine receives - discovered why by encountering deadlock
• Buffered channels don't block until capacity is reached
• Channels maintain FIFO ordering (tested and confirmed)
• Error wrapping with fmt.Errorf("%w", err) for proper error chains
• defer pattern for resource cleanup (file.Close(), resp.Body.Close())

All discovered by writing code that breaks, understanding why, fixing it, and documenting the behavior.

SELF-REVIEW RESULTS

AI code reviews (Claude Sonnet 4.5):

Week 3 Rating: 7.7/10 (B+)
Progression: 7.0 → 8.0 → 7.7

Strengths: Systematic learning methodology (channels.go shows professional approach), advanced queue optimization, mastery of error handling patterns, deep understanding of concurrency primitives, excellent exploratory comments

Gaps: Tests have NO assertions (critical issue - tests always pass!), some files incomplete (select.go, list.go), need to learn WaitGroup for synchronization

The AI called channels.go "professional-level learning" - many senior developers don't explore channels this thoroughly.

PROGRESS TRACKING

Week 1: Data structures (7.0/10)
Week 2: OOP concepts (8.0/10)
Week 3: Concurrency + HTTP (7.7/10)

Posting weekly. Learning publicly.

CODE & REVIEWS

GitHub:
https://github.com/aakku106/Go
Week 3 Code: /0.0009/, /0.0010/, /0.0011/, /0.0012/, /datastructures/
AI Reviews: /review/week3/
for Queue: https://github.com/aakku106/datastructures/blob/main/queue/linearQueue.go

channels.go review explains why the systematic testing approach (T1-T3, B1-B9) demonstrates professional learning methodology.

Learning by breaking things and understanding why. The best code comes from mistakes fixed properly.

#golang #programming #coding #concurrency #softwaredevelopment #learning #CodingChallenge #MemoryOptimization #Queue #SelfHealingCode
