package godom

import (
	"fmt"
	"runtime/debug"
	"syscall/js"
)

type DOM struct {
	document js.Value
}

// Error returned from possible operation failure
// Msg includes human readable details about the operation failure.
// Raw is the raw error object from the browser.
type Error = *rawError

type rawError = struct {
	Msg        string
	Stacktrace string
}

// Create a formatted error message for element errors and return the pointer
func domErrorf(msg string, a ...any) (err Error) {
	// Get data about the error
	Stacktrace := string(debug.Stack()[:])

	// Format provided message with any additional arguments
	msg = fmt.Sprintf(msg, a...)

	// Create the full error message
	formatMsg := fmt.Sprintf("Failed to %s - Did you forget to run DOM.init()", msg)

	return &rawError{
		Msg:        formatMsg,
		Stacktrace: Stacktrace,
	}
}

// Initialize a connection to the DOM.
func (d *DOM) Init() {
	global := js.Global()
	d.document = global.Get("document")
}

// Gets the body element of the document.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/body
func (d *DOM) Body() (elem Element, err Error) {
	defer func() {
		if r := recover(); r != nil {
			err = domErrorf("getting body element")
		}
	}()

	body := d.document.Get("body")
	return Element{body}, err
}

// Create any element in the DOM.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement
func (d *DOM) CreateElement(tag string) (elem Element, err Error) {
	defer func() {
		if r := recover(); r != nil {
			err = domErrorf("creating element %s", tag)
		}
	}()

	element := d.document.Call("createElement", tag)
	return Element{element}, err
}

// Selects a single element in the DOM.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector
func (d *DOM) QuerySelector(selector string) (elem Element, err Error) {
	defer func() {
		if r := recover(); r != nil {
			err = domErrorf("finding element with %s", selector)
		}
	}()

	element := d.document.Call("querySelector", selector)
	return Element{element}, err
}
