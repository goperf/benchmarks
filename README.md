# Benchmarks

## How to Run

### Basic form of go test command

```bash
go test -bench=. -benchmem
```

### All profiles (Not a good practice)
```bash
go test -benchmem -benchtime=100000x -cpuprofile=cpu.profile -memprofile=memory.profile -blockprofile=blocking.profile -run=^$ -bench "^(BenchmarkNotPadded)$" ./...

```

### Take specific profiles
```bash
go test -benchmem -memprofile=memory.profile -bench "^(BenchmarkEscapeAnalysis)$" ./...
```

## Go test command usage

```bash
go test [...flags] <package>
```

-benchmem - Report memort allocations

-benchtime=100000x (Number of iterations)

-cpuprofile=cpu.profile (Collect CPU profile)

-memprofile=memory.profile (Collect memory profile)

-blockprofile=blocking.profile (Collect block profile)

-run=^$ (Run tests matching given regex pattern)

-bench "^(BenchmarkNotPadded)$" (Run banchmarks matching given pattern)
