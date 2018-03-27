# TimeCache
Interface of a timed cache structure in golang



Two different executions of this idea, one as a slice, the other as a dict. 

Test Results:

```
TimeCache(master*)$ go test -v
=== RUN   TestSliceAdd
--- PASS: TestSliceAdd (4.22s)
=== RUN   TestSliceContains1
--- PASS: TestSliceContains1 (0.00s)
=== RUN   TestSliceContains1000000
--- PASS: TestSliceContains1000000 (0.00s)
=== RUN   TestSliceDoesNotContain
--- PASS: TestSliceDoesNotContain (0.03s)
	cache_test.go:55: Current size 10000000
=== RUN   TestSliceExpire
--- PASS: TestSliceExpire (30.17s)
=== RUN   TestDictCreate
--- PASS: TestDictCreate (0.00s)
=== RUN   TestDictAdd
--- PASS: TestDictAdd (7.32s)
=== RUN   TestDictContains1
--- PASS: TestDictContains1 (0.43s)
=== RUN   TestDictContains1000000
--- PASS: TestDictContains1000000 (1.27s)
=== RUN   TestDictDoesNotContain
--- PASS: TestDictDoesNotContain (1.03s)
	cache_test.go:141: Current size 10000000
=== RUN   TestDictExpire
--- PASS: TestDictExpire (31.94s)
PASS
ok  	_/Users/tramer/Go/src/github.com/Tylarb/TimeCache	76.719s

TimeCache(master*)$ go test -bench .
goos: darwin
goarch: amd64
BenchmarkSliceContains-8   	       1	5000418849 ns/op
BenchmarkDictContains-8    	       1	5290397855 ns/op
PASS
ok  	_/Users/tramer/Go/src/github.com/Tylarb/TimeCache	88.918s
```
