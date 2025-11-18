---
name: python-development
description: Master Python for backend development with Django, FastAPI, or Flask. Learn data processing, API development, automation, and machine learning integration.
---

# Python Development

## Quick Start

```python
# FastAPI - Modern Python web framework
from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()

class User(BaseModel):
    name: str
    email: str

@app.post("/users")
async def create_user(user: User):
    return {"id": 1, **user.dict()}

@app.get("/users/{user_id}")
async def get_user(user_id: int):
    return {"id": user_id, "name": "John"}
```

## Core Competencies

### Web Frameworks
- **Django** - Full-featured, batteries-included framework
- **FastAPI** - Modern, fast, with async support
- **Flask** - Lightweight, flexible microframework
- **Bottle** - Micro-framework for simple APIs

### ORM & Databases
- SQLAlchemy for SQL databases
- Django ORM for database operations
- Alembic for database migrations
- PostgreSQL, MySQL, SQLite

### Async & Concurrency
- Async/await syntax
- Asyncio event loop
- Concurrent.futures for threading
- Multiprocessing for CPU-bound tasks

### Data Processing
- NumPy for numerical computing
- Pandas for data analysis
- Data validation and cleaning
- ETL pipeline development

### API Development
- RESTful API design
- GraphQL with Graphene or Strawberry
- API documentation (OpenAPI/Swagger)
- Authentication (JWT, OAuth)
- Rate limiting and caching

### Testing & Quality
- pytest for unit testing
- Fixtures and parametrization
- Coverage analysis
- Integration testing
- Code quality tools (Black, Flake8, Mypy)

## Specialization Areas

- **Web Backend** → Django/FastAPI expert
- **Data Engineering** → Pandas, NumPy, ETL
- **DevOps** → Automation scripts
- **Machine Learning** → TensorFlow, Scikit-learn integration
- **System Administration** → Infrastructure automation

## Project Ideas

- RESTful API backend
- Data pipeline and analytics
- Admin dashboard with Django
- Real-time API with FastAPI
- Web scraping automation

## Ecosystem Tools

- **Poetry/Pip** - Package management
- **Virtual environments** - Project isolation
- **Docker** - Containerization
- **pytest** - Testing framework
- **Uvicorn** - ASGI server
