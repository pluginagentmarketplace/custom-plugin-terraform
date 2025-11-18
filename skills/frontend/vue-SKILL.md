---
name: vue-ecosystem
description: Master Vue 3 with Composition API, Pinia state management, and Nuxt framework. Build reactive, performant applications with elegant syntax and excellent developer experience.
---

# Vue Ecosystem Mastery

## Quick Start

```vue
<script setup>
import { ref, computed } from 'vue';

const count = ref(0);
const doubled = computed(() => count.value * 2);
</script>

<template>
  <div>
    <button @click="count++">Count: {{ count }}</button>
    <p>Doubled: {{ doubled }}</p>
  </div>
</template>

<style scoped>
button { padding: 8px 16px; }
</style>
```

## Core Competencies

### Vue 3 Fundamentals
- Template syntax and directives (v-if, v-for, v-bind)
- Composition API and Options API
- Reactive data binding with ref and reactive
- Computed properties and watchers
- Component lifecycle hooks
- Props, emits, and slots

### Advanced Features
- Provide/Inject for dependency injection
- Composables for logic reuse
- Teleport for DOM manipulation
- Transitions and animations
- Async components and code splitting
- Error handling with error boundaries

### State Management with Pinia
- Store definition and usage
- Actions, getters, and mutations
- Plugins and middleware
- DevTools integration
- TypeScript support

### Nuxt Framework (Full-Stack Vue)
- File-based routing and layouts
- Server-side rendering and static generation
- API routes and middleware
- Auto-imports and components
- Module ecosystem
- Image and font optimization

### Ecosystem Tools
- **Vue Router** - Official routing solution
- **Pinia** - Modern state management
- **Vite** - Fast development server
- **Vitest & Vue Test Utils** - Testing
- **Tailwind CSS** - Styling
- **TypeScript** - Type safety

## Learning Path

1. Master Vue 3 fundamentals and Composition API
2. Learn Pinia for state management
3. Explore Vue Router for complex applications
4. Graduate to Nuxt for full-stack development
5. Optimize performance and deployment

## Project Ideas

- Progressive Todo application
- Real-time collaboration tool
- Content management system
- Admin dashboard
- Full-stack blog with Nuxt

## Community & Resources

- Official Vue documentation (excellent!)
- Awesome Vue ecosystem
- Vue Mastery courses
- Community Discord (vibrant community)
- Official plugins and integrations
