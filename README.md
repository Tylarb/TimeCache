# TimeCache
Interface of a timed cache structure in golang



Two different executions of this idea, one as a slice, the other as a dict. 

Test Results:

```
// Cache size = 10000000
$ go test -v
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


// Cache size = 10000
$ go test -run=XXX -bench=.
goos: darwin
goarch: amd64
BenchmarkSliceDoesNotContainRand-8   	   50000	     33293 ns/op
BenchmarkSliceContainsRand-8         	  100000	     19377 ns/op
BenchmarkSliceContainsLow-8          	 3000000	       425 ns/op
BenchmarkSliceContainsHigh-8         	   50000	     33946 ns/op
BenchmarkDictDoesNotContainRand-8    	    3000	    490225 ns/op
BenchmarkDictContainsRand-8          	    3000	    370982 ns/op
BenchmarkDictContainsLow-8           	    3000	    381490 ns/op
BenchmarkDictContainsHigh-8          	    5000	    398907 ns/op
PASS

// Cache size = 1000000
$ go test -run=XXX -bench=.
goos: darwin
goarch: amd64
BenchmarkSliceDoesNotContainRand-8   	     100	  13947496 ns/op
BenchmarkSliceContainsRand-8         	     100	  26063085 ns/op
BenchmarkSliceContainsLow-8          	 3000000	       481 ns/op
BenchmarkSliceContainsHigh-8         	      30	  58677970 ns/op
BenchmarkDictDoesNotContainRand-8    	       2	 547248957 ns/op
BenchmarkDictContainsRand-8          	       2	 627374031 ns/op
BenchmarkDictContainsLow-8           	       3	 522994370 ns/op
BenchmarkDictContainsHigh-8          	       2	 810512985 ns/op
PASS
```
