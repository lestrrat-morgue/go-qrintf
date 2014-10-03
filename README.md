go-qrintf
=========

Just wondering what happens if you do the same thing as qrintf in go

## Benchmark Oct 3, 2014

```
go test -bench IPv4Addr -benchmem -benchtime=100s
PASS
BenchmarkIPv4Addr_qrintf	10000000000	         0.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkIPv4Addr_sprintf	10000000000	         0.57 ns/op	       0 B/op	       0 allocs/op
ok  	github.com/lestrrat/go-qrintf	45.194s
```
