package godom

import (
	"syscall/js"
)

// Event listener options, optional
type EventListenerOptions struct {
	Capture, Once, Passive bool
}

// Event interface,
//
//	https://developer.mozilla.org/en-US/docs/Web/API/Event
type Event = js.Value

// Event name
//
//	https://developer.mozilla.org/en-US/docs/Web/Events#event_listing
type EventName = string

// Event listener callback function,
//
//	https://developer.mozilla.org/en-US/docs/Web/API/EventListener
type EventHandler = func(this *Element, event Event)

// Cleanup function to remove event liseners, similar to removeEventListener but no arguments are needed
//
//	https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/removeEventListener
type RemoveEventListener = func()

// Add an event listener to the element
//
//	https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// returns a cleanup function like removeEventListener but no arguments are needed
func (e *Element) AddEventListener(event EventName, callback EventHandler) (cleanup RemoveEventListener, err Error) {
	return e.AddEventListenerWithOptions(event, callback, EventListenerOptions{})
}

// Add an event listener to the element with options of type EventListenerOptions,
//
//	https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// returns a cleanup function like removeEventListener but no arguments are needed
func (e *Element) AddEventListenerWithOptions(event EventName, callback EventHandler, eventOptions EventListenerOptions) (cleanup RemoveEventListener, err Error) {
	defer func() {
		if r := recover(); r != nil {
			err = elementErrorf("set event listener of type %s", event)
		}
	}()

	// Create a valid JS function go run Go callback
	jsCallback := js.FuncOf(func(this js.Value, args []js.Value) any {
		callback(e, args[0])
		return nil
	})

	// Convert eventOptions to map[string]interface{}
	eventOptionsMap := map[string]interface{}{}
	if eventOptions.Capture {
		eventOptionsMap["capture"] = true
	}
	if eventOptions.Once {
		eventOptionsMap["once"] = true
	}
	if eventOptions.Passive {
		eventOptionsMap["passive"] = true
	}

	// Create a valid JS object from options
	jsEventOptions := js.Global().Get("Object").Call("create", eventOptionsMap)
	// jsEventOptions.Set("capture", eventOptions.Capture)
	// jsEventOptions.Set("once", eventOptions.Once)
	// jsEventOptions.Set("passive", eventOptions.Passive)

	e.raw.Call("addEventListener", event, jsCallback, jsEventOptions)

	// Create and return a cleanup function to remove the event listener
	return func() {
		e.raw.Call("removeEventListener", event, jsCallback)
		jsCallback.Release()
	}, err
}
