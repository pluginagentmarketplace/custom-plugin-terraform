---
name: Database Management & Optimization
description: Master SQL databases, NoSQL systems, and data modeling. Learn query optimization, indexing strategies, transactions, ACID properties, scaling patterns (replication, sharding), and performance tuning. Design efficient schemas and architect databases for millions of records.
sasmp_version: "1.3.0"
bonded_agent: 01-programming-fundamentals
bond_type: PRIMARY_BOND
---

# 💾 Database Management & Optimization

**Design, optimize, and scale databases from development to production.**

Master relational and non-relational databases, from data modeling to optimization, replication, and sharding strategies.

## Quick Start: Database Fundamentals

### SQL Basics
```sql
-- Creating tables with constraints
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Creating indexes for performance
CREATE INDEX idx_user_email ON users(email);
CREATE INDEX idx_post_user_id ON posts(user_id);
CREATE INDEX idx_post_created_at ON posts(created_at DESC);

-- Query with JOINs and aggregation
SELECT
    u.id,
    u.email,
    COUNT(p.id) as post_count,
    MAX(p.created_at) as last_post_date
FROM users u
LEFT JOIN posts p ON u.id = p.user_id
WHERE u.created_at > DATE_SUB(NOW(), INTERVAL 30 DAY)
GROUP BY u.id
HAVING post_count > 0
ORDER BY post_count DESC
LIMIT 10;

-- Transactions for data consistency
START TRANSACTION;
    UPDATE users SET balance = balance - 100 WHERE id = 1;
    UPDATE users SET balance = balance + 100 WHERE id = 2;
COMMIT;

-- Window functions for analytics
SELECT
    email,
    post_count,
    ROW_NUMBER() OVER (ORDER BY post_count DESC) as rank,
    post_count - LAG(post_count) OVER (ORDER BY post_count DESC) as diff_from_prev
FROM (
    SELECT u.email, COUNT(p.id) as post_count
    FROM users u
    LEFT JOIN posts p ON u.id = p.user_id
    GROUP BY u.id
) ranked;
```

### NoSQL: MongoDB Example
```javascript
// MongoDB schema-less document store
db.users.insertOne({
    email: "user@example.com",
    name: "John Doe",
    profile: {
        bio: "Software engineer",
        location: "San Francisco",
        skills: ["JavaScript", "Python", "Go"]
    },
    posts: [
        { id: 1, title: "First Post", likes: 10 },
        { id: 2, title: "Second Post", likes: 25 }
    ],
    createdAt: new Date()
});

// Indexing for performance
db.users.createIndex({ email: 1 });
db.users.createIndex({ createdAt: -1 });
db.users.createIndex({ "profile.location": 1 });

// Aggregation pipeline for complex queries
db.users.aggregate([
    {
        $match: { createdAt: { $gte: new Date("2024-01-01") } }
    },
    {
        $addFields: { postCount: { $size: "$posts" } }
    },
    {
        $group: {
            _id: "$profile.location",
            avgPosts: { $avg: "$postCount" },
            totalUsers: { $sum: 1 }
        }
    },
    {
        $sort: { avgPosts: -1 }
    }
]);

// Transactions for multiple operations
const session = db.getMongo().startSession();
session.startTransaction();
try {
    db.users.updateOne({ _id: id1 }, { $inc: { balance: -100 } });
    db.users.updateOne({ _id: id2 }, { $inc: { balance: 100 } });
    session.commitTransaction();
} catch (error) {
    session.abortTransaction();
    throw error;
}
```

### PostgreSQL Advanced Features
```sql
-- JSON support for semi-structured data
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title TEXT,
    metadata JSONB DEFAULT '{}'::jsonb,
    tags TEXT[] DEFAULT ARRAY[]::TEXT[]
);

INSERT INTO posts (title, metadata, tags) VALUES
('My Post', '{"views": 100, "likes": 25, "category": "tech"}'::jsonb, ARRAY['tutorial', 'sql']);

-- Query JSON fields
SELECT title, metadata->>'category' as category, metadata->'views' as views
FROM posts
WHERE metadata @> '{"category": "tech"}'::jsonb;

-- Common Table Expressions (CTEs) for readable queries
WITH post_stats AS (
    SELECT
        user_id,
        COUNT(*) as total_posts,
        AVG(views) as avg_views,
        MAX(created_at) as last_post
    FROM posts
    GROUP BY user_id
),
top_users AS (
    SELECT user_id, total_posts
    FROM post_stats
    WHERE total_posts > 10
    ORDER BY total_posts DESC
    LIMIT 10
)
SELECT u.email, tu.total_posts, ps.avg_views
FROM top_users tu
JOIN post_stats ps ON tu.user_id = ps.user_id
JOIN users u ON tu.user_id = u.id;

-- Materialized view for expensive queries
CREATE MATERIALIZED VIEW user_activity_summary AS
SELECT
    u.id,
    u.email,
    COUNT(DISTINCT p.id) as post_count,
    COUNT(DISTINCT c.id) as comment_count,
    MAX(p.created_at) as last_activity
FROM users u
LEFT JOIN posts p ON u.id = p.user_id
LEFT JOIN comments c ON u.id = c.user_id
GROUP BY u.id;

-- Refresh materialized view
REFRESH MATERIALIZED VIEW CONCURRENTLY user_activity_summary;
```

## Data Modeling Principles

### Entity-Relationship Model
```
Users (1) ──── (M) Posts (1) ──── (M) Comments
    |                              |
    └──────── (M) Followers ────────┘

users:
  - id (PK)
  - email (UNIQUE)
  - name
  - created_at

posts:
  - id (PK)
  - user_id (FK to users)
  - title
  - content
  - created_at

comments:
  - id (PK)
  - post_id (FK to posts)
  - user_id (FK to users)
  - content
  - created_at

followers:
  - follower_id (FK to users)
  - following_id (FK to users)
  - created_at
  - PRIMARY KEY (follower_id, following_id)
```

### Normalization vs Denormalization
- **1NF**: Each attribute contains atomic values
- **2NF**: All non-key attributes depend on entire primary key
- **3NF**: No transitive dependencies
- **BCNF**: Every determinant is a candidate key

### Denormalization Trade-offs
```sql
-- Normalized approach (3NF)
SELECT u.email, COUNT(p.id) as post_count
FROM users u
LEFT JOIN posts p ON u.id = p.user_id
GROUP BY u.id;

-- Denormalized approach (single table with cache)
CREATE TABLE user_stats (
    user_id INT PRIMARY KEY,
    post_count INT DEFAULT 0,
    last_updated TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Update statistics when posts are created/deleted
CREATE TRIGGER update_user_stats AFTER INSERT ON posts
FOR EACH ROW
BEGIN
    UPDATE user_stats SET post_count = post_count + 1
    WHERE user_id = NEW.user_id;
END;
```

## Query Optimization Techniques

### Index Strategies
```sql
-- Single column index for WHERE clauses
CREATE INDEX idx_email ON users(email);
SELECT * FROM users WHERE email = 'user@example.com';

-- Composite index for multiple conditions
CREATE INDEX idx_user_date ON posts(user_id, created_at DESC);
SELECT * FROM posts
WHERE user_id = 123 AND created_at > '2024-01-01'
ORDER BY created_at DESC;

-- Partial index for filtered data
CREATE INDEX idx_active_users ON users(id)
WHERE status = 'active';

-- Full-text search index
CREATE FULLTEXT INDEX idx_post_search ON posts(title, content);
SELECT * FROM posts
WHERE MATCH(title, content) AGAINST('database optimization' IN BOOLEAN MODE);

-- Analyze index effectiveness
EXPLAIN ANALYZE
SELECT u.email, COUNT(p.id)
FROM users u
LEFT JOIN posts p ON u.id = p.user_id
WHERE u.created_at > '2024-01-01'
GROUP BY u.id;
```

### Query Optimization Patterns
```sql
-- Avoid N+1 queries: Bad
-- Loop over users and query posts for each
-- Query 1: SELECT * FROM users;
-- Query 2-N: SELECT * FROM posts WHERE user_id = ?;

-- Good: Use JOIN
SELECT u.*, COUNT(p.id) as post_count
FROM users u
LEFT JOIN posts p ON u.id = p.user_id
GROUP BY u.id;

-- Avoid SELECT *
SELECT id, email, name FROM users;  -- Only needed columns

-- Use LIMIT and pagination
SELECT * FROM posts ORDER BY created_at DESC LIMIT 20 OFFSET 0;
-- Better: cursor-based pagination
SELECT * FROM posts
WHERE id > :last_id
ORDER BY id ASC
LIMIT 20;

-- Batch operations
INSERT INTO logs (user_id, action, created_at) VALUES
(1, 'login', NOW()),
(2, 'logout', NOW()),
(3, 'update_profile', NOW());

-- Update with JOIN
UPDATE posts p
INNER JOIN users u ON p.user_id = u.id
SET p.featured = TRUE
WHERE u.premium_status = TRUE AND p.views > 1000;
```

## Scaling Strategies

### Vertical Scaling
- Increase hardware resources (CPU, RAM, disk)
- Single node with more power
- Simple but limited ceiling

### Horizontal Scaling: Replication
```
Master (Write)
    |
    ├── Replica 1 (Read)
    ├── Replica 2 (Read)
    └── Replica 3 (Read)

Benefits:
- Load distribution for reads
- High availability
- Data redundancy
```

### Horizontal Scaling: Sharding
```
Sharding by User ID (hash-based):

Shard 1: Users 1-1000000 ──► DB 1
Shard 2: Users 1000001-2000000 ──► DB 2
Shard 3: Users 2000001-3000000 ──► DB 3
Shard 4: Users 3000001-4000000 ──► DB 4

Algorithm: shard_id = hash(user_id) % num_shards

Challenges:
- Cross-shard queries are complex
- Uneven load distribution
- Shard rebalancing is difficult
```

## Advanced Concepts

### ACID Properties
- **Atomicity**: Transaction is all-or-nothing
- **Consistency**: Database moves from valid state to valid state
- **Isolation**: Concurrent transactions don't interfere
- **Durability**: Committed data persists despite failures

### Transaction Isolation Levels
```sql
-- READ UNCOMMITTED: Dirty reads possible
SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;

-- READ COMMITTED: No dirty reads
SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED;

-- REPEATABLE READ: No dirty/phantom reads
SET SESSION TRANSACTION ISOLATION LEVEL REPEATABLE READ;

-- SERIALIZABLE: Full isolation (slowest)
SET SESSION TRANSACTION ISOLATION LEVEL SERIALIZABLE;
```

### CAP Theorem
- **Consistency**: All nodes see same data
- **Availability**: System always responds
- **Partition Tolerance**: System works despite network splits

**Choose 2 of 3**: Most distributed systems are CP or AP, rarely CA.

## Learning Path

### Phase 1: Foundation (Weeks 1-4)
- SQL basics (SELECT, INSERT, UPDATE, DELETE)
- Table creation and relationships
- Basic JOINs (INNER, LEFT, RIGHT)
- WHERE clauses and filtering
- Simple indexes
- **Projects**:
  - Blog database with users/posts/comments
  - E-commerce product catalog
  - Employee management system

### Phase 2: Intermediate (Weeks 5-12)
- Advanced SQL (CTEs, window functions)
- Query optimization and EXPLAIN
- Transactions and ACID
- Database design and normalization
- Basic replication setup
- Introduction to NoSQL
- **Projects**:
  - Analytics dashboard database
  - Multi-tenant SaaS database
  - Social network with followers
  - Real-time activity feed

### Phase 3: Advanced (Weeks 13-24)
- Sharding strategies
- Performance monitoring
- Backup and recovery
- Distributed transactions
- Database architecture patterns
- Advanced NoSQL design
- **Projects**:
  - Sharded user database
  - High-throughput transaction system
  - Time-series database
  - Document storage system

## 15+ Production Projects

### SQL Projects
1. **Blog Platform Database** - Users, posts, comments, ratings
2. **E-commerce Database** - Products, orders, inventory
3. **Social Network** - Users, posts, followers, notifications
4. **Analytics Dashboard** - Events, aggregations, time-series
5. **Multi-tenant SaaS** - Organizations, users, data isolation
6. **Real-time Activity Feed** - User activities, notifications
7. **Geolocation Service** - Locations, proximity queries

### NoSQL Projects
8. **Document Management System** - MongoDB with hierarchical data
9. **User Profile Store** - Flexible schema with nested data
10. **Real-time Chat Application** - Event logs, message storage
11. **Product Catalog** - Variable attributes per category
12. **Time-series Metrics** - IoT sensor data storage
13. **Content Management System** - Flexible document structure

### Scaling Projects
14. **Sharded User Database** - Horizontal scaling by user ID
15. **Master-Slave Replication** - Read replicas for scale
16. **Data Warehouse** - Dimensional modeling, fact/dimension tables
17. **Search Index** - Elasticsearch integration and optimization

## Best Practices

### Design
- ✅ Use appropriate data types (not everything is STRING)
- ✅ Create indexes on frequently queried columns
- ✅ Normalize to 3NF unless specific reason not to
- ✅ Use foreign keys for referential integrity
- ✅ Plan for growth and partitioning early

### Performance
- ✅ Use EXPLAIN ANALYZE to understand queries
- ✅ Batch operations when possible
- ✅ Use pagination for large result sets
- ✅ Cache frequently accessed data
- ✅ Monitor slow query logs

### Security
- ✅ Use parameterized queries (prevent SQL injection)
- ✅ Encrypt sensitive data at rest
- ✅ Use row-level security for multi-tenant systems
- ✅ Regular backups with verification
- ✅ Principle of least privilege for users

### Operations
- ✅ Monitor disk space and growth
- ✅ Set up automated backups
- ✅ Plan for disaster recovery
- ✅ Document schema and relationships
- ✅ Use version control for schema changes

## 🏆 Success Milestones

- [ ] Design normalized schema for complex application
- [ ] Write efficient queries using JOINs and aggregations
- [ ] Set up appropriate indexes improving query performance by 10x
- [ ] Implement transactions for data consistency
- [ ] Set up master-slave replication
- [ ] Optimize query to 100x improvement using sharding
- [ ] Handle 1M+ records efficiently
- [ ] Implement backup and recovery procedures

---

**Ready to master databases? Start designing your first normalized schema!**
