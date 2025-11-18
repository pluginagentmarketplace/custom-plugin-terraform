---
name: frontend-framework-fundamentals
description: Master HTML5, CSS3, JavaScript, and responsive design fundamentals. Learn DOM manipulation, CSS layouts, and build interactive web applications from the ground up.
---

# Frontend Framework Fundamentals

## Quick Start

Frontend development begins with three core technologies:

```html
<!-- HTML5 provides structure -->
<nav class="header">
  <h1>My App</h1>
</nav>

<!-- CSS provides styling and layout -->
<style>
  .header { display: flex; justify-content: space-between; }
  @media (max-width: 768px) { .header { flex-direction: column; } }
</style>

<!-- JavaScript provides interactivity -->
<script>
  document.querySelector('.header').addEventListener('click', () => {
    console.log('Responsive and interactive!');
  });
</script>
```

## Core Concepts

### HTML5
- Semantic markup (header, nav, main, section, article, footer)
- Form handling and validation
- Accessibility attributes (aria-label, role)
- Web APIs (LocalStorage, Geolocation, Service Workers)

### CSS3 & Layout
- **Flexbox** - One-dimensional layouts, alignment, distribution
- **CSS Grid** - Two-dimensional layouts, complex designs
- **Responsive Design** - Media queries, mobile-first approach
- **Animations** - Transitions, keyframes, performance
- **Preprocessing** - SCSS/SASS for maintainability

### JavaScript Fundamentals
- ES6+ syntax (const/let, arrow functions, destructuring, spread operator)
- DOM Manipulation (querySelector, addEventListener, textContent)
- Event Handling (click, input, scroll, resize)
- Async JavaScript (Promises, async/await, fetch API)
- Array & Object Methods (map, filter, reduce, Object.entries)

## Learning Resources

1. **Interactive Practice**: Build projects immediately (portfolio, todo app, weather app)
2. **Accessibility First**: Learn WCAG 2.1 standards from day one
3. **Responsive Grid**: Master CSS Grid and Flexbox through projects
4. **API Integration**: Practice with real APIs (JSON Placeholder, OpenWeather)
5. **Performance Basics**: Understand critical rendering path, lazy loading

## Practice Projects

- Personal portfolio website with responsive design
- Todo application with local storage
- Weather app with API integration
- E-commerce product page with filters
- Interactive data visualization dashboard

## Tools & Resources

- **Browser DevTools** - Chrome/Firefox developer tools for debugging
- **MDN Web Docs** - Comprehensive documentation for all web technologies
- **Can I Use** - Browser compatibility checking
- **Lighthouse** - Performance and accessibility auditing
