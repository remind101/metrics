# Metrics

A go library for printing metrics in an l2met compatible format.

## Usage

**Sampleing**

```go
Sample("goroutine", 1, "")
```

**Counting**

```go
Count("user.signup", 1)
```

**Measuring**

```go
Measure("request.time.2xx", 12.12, "ms")
```
