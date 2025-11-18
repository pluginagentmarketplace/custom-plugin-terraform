---
name: nodejs-mastery
description: Master Node.js runtime, async programming, event-driven architecture, and build scalable real-time applications with Express, Fastify, or Nest frameworks.
---

# Node.js Mastery

## Quick Start

```javascript
// Node.js async patterns
async function fetchUser(id) {
  const response = await fetch(`/api/users/${id}`);
  const user = await response.json();
  return user;
}

// Event-driven architecture
const EventEmitter = require('events');
const emitter = new EventEmitter();

emitter.on('user:created', (user) => {
  console.log('User created:', user.name);
});

emitter.emit('user:created', { name: 'John' });
```

## Core Competencies

### Runtime & Fundamentals
- Event loop and non-blocking I/O
- Callbacks, Promises, async/await
- Stream processing for large files
- Module system (CommonJS, ESM)
- Global objects and APIs

### Popular Frameworks
- **Express.js** - Lightweight and flexible
- **Fastify** - High performance
- **Nest.js** - Opinionated, enterprise-ready
- **Hapi.js** - Robust and scalable

### Async Patterns
- Promise chaining and error handling
- Async/await for readable code
- Concurrent operations with Promise.all
- Stream processing and backpressure
- Callbacks for legacy patterns

### Database Integration
- MongoDB with Mongoose
- PostgreSQL with Sequelize or Prisma
- Redis for caching
- Connection pooling
- Transactions

### Real-time Features
- WebSockets with Socket.io
- Server-Sent Events
- Message queues (Bull, BullMQ)
- Pub/Sub patterns

### Deployment & DevOps
- Docker containerization
- Environment configuration
- PM2 process management
- Logging and monitoring
- Performance optimization

## Project Ideas

- REST API backend
- Real-time chat application
- Content management system
- E-commerce backend
- Data pipeline and ETL

## Essential Tools

- **npm/yarn** - Package management
- **Postman** - API testing
- **MongoDB Compass** - Database UI
- **Nodemon** - Development auto-reload
- **Jest** - Testing framework
