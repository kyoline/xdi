package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/awesome-gocui/gocui"
)

var (
	clickCount int = 0
)

func quit(*gocui.Gui, *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	// Initialize the GUI.
	gui, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Fatalf("Failed to initialize GUI: %v.", err)
	}
	defer gui.Close()

	gui.SetManagerFunc(layout)

	gui.Cursor = true
	gui.Mouse = true

	// This key binding allows you to exit the application using `Ctrl+C`.
	if err := gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Fatalf("Failed to set exit key combination: %v.", err)
	}

	// Start the application main loop.
	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalf("Failed to start GUI main loop: %v.", err)
	}
}

func layout(gui *gocui.Gui) error {
	button, err := gui.SetView("button", 0, 2, 11, 6, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return errors.Join(errors.New("failed to initialize button"), err)
		}

		button.Frame = false
		button.BgColor = gocui.ColorBlue
		button.FgColor = gocui.ColorWhite
		fmt.Fprint(button, "\n Click me")

		if err := gui.SetKeybinding("button", gocui.MouseLeft, gocui.ModNone, increaseClickCount); err != nil {
			return errors.Join(errors.New("failed to initialize button click listener"), err)
		}
	}

	// This is where we will add our views to our application.
	clickCountView, err := gui.SetView("click-count", 0, 0, 9, 2, 0)
	if err != nil && err != gocui.ErrUnknownView {
		return errors.Join(errors.New("failed to create click count view"), err)
	}

	clickCountView.Clear()
	fmt.Fprintf(clickCountView, "%8d", clickCount)

	return nil
}

func increaseClickCount(gui *gocui.Gui, view *gocui.View) error {
	clickCount++

	if (clickCount & 1) == 0 {
		view.BgColor = gocui.ColorBlue
	} else {
		view.BgColor = gocui.ColorGreen
	}

	return nil
}
