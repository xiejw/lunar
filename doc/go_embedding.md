# Go Programming Language Learning Notes

## Embedding

_Written on 2019-07-13_

According to [Effective
Go](https://tip.golang.org/doc/effective_go.html#embedding): "Go does not
provide the typical, type-driven notion of subclassing, but it does have the
ability to “borrow” pieces of an implementation by embedding types within a
struct or interface."

To define and use it, see example below:

    package main

    import "fmt"

    type Car struct {
      NumOfWheels int
    }

    func (c *Car) Print() {
      fmt.Println("I am Car.")
    }

    type Tesla struct {
      Car

      BatteryCapacity float64
    }

    func (t *Tesla) Print() {
      fmt.Println("I am Tesla.")
    }

    func main() {
      // Both valid. For *pkg.T type, use T only (no pointer, no pkg)
      a := &Tesla{Car{NumOfWheels: 4}, 1000}
      _ = Tesla{Car: Car{NumOfWheels: 4}, BatteryCapacity: 1000}

      _ = a.Car
      _ = a.NumOfWheels

      a.Print()     // Prints: I am Tesla.
      a.Car.Print() // Prints: I am Car.
    }
