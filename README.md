# Logger - Lightweight Structured Logging for Go

[![Go Report Card](https://goreportcard.com/badge/github.com/hamedcodelab/logger)](https://goreportcard.com/report/github.com/hamedcodelab/logger)
[![GoDoc](https://pkg.go.dev/badge/github.com/hamedcodelab/logger)](https://pkg.go.dev/github.com/hamedcodelab/logger)

## Overview
Logger is a lightweight and extensible structured logging package for Go. It provides easy-to-use, performant logging with support for different log levels and output formats.

## Features
- Simple and structured logging
- Configurable log levels (Debug, Info, Warn, Error, Fatal)
- JSON and plain text output formats
- Extensible with custom log handlers

## Installation
To install Logger, run:

```sh
go get github.com/hamedcodelab/logger
```

## Usage
### Basic Logging
```go
package main

import (
	"github.com/hamedcodelab/logger"
)

func main() {
	log := logger.New()
	log.Info("Application started")
	log.Error("Something went wrong", "error", "invalid input")
}
```

### Using JSON Format
```go
log := logger.New(logger.WithJSONFormat())
log.Info("Processing request", "method", "POST", "endpoint", "/api/user")
```

### Custom Log Level
```go
log := logger.New(logger.WithLevel(logger.DebugLevel))
log.Debug("Debugging mode enabled")
```

## Contributing
We welcome contributions! To get started:

1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Commit your changes (`git commit -m "Add new feature"`)
4. Push to your fork (`git push origin feature-branch`)
5. Open a pull request

### Guidelines
- Follow Go best practices
- Ensure compatibility with structured logging
- Write unit tests for new features
- Keep documentation up-to-date

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact
For questions or support, open an issue or reach out to [@hamedcodelab](https://github.com/hamedcodelab).

