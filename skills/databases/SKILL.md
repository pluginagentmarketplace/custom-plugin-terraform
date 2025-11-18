---
name: database-design
description: Master database design, schema creation, normalization, query optimization, and relational database fundamentals for building efficient data systems.
---

# Database Design

## Quick Start

```sql
-- Create normalized schema
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE posts (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  title VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create index for performance
CREATE INDEX idx_posts_user_id ON posts(user_id);

-- Query with JOIN
SELECT u.email, COUNT(p.id) as post_count
FROM users u
LEFT JOIN posts p ON u.id = p.user_id
GROUP BY u.id;
```

## Core Concepts

### Data Modeling
- Entity-relationship diagrams
- Relationships (one-to-one, one-to-many, many-to-many)
- Keys and constraints
- Referential integrity

### Normalization
- First normal form (1NF) - atomic values
- Second normal form (2NF) - no partial dependencies
- Third normal form (3NF) - no transitive dependencies
- BCNF for edge cases
- Denormalization trade-offs

### SQL Fundamentals
- SELECT, INSERT, UPDATE, DELETE
- JOINs (INNER, LEFT, RIGHT, FULL)
- Aggregation (GROUP BY, HAVING)
- Subqueries and CTEs
- Window functions

### Indexing Strategy
- B-tree indexes (most common)
- Hash indexes for equality
- Bitmap indexes
- Composite indexes
- Index performance analysis (EXPLAIN)

### Query Optimization
- Understanding execution plans
- Index selection
- Query rewriting
- Avoiding N+1 queries
- Batch operations

### Transactions
- ACID properties
- Isolation levels
- Deadlock prevention
- Lock management

## Database Types

- PostgreSQL - robust, feature-rich
- MySQL - popular, reliable
- MariaDB - MySQL compatible
- SQLite - embedded, simple

## Design Patterns

- Star schema for analytics
- Snowflake schema for complex warehouses
- Time-series optimization
- Event sourcing patterns

## Tools & Utilities

- **psql/mysql CLI** - SQL interface
- **pgAdmin/MySQL Workbench** - GUI tools
- **DBeaver** - Universal database tool
- **Liquibase/Flyway** - Migration tools
- **DataGrip** - IDE for databases
