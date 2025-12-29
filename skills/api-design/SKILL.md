---
name: API Design & Development
description: Master REST and GraphQL API design, authentication, security, error handling, documentation, and deployment. Learn HTTP semantics, resource modeling, rate limiting, versioning strategies, and build production-grade APIs serving millions of requests.
sasmp_version: "1.3.0"
bonded_agent: 01-programming-fundamentals
bond_type: PRIMARY_BOND
---

# 🔌 API Design & Development

**Design, secure, and deploy APIs that power modern applications.**

Master REST and GraphQL APIs with production-grade security, documentation, and performance optimization.

## Quick Start: Building Your First API

### REST API with Express.js
```javascript
const express = require('express');
const app = express();

// Middleware
app.use(express.json());
app.use((req, res, next) => {
    console.log(`${req.method} ${req.path}`);
    next();
});

// Routes with proper resource naming
app.get('/api/v1/users', (req, res) => {
    // Query parameters for filtering
    const { status = 'active', limit = 10, offset = 0 } = req.query;

    res.json({
        data: [],
        pagination: { limit, offset, total: 0 },
        meta: { timestamp: new Date() }
    });
});

app.get('/api/v1/users/:id', (req, res) => {
    res.json({ id: req.params.id, name: 'John' });
});

app.post('/api/v1/users', (req, res) => {
    const { email, name } = req.body;

    if (!email || !name) {
        return res.status(400).json({
            error: 'VALIDATION_ERROR',
            message: 'Missing required fields',
            details: { required: ['email', 'name'] }
        });
    }

    res.status(201).json({
        id: 1,
        email,
        name,
        created_at: new Date()
    });
});

app.patch('/api/v1/users/:id', (req, res) => {
    res.json({ id: req.params.id, ...req.body, updated_at: new Date() });
});

app.delete('/api/v1/users/:id', (req, res) => {
    res.status(204).send();
});

// Error handling middleware
app.use((err, req, res, next) => {
    console.error(err);
    res.status(err.status || 500).json({
        error: err.code || 'INTERNAL_SERVER_ERROR',
        message: err.message
    });
});

app.listen(3000);
```

### GraphQL API with Apollo
```javascript
const { ApolloServer, gql } = require('apollo-server-express');

// Schema definition
const typeDefs = gql`
    type User {
        id: ID!
        email: String!
        name: String!
        posts: [Post!]!
        createdAt: DateTime!
    }

    type Post {
        id: ID!
        title: String!
        content: String!
        author: User!
        likes: Int!
        createdAt: DateTime!
    }

    type Query {
        user(id: ID!): User
        users(limit: Int = 10, offset: Int = 0): [User!]!
        post(id: ID!): Post
        searchPosts(query: String!): [Post!]!
    }

    type Mutation {
        createUser(email: String!, name: String!): User!
        updateUser(id: ID!, name: String): User
        deleteUser(id: ID!): Boolean!

        createPost(title: String!, content: String!, authorId: ID!): Post!
        updatePost(id: ID!, title: String, content: String): Post
        likePost(id: ID!): Post!
    }

    scalar DateTime
`;

// Resolvers
const resolvers = {
    Query: {
        user: async (parent, { id }) => {
            return { id, email: 'user@example.com', name: 'John' };
        },
        users: async (parent, { limit, offset }) => {
            return [];
        },
        searchPosts: async (parent, { query }, context) => {
            // Use DataLoader to prevent N+1 queries
            return [];
        }
    },
    Mutation: {
        createUser: async (parent, { email, name }, context) => {
            // Validation
            if (!email || !name) throw new Error('Missing required fields');
            return { id: '1', email, name, createdAt: new Date() };
        }
    },
    User: {
        posts: async (user, args, { loaders }) => {
            // Batch load posts for this user
            return loaders.postsByUserId.load(user.id);
        }
    }
};

const server = new ApolloServer({ typeDefs, resolvers });
```

### REST API Best Practices
```javascript
// 1. Resource-oriented URLs
GET    /api/v1/users              // List all
POST   /api/v1/users              // Create
GET    /api/v1/users/:id          // Get one
PATCH  /api/v1/users/:id          // Update
DELETE /api/v1/users/:id          // Delete

// 2. Proper HTTP status codes
200 OK              // Successful GET, PATCH, DELETE
201 Created         // Successful POST
204 No Content      // Successful DELETE with no body
400 Bad Request     // Invalid input
401 Unauthorized    // Missing authentication
403 Forbidden       // Lacks permission
404 Not Found       // Resource doesn't exist
409 Conflict        // Conflict (e.g., duplicate email)
422 Unprocessable   // Validation failed
429 Too Many Requests // Rate limited
500 Internal Error  // Server error

// 3. Consistent error responses
{
    "error": "VALIDATION_ERROR",
    "message": "Input validation failed",
    "details": {
        "email": "Invalid email format",
        "age": "Must be 18 or older"
    }
}

// 4. Pagination for list endpoints
{
    "data": [...],
    "pagination": {
        "limit": 20,
        "offset": 0,
        "total": 1000,
        "pages": 50
    }
}

// 5. Include metadata
{
    "data": {...},
    "meta": {
        "timestamp": "2024-01-15T10:30:00Z",
        "version": "1.0",
        "requestId": "req_abc123"
    }
}
```

## API Security

### Authentication Methods

#### API Keys
```javascript
// Simple API key authentication
app.use((req, res, next) => {
    const apiKey = req.headers['x-api-key'];

    if (!apiKey || !validateApiKey(apiKey)) {
        return res.status(401).json({ error: 'Invalid API key' });
    }

    req.user = getUserFromApiKey(apiKey);
    next();
});
```

#### JWT (JSON Web Tokens)
```javascript
const jwt = require('jsonwebtoken');

// Generate token
const token = jwt.sign(
    { userId: 123, email: 'user@example.com' },
    process.env.JWT_SECRET,
    { expiresIn: '24h' }
);

// Verify token
app.use((req, res, next) => {
    const token = req.headers.authorization?.replace('Bearer ', '');

    if (!token) {
        return res.status(401).json({ error: 'Missing token' });
    }

    try {
        const decoded = jwt.verify(token, process.env.JWT_SECRET);
        req.user = decoded;
        next();
    } catch (error) {
        res.status(401).json({ error: 'Invalid token' });
    }
});
```

#### OAuth 2.0
```javascript
// OAuth 2.0 flow for third-party authorization
app.get('/auth/google', (req, res) => {
    const redirectUrl = `https://accounts.google.com/o/oauth2/v2/auth?
        client_id=${process.env.GOOGLE_CLIENT_ID}&
        redirect_uri=${process.env.REDIRECT_URI}&
        response_type=code&
        scope=openid+email+profile`;

    res.redirect(redirectUrl);
});

app.get('/auth/callback', async (req, res) => {
    const { code } = req.query;

    // Exchange code for token
    const token = await exchangeCodeForToken(code);
    const user = await getUserInfo(token);

    res.json({ token: generateJWT(user) });
});
```

### Security Headers
```javascript
app.use((req, res, next) => {
    res.setHeader('Content-Security-Policy', "default-src 'self'");
    res.setHeader('X-Content-Type-Options', 'nosniff');
    res.setHeader('X-Frame-Options', 'DENY');
    res.setHeader('X-XSS-Protection', '1; mode=block');
    res.setHeader('Strict-Transport-Security', 'max-age=31536000');
    next();
});
```

### Rate Limiting
```javascript
const rateLimit = require('express-rate-limit');

const limiter = rateLimit({
    windowMs: 15 * 60 * 1000, // 15 minutes
    max: 100, // 100 requests per window
    message: 'Too many requests, please try again later',
    standardHeaders: true, // Return rate limit info in headers
    legacyHeaders: false
});

app.use('/api/', limiter);

// Per-user rate limiting
const perUserLimiter = rateLimit({
    keyGenerator: (req) => req.user.id,
    max: 50 // 50 per user
});

app.use('/api/expensive-operation', perUserLimiter);
```

### Input Validation
```javascript
const { body, validationResult } = require('express-validator');

app.post('/api/v1/users',
    body('email').isEmail().normalizeEmail(),
    body('name').trim().isLength({ min: 1, max: 100 }),
    body('age').optional().isInt({ min: 0, max: 150 }),
    (req, res) => {
        const errors = validationResult(req);
        if (!errors.isEmpty()) {
            return res.status(422).json({ errors: errors.array() });
        }

        // Process validated data
        res.json({ success: true });
    }
);
```

## API Versioning Strategies

### URL Path Versioning
```
GET /api/v1/users
GET /api/v2/users (different response format)
```

### Header Versioning
```javascript
app.get('/api/users', (req, res) => {
    const version = req.headers['api-version'] || '1';

    if (version === '2') {
        return res.json({ data: users, meta: {} });
    }

    res.json(users); // v1 format
});
```

### Media Type Versioning
```
GET /api/users
Accept: application/vnd.myapi.v2+json
```

## API Documentation

### OpenAPI/Swagger
```yaml
openapi: 3.0.0
info:
    title: User API
    version: 1.0.0
    description: User management API

servers:
    - url: https://api.example.com/v1

paths:
    /users:
        get:
            summary: List all users
            parameters:
                - name: limit
                  in: query
                  schema:
                    type: integer
                    default: 10
            responses:
                '200':
                    description: Success
                    content:
                        application/json:
                            schema:
                                type: array
                                items:
                                    $ref: '#/components/schemas/User'
        post:
            summary: Create new user
            requestBody:
                required: true
                content:
                    application/json:
                        schema:
                            $ref: '#/components/schemas/CreateUserRequest'
            responses:
                '201':
                    description: User created
                    content:
                        application/json:
                            schema:
                                $ref: '#/components/schemas/User'

components:
    schemas:
        User:
            type: object
            properties:
                id:
                    type: integer
                email:
                    type: string
                name:
                    type: string
                createdAt:
                    type: string
                    format: date-time
```

## Performance Optimization

### Caching Strategies
```javascript
// HTTP caching headers
app.get('/api/posts/:id', (req, res) => {
    res.setHeader('Cache-Control', 'public, max-age=3600'); // 1 hour
    res.json(post);
});

// ETags for conditional requests
const etag = crypto.createHash('md5').update(JSON.stringify(post)).digest('hex');
res.setHeader('ETag', etag);

if (req.headers['if-none-match'] === etag) {
    return res.status(304).send(); // Not Modified
}

res.json(post);

// Redis caching
const redis = require('redis');
const client = redis.createClient();

app.get('/api/posts/:id', async (req, res) => {
    const cached = await client.get(`post:${req.params.id}`);
    if (cached) return res.json(JSON.parse(cached));

    const post = await db.posts.findOne({ id: req.params.id });
    await client.setEx(`post:${req.params.id}`, 3600, JSON.stringify(post));

    res.json(post);
});
```

### Response Compression
```javascript
const compression = require('compression');
app.use(compression());
```

## Learning Path

### Phase 1: Foundation (Weeks 1-4)
- HTTP methods and status codes
- REST principles and conventions
- Basic Express/framework routing
- Simple endpoint creation
- Query parameters and filtering
- **Projects**:
  - Todo API with CRUD
  - Weather API wrapper
  - Simple blog API

### Phase 2: Intermediate (Weeks 5-12)
- Authentication (JWT, OAuth)
- Input validation and error handling
- Database integration
- Rate limiting and security
- API documentation (Swagger)
- Testing (unit and integration)
- **Projects**:
  - User management API with auth
  - E-commerce API with payments
  - Social network API

### Phase 3: Advanced (Weeks 13-24)
- GraphQL implementation
- Advanced caching strategies
- Performance optimization
- Monitoring and analytics
- Multi-tenant APIs
- Versioning strategies
- **Projects**:
  - Production API with 1M+ users
  - GraphQL server with complex schema
  - Microservices communication layer

## 18+ Production Projects

### Foundation
1. **Todo API** - Create, read, update, delete todos
2. **Weather API Wrapper** - Fetch and serve weather data
3. **Blog API** - Posts with CRUD operations
4. **Calculator API** - Mathematical operations

### Intermediate
5. **User Management API** - User registration, auth, profiles
6. **E-commerce API** - Products, shopping cart, orders
7. **Social Network API** - Users, posts, followers, likes
8. **Real-time Chat API** - Messages, conversations
9. **Task Management** - Projects, tasks, assignments, progress
10. **Analytics API** - Event tracking, aggregation

### Advanced
11. **Multi-tenant SaaS API** - Organizations, data isolation
12. **Payment Processing API** - Stripe/PayPal integration
13. **GraphQL Server** - Complex schema with resolvers
14. **Microservices Architecture** - Multiple coordinated services
15. **Real-time Notification Service** - WebSockets, subscriptions
16. **API Gateway** - Request routing, auth, rate limiting
17. **Data Pipeline API** - ETL operations
18. **Search API** - Elasticsearch integration

## Best Practices Checklist

### Design
- ✅ Use resource-oriented URLs
- ✅ Use proper HTTP methods and status codes
- ✅ Design for pagination and filtering
- ✅ Version your APIs
- ✅ Include request IDs for tracking

### Security
- ✅ Always use HTTPS/TLS
- ✅ Implement authentication
- ✅ Validate all inputs
- ✅ Use rate limiting
- ✅ Set security headers
- ✅ Log security events
- ✅ Never expose sensitive data

### Performance
- ✅ Use caching (HTTP, Redis)
- ✅ Compress responses
- ✅ Optimize database queries
- ✅ Use pagination for large datasets
- ✅ Monitor performance metrics

### Documentation
- ✅ Use OpenAPI/Swagger
- ✅ Include code examples
- ✅ Document error responses
- ✅ Provide getting started guide
- ✅ Keep documentation updated

## 🏆 Success Milestones

- [ ] Build REST API with full CRUD operations
- [ ] Implement JWT authentication
- [ ] Create comprehensive API documentation
- [ ] Achieve 95%+ test coverage
- [ ] Implement rate limiting and security
- [ ] Handle 1000+ requests per second
- [ ] Deploy API to production
- [ ] Monitor and optimize API performance

---

**Ready to build APIs? Start with a simple REST API and work up to GraphQL!**
