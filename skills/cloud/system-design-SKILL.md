---
name: system-design
description: Master large-scale distributed system design with caching, databases, load balancing, microservices, and patterns for handling millions of users.
---

# System Design

## Quick Start - Basic Architecture

```
                     [Users]
                       |
              [Load Balancer]
                       |
        ____________/  |  \_____________
       /              |              \
   [Web 1]        [Web 2]        [Web 3]
       \              |              /
        \_____________|_____________/
                  [Cache Layer]
                  (Redis/Memcached)
                       |
                  [Database]
                  (PostgreSQL)
```

## Core Concepts

### Load Balancing
- Round-robin, least connections, IP hash
- Sticky sessions
- Health checks and failover
- Global load balancing

### Caching Strategies
- **Cache-Aside** - Check cache, fetch from DB
- **Write-Through** - Write to cache and DB
- **Write-Behind** - Write to cache, async to DB
- Cache invalidation patterns

### Database Design
- SQL vs NoSQL decision factors
- Replication (master-slave, master-master)
- Sharding and partitioning strategies
- Read replicas for scaling

### Consistency Models
- Strong consistency (ACID)
- Eventual consistency (BASE)
- Causal consistency
- CAP theorem trade-offs

### Messaging & Queues
- Message brokers (Kafka, RabbitMQ)
- Publish-subscribe patterns
- Asynchronous processing
- Event-driven architecture

### API Design
- Pagination and filtering
- Rate limiting
- Versioning strategies
- API gateway patterns

## Scalability Patterns

- Horizontal scaling (multiple servers)
- Vertical scaling (bigger servers)
- Caching layers
- Database optimization
- CDN for static content
- Microservices architecture

## System Design Interview Patterns

- Functional requirements clarification
- Non-functional requirements (scale, latency)
- High-level design
- Detailed design
- Trade-offs and optimization

## Real-World Examples

- Twitter-scale systems
- Netflix streaming architecture
- YouTube video delivery
- Facebook's infrastructure
- Uber's real-time systems

## Performance Metrics

- Latency (P50, P95, P99)
- Throughput (requests/second)
- Availability and reliability
- Consistency levels

## Monitoring & Observability

- Metrics collection
- Distributed tracing
- Log aggregation
- Alerting strategies
- Dashboard design
