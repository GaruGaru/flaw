# Flaw 

## Inject failures on api calls for local chaos engineering

### Proxy 

```bash
flaw https://jsonplaceholder.typicode.com/users
```

### Proxy injecting 50% of http (status 500) failures 

```bash
flaw --run=perc --percentage=50 --host=https://jsonplaceholder.typicode.com/users
```

### Proxy injecting 30% of http rate limit (status 429) failures 

```bash
flaw --status-code=429 --run=perc --percentage=30 --host=https://jsonplaceholder.typicode.com/users
```

### Proxy injecting 30% of http failures + 1s latency 

```bash
flaw --latency=1000 --run=perc --percentage=30 --host=https://jsonplaceholder.typicode.com/users
```

### Installation

```bash
go get -u github.com/GaruGaru/flaw
```