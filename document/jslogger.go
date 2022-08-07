package godom

import (
	"fmt"
	"syscall/js"
)

// Log a message (msg) to the console.
// logType can be any type listed here (log, error, warn, etc)
// https://developer.mozilla.org/en-US/docs/Web/API/console
func Log(msg string, logType string) {
	js.Global().Get("console").Call(logType, msg)
}

// Log a custom format message (msg) to the console, works like fmt.Printf
// logType can be any type listed here (log, error, warn, etc)
// https://developer.mozilla.org/en-US/docs/Web/API/console/debug
// Just like fmt.Printf, you can provide format arguments to print variables
// Example:
// 	Logf("Hello %s", "log", "John Snow")
// If you don't want to specify the logType, other logging functions also format option
func Logf(msg string, logType string, f ...any) {
	formatMsg := fmt.Sprintf(msg, f...)
	js.Global().Get("console").Call(logType, formatMsg)
}

// Log a error message (msg) to the console.
// https://developer.mozilla.org/en-US/docs/Web/API/console/error
func LogError(msg string) {
	Log(msg, "error")
}

// Log a custom format error message (msg) to the console, works like Logf but with specified "error" logType
// Tip: This log is not fatal
// https://developer.mozilla.org/en-US/docs/Web/API/console/error
func LogErrorf(msg string, f ...any) {
	Logf(msg, "error", f...)
}

// Log a warning message (msg) to the console.
// https://developer.mozilla.org/en-US/docs/Web/API/console/warn
func LogWarn(msg string) {
	Log(msg, "warn")
}

// Log a custom format warning message (msg) to the console, works like Logf but with specified "warn" logType
// https://developer.mozilla.org/en-US/docs/Web/API/console/warn
func LogWarnf(msg string, f ...any) {
	Logf(msg, "warn", f...)
}

// Log an info message (msg) to the console.
// https://developer.mozilla.org/en-US/docs/Web/API/console/info
func LogInfo(msg string) {
	Log(msg, "info")
}

// Log a custom format info message (msg) to the console, works like Logf but with specified "info" logType
// https://developer.mozilla.org/en-US/docs/Web/API/console/info
func LogInfof(msg string, f ...any) {
	Logf(msg, "info", f...)
}

// Log a debug message (msg) to the console.
// https://developer.mozilla.org/en-US/docs/Web/API/console/debug
func LogDebug(msg string) {
	Log(msg, "debug")
}

// Log a custom format debug message (msg) to the console, works like Logf but with specified "debug" logType
// https://developer.mozilla.org/en-US/docs/Web/API/console/debug
func LogDebugf(msg string, f ...any) {
	Logf(msg, "debug", f...)
}

