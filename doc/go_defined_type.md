# Go Programming Language Learning Notes

## Defined Type

_Written on 2018-07-22_

The new type, defined by `type` is called a defined type. It is different from
any other type, including the type it is created from.

However, it is not easy to understand why you can pass a constant (literal) to
defined type. This is related to [Constant
Representability](https://golang.org/ref/spec\#Representability).


    package main

    import "fmt"

    type AnotherBool bool

    func SayHi(b bool) { fmt.Println("h") }

    func SayHello(b AnotherBool) { fmt.Println("g") }

    func main() {
      var b bool
      b = false

      SayHi(true)
      SayHi(b)

      var ab AnotherBool
      ab = true
      SayHello(ab)
      SayHello(false)  // This is valid

      // SayHello(b)  // Invalid
      // SayHi(AnotherBool(true))  // invalid
      // SayHi(ab)  // invalid
    }
