---
name: react-native-cross-platform
description: Master React Native for building iOS and Android apps with JavaScript/TypeScript. Learn navigation, state management, native modules, and performance optimization.
---

# React Native Cross-Platform

## Quick Start - Navigation

```javascript
import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';

const Stack = createNativeStackNavigator();
const Tab = createBottomTabNavigator();

export default function App() {
  return (
    <NavigationContainer>
      <Tab.Navigator>
        <Tab.Screen name="Home" component={HomeScreen} />
        <Tab.Screen name="Profile" component={ProfileScreen} />
      </Tab.Navigator>
    </NavigationContainer>
  );
}
```

## Core Competencies

### React Native Basics
- Components and APIs
- Platform-specific code
- Styling with StyleSheet
- Flexbox layout system
- State and props

### Navigation Solutions
- **React Navigation** - Most popular
- **Native Navigation** - Platform-native feel
- Stack, tab, and drawer navigation
- Deep linking and navigation params

### State Management
- Redux for complex state
- Context API for simpler cases
- Zustand for lightweight alternative
- MobX for reactive state

### Native Modules
- Bridging JavaScript and native code
- Accessing device APIs
- Creating custom modules
- Third-party native libraries

### Networking & Data
- Fetch API for HTTP requests
- AsyncStorage for local data
- SQLite for local database
- State management patterns

### Performance Optimization
- FlatList vs ScrollView
- Code splitting and lazy loading
- Performance monitoring
- Memory leak prevention
- Image optimization

### Testing
- Jest for unit testing
- React Native Testing Library
- Detox for E2E testing
- Platform-specific testing

## Platform-Specific Development

- iOS-specific APIs and permissions
- Android-specific APIs and permissions
- Running on physical devices
- Platform channels for native code

## Debugging Tools

- React Native Debugger
- Flipper for advanced debugging
- Chrome DevTools
- Profiling and performance analysis

## Deployment

- iOS App Store submission
- Google Play Store submission
- Over-the-air updates (EAS Updates)
- Version management
- Release workflow

## Popular Libraries

- **react-native-reanimated** - Smooth animations
- **react-native-gesture-handler** - Touch handling
- **react-native-svg** - Vector graphics
- **react-native-maps** - Mapping
- **firebase** - Backend services
