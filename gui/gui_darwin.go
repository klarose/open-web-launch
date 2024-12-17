package gui

import "fmt"


func New() *GUI {
	return nil
}

func (gui *GUI) Start(windowTitle string) error {
	fmt.Printf("Starting gui, windowTitle: %s", windowTitle) 
	return nil
}

func (gui *GUI) SendTextMessage(text string) error {
	fmt.Printf("Sending gui text message: %s", text) 
	return nil
}

func (gui *GUI) SendErrorMessage(err error) error {
	fmt.Printf("Sending gui error message: %v", err) 
	return nil
}

func (gui *GUI) SendCloseMessage() error {
	fmt.Printf("Send Gui closed message")
	return nil
}

func (gui *GUI) SetTitle(title string) error {
	fmt.Printf("Seting Gui title: %s", title) 
	return nil
}

func (gui *GUI) SetProgressMax(val int) {
	fmt.Printf("Setting Gui Progress max")
	return
}

func (gui *GUI) ProgressStep() {
	fmt.Printf("Progressing Gui step")
	return
}

func (gui *GUI) Closed() bool {
	fmt.Printf("Closing Gui")
	return false
}

func (gui *GUI) Terminate() error {
	fmt.Printf("Terminating Gui")
	return nil
}

