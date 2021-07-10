This is my Go repository. It contains several common modules I used in other
projects.

# Set Up

```go
import (
  "github.com/xiejw/lunar/base"
)

func main() {
  base.Init(true/*parseFlag*/)
  defer base.FinishUp()

}
```

# APIs

## Package `cfmt`

Usage:

```go
cfmt.Red("Hello")
```

## Package `glog`

Fork of [golang/glog](https://github.com/golang/glog). Usage

```go
glog.Info("Prepare to repel boarders")

glog.Fatalf("Initialization failed: %s", err)

if glog.V(2) {
	glog.Info("Starting transaction...")
}
glog.V(2).Infoln("Processed", nItems, "elements")
```

## Package `base/exec`

Usage:

```go
results, _ := RunCmd("whoami")
fmt.Println(results[0])
```

## Package `base/crypto`

Usage:

```go
hash, _ := Sha256Sum("~/.profile")
fmt.Println(hash)
```
