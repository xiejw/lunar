## Reading Notes on Memory Model

### Ref

- [Go Memory Model](https://golang.org/ref/mem)
- [Memory Models](https://research.swtch.com/mm)

### Examples

**Single thread**

```
x = 1
y = x
```

The rule for data racing free is the y must observe the new value of `x` with
programming language compiler reordering or hardware reording.

Here, the happen-before relationship is
```
x=1 -> y=x.
```

To make this example more complicated, we can have:

```
x = 1
a = 2
y = x
b = a
```

This example is allowed to be executed as

```
x = 1
y = x
a = 2
b = 3
```

For happen-before relationship, it is

```
x=1 -> a=2 -> y=x -> b=a.
```
But as the reorder example above is safe for data trace, it is allowed to do that.

**Multi Threads**
