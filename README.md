# waitfor-http

[![Build](https://github.com/go-waitfor/waitfor-http/actions/workflows/build.yml/badge.svg)](https://github.com/go-waitfor/waitfor-http/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-waitfor/waitfor-http)](https://goreportcard.com/report/github.com/go-waitfor/waitfor-http)
[![codecov](https://codecov.io/gh/go-waitfor/waitfor-http/branch/main/graph/badge.svg)](https://codecov.io/gh/go-waitfor/waitfor-http)

HTTP/HTTPS resource readiness assertion library for the [waitfor](https://github.com/go-waitfor/waitfor) framework.

## Overview

`waitfor-http` is a plugin for the [waitfor](https://github.com/go-waitfor/waitfor) library that provides HTTP and HTTPS resource testing capabilities. It allows you to wait for web services, APIs, and other HTTP endpoints to become available before proceeding with your application logic.

The library performs GET requests to specified URLs and considers the resource ready when it receives a successful HTTP response (status codes 200-399). This is particularly useful for:

- Waiting for web services to start up in Docker containers
- Ensuring API dependencies are available before starting your application
- Health checks in microservice architectures
- Integration testing scenarios

## Installation

```bash
go get github.com/go-waitfor/waitfor-http
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-waitfor/waitfor"
	"github.com/go-waitfor/waitfor-http"
	"os"
)

func main() {
	runner := waitfor.New(http.Use())

	err := runner.Test(
		context.Background(),
		[]string{"http://localhost:8080", "https://api.example.com"},
		waitfor.WithAttempts(5),
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

## Usage Examples

### Basic HTTP Testing

```go
import (
	"context"
	"github.com/go-waitfor/waitfor"
	"github.com/go-waitfor/waitfor-http"
)

func waitForService() error {
	runner := waitfor.New(http.Use())
	
	ctx := context.Background()
	return runner.Test(ctx, []string{"http://localhost:3000"})
}
```

### Testing Multiple Endpoints

```go
func waitForMultipleServices() error {
	runner := waitfor.New(http.Use())
	
	endpoints := []string{
		"http://localhost:8080/health",
		"https://api.service.com/status",
		"http://database:5432",
	}
	
	ctx := context.Background()
	return runner.Test(ctx, endpoints, waitfor.WithAttempts(10))
}
```

### With Timeout and Custom Configuration

```go
import (
	"context"
	"time"
	"github.com/go-waitfor/waitfor"
	"github.com/go-waitfor/waitfor-http"
)

func waitWithTimeout() error {
	runner := waitfor.New(http.Use())
	
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	return runner.Test(
		ctx,
		[]string{"http://slow-starting-service:8080"},
		waitfor.WithAttempts(15),
		waitfor.WithInterval(2*time.Second),
	)
}
```

## API Reference

### `http.Use() waitfor.ResourceConfig`

Returns a resource configuration for HTTP/HTTPS testing that can be used with the waitfor framework.

### Supported URL Schemes

- `http://` - HTTP endpoints
- `https://` - HTTPS endpoints

### Status Code Handling

The library considers the following HTTP status codes as successful:
- **2xx Success**: 200-299 (OK, Created, Accepted, etc.)
- **3xx Redirection**: 300-399 (Moved Permanently, Found, etc.)

Status codes 400 and above (4xx Client Error, 5xx Server Error) are considered failures.

## Error Handling

The library will return errors in the following cases:

- **Network errors**: Connection refused, timeout, DNS resolution failures
- **HTTP errors**: 4xx and 5xx status codes
- **Invalid URLs**: Malformed or nil URLs

Example error handling:

```go
runner := waitfor.New(http.Use())
err := runner.Test(ctx, []string{"http://localhost:8080"})

if err != nil {
	// Handle different types of errors
	fmt.Printf("Service not ready: %v\n", err)
	// Implement retry logic or fail gracefully
}
```

## Integration with waitfor

This library is designed to work with the [waitfor](https://github.com/go-waitfor/waitfor) framework. For more advanced configuration options like custom retry intervals, backoff strategies, and attempt limits, refer to the [waitfor documentation](https://github.com/go-waitfor/waitfor).

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.