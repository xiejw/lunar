This is my Go repository. It contains several common modules I used in other
projects.

# Set Up

```go
import (
  "github.com/xiejw/lunar"
)

func main() {
  lunar.Init(true/*parseFlag*/)
  defer lunar.FinishUp()

}
```

# APIs

## Package `cfmt`

Usage:

    cfmt.Red("Hello")

## Package `base/exec`

Usage:

    results, _ := RunCmd("whoami")
    fmt.Println(results[0])

## Package `base/crypto`

Usage:

    hash, _ := Sha256Sum("~/.profile")
    fmt.Println(hash)

