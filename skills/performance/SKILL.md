---
name: Performance Optimization & Caching
description: Master profiling, benchmarking, and optimization techniques. Learn caching strategies, database optimization, code optimization, and system-level tuning to build lightning-fast applications. Identify bottlenecks and achieve 10x-100x performance improvements.
sasmp_version: "1.3.0"
bonded_agent: 01-programming-fundamentals
bond_type: PRIMARY_BOND
---

# ⚡ Performance Optimization & Caching

**Build lightning-fast, scalable systems through profiling and optimization.**

Master performance profiling, caching strategies, and system optimization to handle millions of users efficiently.

## Quick Start: Profiling Your Application

### Python Profiling
```python
import cProfile
import pstats
from functools import wraps
import time

# Basic timing decorator
def measure_time(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        start = time.perf_counter()
        result = func(*args, **kwargs)
        elapsed = time.perf_counter() - start
        print(f"{func.__name__} took {elapsed:.4f} seconds")
        return result
    return wrapper

@measure_time
def slow_function():
    total = 0
    for i in range(10000000):
        total += i
    return total

# cProfile for detailed analysis
def profile_code():
    profiler = cProfile.Profile()
    profiler.enable()

    slow_function()

    profiler.disable()
    stats = pstats.Stats(profiler)
    stats.sort_stats('cumulative')
    stats.print_stats(10)  # Top 10 functions

profile_code()
```

### JavaScript Profiling
```javascript
// Using Performance API
function measureFunction() {
    const start = performance.now();

    // Code to measure
    let sum = 0;
    for (let i = 0; i < 10000000; i++) {
        sum += i;
    }

    const end = performance.now();
    console.log(`Function took ${end - start}ms`);
}

// Chrome DevTools profiling
console.profile('myProfile');
// Code to profile
console.profileEnd('myProfile');

// Benchmarking with statistics
const benchmark = async (fn, iterations = 1000) => {
    const times = [];

    for (let i = 0; i < iterations; i++) {
        const start = performance.now();
        await fn();
        const end = performance.now();
        times.push(end - start);
    }

    times.sort((a, b) => a - b);
    const avg = times.reduce((a, b) => a + b, 0) / times.length;
    const median = times[Math.floor(times.length / 2)];
    const p95 = times[Math.floor(times.length * 0.95)];
    const p99 = times[Math.floor(times.length * 0.99)];

    return { avg, median, p95, p99, min: times[0], max: times[times.length - 1] };
};

const result = await benchmark(myFunction);
console.log(result);
```

### Go Profiling
```go
package main

import (
    "fmt"
    "runtime"
    "runtime/pprof"
    "os"
    "time"
)

func slowFunction() int {
    total := 0
    for i := 0; i < 10000000; i++ {
        total += i
    }
    return total
}

func main() {
    // CPU Profiling
    cpuFile, _ := os.Create("cpu.prof")
    pprof.StartCPUProfile(cpuFile)
    defer pprof.StopCPUProfile()

    // Memory Profiling
    memFile, _ := os.Create("mem.prof")
    defer func() {
        runtime.GC()
        pprof.WriteHeapProfile(memFile)
    }()

    // Measure execution time
    start := time.Now()
    result := slowFunction()
    elapsed := time.Since(start)

    fmt.Printf("Result: %d, took %v\n", result, elapsed)
}
```

## Caching Strategies

### Cache-Aside (Lazy Loading)
```python
import redis

cache = redis.Redis(host='localhost', port=6379)
db = Database()

def get_user(user_id):
    # Check cache first
    cached_user = cache.get(f"user:{user_id}")
    if cached_user:
        return json.loads(cached_user)

    # Cache miss, fetch from DB
    user = db.query(f"SELECT * FROM users WHERE id = {user_id}")
    if user:
        # Store in cache for 1 hour
        cache.setex(f"user:{user_id}", 3600, json.dumps(user))

    return user
```

### Write-Through Cache
```python
def update_user(user_id, data):
    # Update cache AND database (synchronously)
    user = {**get_user(user_id), **data}

    # Write to cache first
    cache.setex(f"user:{user_id}", 3600, json.dumps(user))

    # Then write to database
    db.update("users", {"id": user_id}, data)

    return user
```

### Write-Behind (Write-Back) Cache
```python
import queue
import threading

write_queue = queue.Queue()

def update_user_async(user_id, data):
    user = {**get_user(user_id), **data}

    # Update cache immediately
    cache.setex(f"user:{user_id}", 3600, json.dumps(user))

    # Queue database update (asynchronous)
    write_queue.put({"user_id": user_id, "data": data})

    return user

def background_writer():
    while True:
        item = write_queue.get()
        try:
            db.update("users", {"id": item["user_id"]}, item["data"])
        except Exception as e:
            print(f"Write failed: {e}")
            write_queue.put(item)  # Retry

# Start background writer thread
writer_thread = threading.Thread(target=background_writer, daemon=True)
writer_thread.start()
```

### Multi-Level Caching
```python
# L1: In-memory cache (fast, small)
class InMemoryCache:
    def __init__(self):
        self.data = {}

    def get(self, key):
        return self.data.get(key)

    def set(self, key, value):
        self.data[key] = value

# L2: Redis cache (faster than DB, larger)
# L3: Database (slowest, source of truth)

class CacheManager:
    def __init__(self):
        self.l1 = InMemoryCache()
        self.l2 = redis.Redis()
        self.db = Database()

    def get(self, key):
        # Try L1 (in-memory)
        value = self.l1.get(key)
        if value:
            return value

        # Try L2 (Redis)
        value = self.l2.get(key)
        if value:
            self.l1.set(key, value)
            return value

        # Fetch from database
        value = self.db.query(f"SELECT * FROM data WHERE key = '{key}'")
        if value:
            self.l1.set(key, value)
            self.l2.setex(key, 3600, json.dumps(value))

        return value
```

## Database Optimization

### Indexing Strategies
```sql
-- Single column index for WHERE clauses
CREATE INDEX idx_email ON users(email);

-- Composite index for multiple conditions
CREATE INDEX idx_user_date ON posts(user_id, created_at DESC);

-- Partial index for filtered data
CREATE INDEX idx_active_users ON users(id) WHERE status = 'active';

-- Index statistics
ANALYZE TABLE users;
EXPLAIN SELECT * FROM users WHERE email = 'user@example.com';
```

### Query Optimization
```sql
-- Bad: Fetches unnecessary columns
SELECT * FROM users;  -- 50 columns

-- Good: Fetch only needed columns
SELECT id, email, name FROM users;

-- Bad: N+1 queries (loop over users, query posts for each)
-- SELECT * FROM users;
-- Loop: SELECT * FROM posts WHERE user_id = ?;

-- Good: Single JOIN query
SELECT u.*, COUNT(p.id) as post_count
FROM users u
LEFT JOIN posts p ON u.id = p.user_id
GROUP BY u.id;

-- Bad: Complex calculations in WHERE
SELECT * FROM orders
WHERE YEAR(order_date) = YEAR(CURDATE())
  AND MONTH(order_date) = MONTH(CURDATE());

-- Good: Use date ranges
SELECT * FROM orders
WHERE order_date >= CURDATE() - INTERVAL 1 MONTH
  AND order_date < CURDATE();
```

### Connection Pooling
```python
from sqlalchemy.pool import QueuePool
from sqlalchemy import create_engine

# Create connection pool
engine = create_engine(
    'postgresql://user:password@localhost/db',
    poolclass=QueuePool,
    pool_size=20,  # Keep 20 connections open
    max_overflow=40,  # Allow up to 40 overflow connections
    pool_recycle=3600,  # Recycle connections after 1 hour
    pool_pre_ping=True  # Verify connections before using
)

# Connections are automatically managed
with engine.connect() as conn:
    result = conn.execute("SELECT * FROM users")
```

## Code Optimization

### Algorithm Optimization
```python
# Bad: O(n²) algorithm
def find_duplicates_slow(arr):
    duplicates = []
    for i in range(len(arr)):
        for j in range(i + 1, len(arr)):
            if arr[i] == arr[j]:
                duplicates.append(arr[i])
    return duplicates

# Good: O(n) algorithm using set
def find_duplicates_fast(arr):
    seen = set()
    duplicates = set()
    for num in arr:
        if num in seen:
            duplicates.add(num)
        seen.add(num)
    return list(duplicates)

# Benchmark
import timeit
slow_time = timeit.timeit(lambda: find_duplicates_slow(range(1000)), number=100)
fast_time = timeit.timeit(lambda: find_duplicates_fast(range(1000)), number=100)
print(f"Slow: {slow_time:.4f}s, Fast: {fast_time:.4f}s, Improvement: {slow_time/fast_time:.0f}x")
```

### Memory Optimization
```python
# Bad: Unnecessary list copies
def process_data_slow(data):
    data_copy = list(data)  # Copy entire list
    filtered = [x for x in data_copy if x > 0]
    return [x * 2 for x in filtered]

# Good: Use generators
def process_data_fast(data):
    filtered = (x for x in data if x > 0)
    return (x * 2 for x in filtered)

# Use generators for large datasets
def read_large_file(filename):
    with open(filename) as f:
        for line in f:  # Lazy reading
            yield line.strip()

# Process without loading entire file
for line in read_large_file('large_file.txt'):
    process(line)
```

## System-Level Optimization

### Load Balancing
```
┌─────────────────────┐
│  Load Balancer      │
│  (Round Robin)      │
└──────────┬──────────┘
           │
    ┌──────┼──────┐
    │      │      │
    ▼      ▼      ▼
  Server1 Server2 Server3
  (CPU: 50%) (CPU: 50%) (CPU: 50%)
```

### Content Delivery Network (CDN)
```
User in New York
    ↓
CDN Edge Node (NY)
    ↓
CDN Cache Hit → Serve immediately

User in Tokyo
    ↓
CDN Edge Node (Tokyo)
    ↓
CDN Cache Miss → Fetch from origin → Serve & Cache
```

### Compression
```javascript
// Gzip compression
const compression = require('compression');
app.use(compression());

// Response: 100KB
// After compression: 10KB (90% reduction)
```

## Learning Path

### Phase 1: Foundation (Weeks 1-4)
- Profiling basics
- Simple caching strategies
- Database indexing
- **Projects**:
  - Profile slow application
  - Add basic caching
  - Index database queries

### Phase 2: Intermediate (Weeks 5-12)
- Advanced profiling
- Multi-level caching
- Query optimization
- Performance monitoring
- **Projects**:
  - Optimize system under load
  - Implement caching strategy
  - Reduce query time by 10x

### Phase 3: Advanced (Weeks 13-24)
- System-wide optimization
- Distributed caching
- Performance architecture
- Optimization at scale
- **Projects**:
  - Optimize system to 1M requests/sec
  - Design for billion-scale data
  - 100x performance improvement

## 14+ Production Projects

### Profiling Projects
1. **CPU Profiler Tool** - Identify bottlenecks
2. **Memory Analyzer** - Detect leaks and over-allocation
3. **Query Analyzer** - Slow query detection
4. **Latency Tracker** - Request latency monitoring

### Caching Projects
5. **Distributed Cache** - Redis cluster setup
6. **Cache Invalidation System** - Smart cache management
7. **Cache Warming** - Pre-populate caches
8. **Cache Analytics** - Hit rate monitoring

### Optimization Projects
9. **Query Optimizer** - Automatic index suggestions
10. **Database Tuning** - Connection pooling, settings
11. **API Response Optimization** - Compression, pagination
12. **Image Optimization** - Resize, compress, serve optimized
13. **Code Optimizer** - Identify inefficient algorithms
14. **Full-stack Performance** - End-to-end optimization

## Performance Monitoring

### Key Metrics
```
- Response Time (p50, p95, p99)
- Throughput (requests/second)
- Error Rate (%)
- CPU Usage (%)
- Memory Usage (%)
- Disk I/O (operations/sec)
- Network Bandwidth (Mbps)
- Cache Hit Rate (%)
```

### Monitoring Tools
- **APM**: New Relic, Datadog, Dynatrace
- **Metrics**: Prometheus, InfluxDB
- **Logging**: ELK Stack, Splunk
- **Tracing**: Jaeger, Zipkin

## 🏆 Success Milestones

- [ ] Profile code and identify bottlenecks
- [ ] Implement caching improving response by 10x
- [ ] Optimize database queries by 100x
- [ ] Reduce memory usage by 50%
- [ ] Handle 10,000 requests/second
- [ ] Implement monitoring dashboard
- [ ] Design system for million-scale performance
- [ ] Achieve sub-100ms response times at scale

---

**Ready to optimize? Start profiling your application and find the biggest bottlenecks!**
