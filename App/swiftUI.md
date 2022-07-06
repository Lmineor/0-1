## Making a Class Observable

You can make a class observable by adopting the `ObservableObject` protocol. Choose the properties in the class that need to trigger a UI update when they change. Then, add the `@Published` attribute to each of these properties.

The starter project for the next tutorial includes a `ScrumTimer` class that keeps time for a daily scrum meeting:

```swift
class ScrumTimer: ObservableObject {
   @Published var activeSpeaker = ""
   @Published var secondsElapsed = 0
   @Published var secondsRemaining = 0
   // ...
}```

The class declares several published properties that update frequently during a scrum session. `ScrumTimer` notifies any observers when the values of its published properties change.

## Monitoring an Object for Changes

You can tell SwiftUI to monitor an observable object by adding one of the following attributes to the property’s declaration: `ObservedObject`, `StateObject`, or `EnvironmentObject`. A view property declared with one of these wrappers creates a new source of truth for your view hierarchy.

The `@StateObject` wrapper creates an observable object in a view. The system initializes the object when creating the view, and it keeps the object available to use in that view or in other views that you pass the object to.

```swift
struct MeetingView: View {
   @StateObject var scrumTimer = ScrumTimer()
   // ...
}
```

Use the `@ObservedObject` property wrapper to indicate that you passed in an object from another view. Because you create the object in another view, you don’t provide an initial value for an `ObservedObject`:

```swift
struct ChildView: View {
   @ObservedObject var timer: ScrumTimer
   // ...
}
```

Then, you pass an instance of the observable object to the view’s initializer:
```swift
struct MeetingView: View {
   @StateObject var scrumTimer = ScrumTimer()
   var body: some View {
      VStack {
         ChildView(timer: scrumTimer)
      }
   }
   // ...
}
```

Instead of passing objects into individual views, you can place objects into the environment. The `environmentObject(_:)` view modifier places an object in the environment in a parent view. Any ancestor view can then access the object without an explicit injection chain.

```swift
struct ParentView: View {
   @StateObject var scrumTimer = ScrumTimer()
   var body: some View {
      VStack {
         ChildView()
            .environmentObject(scrumTimer)
      }
   }
   // ...
}
```

Then, you can use the `@EnvironmentObject` property wrapper to access the object in any descendent of the parent view, even if intermediate views in the hierarchy don’t have references to the object. SwiftUI tracks only a dependency in views that read the data.

```swift
struct ChildView: View {
   @EnvironmentObject var timer: ScrumTimer
   // ...
}
```

You won’t use `@EnvironmentObject` in this module, but you can learn more about it, as well as the other property wrappers, by visiting [Managing Model Data in Your App](https://developer.apple.com/documentation/swiftui/managing-model-data-in-your-app).


## 生命周期

SwiftUI有三个modifiers that respond to view life cycle events:
-   `onAppear(perform:)` triggers actions any time the view appears on screen, even if it’s not the first time.
-   `onDisappear(perform:)` triggers actions when a view disappears from screen.
-   `task(priority:_:)` triggers actions that execute asynchronously when the view appears on screen.

