package main

import (
	"fmt"

	"github.com/egosown/gogetwindow"
)

func main() {
	win, err := getwindow.GetWindow()
	if err != nil {
		fmt.Println("Could not get window")
	} else{
		fmt.Println("Your active window is " + win)
	}

}
