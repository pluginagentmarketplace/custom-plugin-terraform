---
name: Testing, Security & Observability
description: Master testing strategies (unit, integration, E2E), security best practices, OWASP Top 10, secure coding, and system monitoring. Learn to write secure code, comprehensive tests, and build resilient systems with proper incident response and compliance.
---

# 🔒 Testing, Security & Observability

**Build secure, reliable systems with comprehensive testing and monitoring.**

Master testing strategies, security hardening, monitoring, and incident response to ensure code quality and user safety.

## Quick Start: Testing Fundamentals

### Unit Testing with Jest
```javascript
// calculator.js
export const add = (a, b) => a + b;
export const subtract = (a, b) => a - b;
export const multiply = (a, b) => a * b;
export const divide = (a, b) => {
    if (b === 0) throw new Error('Division by zero');
    return a / b;
};

// calculator.test.js
import { add, subtract, multiply, divide } from './calculator';

describe('Calculator', () => {
    describe('add', () => {
        it('should add two positive numbers', () => {
            expect(add(2, 3)).toBe(5);
        });

        it('should add negative numbers', () => {
            expect(add(-2, -3)).toBe(-5);
        });

        it('should handle zero', () => {
            expect(add(0, 5)).toBe(5);
        });
    });

    describe('divide', () => {
        it('should divide two numbers', () => {
            expect(divide(10, 2)).toBe(5);
        });

        it('should throw error for division by zero', () => {
            expect(() => divide(10, 0)).toThrow('Division by zero');
        });
    });
});

// Run tests
// npm test
```

### Integration Testing
```javascript
// api.test.js
const request = require('supertest');
const app = require('../app');
const db = require('../db');

describe('User API', () => {
    beforeAll(async () => {
        await db.initialize();
    });

    afterEach(async () => {
        await db.clearUsers();
    });

    afterAll(async () => {
        await db.close();
    });

    describe('POST /api/users', () => {
        it('should create a new user', async () => {
            const response = await request(app)
                .post('/api/users')
                .send({ email: 'user@example.com', name: 'John' })
                .expect(201);

            expect(response.body).toHaveProperty('id');
            expect(response.body.email).toBe('user@example.com');
        });

        it('should reject duplicate email', async () => {
            await request(app)
                .post('/api/users')
                .send({ email: 'user@example.com', name: 'John' });

            const response = await request(app)
                .post('/api/users')
                .send({ email: 'user@example.com', name: 'Jane' })
                .expect(409);

            expect(response.body.error).toBe('DUPLICATE_EMAIL');
        });
    });

    describe('GET /api/users/:id', () => {
        it('should retrieve user by id', async () => {
            const createRes = await request(app)
                .post('/api/users')
                .send({ email: 'user@example.com', name: 'John' });

            const response = await request(app)
                .get(`/api/users/${createRes.body.id}`)
                .expect(200);

            expect(response.body.email).toBe('user@example.com');
        });

        it('should return 404 for non-existent user', async () => {
            await request(app)
                .get('/api/users/999')
                .expect(404);
        });
    });
});
```

### Python Testing with Pytest
```python
# test_user_service.py
import pytest
from user_service import UserService, InvalidEmailError

class TestUserService:
    @pytest.fixture
    def service(self):
        """Create service instance for each test"""
        return UserService(in_memory=True)

    def test_create_user(self, service):
        user = service.create_user('user@example.com', 'John')
        assert user.email == 'user@example.com'
        assert user.name == 'John'

    def test_invalid_email(self, service):
        with pytest.raises(InvalidEmailError):
            service.create_user('invalid-email', 'John')

    def test_duplicate_email(self, service):
        service.create_user('user@example.com', 'John')
        with pytest.raises(ValueError, match='Email already exists'):
            service.create_user('user@example.com', 'Jane')

    @pytest.mark.parametrize('email,valid', [
        ('user@example.com', True),
        ('invalid.email', False),
        ('test@domain.co.uk', True),
    ])
    def test_email_validation(self, service, email, valid):
        if valid:
            user = service.create_user(email, 'John')
            assert user.email == email
        else:
            with pytest.raises(InvalidEmailError):
                service.create_user(email, 'John')
```

## Security Best Practices

### Input Validation
```python
from email_validator import validate_email, EmailNotValidError
import re

def validate_user_input(email, name, age):
    """Validate user input"""
    errors = {}

    # Email validation
    try:
        valid = validate_email(email)
        email = valid.email
    except EmailNotValidError as e:
        errors['email'] = str(e)

    # Name validation
    if not name or len(name) < 2 or len(name) > 100:
        errors['name'] = 'Name must be 2-100 characters'

    if not re.match(r'^[a-zA-Z\s-]+$', name):
        errors['name'] = 'Name contains invalid characters'

    # Age validation
    try:
        age = int(age)
        if age < 0 or age > 150:
            errors['age'] = 'Age must be 0-150'
    except ValueError:
        errors['age'] = 'Age must be a number'

    if errors:
        raise ValueError(errors)

    return email, name, age
```

### Password Security
```python
from argon2 import PasswordHasher
from argon2.exceptions import VerifyMismatchError
import secrets
import string

# Hash passwords securely (never store plain text)
hasher = PasswordHasher()

def hash_password(password: str) -> str:
    """Hash password with Argon2"""
    if len(password) < 12:
        raise ValueError('Password must be at least 12 characters')

    return hasher.hash(password)

def verify_password(password: str, hash: str) -> bool:
    """Verify password against hash"""
    try:
        hasher.verify(hash, password)
        return True
    except VerifyMismatchError:
        return False

def generate_secure_token(length: int = 32) -> str:
    """Generate cryptographically secure token"""
    alphabet = string.ascii_letters + string.digits
    return ''.join(secrets.choice(alphabet) for _ in range(length))
```

### SQL Injection Prevention
```python
# Bad: SQL injection vulnerability
def get_user_bad(user_id):
    query = f"SELECT * FROM users WHERE id = {user_id}"
    return db.execute(query)

# Good: Parameterized queries (safe)
def get_user_good(user_id):
    query = "SELECT * FROM users WHERE id = ?"
    return db.execute(query, (user_id,))

# SQLAlchemy ORM (prevents injection)
from sqlalchemy import select
from models import User

def get_user_orm(user_id):
    stmt = select(User).where(User.id == user_id)
    return db.session.execute(stmt).scalar_one_or_none()
```

### XSS Prevention
```javascript
// Bad: XSS vulnerability
const userName = getUserInput();
document.getElementById('greeting').innerHTML = `Hello ${userName}`; // Unsafe!

// Good: Use textContent instead of innerHTML
document.getElementById('greeting').textContent = `Hello ${userName}`; // Safe

// Or sanitize HTML
const DOMPurify = require('dompurify');
const sanitized = DOMPurify.sanitize(userInput);
document.getElementById('greeting').innerHTML = `Hello ${sanitized}`;

// React automatically escapes by default
const greeting = <div>Hello {userName}</div>; // Safe by default
```

### CORS Configuration
```python
from flask import Flask
from flask_cors import CORS

app = Flask(__name__)

# Allow specific origins only
CORS(app, resources={
    r"/api/*": {
        "origins": ["https://example.com", "https://app.example.com"],
        "methods": ["GET", "POST", "PUT", "DELETE"],
        "allow_headers": ["Content-Type", "Authorization"],
        "supports_credentials": True
    }
})

# Or per-route
@app.route('/api/public')
@cross_origin(origins="*")
def public():
    return {"data": "public"}

@app.route('/api/private')
@cross_origin(origins=["https://example.com"])
def private():
    return {"data": "private"}
```

## OWASP Top 10

### 1. Injection (SQL, NoSQL, Command)
```python
# Vulnerable
os.system(f"ls {user_directory}")  # Command injection!

# Safe
import subprocess
subprocess.run(["ls", user_directory], check=True)
```

### 2. Broken Authentication
```python
# Bad: Weak session management
session_id = str(user_id)  # Predictable!

# Good: Secure tokens
session_id = secrets.token_urlsafe(32)

# Use established auth libraries
from flask_login import LoginManager, UserMixin
from werkzeug.security import generate_password_hash, check_password_hash
```

### 3. Sensitive Data Exposure
```python
# Bad: Logging sensitive data
logger.info(f"User {user} logged in with password {password}")

# Good: Log safely
logger.info(f"User {user} logged in")  # Never log passwords

# Encrypt sensitive data
from cryptography.fernet import Fernet
cipher = Fernet(key)
encrypted = cipher.encrypt(sensitive_data.encode())
decrypted = cipher.decrypt(encrypted).decode()
```

### 4. XML External Entities (XXE)
```python
# Vulnerable
import xml.etree.ElementTree as ET
tree = ET.parse(user_uploaded_file)  # XXE vulnerability!

# Safe
from defusedxml import ElementTree
tree = ElementTree.parse(user_uploaded_file)  # Protected
```

### 5. Broken Access Control
```python
# Bad: User can access any user's data
@app.route('/api/users/<user_id>')
def get_user(user_id):
    return User.query.get(user_id)

# Good: Verify ownership
@app.route('/api/users/<user_id>')
@login_required
def get_user(user_id):
    user = User.query.get(user_id)
    if user.id != current_user.id:
        abort(403)  # Forbidden
    return user
```

## Monitoring & Logging

### Structured Logging
```javascript
const winston = require('winston');

const logger = winston.createLogger({
    format: winston.format.json(),
    transports: [
        new winston.transports.File({ filename: 'error.log', level: 'error' }),
        new winston.transports.File({ filename: 'combined.log' })
    ]
});

// Good: Structured logs
logger.info('User login', {
    userId: 123,
    email: 'user@example.com',
    timestamp: new Date(),
    ipAddress: req.ip
});

// Bad: Unstructured
logger.info('User login: user@example.com');
```

### Key Metrics to Monitor
```
- Error rates (% of failed requests)
- Response times (p50, p95, p99)
- Failed authentication attempts
- Unusual access patterns
- Database query performance
- API rate limit violations
- Security events (failed logins, permission denials)
```

## Learning Path

### Phase 1: Foundation (Weeks 1-4)
- Testing basics (unit tests)
- OWASP Top 10 overview
- Secure coding fundamentals
- Monitoring basics
- **Projects**:
  - Add unit tests to project
  - Fix OWASP vulnerabilities
  - Set up basic monitoring

### Phase 2: Intermediate (Weeks 5-12)
- Integration and E2E testing
- Advanced security practices
- Penetration testing basics
- Monitoring setup
- **Projects**:
  - Comprehensive test suite
  - Secure API implementation
  - Monitoring dashboard

### Phase 3: Advanced (Weeks 13-24)
- Security architecture
- Incident response procedures
- Compliance (GDPR, HIPAA)
- Advanced monitoring
- **Projects**:
  - Production security audit
  - Incident response plan
  - Compliance implementation

## 18+ Production Projects

### Testing Projects
1. **Unit Test Suite** - 80%+ coverage
2. **Integration Tests** - API endpoints
3. **E2E Tests** - User workflows
4. **Load Testing** - Performance under stress
5. **Security Testing** - Vulnerability scanning

### Security Projects
6. **Secure Authentication** - JWT, OAuth
7. **Input Validation** - OWASP rules
8. **SQL Injection Prevention** - Parameterized queries
9. **XSS Prevention** - HTML sanitization
10. **CSRF Protection** - Token validation

### Monitoring Projects
11. **Log Aggregation** - ELK Stack
12. **Metrics Dashboard** - Prometheus + Grafana
13. **Alerting System** - Notify on anomalies
14. **Error Tracking** - Sentry integration
15. **APM System** - Request tracing

### Compliance Projects
16. **GDPR Compliance** - Data protection
17. **HIPAA Compliance** - Healthcare data
18. **Audit Logging** - Compliance tracking

## Best Practices Checklist

### Testing
- ✅ Unit tests for all functions
- ✅ Integration tests for APIs
- ✅ E2E tests for user workflows
- ✅ Test error scenarios
- ✅ Aim for 80%+ coverage

### Security
- ✅ Use HTTPS always
- ✅ Validate all inputs
- ✅ Hash passwords (Argon2, bcrypt)
- ✅ Use parameterized queries
- ✅ Implement rate limiting
- ✅ Log security events
- ✅ Regular security audits

### Monitoring
- ✅ Log all important events
- ✅ Track key metrics
- ✅ Set up alerting
- ✅ Monitor error rates
- ✅ Track performance
- ✅ Analyze logs for anomalies

## 🏆 Success Milestones

- [ ] Write comprehensive unit tests (80%+ coverage)
- [ ] Secure API against OWASP Top 10
- [ ] Implement JWT authentication
- [ ] Set up monitoring and alerting
- [ ] Perform security audit
- [ ] Pass penetration test
- [ ] Implement disaster recovery
- [ ] Achieve compliance certifications

---

**Ready to build secure systems? Start with comprehensive testing and work your way to security audits!**
