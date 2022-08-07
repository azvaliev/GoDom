package godom

import (
	"fmt"
	"syscall/js"
)

// An HTML element.
type Element struct {
	raw js.Value
}

func handleElementError(msg string) {
	if err := recover(); err != nil {
		LogError(fmt.Sprintf("Error %s - Does this element exist? You can check with Element.Exists():\n" + err.(error).Error(), msg))
	}
}

// Check if DOM element exists
func (e *Element) Exists() bool {
	return !e.raw.IsNull()
}

// Get the inner text of the element.
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/innerText
func (e *Element) GetInnerText() string {
	defer handleElementError("getting inner text")
	return e.raw.Get("innerText").String()
}

// Set the inner text of the element
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/innerText
func (e *Element) SetInnerText(text string) {
	defer handleElementError("setting inner text")
	e.raw.Set("innerText", text)
}

// Add a child element to the element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/appendChild
func (e *Element) AppendChild(child Element) {
	defer handleElementError("appending child")
	e.raw.Call("appendChild", child.raw)
}

// Get any attribute on the element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getAttribute
func (e *Element) GetAttribute(name string) string {
	defer handleElementError(fmt.Sprintf("getting attribute %s", name))
	return e.raw.Get(name).String()
}

// Set any attribute on the element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/setAttribute
func (e *Element) SetAttribute(name string, value string) {
	defer handleElementError(fmt.Sprintf("setting attribute %s", name))
	e.raw.Set(name, value)
}

