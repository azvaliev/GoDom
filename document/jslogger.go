package godom

import (
	"syscall/js"
)

// Log a message to the console.
func Log(msg string, logType string) {
	js.Global().Get("console").Call(logType, msg)
}

// Log a error message to the console.
func LogError(msg string) {
	js.Global().Get("console").Call("error", msg)
}

// Log a warning message to the console.
func LogWarn(msg string) {
	js.Global().Get("console").Call("warn", msg)
}

// Log an info message to the console.
func LogInfo(msg string) {
	js.Global().Get("console").Call("info", msg)
}

// Log a debug message to the console.
func LogDebug(msg string) {
	js.Global().Get("console").Call("debug", msg)
}
