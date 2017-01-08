# Numbers

This Markdown is intended to have brotips for using go in data and numbers analysis.

## Random Numbers

__Setting Seeds__

``` go
import "math/rand"

rand.Seed(time.Now().UTC().UnixNano())

```
