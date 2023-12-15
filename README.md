# Go Cantrips

Cantrips is a collection of useful functions in the [Go Programming Language](https://go.dev/)

[![Go Reference](https://pkg.go.dev/badge/github.com/banaaron/cantrips.svg)](https://pkg.go.dev/github.com/banaaron/cantrips)

## What are Cantrips?

> **NOTE**
>
> / ˈkɑn trɪp / _noun_
>
> 1. `chiefly Scotland`: a magic spell; trick by sorcery.
> 2. `chiefly British`: hocus-pocus. <br><br>

Cantrips are small (_magical_) utility functions to achieve common tasks in Go.
The name was inspired by Svelte 5's `Runes` and Baldur's Gate 3.

For example the `Pop()` function will remove the last element from a slice and
return that element.

```go
package main

import (
	"fmt"
	"github.com/banaaron/cantrips"
)

func main() {
	names := []string{"aaron", "chris", "drew"}
	n := cantrips.Pop(&names)
	fmt.Println(n)     // "drew"
	fmt.Println(names) // ["aaron", "chris"]

	// you can provide an index to remove an element at a specific index
	p := cantrips.Pop(&names, 1)
	fmt.Println(p)     // "chris"
	fmt.Println(names) // ["aaron"]
}
```
