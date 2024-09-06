## What is a Log, Logging, a Log File and a Logger, in programming??
- A log is a record of events during the execution of a program.
- A Log File is one of the most common places where the logs are recorded. It is the most common output stream used during logging.
- Logging is the process of recording the logs (events that happen during execution) to a suitable output stream. There are various locations where logs are written, to log files, to databases, to terminals or command lines, to a dedicated UI (page) in the application when running (example CI tools), etc.
- Then finally, a Logger is the piece of code which is responsible in handling the logging process in a software application (or in programming).

## Purpose of Logging
- Loggers or Logging is really helpful and is very much a recommended practice in the context of programming. The main purpose of Logging is to generate a history of events executed in a software application and it helps any developer to go through the history of events logged and identify what has exactly happened, and it is most useful when an error or unexpected event occurs, developers are able to easily identify the error and trace backwards to identify the root cause of the error. In other words, logging makes the error identification process or debugging process much easier and straight forward for developers.

## What to Log?
 An interesting question, different people may have different opinions. You can log anything and everything, but you got to understand the purpose of why you need a log for your application and then decide what to log. No real straight forward answer but, depends on the requirement and use-case.

 Typically, when it comes to logging the general purpose as mentioned above is to trace back errors. Hence, it is important to log specific and useful event in the program and also focus on not to flood the log with not so useful information and also not to miss-out on useful information.

 Formatting the logger, i.e. deciding what information to log, can also be tricky. According to me, it is very helpful to log, the type/level of logger, date and time at which the event occurred, file in which the event occurred, function in which the event occurred (not should have, but nice to have), line of code responsible, developer message, error message given by the program (if logging an error), full error stack-trace given by the program (if logging an error). è Basically says; when, where and what event took place.


## Implementation

### LogLevel Type and Constants

```go
type LogLevel int

const (
	INFO LogLevel = iota
	WARN
	ERROR
	DEBUG
)
```

- **LogLevel**: This custom type (`LogLevel`) represents the different log levels our logger can handle. 
- **Constants**: Using `iota`, we define log levels (`INFO`, `WARN`, `ERROR`, and `DEBUG`) as constants. `iota` automatically assigns successive integer values starting from 0.

### Logger Struct

```go
type Logger struct {
	mu     sync.Mutex
	level  LogLevel
	output io.Writer
}
```

- **Logger Struct**: It represents the core of the logging system.
  - `mu sync.Mutex`: Ensures thread-safe logging by preventing concurrent writes from overlapping and corrupting the log output.
  - `level LogLevel`: Defines the minimum log level that the logger will handle. Messages below this level are ignored.
  - `output io.Writer`: Specifies where the logs will be written (e.g., console, file). This allows flexibility in directing log output.

### NewLogger Function

```go
func NewLogger(output io.Writer, level loglevel) *Logger {
	return &Logger{
		output: output,
		level: level,
	}
}

```

- **NewLogger**: This constructor function initializes a new `Logger` instance.
  - `level`: Sets the log level filter.
  - `output`: Sets the output destination (`os.Stdout` for console, file, etc.).

### log Method

```go
func (l *Logger) log(level loglevel, msg string) error{
	l.mu.Lock()
	defer l.mu.Unlock()
	if level < l.level {
		return nil
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	levelStr := [...]string{"INFO", "WARN", "ERROR", "DEBUG"}[level]
	_, err := l.output.Write([]byte(timestamp + " [" + levelStr + "] " + msg + "\n"))
	return err
}
```

- **`log` Method**: Handles the actual logging process and is called internally by other logging methods (`Info`, `Warn`, `Error`, `Debug`).
  - **Locking (`l.mu.Lock()`)**: Ensures thread safety by locking access until the log operation is complete.
  - **Level Check (`if level < l.level`)**: Only logs messages that meet or exceed the configured log level. If the message’s level is lower, it is skipped.
  - **Timestamp**: Uses the current time formatted as `YYYY-MM-DD HH:MM:SS`.
  - **Level String Mapping**: Converts the `LogLevel` constant to a human-readable string (`INFO`, `WARN`, etc.).
  - **Log Message Formatting**: Constructs the final log message string, combining the timestamp, log level, and actual message.
  - **Writing to Output**: Sends the formatted message to the designated output (`l.output`).

### Logging Methods (`Info`, `Warn`, `Error`, `Debug`)

```go

func (l *Logger) Info(msg string) error {
	return l.log(INFO, msg)
}
func (l *Logger) Warn(msg string) error {
	return l.log(WARN, msg)
}
func (l *Logger) Error(msg string) error {
	return l.log(ERROR, msg)
}
func (l *Logger) Debug(msg string) error {
	return l.log(DEBUG, msg)
}

```

- **Logging Methods**: These methods provide convenient, level-specific logging functions.
  - **Info**: Logs an `INFO` level message.
  - **Warn**: Logs a `WARN` level message.
  - **Error**: Logs an `ERROR` level message.
  - **Debug**: Logs a `DEBUG` level message.

Each method calls the internal `log` method with the appropriate level, ensuring that the correct messages are logged according to the logger's configuration.

