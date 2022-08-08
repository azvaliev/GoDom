package godom

// Get the title of the HTML document.
func (d DOM) GetTitle() (title string, err Error) {
	defer func() {
		if r := recover(); r != nil {
			err = domErrorf("getting document title")
		}
	}()

	return d.document.Get("title").String(), err
}

// Set the title of the HTML document.
func (d DOM) SetTitle(title string) (err Error) {
	defer func() {
		if r := recover(); r != nil {
			err = domErrorf("setting document title to %s", title)
		}
	}()

	d.document.Set("title", title)
	return err
}
