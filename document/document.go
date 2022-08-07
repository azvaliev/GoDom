package godom

import (
	"fmt"
	"syscall/js"
)

type DOM struct {
	document js.Value
}

func handleDOMError(msg string) {
	if err := recover(); err != nil {
		LogError(fmt.Sprintf("Error %s - Did you forget to run DOM.init():\n" + err.(error).Error(), msg))
	}
}

// Initialize a connection to the DOM.
func (d *DOM) Init() {
	global := js.Global()
	d.document = global.Get("document")
}

// Gets the body element of the document.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/body
func (d *DOM) Body() Element {
	defer handleDOMError("getting body element")
	body := d.document.Get("body")
	return Element{body}
}

// Create any element in the DOM.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement
func (d *DOM) CreateElement(tag string) Element {
	defer handleDOMError(fmt.Sprintf("creating element %s", tag))
	element := d.document.Call("createElement", tag)
	return Element{element}
}
