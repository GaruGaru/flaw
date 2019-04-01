# Flaw 

## Inject failures on api calls for local chaos engineering


[![Build Status](https://travis-ci.org/GaruGaru/flaw.svg?branch=master)](https://travis-ci.org/GaruGaru/flaw)
[![Go Report Card](https://goreportcard.com/badge/github.com/GaruGaru/flaw)](https://goreportcard.com/report/github.com/GaruGaru/flaw)
![license](https://img.shields.io/github/license/GaruGaru/flaw.svg)
 
 
Flaw works as a proxy trought an external http service injecting failures 
randomly 

<img src="https://github.com/garugaru/flaw/raw/master/res/example-00.png" width="1000">


### Proxy 

```bash
flaw --host=https://jsonplaceholder.typicode.com/users
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