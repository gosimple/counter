# counter

[![Go Reference](https://pkg.go.dev/badge/github.com/gosimple/counter.svg)](https://pkg.go.dev/github.com/gosimple/counter)
[![Tests](https://github.com/gosimple/counter/actions/workflows/tests.yml/badge.svg)](https://github.com/gosimple/counter/actions/workflows/tests.yml)

Package `counter` is a simple thread-safe atomic counter. It is basically just
a wrapper of convenience functions around "sync/atomic" compare and swap
operations (CAS) for `uint64` and `int64` (without ability to go into negative
numbers).

Fork of <https://github.com/Avalanche-io/counter>

## Example

```go
package main

import (
    "fmt"
    "sync"

    "github.com/gosimple/counter"
)

func main() {

    c := counter.New()
    var wg sync.WaitGroup
    for i := 0; i < 8; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for i := 0; i < 10000; i++ {
                c.Up()
            }
        }()
    }
    wg.Wait()

    fmt.Printf("c = %d\n", c.Get())

}
```

output: `c = 80000`

### Functions

- `Add(val uint64) uint64` - Add val to the counter and return the total.

- `Up() uint64` - Add 1

- `Down() uint64` - Subtract 1

- `Subtract(val uint64) uint64` - Subtract val from the counter and return
  the total. If the value would be below zero then then the counter is set
  to zero and zero is returned.

- `Set(val uint64)` - Set the counter to val.

- `Get() uint64` - Get the value.

All methods are tread-safe non-blocking operations on the underlying uint64.
