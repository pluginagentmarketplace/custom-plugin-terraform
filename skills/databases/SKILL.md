---
name: databases
description: Master database design and SQL. Learn relational databases, NoSQL, query optimization, data modeling, indexing, and build scalable data systems.
---

# Database Design & SQL

## Quick Start

```sql
-- Create table
CREATE TABLE users (
  id INT PRIMARY KEY,
  name VARCHAR(255),
  email VARCHAR(255) UNIQUE
);

-- Query with JOIN
SELECT u.name, COUNT(p.id) as posts
FROM users u
LEFT JOIN posts p ON u.id = p.user_id
GROUP BY u.id;
```

## Database Types

**Relational**: PostgreSQL, MySQL (structured data)
**NoSQL**: MongoDB, Redis (flexible data)
**Specialized**: Elasticsearch, Cassandra

## Core Topics

- Database design and normalization
- SQL fundamentals and advanced queries
- Indexing and query optimization
- Transactions and ACID properties
- Backup and recovery
- Scaling strategies

## Learning Path

1. SQL basics (1-2 weeks)
2. Database design (1-2 weeks)
3. Query optimization (1-2 weeks)
4. Advanced topics (NoSQL, scaling)
5. Real-world projects

## Key Concepts

- Normalization vs Denormalization
- Index strategies
- Query execution plans
- Replication and sharding
- Eventual consistency
