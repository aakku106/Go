# Week 5 Fixes

## Critical

- [ ] Add error handling to try1/main.go
- [ ] Add error handling to try2_mux/main.go
- [ ] Add error handling to try2_mux/Eg2/main.go
- [ ] Add error handling to try2_mux/Eg2/eg2.2/main.go
- [ ] Add error handling to try4/main.go
- [ ] Fix test name: testCreate → TestCreate
- [ ] Fix test name: testInsertAtBeginning → TestInsertAtBeginning
- [ ] Remove PrintList() from all tests
- [ ] Add assertions to TestCreate
- [ ] Add assertions to TestInsertAtBeginning
- [ ] Add assertions to TestInsertAtLast
- [ ] Fix InsertAtLast: return l (not return head)
- [ ] Fix try4: parse templates once at startup, not per request
- [ ] Fix try3_POST: use ID counter instead of len()

## Major

- [ ] Complete or delete: InsertAfter(), InsertBefore(), InsertAt()
- [ ] Fix LinkList interface: add return types
- [ ] Remove Create() from LinkList interface
- [ ] Fix DoublyLinkList: add prev pointer
- [ ] Fix try2_mux/Eg2: set server.Addr = ":8080"
- [ ] Fix content-type: "app/json" → "application/json"
- [ ] Delete stripedDown/main.go or rename to anti_pattern.go
- [ ] Add nil checks to linked list methods
- [ ] Change PrintList to pointer receiver

## Minor

- [ ] Install Code Spell Checker extension
- [ ] Fix: "Initilizating" → "Initializing"
- [ ] Fix: "simpally" → "simply"
- [ ] Fix: "lest" → "let's"
- [ ] Fix: "SingelyLinkList" → "SinglyLinkedList"
- [ ] Delete commented placeholder functions
- [ ] Reduce comments in try1/main.go

## Learn

- [ ] POST vs PUT vs PATCH
- [ ] HTTP status codes (200, 201, 204, 400, 404, 500)
- [ ] Table-driven tests
- [ ] go test -cover
- [ ] Context usage
- [ ] Middleware pattern
