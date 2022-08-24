Lazy
====

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/mono83/lazy.svg)](https://github.com/mono83/lazy)
[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/mono83/lazy)
[![Go Report Card](https://goreportcard.com/badge/github.com/mono83/lazy)](https://goreportcard.com/report/github.com/mono83/lazy)
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/mono83/lazy/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/mono83/lazy/tree/main)

Simple generic concurrent-safe lazy initialization container for Go. 

## Installation

```bash
go get -u github.com/mono83/lazy
```


## Usage

```go
import "github.com/mono83/lazy"

// Static value supplier
greet := lazy.Const("Hello, world")
fmt.Println(greet())

// With supplier
some := lazy.New(func() any { /* Some heavy operations be execute once */ })
value := some()

// Also there are constructors for value + error tuples
lazy.ConstE[T](42)                // -> func() (T,error)
lazy.Error[T](nil)                // -> func() (T,error)
lazy.NewE(func(T any, err error)) // -> func() (T,error)
```

### Benchmarks

```
$ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/mono83/lazy
cpu: Intel(R) Core(TM) i5-10400F CPU @ 2.90GHz
BenchmarkConst-12       61397071                20.70 ns/op           24 B/op          1 allocs/op
BenchmarkNew-12         18599815                66.76 ns/op           72 B/op          3 allocs/op
BenchmarkAll/Read-const-1times-12               522478002                2.310 ns/op           0 B/op          0 allocs/op
BenchmarkAll/Read-new-1times-12                 422763762                2.863 ns/op           0 B/op          0 allocs/op
BenchmarkAll/Read-const-100times-12              7491662               162.9 ns/op             0 B/op          0 allocs/op
BenchmarkAll/Read-new-100times-12                4521046               260.5 ns/op             0 B/op          0 allocs/op
BenchmarkAll/Read-const-1000times-12              748990              1563 ns/op               0 B/op          0 allocs/op
BenchmarkAll/Read-new-1000times-12                476796              2619 ns/op               0 B/op          0 allocs/op
BenchmarkAll/Read-const-10000times-12              78136             15567 ns/op               0 B/op          0 allocs/op
BenchmarkAll/Read-new-10000times-12                46693             26274 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/mono83/lazy  14.628s
```