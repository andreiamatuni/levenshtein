# levenshtein

Compute Levenshtein distances between strings using the Wagner-Fischer algorithm.

There are 4 functions, EditDistance() and BufferedEditDistance (and their compact counterparts).
If you're going to be computing a lot of distances in a loop, you should use the buffered version
to avoid all the allocation/deallocation costs. Buffered functions **do not allocate**. Here's the performance difference on my machine. Each run is a 1000 pair comparison:

```

BenchmarkUnbuffered-4                   	    3000	    552195 ns/op	  628480 B/op	    8120 allocs/op
BenchmarkBuffered-4                     	   10000	    161841 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnbufferedCompact-4            	    5000	    237236 ns/op	  128000 B/op	    2000 allocs/op
BenchmarkBufferedCompact-4              	   10000	    172415 ns/op	       0 B/op	       0 allocs/op
BenchmarkBasePairsUnBuffered-4          	     200	   9482169 ns/op	12736000 B/op	   38000 allocs/op
BenchmarkBasePairsBuffered-4            	     200	   7109534 ns/op	      63 B/op	       0 allocs/op
BenchmarkBasePairsUnbufferedCompact-4   	     200	   6832891 ns/op	  640000 B/op	    2000 allocs/op
BenchmarkBasePairsBufferedCompact-4     	     300	   5796142 ns/op	       0 B/op	       0 allocs/op

```

The BasePair benchmarks use 36x36 character string comparisons (DNA/RNA base pairs). They highlight the benefit of the compact functions over building the whole distance matrix when inputs are large.

The compact versions just use the current and previous rows, so they run in O(m) memory rather than O(mn)


You set edit weights by passing in a Weights{} struct as a configuration parameter to the functions. Benchmarks were run with all weights set to 1.

<!-- Speed gain between buffered and unbuffered compact versions is about 38%

Buffered whole matrix is slightly faster than the buffered compact on small input strings (about 7%)

When input strings are large, buffered compact is about 22% faster than

Speed gain between buffered compact vs. buffered whole matrix versions is about 4%

Speed gain between buffered compact and unbuffered whole matrix is about 325% -->
