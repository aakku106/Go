# Recommended Go Projects - Based on Your Current Level

**Your Current Level**: Beginner → Intermediate Transition  
**Skills Demonstrated**: Data structures, CLI apps, algorithm implementation  
**Recommendation Date**: December 14, 2025

---

## 🎯 How to Use This Guide

This guide contains **12 projects** organized in 3 tiers:

- **Tier 1**: Consolidate what you know (1-2 weeks each)
- **Tier 2**: Stretch your skills (2-4 weeks each)
- **Tier 3**: Challenge projects (1-2 months each)

Start with Tier 1, complete at least 2 projects before moving up.

---

## 📊 Tier 1: Consolidation Projects

**Goal**: Master what you've learned, add testing and best practices

### Project 1: Data Structures Library ⭐ **START HERE**

**Duration**: 1-2 weeks  
**Difficulty**: ★★☆☆☆

**Description**:
Refactor all your existing data structures into a proper Go library with tests.

**What You'll Learn**:

- Proper Go package structure
- Unit testing with `testing` package
- Documentation with godoc
- Interfaces and polymorphism
- Benchmarking

**Requirements**:

```
datastructures/
├── go.mod
├── README.md
├── stack/
│   ├── stack.go
│   ├── stack_test.go
│   └── example_test.go
├── queue/
│   ├── linear.go
│   ├── linear_test.go
│   ├── circular.go
│   ├── circular_test.go
│   ├── priority.go
│   └── priority_test.go
└── list/
    ├── linked.go
    ├── linked_test.go
    ├── doubly.go
    └── doubly_test.go
```

**Specific Tasks**:

1. Create struct-based implementations (no global variables)
2. Write comprehensive tests (aim for 80%+ coverage)
3. Add benchmarks for all operations
4. Document all exported functions
5. Create examples in `example_test.go`
6. Add a Queue interface that all queues implement:
   ```go
   type Queue interface {
       Enqueue(value int) error
       Dequeue() (int, error)
       Peek() (int, error)
       IsEmpty() bool
       Size() int
   }
   ```

**Success Criteria**:

- All tests pass: `go test ./...`
- No race conditions: `go test -race ./...`
- Good coverage: `go test -cover ./...`
- Documentation works: `go doc datastructures/stack`
- Can be imported by other projects

**Skills Gained**:

- Go modules
- Unit testing
- Interface design
- Package documentation
- Code organization

---

### Project 2: CLI Task Manager 📝

**Duration**: 1 week  
**Difficulty**: ★★☆☆☆

**Description**:
Build a command-line TODO application with persistent storage.

**Features**:

```bash
# Commands
todo add "Buy groceries"          # Add task
todo list                         # List all tasks
todo complete 1                   # Mark task 1 as done
todo delete 2                     # Delete task 2
todo search "groceries"           # Search tasks
todo clear                        # Clear completed tasks
```

**What You'll Learn**:

- File I/O (reading/writing JSON)
- Command-line flag parsing (`flag` package)
- Struct marshaling/unmarshaling
- Error handling patterns
- Time handling (for due dates)

**Data Structure**:

```go
type Task struct {
    ID          int       `json:"id"`
    Description string    `json:"description"`
    Completed   bool      `json:"completed"`
    CreatedAt   time.Time `json:"created_at"`
    DueDate     *time.Time `json:"due_date,omitempty"`
}

type TodoList struct {
    Tasks []Task `json:"tasks"`
    file  string
}
```

**Bonus Features**:

- Add priorities (use your priority queue knowledge!)
- Add tags/categories
- Export to CSV
- Due date reminders
- Colored output (use `github.com/fatih/color`)

**Skills Gained**:

- CLI applications
- JSON handling
- File I/O
- Time management
- User input validation

---

### Project 3: Simple HTTP API Server 🌐

**Duration**: 1-2 weeks  
**Difficulty**: ★★★☆☆

**Description**:
Create a REST API for a simple bookmark manager.

**Endpoints**:

```
GET    /bookmarks           # List all bookmarks
GET    /bookmarks/:id       # Get specific bookmark
POST   /bookmarks           # Create bookmark
PUT    /bookmarks/:id       # Update bookmark
DELETE /bookmarks/:id       # Delete bookmark
GET    /bookmarks/search?q= # Search bookmarks
```

**What You'll Learn**:

- HTTP server with `net/http`
- RESTful API design
- JSON encoding/decoding
- Routing (use `gorilla/mux` or `chi`)
- Middleware concepts
- Error responses

**Code Structure**:

```go
type Bookmark struct {
    ID          string    `json:"id"`
    URL         string    `json:"url"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Tags        []string  `json:"tags"`
    CreatedAt   time.Time `json:"created_at"`
}

type Server struct {
    store  Store
    router *http.ServeMux
}

type Store interface {
    Create(b Bookmark) error
    Get(id string) (Bookmark, error)
    List() ([]Bookmark, error)
    Update(id string, b Bookmark) error
    Delete(id string) error
}
```

**Bonus Features**:

- Add authentication (JWT tokens)
- Add rate limiting
- Add pagination
- Add database (SQLite or PostgreSQL)
- Add logging middleware
- Write API tests

**Skills Gained**:

- Web development
- HTTP protocol
- REST principles
- Middleware patterns
- Testing HTTP handlers

---

### Project 4: File System Organizer 📁

**Duration**: 1 week  
**Difficulty**: ★★☆☆☆

**Description**:
Build a tool that organizes messy folders by file type, date, or size.

**Features**:

```bash
organize --input ~/Downloads --output ~/Organized --by type
organize --input ~/Photos --by date --format 2006-01
organize --input ~/Documents --by size --dry-run
```

**What You'll Learn**:

- File system operations (`os`, `path/filepath`)
- Walking directory trees
- File metadata (size, modified time)
- Command-line flags
- Concurrent operations (goroutines)

**Example Organization**:

```
Before:
Downloads/
  ├── photo.jpg
  ├── doc.pdf
  ├── song.mp3
  └── code.go

After (by type):
Organized/
  ├── Images/
  │   └── photo.jpg
  ├── Documents/
  │   └── doc.pdf
  ├── Audio/
  │   └── song.mp3
  └── Code/
      └── code.go
```

**Bonus Features**:

- Duplicate file detection (hash files)
- Archive old files (compress to .zip)
- Watch mode (monitor folder for changes)
- Config file for custom rules
- Progress bar (use `github.com/schollz/progressbar`)

**Skills Gained**:

- File operations
- Concurrent programming
- Hash functions
- Configuration management

---

## 📊 Tier 2: Skill Expansion Projects

**Goal**: Learn new Go concepts and build more complex applications

### Project 5: Web Scraper with Concurrency 🕷️

**Duration**: 2 weeks  
**Difficulty**: ★★★☆☆

**Description**:
Build a concurrent web scraper that extracts data from websites.

**Features**:

- Scrape product prices from e-commerce sites
- Concurrent fetching with goroutines
- Rate limiting to be polite
- Export to CSV/JSON
- Retry failed requests
- Track scraping progress

**What You'll Learn**:

- Goroutines and channels
- HTTP client programming
- HTML parsing (`golang.org/x/net/html` or `goquery`)
- Concurrency patterns (worker pools)
- Context for cancellation
- Rate limiting

**Example Code Structure**:

```go
type Scraper struct {
    client      *http.Client
    rateLimit   time.Duration
    maxWorkers  int
    results     chan Result
}

type Result struct {
    URL   string
    Title string
    Price float64
    Error error
}

func (s *Scraper) Scrape(urls []string) []Result {
    // Use worker pool pattern
}
```

**Skills Gained**:

- Concurrency in Go
- HTTP client usage
- HTML parsing
- Channel patterns
- Context usage

---

### Project 6: Chat Application (CLI or Web) 💬

**Duration**: 2-3 weeks  
**Difficulty**: ★★★★☆

**Description**:
Build a real-time chat system with Go.

**Two Options**:

**Option A: CLI Chat** (Easier)

- TCP server/client
- Multiple chat rooms
- Private messages
- User nicknames
- Command system (/join, /leave, /msg)

**Option B: Web Chat** (Harder)

- WebSocket server
- Browser-based client (HTML/JS)
- User authentication
- Chat history
- Online user list

**What You'll Learn**:

- TCP/WebSocket programming
- Goroutines for handling connections
- Broadcasting messages
- Connection management
- Protocol design

**Example Server Structure**:

```go
type Server struct {
    clients   map[*Client]bool
    broadcast chan Message
    register  chan *Client
    unregister chan *Client
}

type Client struct {
    conn *websocket.Conn
    send chan Message
    name string
    room string
}

type Message struct {
    Type    string    `json:"type"`
    From    string    `json:"from"`
    To      string    `json:"to,omitempty"`
    Content string    `json:"content"`
    Time    time.Time `json:"time"`
}
```

**Skills Gained**:

- Network programming
- Real-time communication
- Goroutine coordination
- WebSocket protocol
- State management

---

### Project 7: Markdown Blog Generator 📝

**Duration**: 2 weeks  
**Difficulty**: ★★★☆☆

**Description**:
Build a static site generator that converts Markdown files to HTML.

**Features**:

```bash
blog new "My First Post"              # Create new post
blog build                             # Generate HTML
blog serve                             # Local server
blog deploy                            # Deploy to GitHub Pages
```

**What You'll Learn**:

- File processing
- Template system (`text/template`, `html/template`)
- Markdown parsing (`github.com/russross/blackfriday`)
- Static site generation
- File watching for auto-rebuild

**Project Structure**:

```
myblog/
├── content/
│   ├── posts/
│   │   ├── 2025-01-15-first-post.md
│   │   └── 2025-01-16-second-post.md
│   └── pages/
│       └── about.md
├── templates/
│   ├── base.html
│   ├── post.html
│   └── index.html
├── static/
│   ├── css/
│   └── js/
└── public/          # Generated files
    └── index.html
```

**Skills Gained**:

- Template systems
- File generation
- Markdown processing
- HTML generation
- Static site concepts

---

### Project 8: Simple Database (Key-Value Store) 🗄️

**Duration**: 2-3 weeks  
**Difficulty**: ★★★★☆

**Description**:
Implement a simple in-memory database with persistence.

**Features**:

```go
db.Set("user:1", "John")              // Set value
db.Get("user:1")                      // Get value
db.Delete("user:1")                   // Delete
db.Expire("session:123", 30*time.Minute)  // TTL
db.Save("data.db")                    // Persist to disk
db.Load("data.db")                    // Load from disk
```

**What You'll Learn**:

- Hash tables and data structures
- Binary encoding/decoding
- File formats
- Concurrent safe operations (sync.RWMutex)
- Serialization

**Advanced Features**:

- Transaction support
- Multiple data types (string, list, set)
- Pattern matching (GET user:\*)
- Pub/Sub system
- Replication (master-slave)

**Skills Gained**:

- Low-level data storage
- Concurrency safety
- Binary formats
- Performance optimization
- Database concepts

---

## 📊 Tier 3: Challenge Projects

**Goal**: Build production-ready applications

### Project 9: Load Balancer 🔄

**Duration**: 3-4 weeks  
**Difficulty**: ★★★★★

**Description**:
Implement a reverse proxy load balancer that distributes traffic across multiple backends.

**Features**:

- Round-robin, least connections, weighted algorithms
- Health checking of backends
- SSL/TLS termination
- Request logging
- Circuit breaker pattern
- Admin API for configuration

**What You'll Learn**:

- Reverse proxy concepts
- Load balancing algorithms
- HTTP proxying
- Health checks
- Production patterns

**Skills Gained**:

- Advanced networking
- System design
- Production operations
- Performance testing
- Observability

---

### Project 10: Container Runtime (Mini Docker) 🐳

**Duration**: 1-2 months  
**Difficulty**: ★★★★★

**Description**:
Build a simplified container runtime using Linux namespaces.

**Features**:

- Create isolated processes
- Filesystem isolation (chroot)
- Network namespaces
- Resource limits (cgroups)
- Image layers

**What You'll Learn**:

- Linux namespaces
- Filesystem operations
- System calls
- Process management
- Container concepts

**Skills Gained**:

- Systems programming
- Low-level Linux
- Container internals
- Advanced Go

---

### Project 11: Distributed Task Queue 📬

**Duration**: 1-2 months  
**Difficulty**: ★★★★★

**Description**:
Build a distributed job queue system (like Celery/RabbitMQ).

**Features**:

- Multiple workers
- Task priorities
- Retry logic
- Dead letter queue
- Dashboard for monitoring
- Persistence

**What You'll Learn**:

- Distributed systems
- Message queuing
- Worker pools
- Job scheduling
- System resilience

**Skills Gained**:

- Distributed systems
- Queue systems
- Reliability patterns
- Monitoring
- Scale considerations

---

### Project 12: Terminal Text Editor 📝

**Duration**: 1-2 months  
**Difficulty**: ★★★★★

**Description**:
Build a terminal-based text editor (like nano or vim).

**Features**:

- File editing with cursor movement
- Syntax highlighting
- Search and replace
- Multiple buffers
- Command mode
- Copy/paste
- Undo/redo

**What You'll Learn**:

- Terminal control
- Text rendering
- Buffer management
- Rope data structure
- Event handling

**Skills Gained**:

- Terminal programming
- Complex UI in terminal
- Text processing
- Performance optimization
- User experience

---

## 🎯 Project Selection Guide

### If you want to focus on...

**Testing & Best Practices**:
→ Start with Project 1 (Data Structures Library)

**CLI Tools**:
→ Projects 2, 4

**Web Development**:
→ Projects 3, 7

**Concurrency**:
→ Projects 5, 6

**Systems Programming**:
→ Projects 8, 9, 10

**Distributed Systems**:
→ Project 11

**Complex Algorithms**:
→ Projects 8, 12

---

## 📚 Learning Path Recommendation

### Month 1-2: Foundation

1. ✅ **Project 1**: Data Structures Library (MUST DO)
2. ✅ **Project 2**: CLI Task Manager
3. ✅ **Project 3**: HTTP API Server

**Why**: These will fix your current gaps (testing, organization, errors) while teaching new skills.

### Month 3-4: Growth

4. ✅ **Project 5**: Web Scraper (learn concurrency)
5. ✅ **Project 6 or 7**: Chat App OR Blog Generator

**Why**: Introduces concurrency and real-world patterns.

### Month 5+: Mastery

6. ✅ **Pick from Tier 3**: Based on your interests

**Why**: Build something impressive for your portfolio.

---

## 💡 Tips for Success

### For Every Project:

1. **Write Tests First**: Practice TDD
2. **Use Git**: Commit frequently
3. **Write README**: Document setup and usage
4. **Code Review**: Ask for feedback (Reddit, Discord)
5. **Refactor**: Don't stop at "working" - make it clean
6. **Deploy**: Put it somewhere (GitHub, personal server)

### When Stuck:

1. **Read Documentation**: Official Go docs are excellent
2. **Check Examples**: Look at similar projects on GitHub
3. **Ask Questions**: r/golang, Go Forum, Stack Overflow
4. **Pair Program**: Find a Go buddy
5. **Take Breaks**: Sometimes you need distance

### Track Progress:

Create a learning journal:

```markdown
# Go Learning Journal

## Project 1: Data Structures Library

- Started: 2025-01-15
- Completed: 2025-01-30
- Skills: Testing, interfaces, documentation
- Challenges: Understanding table-driven tests
- Proud of: 95% test coverage!
- GitHub: github.com/yourusername/datastructures
```

---

## 🚀 Next Steps

1. ⭐ **START TODAY**: Begin Project 1 (Data Structures Library)
2. 📅 **Schedule**: Dedicate 1-2 hours daily
3. 🎯 **Goal**: Complete 2 Tier 1 projects in next month
4. 📝 **Track**: Keep notes on what you learn
5. 🤝 **Share**: Post your progress, get feedback

---

## 🎊 Final Motivation

You've done **excellent** work in your first week! Your understanding of slices and data structures is solid. Now it's time to:

1. **Clean up** your existing code (use these reviews!)
2. **Build** your first properly structured project
3. **Test** everything you write
4. **Share** your work and get feedback

**Remember**:

- Don't rush to finish
- Quality > Quantity
- Testing is not optional
- Refactoring is part of the process
- Asking questions shows strength, not weakness

**You've got this! Happy coding! 🚀**

---

## 📖 Additional Resources

### Communities

- r/golang (Reddit)
- Gophers Slack (invite.slack.golangbridge.org)
- Go Forum (forum.golangbridge.org)

### Blogs

- The Go Blog (go.dev/blog)
- Dave Cheney (dave.cheney.net)
- Ardan Labs Blog

### GitHub Projects to Study

- **Simple**: github.com/gorilla/mux
- **Medium**: github.com/gin-gonic/gin
- **Advanced**: github.com/kubernetes/kubernetes

### Practice Platforms

- Exercism.org (Go track)
- LeetCode (filter by Go)
- Codewars (Go kata)
- HackerRank

**Good luck! 🎉**
