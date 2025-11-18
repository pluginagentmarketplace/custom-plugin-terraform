---
name: System Architecture & Design Patterns
description: Master 23 Gang of Four design patterns, architectural styles, and scalable system design. Learn microservices, event-driven architecture, distributed systems, and design systems handling millions of users with proper trade-off analysis.
---

# 🏗️ System Architecture & Design Patterns

**Design scalable systems and implement proven patterns.**

Master architectural patterns and system design principles for systems ranging from thousands to millions of users.

## Quick Start: Common Design Patterns

### Creational Patterns

#### Singleton Pattern
```python
class DatabaseConnection:
    _instance = None

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super().__new__(cls)
            cls._instance._connect()
        return cls._instance

    def _connect(self):
        print("Connecting to database...")

# Usage
db1 = DatabaseConnection()
db2 = DatabaseConnection()
assert db1 is db2  # Same instance
```

#### Factory Pattern
```python
class DataProcessor:
    def process(self, data):
        pass

class JsonProcessor(DataProcessor):
    def process(self, data):
        return json.loads(data)

class CsvProcessor(DataProcessor):
    def process(self, data):
        return list(csv.reader(data.splitlines()))

class ProcessorFactory:
    @staticmethod
    def create(file_type: str) -> DataProcessor:
        processors = {
            'json': JsonProcessor,
            'csv': CsvProcessor
        }
        processor_class = processors.get(file_type)
        if not processor_class:
            raise ValueError(f"Unknown type: {file_type}")
        return processor_class()

# Usage
factory = ProcessorFactory()
processor = factory.create('json')
```

#### Builder Pattern
```python
class QueryBuilder:
    def __init__(self):
        self.query = {}

    def select(self, *fields):
        self.query['select'] = fields
        return self

    def where(self, condition):
        self.query['where'] = condition
        return self

    def limit(self, n):
        self.query['limit'] = n
        return self

    def build(self):
        return self.query

# Usage
query = (QueryBuilder()
    .select('id', 'name', 'email')
    .where('age > 18')
    .limit(10)
    .build())
```

### Structural Patterns

#### Adapter Pattern
```python
# Old interface
class LegacyPaymentSystem:
    def pay(self, amount):
        print(f"Legacy payment: ${amount}")

# New interface we want
class PaymentProcessor:
    def process_payment(self, amount):
        pass

# Adapter
class PaymentAdapter(PaymentProcessor):
    def __init__(self, legacy_system):
        self.legacy = legacy_system

    def process_payment(self, amount):
        self.legacy.pay(amount)

# Usage
legacy = LegacyPaymentSystem()
adapter = PaymentAdapter(legacy)
adapter.process_payment(100)
```

#### Decorator Pattern
```python
def timing_decorator(func):
    def wrapper(*args, **kwargs):
        import time
        start = time.time()
        result = func(*args, **kwargs)
        elapsed = time.time() - start
        print(f"{func.__name__} took {elapsed:.2f}s")
        return result
    return wrapper

def logging_decorator(func):
    def wrapper(*args, **kwargs):
        print(f"Calling {func.__name__}")
        return func(*args, **kwargs)
    return wrapper

@timing_decorator
@logging_decorator
def fetch_data():
    import time
    time.sleep(1)
    return "data"

fetch_data()
# Output:
# Calling fetch_data
# fetch_data took 1.00s
```

#### Facade Pattern
```python
# Complex subsystems
class EmailService:
    def send(self, to, subject, body):
        print(f"Sending email to {to}")

class SmsService:
    def send(self, phone, message):
        print(f"Sending SMS to {phone}")

class NotificationService:
    def send_notification(self, channel, recipient, message):
        print(f"Sending {channel} notification to {recipient}")

# Facade simplifies interaction
class NotificationFacade:
    def __init__(self):
        self.email = EmailService()
        self.sms = SmsService()
        self.push = NotificationService()

    def notify_user(self, user, message):
        if user.prefers_email():
            self.email.send(user.email, "Notification", message)
        elif user.prefers_sms():
            self.sms.send(user.phone, message)
        else:
            self.push.send_notification('push', user.id, message)

# Usage
facade = NotificationFacade()
facade.notify_user(user, "Your order is ready")
```

### Behavioral Patterns

#### Observer Pattern
```python
class EventEmitter:
    def __init__(self):
        self.listeners = {}

    def on(self, event, callback):
        if event not in self.listeners:
            self.listeners[event] = []
        self.listeners[event].append(callback)

    def emit(self, event, data):
        if event in self.listeners:
            for callback in self.listeners[event]:
                callback(data)

# Usage
emitter = EventEmitter()

def on_user_created(user):
    print(f"Sending welcome email to {user['email']}")

def on_user_created_analytics(user):
    print(f"Tracking user creation: {user['id']}")

emitter.on('user:created', on_user_created)
emitter.on('user:created', on_user_created_analytics)

emitter.emit('user:created', {'id': 1, 'email': 'user@example.com'})
```

#### Strategy Pattern
```python
class PaymentStrategy:
    def pay(self, amount):
        pass

class CreditCardPayment(PaymentStrategy):
    def __init__(self, card_number):
        self.card_number = card_number

    def pay(self, amount):
        print(f"Charging ${amount} to card {self.card_number[-4:]}")

class PayPalPayment(PaymentStrategy):
    def __init__(self, email):
        self.email = email

    def pay(self, amount):
        print(f"Charging ${amount} via PayPal {self.email}")

class Order:
    def __init__(self, total, payment_strategy):
        self.total = total
        self.payment_strategy = payment_strategy

    def checkout(self):
        self.payment_strategy.pay(self.total)

# Usage
credit_card = CreditCardPayment("1234-5678-9012-3456")
order = Order(99.99, credit_card)
order.checkout()
```

#### Chain of Responsibility Pattern
```python
class LogLevel:
    ERROR = 3
    WARNING = 2
    INFO = 1

class Logger:
    def __init__(self, name, level):
        self.name = name
        self.level = level
        self.next_logger = None

    def set_next(self, logger):
        self.next_logger = logger
        return logger

    def log(self, message, level):
        if self.level <= level:
            self._write_log(message)
        if self.next_logger:
            self.next_logger.log(message, level)

    def _write_log(self, message):
        print(f"[{self.name}] {message}")

# Usage
logger1 = Logger("INFO", LogLevel.INFO)
logger2 = logger1.set_next(Logger("WARNING", LogLevel.WARNING))
logger2.set_next(Logger("ERROR", LogLevel.ERROR))

logger1.log("Informational message", LogLevel.INFO)
logger1.log("Warning message", LogLevel.WARNING)
```

## Architectural Styles

### Monolithic Architecture
```
User Requests
    ↓
    Load Balancer
    ↓
Web Server (Single Codebase)
├── User Service
├── Order Service
├── Payment Service
└── Notification Service
    ↓
Database
```

**Advantages**: Simple deployment, easy debugging, good performance
**Disadvantages**: Hard to scale individual services, tightly coupled

### Microservices Architecture
```
User Requests
    ↓
    API Gateway
    ├── User Service (Port 3001)
    ├── Order Service (Port 3002)
    ├── Payment Service (Port 3003)
    └── Notification Service (Port 3004)
    ↓
Service Registry → Load Balancing
    ↓
Databases (Per-service)
```

**Advantages**: Independent scaling, technology flexibility, deployment independence
**Disadvantages**: Complex debugging, distributed transactions, network overhead

### Event-Driven Architecture
```
Services communicate via events:

User Service publishes "user:created" event
    ↓
Message Broker (RabbitMQ, Kafka)
    ↓
Subscribed Services:
├── Email Service (sends welcome email)
├── Analytics Service (tracks signup)
└── Notification Service (creates user notification)
```

### Serverless Architecture
```
AWS Lambda Functions:
- POST /users → CreateUserFunction
- GET /users/:id → GetUserFunction
- DELETE /users/:id → DeleteUserFunction

Benefits:
- No server management
- Auto-scaling
- Pay per execution
```

## Large-Scale System Design

### Twitter-Scale Architecture
**Requirements**: 300M users, 500M tweets/day, 1000 requests/sec

```
┌─────────────────┐
│  Load Balancer  │
└────────┬────────┘
         │
┌────────┴─────────┬─────────────┐
│                  │             │
▼                  ▼             ▼
Tweet Service  User Service  Timeline Service
│              │              │
└─────────────┬─────────────┘
              │
         ┌────┴─────┐
         │           │
         ▼           ▼
    Master DB   Replica DBs
    (Write)     (Read)
         │
    ┌────┴─────┐
    │           │
    ▼           ▼
  Cache      Search Index
(Redis)     (Elasticsearch)

Timeline Caching:
- Fetch user's followers
- Get latest tweets from each follower
- Sort and paginate
- Cache result
```

### Database Scaling Patterns
```
Vertical Scaling:
Server 1 (4GB RAM) → Server 1 (64GB RAM)

Horizontal Scaling - Replication:
Master (Write)
├── Read Replica 1
├── Read Replica 2
└── Read Replica 3

Horizontal Scaling - Sharding:
Hash(user_id) % 4 = Shard
├── Shard 1: Users 1-250M (DB1)
├── Shard 2: Users 250M-500M (DB2)
├── Shard 3: Users 500M-750M (DB3)
└── Shard 4: Users 750M-1B (DB4)
```

## Learning Path

### Phase 1: Foundation (Weeks 1-4)
- Design patterns basics
- SOLID principles
- Component design
- Simple system design
- **Projects**:
  - Design a simple blog system
  - Build todo app with patterns
  - Design a cache implementation

### Phase 2: Intermediate (Weeks 5-12)
- Advanced design patterns
- Microservices architecture
- Event-driven systems
- Database scaling basics
- **Projects**:
  - Design medium-scale system (100K users)
  - Multi-tier application
  - Event-driven task system

### Phase 3: Advanced (Weeks 13-24)
- Large-scale system design
- Distributed systems
- Trade-off analysis
- Advanced scaling patterns
- **Projects**:
  - Design Twitter-scale system (300M users)
  - Design Netflix-scale system (200M users)
  - Design Uber-scale system (real-time matching)

## 16+ Production Projects

### Design Pattern Projects
1. **Payment System** - Multiple payment strategies
2. **Logger System** - Chain of responsibility
3. **Observer-based Event System** - Publish-subscribe
4. **Cache Implementation** - Singleton with eviction
5. **Factory for Database Adapters** - Multiple DB types

### Architectural Projects
6. **Monolithic Blog Platform** - Single codebase
7. **Microservices E-commerce** - Multiple services
8. **Event-driven Notification System** - Async events
9. **Real-time Analytics** - Stream processing
10. **Distributed Cache** - Redis cluster

### System Design Projects
11. **Twitter Clone** - 1M users, 10K tweets/sec
12. **Netflix Clone** - Video streaming architecture
13. **Uber Clone** - Real-time matching and tracking
14. **YouTube Clone** - Video upload, transcoding, streaming
15. **Slack Clone** - Real-time messaging, channels
16. **Airbnb Clone** - Search, booking, payment

## SOLID Principles

### Single Responsibility Principle
```python
# Bad: Class has multiple responsibilities
class User:
    def save_to_db(self):
        pass

    def send_email(self):
        pass

    def generate_report(self):
        pass

# Good: Each class has one responsibility
class User:
    pass

class UserRepository:
    def save(self, user):
        pass

class EmailService:
    def send_email(self, user):
        pass
```

### Open/Closed Principle
```python
# Closed for modification, open for extension
class PaymentProcessor:
    def process(self, payment):
        pass

class StripeProcessor(PaymentProcessor):
    def process(self, payment):
        # Stripe-specific logic
        pass

class PayPalProcessor(PaymentProcessor):
    def process(self, payment):
        # PayPal-specific logic
        pass
```

## 🏆 Success Milestones

- [ ] Implement all 23 Gang of Four patterns
- [ ] Design monolithic system with 100K users
- [ ] Design microservices architecture
- [ ] Design event-driven system
- [ ] Design database sharding strategy
- [ ] Design Twitter-scale system
- [ ] Analyze architectural trade-offs
- [ ] Mentor others on system design

---

**Ready to design systems? Start with design patterns and work up to large-scale architecture!**
