# Metrics

A go library for printing metrics in an l2met compatible format.

## Usage

**Sampleing**

```go
metrics.Sample("goroutine", 1, "")
```

**Counting**

```go
metrics.Count("user.signup", 1)
```

**Measuring**

```go
metrics.Measure("request.time.2xx", 12.12, "ms")
```

**Timing**

```go
t := metrics.Time("request.time")
time.Sleep(1000)
t.Done()
```
