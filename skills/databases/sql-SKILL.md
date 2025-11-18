---
name: sql-mastery
description: Master SQL from basics to advanced queries. Learn query optimization, window functions, common table expressions, and best practices for relational databases.
---

# SQL Mastery

## Quick Start - Advanced SQL

```sql
-- Window functions for ranking
SELECT
  user_id,
  order_date,
  amount,
  ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY order_date) as order_num,
  SUM(amount) OVER (PARTITION BY user_id) as total_spent
FROM orders;

-- Common table expression for complex logic
WITH monthly_revenue AS (
  SELECT DATE_TRUNC('month', order_date) as month,
         SUM(amount) as revenue
  FROM orders
  GROUP BY DATE_TRUNC('month', order_date)
)
SELECT month, revenue,
       LAG(revenue) OVER (ORDER BY month) as prev_month_revenue
FROM monthly_revenue;

-- Recursive CTE for hierarchies
WITH RECURSIVE org_hierarchy AS (
  SELECT id, name, manager_id, 1 as level
  FROM employees WHERE manager_id IS NULL
  UNION ALL
  SELECT e.id, e.name, e.manager_id, h.level + 1
  FROM employees e
  JOIN org_hierarchy h ON e.manager_id = h.id
)
SELECT * FROM org_hierarchy;
```

## Advanced Concepts

### Window Functions
- ROW_NUMBER, RANK, DENSE_RANK
- LAG, LEAD for time series
- SUM, AVG for running totals
- PARTITION BY and ORDER BY
- Sliding windows

### Common Table Expressions (CTE)
- Non-recursive CTEs
- Recursive CTEs for hierarchies
- Multiple CTEs in single query
- Readability and maintainability

### Optimization Techniques
- EXPLAIN PLAN analysis
- Query rewriting
- Index selection
- Avoiding full table scans
- Cost-based optimization

### Advanced JOINs
- Outer joins with NULLs
- Self-joins for hierarchies
- Cross joins
- Lateral joins (if supported)

### Set Operations
- UNION and UNION ALL
- INTERSECT for common rows
- EXCEPT for difference
- Performance implications

### Subqueries & Scalar Functions
- Correlated subqueries
- Subqueries in WHERE, FROM, SELECT
- Existence checks with EXISTS
- Aggregate functions

### Transactions
- Atomicity guarantees
- Isolation levels
- ACID properties
- Rollback scenarios

## Performance Tuning

- Index creation and maintenance
- Execution plan analysis
- Query optimization
- Statistics and cardinality
- Connection pooling

## Best Practices

- Use parameterized queries (prevent SQL injection)
- Write readable SQL with proper formatting
- Use views for complex logic
- Modularize with stored procedures
- Comment complex logic

## Tools for SQL Development

- **SQL IDE** - DataGrip, DBeaver, vs Code
- **Query Analyzer** - Execution plans
- **Profilers** - Performance analysis
- **Version Control** - Migration tracking
- **Testing** - Query validation
