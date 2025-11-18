---
name: flutter-cross-platform
description: Master Flutter framework for building beautiful, performant apps across iOS, Android, web, and desktop from a single codebase. Learn Dart, widgets, state management, and deployment.
---

# Flutter Cross-Platform

## Quick Start

```dart
import 'package:flutter/material.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: HomePage(),
    );
  }
}

class HomePage extends StatefulWidget {
  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  int _counter = 0;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Counter')),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text('Count: $_counter'),
            ElevatedButton(
              onPressed: () => setState(() => _counter++),
              child: const Text('Increment'),
            ),
          ],
        ),
      ),
    );
  }
}
```

## Core Competencies

### Dart Language
- Type system and null safety
- Classes and inheritance
- Async/await for async operations
- Extension methods
- Mixins for code reuse

### Flutter Fundamentals
- Widgets and widget tree
- Stateful vs stateless widgets
- Material Design and Cupertino (iOS)
- Layout widgets (Row, Column, Stack)
- Navigation and routing

### State Management
- **Provider** - Recommended by Google
- **Riverpod** - Modern alternative
- **GetX** - Full-featured solution
- **BLoC** - Business logic components
- **MobX** - Reactive state

### UI Development
- Material Design components
- Cupertino (iOS) components
- Custom widgets
- Animations and transitions
- Responsive design for multiple screens

### Platform Channels
- Calling native iOS code (Swift)
- Calling native Android code (Kotlin)
- Bi-directional communication
- Plugin development

### Networking & Data
- HTTP client for APIs
- JSON serialization with Freezed
- Firebase integration
- SQLite for local storage
- Shared preferences

### Performance Optimization
- Build optimization
- Memory management
- Rendering optimization
- Profiling tools
- Performance monitoring

## Platform Support

- **Mobile** - iOS and Android (primary)
- **Web** - Flutter Web (growing)
- **Desktop** - Windows, macOS, Linux
- **Embedded** - Fuchsia support

## Testing Strategies

- Unit testing with test package
- Widget testing for UI
- Integration testing
- Test coverage analysis

## Development Tools

- **Flutter SDK** - Framework and tools
- **Android Studio/VS Code** - IDEs
- **DevTools** - Performance debugging
- **Emulators** - Testing environments

## Deployment

- iOS App Store (TestFlight beta)
- Google Play Store
- Web deployment
- Desktop distribution
- Versioning and releases

## Ecosystem & Packages

- **pub.dev** - Package repository
- Popular packages (http, dio, provider, get)
- Custom plugin development
- Open-source contributions
