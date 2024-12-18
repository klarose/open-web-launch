package gui

import (
	"os"
	"fmt"
)


func New() *GUI {
	return nil
}

func (gui *GUI) Start(windowTitle string) error {
	fmt.Printf("Starting gui, windowTitle: %s\n", windowTitle) 
	return nil
}

func (gui *GUI) SendTextMessage(text string) error {
	fmt.Printf("%s\n", text) 
	return nil
}

func (gui *GUI) SendErrorMessage(err error) error {
	fmt.Fprintln(os.Stderr, err)
	return nil
}

func (gui *GUI) SendCloseMessage() error {
	return nil
}

func (gui *GUI) SetTitle(title string) error {
	fmt.Printf("Setting Gui title: %s\n", title) 
	return nil
}

func (gui *GUI) SetProgressMax(val int) {
	return
}

func (gui *GUI) ProgressStep() {
	return
}

func (gui *GUI) Closed() bool {
	return false
}

func (gui *GUI) Terminate() error {
	return nil
}

