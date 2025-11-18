---
name: mobile-development
description: Master mobile app development fundamentals for iOS, Android, and cross-platform frameworks. Learn native development, mobile UI/UX, performance, and app deployment.
---

# Mobile Development

## Quick Start - React Native

```javascript
import React, { useState } from 'react';
import { View, Text, TouchableOpacity, StyleSheet } from 'react-native';

export default function Counter() {
  const [count, setCount] = useState(0);

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Count: {count}</Text>
      <TouchableOpacity
        style={styles.button}
        onPress={() => setCount(count + 1)}
      >
        <Text style={styles.buttonText}>Increment</Text>
      </TouchableOpacity>
    </View>
  );
}

const styles = StyleSheet.create({
  container: { flex: 1, justifyContent: 'center', alignItems: 'center' },
  title: { fontSize: 24, fontWeight: 'bold' },
  button: { backgroundColor: '#007AFF', padding: 10, borderRadius: 5 },
  buttonText: { color: '#fff', fontSize: 16 }
});
```

## Core Concepts

### Mobile Platforms
- **iOS** - Apple ecosystem, Swift/Objective-C
- **Android** - Google ecosystem, Kotlin/Java
- **Web** - Cross-platform option
- **Desktop** - Windows, macOS, Linux

### Native Development
- Platform-specific SDKs and APIs
- Native performance and features
- Access to device capabilities
- Platform guidelines and best practices

### Cross-Platform Frameworks
- **React Native** - JavaScript/TypeScript
- **Flutter** - Dart language
- **Xamarin** - C# and .NET
- Trade-offs between native and cross-platform

### Mobile UI/UX
- Material Design (Android)
- iOS Human Interface Guidelines
- Responsive layouts
- Touch interactions
- Navigation patterns

### Performance Optimization
- App startup time
- Battery efficiency
- Memory management
- Frame rate and animations
- Network optimization

### Data Persistence
- Local databases (SQLite, Realm)
- File storage
- Caching strategies
- Offline-first approach

## App Lifecycle

- Initialization and setup
- Background/foreground transitions
- Memory management
- App suspension and termination

## Testing Strategies

- Unit testing (logic)
- Widget/Component testing
- Integration testing
- End-to-end testing
- Manual testing

## Distribution & Deployment

- iOS App Store requirements
- Google Play Store requirements
- Beta testing (TestFlight, Google Play Beta)
- Release management
- Versioning strategy

## Essential Tools

- **Xcode** (iOS)
- **Android Studio** (Android)
- **Visual Studio Code**
- **Device Simulators/Emulators**
- **Firebase** (analytics, push notifications)
