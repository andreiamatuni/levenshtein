# levenshtein

Compute Levenshtein distance between strings using Wagner-Fischer algorithm.

There are two main functions, EditDistance() and BufferedEditDistance. If you're going to be computing
a lot of distances in a loop, you should use the buffered version. Here's the performance difference on my machine:

```

BenchmarkUnBufferedDistance-4   	  200000	     10580 ns/op	   12496 B/op	     163 allocs/op
BenchmarkBufferedDistance-4     	  500000	      3635 ns/op	       0 B/op	       0 allocs/op

```
