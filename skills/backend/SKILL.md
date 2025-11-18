---
name: backend-fundamentals
description: Learn server-side development fundamentals including REST APIs, databases, authentication, and system design principles required for all backend languages.
---

# Backend Fundamentals

## Quick Start

```javascript
// Express.js - Simple backend API
const express = require('express');
const app = express();

app.get('/api/users', (req, res) => {
  res.json({ users: [] });
});

app.post('/api/users', (req, res) => {
  // Validate input, save to database
  res.status(201).json({ id: 1 });
});

app.listen(3000);
```

## Core Concepts

### REST API Design
- HTTP methods (GET, POST, PUT, DELETE, PATCH)
- Status codes (200, 201, 400, 401, 404, 500)
- Request/response handling
- JSON serialization
- API versioning strategies

### Databases
- ACID properties and transactions
- Relational databases (SQL)
- Schema design and normalization
- Query writing and optimization
- Connection pooling

### Authentication & Security
- Passwords (hashing, salting)
- JWT (JSON Web Tokens)
- Sessions vs token-based auth
- CORS and CSRF protection
- Input validation and sanitization
- SQL injection prevention

### Patterns & Architecture
- MVC/MVVM architecture
- Repository pattern
- Dependency injection
- Service layers
- Error handling

### Tools & Technologies
- Package managers (npm, pip, cargo)
- Databases (PostgreSQL, MongoDB)
- Testing (Jest, pytest, unit tests)
- Version control (Git)
- Environment management (.env files)

## Learning Progression

1. Learn one language and framework deeply
2. Master HTTP and REST principles
3. Build database-backed applications
4. Implement authentication
5. Scale and optimize

## Essential Skills

- Understanding request-response cycle
- Writing secure, maintainable code
- Debugging and testing
- Performance optimization
- Deployment and DevOps basics
