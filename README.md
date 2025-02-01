# Go Slice Modification Bug

This repository demonstrates a common but subtle bug in Go related to how slices modify the underlying array.  The `bug.go` file contains code that showcases the unexpected behavior.  The `bugSolution.go` file offers corrected code to avoid the issue.

The core problem stems from the fact that slices in Go are views into the underlying array.  Modifying a slice directly modifies the original array. This is usually what you intend. But unexpected behavior can arise when you are slicing the array and then passing that slice to a function for modification. When the function modifies the slice, it's modifying the original array and the behavior may not be what you are expecting.

This example highlights the importance of understanding how slices and arrays interact in Go to write robust and predictable code.