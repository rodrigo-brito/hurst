# hurst

Hurst exponent estimation in Go

## Installation

```bash
go get -u github.com/rodrigo-brito/hurst
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/rodrigo-brito/hurst"
)

func main() {
    data := []float64{
		// your data here....
    }
    fmt.Println(hurst.Estimate(data, 2, 100)) // where 2 is the min lag, and 100 max lag
}
```

## References

- [Find Your Best Market to Trade With the Hurst Exponent](https://raposa.trade/blog/find-your-best-market-to-trade-with-the-hurst-exponent/)