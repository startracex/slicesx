# slicesx

## Installation

```sh
go get -u github.com/startracex/slicesx
```

## Example

```go
package main

import (
  "fmt"
  "github.com/startracex/slicesx"
)

func main() {
  nums := []int{1, 2, 3}

  doubled := slicesx.Map(nums, func(v int, _ int) int {
    return v * 2
  })

  fmt.Println(doubled) // [2 4 6]
}
```
