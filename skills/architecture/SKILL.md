---
name: system-architecture
description: Master system architecture and design patterns. Learn to design scalable systems, choose appropriate architectures, and solve complex technical challenges.
---

# System Architecture & Design

## Quick Start - Monolithic vs Microservices

```
Monolithic:
[Web] [API] [DB] [Cache] - Simple, coupled

Microservices:
[Gateway] → [User Service] → [DB]
         → [Post Service] → [DB]
         → [Comment Service] → [DB]
```

## Architectural Styles

**Monolithic**: Single codebase, simple but coupled
**Microservices**: Multiple services, complex but scalable
**Serverless**: Function-based, pay-per-use
**Event-driven**: Asynchronous, loosely coupled

## Design Patterns

**Creational**: Singleton, Factory, Builder
**Structural**: Adapter, Decorator, Facade
**Behavioral**: Observer, Strategy, Command

## Core Concepts

- Scalability (horizontal vs vertical)
- Load balancing
- Caching strategies
- Database sharding
- Event-driven architecture
- Service communication

## Learning Path

1. Design patterns fundamentals
2. Architecture basics
3. System design interviews
4. Large-scale system design
5. Enterprise patterns

## Key Skills

- Trade-off analysis
- Capacity planning
- Bottleneck identification
- Technology selection
- Documentation
