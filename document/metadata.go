package godom

// Get the title of the HTML document.
func (d DOM) GetTitle() string {
	defer func() {
		if err := recover(); err != nil {
			LogError("Error getting title - Did you forget to run DOM.init():\n" + err.(error).Error())

		}
	}()
	return d.document.Get("title").String()
}

// Set the title of the HTML document.
func (d DOM) SetTitle(title string) {
	defer func() {
		if err := recover(); err != nil {
			LogError("Error setting title - Did you forget to run DOM.init():\n" + err.(error).Error())

		}
	}()
	d.document.Set("title", title)
}
