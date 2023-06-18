# ranger

![GitHub All Releases](https://img.shields.io/github/downloads/raspi/ranger/total?style=for-the-badge)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/raspi/ranger?style=for-the-badge)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/raspi/ranger?style=for-the-badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/raspi/ranger)](https://goreportcard.com/report/github.com/raspi/ranger)

Generate number ranges from given numbers.

For example: **1,2,3,5** generates **1-3,5-5**.

## Example

```go
package main

import (
	"fmt"
	"github.com/raspi/ranger"
)

func main() {
	r := []int8{-1, 0, 1, 3, 4, 5, 10, 11, 12, 13, 14, 15, 20, 30, 21, 25, 24, 21}

	ranges, err := ranger.GetIntegerRanges(r)
	if err != nil {
		panic(err)
	}

	for _, r := range ranges {
		fmt.Printf(`%v - %v`+"\n", r.Start, r.End)
	}
}
```

Outputs:

```
-1 - 1
3 - 5
10 - 15
20 - 21
24 - 25
30 - 30
```
