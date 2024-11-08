# ⚠️ This package is no longer supported

Most of the functionality in this package is now provided by the standard library
or does not work due to changes made to the way concurrency is handled in the new feature.

# Loop

Loop is a package that provides commonly used functions
for ranging.

In any case, this package should not be used moving forward


## Requirements

This package requires Go 1.22 with the `GOEXPERIMENT=rangefunc` env
var enabled.

## Usage

The package provides a number of different functions for ranging.

### Parallel

⚠️ Since the features official release in 1.23, this method will now panic.

A commonly used pattern in Go is to iterate over a slice of elements in parallel with a 
wait group.

The parallel iterator provides this functionality in an easy to use interface

```go
package main

import "github.com/dreamsofcode-io/loop"

func main() {
    xs := []int{1,2,3,4,5}
    squares := make([]int{}, len(xs))

    // Each iteration runs in a goroutine
    for i, x := range loop.Parallel(xs) {
        // Simulate a long running task
        time.Sleep(time.Second)
        squares[i] = x * x
    }

    fmt.Println(squares) // [2, 4, 9, 16, 25]
}
```

The above task will run in parallel, which means the total operation will only take 1 second, 
instead of the 5 it would take otherwise. 

⚠️ One thing to be aware of is that teach iteration runs in a separate goroutine. Therefore
you'll want to make sure you are performing thread safe operations.

The parallel task won't speed up any compute heavy operations, in that case, you're better
off using a normal loop. However, in the event of performing network requests or async
tasks, then using loop.Parallel will improve performance.

```go
import (
    "slog"
    "net/http"

    "github.com/dreamsofcode-io/loop"
)

func main() {
    colors := []string{"green", "yellow", "blue"}

    results := make([]*http.Response{}, len(colors))
    for _, color := range loop.Parallel(colors) {
        _, err := http.Post("http://example.com/colors", "text/plain", strings.NewReader(color))
        if err != nil {
            slog.Error("oops", slog.Any(err))
        }
    }
}
```

### Pool
The pool function is very similar to `loop.Parallel`, however it allows to caller to set the
concurrency amount with the second argument.

This is useful in the event you want bounded concurrency.

```go
package main

import "github.com/dreamsofcode-io/loop"

func main() {
    xs := []int{1,2,3,4,5}

    // Each iteration runs in a goroutine
    for _, x := range loop.Pool(xs, 2) {
        // Simulate a long running task
        time.Sleep(time.Second)
    }
}
```

In the above example, only 2 elements will be performed at a time.

### Batch

The Batch function provides the ability to range over elements in batches. The size of each batch
is decided by the given size argument, in which a batch will either be the same size or less than.

The `loop.Batch` method runs in a single goroutine

```go
import "github.com/dreamsofcode-io/loop"

func main() {
    nums := []int{1, 2, 3, 4, 5}

    for i, batch := range loop.Batch(nums, 2) {
        fmt.Println(i, batch)
    }
}
```

The above code will print the following output:

```
0 [1, 2]
1 [3, 4]
2 [5]
```

If a batch size of 0 is passed in, then no iterations of the loop are performed. This behavior
may change instead to panic as it's effectively a divide by 0.

### Range

The range function allows you to iterate over a range of integer types.

```go
import "github.com/dreamsofcode-io/loop"

func main() {
    for i := range loop.Range(0, 5) {
        fmt.Println(i)
    }
}
```

The above code will print out

```
0
1
2
3
4
```

The `loop.Range` method includes the starting value, but excludes the stop value.

### ParallelTimes

This method allows to perform a parallel operation for a given number of times.

For example

```go
import "github.com/dreamsofcode-io/loop"

func main() {
    for i := range loop.ParallelTimes(10) {
        time.Sleep(time.Second)
    }
}
```

will cause time.Sleep to be called 10 times in parallel.
