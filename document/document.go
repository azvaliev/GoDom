package godom

import (
	"syscall/js"
)

type DOM struct {
	document js.Value
}

// Initialize a connection to the DOM.
func (d *DOM) Init() {
	global := js.Global()
	d.document = global.Get("document")
}
