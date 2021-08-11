[![Build Status](https://github.com/azazeal/pause/actions/workflows/build.yml/badge.svg)](https://github.com/azazeal/pause/actions/workflows/build.yml)
[![Coverage Report](https://coveralls.io/repos/github/pause/fly/badge.svg?branch=master)](https://coveralls.io/github/azazeal/pause?branch=master)
[![Go Reference](https://pkg.go.dev/badge/github.com/pause/fly.svg)](https://pkg.go.dev/github.com/azazeal/pause)

# pause

Package pause implements a method which suspends execution until its given
Context is done or it's waited its given timeout.

## Usage

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/azazeal/pause"
)

func main() {
	pause.For(context.TODO(), time.Second)
	fmt.Println("about a second elapsed since we called pause.For")

	ctx, cancel := context.WithCancel(context.TODO())
	cancel()
	pause.For(ctx, time.Hour)
	fmt.Println("almost no time elapsed since we called pause.For")
}
```
