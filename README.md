# INFORMATION_SCHEMA vs SHOW TABLE STATUS benchmarks

## MYSQL 5.6
### Run #1
```
PORT=3306 go test -bench=.
goos: linux
goarch: amd64
BenchmarkGetISStats-8                 50          27481419 ns/op
BenchmarkShowTableStatus-8            50          28991959 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      3.338s
```
### Run #2
```
PORT=3306 go test -bench=.
goos: linux
goarch: amd64
BenchmarkGetISStats-8                 50          27405288 ns/op
BenchmarkShowTableStatus-8            50          28848221 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      2.874s
```
### Run #3
```
PORT=3306 go test -bench=.
goos: linux
goarch: amd64
BenchmarkGetISStats-8                 50          27564682 ns/op
BenchmarkShowTableStatus-8            50          28978764 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      2.886s
```
  
## MySQL 5.7
### Run #1
```
PORT=3307 go test -bench=.
goos: linux
goarch: amd64
BenchmarkGetISStats-8                 10         162660270 ns/op
BenchmarkShowTableStatus-8            10         161919889 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      4.373s
```
### Run #2
```
PORT=3307 go test -bench=.
goos: linux
goarch: amd64
BenchmarkGetISStats-8                 10         160471455 ns/op
BenchmarkShowTableStatus-8            10         154630866 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      3.461s
```
### Run #3
```
PORT=3307 go test -bench=.
goos: linux
goarch: amd64
BenchmarkGetISStats-8                 10         170686562 ns/op
BenchmarkShowTableStatus-8            10         164771823 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      3.702s
```
  
## MySQL 5.8
### Run #1
```
PORT=3308 go test -bench=.
goos: linux
goarch: amd64
BenchmarkGetISStats-8                  1        7632480754 ns/op
BenchmarkShowTableStatus-8            20          72842002 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      9.965s
```
### Run #2
```
PORT=3308 go test -bench=.
goos: linux
goarch: amd64
BenchmarkGetISStats-8                 30          49997142 ns/op
BenchmarkShowTableStatus-8            20          72326338 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      3.074s
```
### Run #3
```
PORT=3308 go test -bench=.
goos: linux
goarch: amd64
BenchmarkGetISStats-8                 30          50578023 ns/op
BenchmarkShowTableStatus-8            20          72491003 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      3.099s
```
### Run #4
```
PORT=3308 go test -bench=.
goos: linux
goarch: amd64
BenchmarkGetISStats-8                 30          50446703 ns/op
BenchmarkShowTableStatus-8            20          72131903 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      3.084s
```
### Run #5
```
PORT=3308 go test -bench=.
goos: linux
goarch: amd64
BenchmarkGetISStats-8                 30          50335986 ns/op
BenchmarkShowTableStatus-8            20          72521629 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      3.089s
```
### Run #6
```
karl@karl-OMEN ~/tmp/is_tests/a $ PORT=3308 go test -bench=.
goos: linux
goarch: amd64
BenchmarkGetISStats-8                 30          50151842 ns/op
BenchmarkShowTableStatus-8            20          72588632 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      3.085s
```
