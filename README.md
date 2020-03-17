# diff

Fast diff library for Myers algorithm.
The algorithm is described in "An O(ND) Difference Algorithm and its Variations", Eugene Myers, Algorithmica Vol. 1 No. 2, 1986, pp. 251-266.

```
BenchmarkBytesDiff/bytestrings-12         	      74	  16075247 ns/op	   91153 B/op	      12 allocs/op
BenchmarkDiff-12                          	  638715	      1882 ns/op	     680 B/op	       6 allocs/op
BenchmarkInts-12                          	  582735	      1971 ns/op	     728 B/op	       7 allocs/op
BenchmarkDiffRunes-12                     	  574765	      1919 ns/op	     728 B/op	       7 allocs/op
BenchmarkDiffBytes-12                     	   66489	     19385 ns/op	    3408 B/op	       8 allocs/op
BenchmarkDiffByteStrings-12               	   72484	     16184 ns/op	    3392 B/op	       8 allocs/op
```