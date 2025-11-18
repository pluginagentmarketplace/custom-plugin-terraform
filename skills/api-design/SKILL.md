---
name: api-design
description: Design and build REST APIs and GraphQL. Learn API security, documentation, versioning, rate limiting, and create professional, scalable APIs.
---

# API Design & Development

## Quick Start

```javascript
// Express.js REST API
app.get('/users/:id', async (req, res) => {
  const user = await User.findById(req.params.id);
  res.json(user);
});

// GraphQL Query
{
  user(id: 1) {
    name
    posts { title }
  }
}
```

## API Types

**REST**: Resource-based, stateless, widely used
**GraphQL**: Query language, flexible, efficient
**RPC**: Function calls, simple, specific use

## Core Topics

- HTTP methods and status codes
- Resource design and naming
- Error handling and validation
- API versioning strategies
- Authentication and authorization
- Rate limiting and throttling
- Documentation (OpenAPI/Swagger)
- Testing APIs

## Learning Path

1. REST fundamentals (1 week)
2. Build first API (2 weeks)
3. Security and authentication (1 week)
4. GraphQL basics (1 week)
5. Advanced patterns (scaling, caching)

## Best Practices

- Use proper HTTP methods
- Consistent naming conventions
- Meaningful status codes
- Clear error messages
- API documentation
- Security first
- Backwards compatibility
