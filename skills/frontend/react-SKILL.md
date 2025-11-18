---
name: react-ecosystem
description: Master React 18+ with Hooks, state management, server components, and production-grade patterns. Build scalable applications with Next.js, TypeScript, and modern tooling.
---

# React Ecosystem Mastery

## Quick Start

```jsx
// Modern React with Hooks
import { useState, useEffect } from 'react';

export function Counter() {
  const [count, setCount] = useState(0);

  useEffect(() => {
    document.title = `Count: ${count}`;
  }, [count]);

  return (
    <button onClick={() => setCount(c => c + 1)}>
      Click me: {count}
    </button>
  );
}
```

## Core Competencies

### React Fundamentals
- Functional components and Hooks (useState, useEffect, useContext)
- Props, state, and component lifecycle
- Event handling and form management
- Conditional rendering and lists
- Custom hooks for logic reuse
- Performance optimization (memo, useCallback, useMemo)

### Advanced Patterns
- Server Components and Suspense
- Error boundaries for error handling
- Render optimization techniques
- Code splitting and lazy loading
- Provider patterns for data sharing
- Controlled vs uncontrolled components

### Ecosystem Integration
- **State Management**: Redux, Zustand, Jotai, Context API
- **Data Fetching**: TanStack Query, SWR, Axios
- **Routing**: React Router v6+, TanStack Router
- **Styling**: Styled Components, Tailwind CSS, CSS Modules
- **Forms**: React Hook Form, Formik, Zod validation
- **Testing**: Jest, React Testing Library, Vitest

### Next.js Framework
- File-based routing and API routes
- SSR, SSG, and ISR strategies
- Image optimization and Web Vitals
- Middleware and server actions
- Database integration patterns
- Deployment strategies (Vercel, AWS, self-hosted)

## Learning Path

1. Master React Hooks and component patterns
2. Implement state management solution
3. Build data fetching patterns with TanStack Query
4. Learn Next.js for full-stack development
5. Master performance optimization and deployment

## Real-World Projects

- E-commerce platform with Next.js
- Real-time chat application
- Analytics dashboard
- Content management system
- Social media application

## Essential Tools

- **Vite/Create React App** - Development environment
- **React DevTools** - Browser extension for debugging
- **Storybook** - Component development and documentation
- **Jest & React Testing Library** - Testing
- **TypeScript** - Type safety
