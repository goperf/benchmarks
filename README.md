# Benchmarks

## How to Run

```bash
go test -benchmem -benchtime=100000x -cpuprofile=cpu.profile -memprofile=memory.profile -blockprofile=blocking.profile -run=^$ -bench "^(BenchmarkNotPadded)$" ./...

```
