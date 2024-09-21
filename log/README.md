# Logger Library
A simple logger library for Go applications.

This logger library provides an easy-to-use interface for logging messages at different log levels in your Go applications. It offers options to set the log output file and log level based on your needs.

## Installation
To use this logger library in your Go project, you need to import it:

```go
import "github.com/rahulshewale153/common-lib/log"
```

## Usage
### Initialize Logger
The logger is initialized automatically upon importing the library. You can access the logger through the exported functions.

```go
log.SetOutputFile("mylog.log") // Set log output file
log.SetLogLevel(log.DebugLevel) // Set log level
```

### Log Levels
The logger supports different log levels:

- Trace
- Debug
- Info
- Warn
- Error
- Critical
- Fatal

> Fatal level logs the message and exits the process, while Critical level does not. Use the Critical level when the app can continue running, but it will have a limited impact on the server and requires immediate attention.

Use these functions to log messages at various levels.

```go
log.Info("This is an info message.")
log.Warn("This is a warning message.")
```

### Example
Here's a simple example of using the logger in your application:

```go
package main

import (
    "github.com/rahulshewale153/common-lib/log"
)

func main() {
    log.SetOutputFile("app.log")
    log.SetLogLevel(log.DebugLevel)

    log.Info("Starting the application...")
    log.Debug("Debugging information...")
    log.Error("An error occurred.")

    log.Info("Application finished.")
}
```