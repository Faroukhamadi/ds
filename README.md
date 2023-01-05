# ds

## This is my data-structures implementation in Go

---

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/Faroukhamadi/ds
```

## Basic Linked List Example

```go
func main() {
    l := ll.New(1)
    l = ll.Append(l, 2)
    l = ll.Append(l, 3)

    l.Print()
}
```
