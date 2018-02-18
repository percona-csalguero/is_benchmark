# INFORMATION_SCHEMA vs SHOW TABLE STATUS benchmarks

## MYSQL 5.6.38 at port 3306
### Testing instance at port 3306. Run # 1
```
goos: linux
goarch: amd64
BenchmarkGetISStats-8                  3         359120493 ns/op
BenchmarkShowTableStatus-8             3         387891153 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      4.469s
```
### Testing instance at port 3306. Run # 2
```
goos: linux
goarch: amd64
BenchmarkGetISStats-8                  3         375187488 ns/op
BenchmarkShowTableStatus-8             3         348196759 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      4.318s
```
### Testing instance at port 3306. Run # 3
```
goos: linux
goarch: amd64
BenchmarkGetISStats-8                  3         358736163 ns/op
BenchmarkShowTableStatus-8             3         356398255 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      4.271s
```
### Testing instance at port 3306. Run # 4
```
goos: linux
goarch: amd64
BenchmarkGetISStats-8                  3         372290582 ns/op
BenchmarkShowTableStatus-8             3         357169090 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      4.287s
```
### Testing instance at port 3306. Run # 5
```
goos: linux
goarch: amd64
BenchmarkGetISStats-8                  3         366820981 ns/op
BenchmarkShowTableStatus-8             3         351896857 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      4.218s
```
## MySQL 5.7.20 at port 3307
### Testing instance at port 3307. Run # 1
```
goos: linux
goarch: amd64
BenchmarkGetISStats-8                  5         320528250 ns/op
BenchmarkShowTableStatus-8             5         309795997 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      6.362s
```
### Testing instance at port 3307. Run # 2
```
goos: linux
goarch: amd64
BenchmarkGetISStats-8                  3         349692895 ns/op
BenchmarkShowTableStatus-8             5         314158521 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      4.208s
```
### Testing instance at port 3307. Run # 3
```
goos: linux
goarch: amd64
BenchmarkGetISStats-8                  5         321753271 ns/op
BenchmarkShowTableStatus-8             5         321569540 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      5.826s
```
### Testing instance at port 3307. Run # 4
```
goos: linux
goarch: amd64
BenchmarkGetISStats-8                  5         320067480 ns/op
BenchmarkShowTableStatus-8             3         337697981 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      4.207s
```
### Testing instance at port 3307. Run # 5
```
goos: linux
goarch: amd64
BenchmarkGetISStats-8                  5         317414181 ns/op
BenchmarkShowTableStatus-8             5         327111794 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      5.822s
```
## MySQL 8.0.3-rc-log at port 3308
### Testing instance at port 3308. Run # 1
```
goos: linux
goarch: amd64
BenchmarkGetISStats-8                 20          98974508 ns/op
BenchmarkShowTableStatus-8            10         142689304 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      3.655s
```
### Testing instance at port 3308. Run # 2
```
goos: linux
goarch: amd64
BenchmarkGetISStats-8                 20          98549707 ns/op
BenchmarkShowTableStatus-8            10         143455072 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      3.653s
```
### Testing instance at port 3308. Run # 3
```
goos: linux
goarch: amd64
BenchmarkGetISStats-8                 20          98699462 ns/op
BenchmarkShowTableStatus-8            10         142670109 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      3.649s
```
### Testing instance at port 3308. Run # 4
```
goos: linux
goarch: amd64
BenchmarkGetISStats-8                 20          99339057 ns/op
BenchmarkShowTableStatus-8            10         142692422 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      4.655s
```
### Testing instance at port 3308. Run # 5
```
goos: linux
goarch: amd64
BenchmarkGetISStats-8                 20          99505314 ns/op
BenchmarkShowTableStatus-8            10         142655675 ns/op
PASS
ok      _/home/karl/tmp/is_tests/a      3.663s
```

