package godom

import (
	"fmt"
	"runtime/debug"
	"syscall/js"
)

// An HTML element.
type Element struct {
	raw js.Value
}

// Create a formatted error message for element errors and return the pointer
func elementErrorf(msg string, a ...any) (err Error) {
	// Get data about the error
	Stacktrace := string(debug.Stack()[:])

	// Format provided message with any additional arguments
	msg = fmt.Sprintf(msg, a...)

	// Create the full error message
	formatMsg := fmt.Sprintf("Failed to %s - Are you sure this element exists? You can check with Element.Exists()", msg)

	return &rawError{
		Msg:        formatMsg,
		Stacktrace: Stacktrace,
	}
}

// Check if DOM element exists
func (e *Element) Exists() bool {
	return !e.raw.IsNull()
}

// Get the inner text of the element.
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/innerText
func (e *Element) GetInnerText() (innerText string, err Error) {
	defer func() {
		if r := recover(); r != nil {
			err = elementErrorf("get inner text of element")
		}
	}()

	innerText = e.raw.Get("innerText").String()
	return innerText, err
}

// Set the inner text of the element
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/innerText
func (e *Element) SetInnerText(text string) (err Error) {
	defer func() {
		if r := recover(); r != nil {
			err = elementErrorf("set inner text of element to ", text)
		}
	}()

	e.raw.Set("innerText", text)
	return err
}

// Add a child element to the element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/appendChild
func (e *Element) AppendChild(child Element) (err Error) {
	defer func() {
		if r := recover(); r != nil {
			err = elementErrorf("append child to element")
		}
	}()

	e.raw.Call("appendChild", child.raw)
	return err
}

// Get any attribute on the element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getAttribute
func (e *Element) GetAttribute(attributeName string) (attributeValue string, err Error) {
	defer func() {
		if r := recover(); r != nil {
			err = elementErrorf("get attribute %s", attributeName)
		}
	}()

	return e.raw.Get(attributeName).String(), err
}

// Set any attribute on the element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/setAttribute
func (e *Element) SetAttribute(attributeName string, value string) (err Error) {
	defer func() {
		if r := recover(); r != nil {
			err = elementErrorf("set attribute %s", attributeName)
		}
	}()

	e.raw.Set(attributeName, value)
	return err
}
