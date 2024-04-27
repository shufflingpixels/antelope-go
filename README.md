antelope-go
===========

[![Go Report Card](https://goreportcard.com/badge/github.com/pnx/antelope-go)](https://goreportcard.com/report/github.com/pnx/antelope-go)
[![Test](https://github.com/pnx/antelope-go/actions/workflows/test.yml/badge.svg?branch=master&event=push)](https://github.com/pnx/antelope-go/actions/workflows/test.yml)

Library for interacting with antelope

Benchmarks
----------

```
go version: 1.21.8
goos: linux
goarch: amd64
pkg: github.com/pnx/antelope-go/internal/benchmarks
cpu: Intel(R) Xeon(R) CPU E5-1650 v3 @ 3.50GHz
Benchmark_Decode_AbiDef-12              	 516951	     2340 ns/op	   1264 B/op	     42 allocs/op
Benchmark_Decode_AbiDef_EosCanada-12    	 342844	     3235 ns/op	   1129 B/op	     30 allocs/op
Benchmark_Encode_AbiDef-12              	 696310	     1817 ns/op	    744 B/op	     29 allocs/op
Benchmark_Encode_AbiDef_EosCanada-12    	 215685	     6254 ns/op	   1513 B/op	     56 allocs/op
Benchmark_Decode-12                     	 555462	     1986 ns/op	   1016 B/op	     52 allocs/op
Benchmark_Decode_NoOptimize-12          	 154348	     7013 ns/op	   1352 B/op	     92 allocs/op
Benchmark_Decode_EosCanada-12           	 105522	    11338 ns/op	   2376 B/op	     95 allocs/op
Benchmark_Encode-12                     	1387473	    850.7 ns/op	    392 B/op	     39 allocs/op
Benchmark_Encode_NoOptimize-12          	 256351	     4449 ns/op	   1056 B/op	     88 allocs/op
Benchmark_Encode_EosCanada-12           	 116320	     9305 ns/op	   1744 B/op	    135 allocs/op
```

```
go version: 1.17.13
goos: linux
goarch: amd64
pkg: github.com/pnx/antelope-go/internal/benchmarks
cpu: AMD EPYC 7763 64-Core Processor
Benchmark_Decode_AbiDef-4             	  470780	      2262 ns/op	    1760 B/op	      52 allocs/op
Benchmark_Decode_AbiDef_EosCanada-4   	  342249	      3635 ns/op	    1840 B/op	      43 allocs/op
Benchmark_Encode_AbiDef-4             	  643518	      1809 ns/op	    1240 B/op	      39 allocs/op
Benchmark_Encode_AbiDef_EosCanada-4   	  239196	      5057 ns/op	    2048 B/op	      66 allocs/op
Benchmark_Decode-4                    	  816806	      1481 ns/op	    1016 B/op	      52 allocs/op
Benchmark_Decode_NoOptimize-4         	  202341	      5934 ns/op	    1352 B/op	      92 allocs/op
Benchmark_Decode_EosCanada-4          	  121315	      9902 ns/op	    2376 B/op	      95 allocs/op
Benchmark_Encode-4                    	 1545202	       694.9 ns/op	     392 B/op	      39 allocs/op
Benchmark_Encode_NoOptimize-4         	  320676	      3768 ns/op	    1056 B/op	      88 allocs/op
Benchmark_Encode_EosCanada-4          	  166828	      7241 ns/op	    1744 B/op	     135 allocs/op
```

[All benchmark runs](https://github.com/pnx/antelope-go/actions/workflows/benchmark.yml)


License
-------

```
Copyright (C) 2024  Henrik Hautakoski <henrik.hautakoski@gmail.com>
Copyright (C) 2021  Greymass Inc.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
```
