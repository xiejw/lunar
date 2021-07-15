## Reading Notes on Memory Model

### Ref

- [Go Memory Model](https://golang.org/ref/mem)
- [Memory Models](https://research.swtch.com/mm)


### Happens-Before and Reordering

Happens-Before defines the ordering relationship and thereby defines what
writes must be observed by the reads later as a contract.  This draws a hard
requirement on multi-threads program for compiler and hareware support, as a
natual extention of the single-thread program, as well as for user's mental
models.

### Examples of Happens-Before

**Single thread**

```
x = 1
y = x
```

The rule for data racing free or correctness is the `y` must observe the new
value of `x` with programming language compiler reordering or hardware reording.

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

```
(x,y started with 0)

Thread 1    Thread 2
--------    --------
x = 1       a = x
y = 2       b = y
```

Here, there is no happens-before relationship between two threads, there will be
no guarantee the final results might look like. A surprising result could be
```
a = 0
b = 2
```
due to multi-layer recording happen in the system. For multi-word size data
types, the corruption could be unpredictable.


A promising example could be
```
(x,y started with 0)

Thread 1    Thread 2
--------    --------
x = 1
y = 2
S(a)     -> S(a)
            a = x
            b = y
```
where `S(a)` is a sync point, which defines Happens-Before relationship as

```
x = 1  ->
y = 2  ->
a = x  ->
b = y
```
across the multi-threads. With this, it is clear that, `a` and `b` must see the
expected values.

A confusing example could be
```
(x,y started with 0)

Thread 1    Thread 2    Thread 3
--------    --------    --------
x = 1       x =2
S(a)                    -> S(a)
            S(b)        -> S(b)
                        a = x
                        b = x
```
