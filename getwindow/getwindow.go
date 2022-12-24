//go:build linux
// +build linux

package getwindow

import "os/exec"

func GetWindow() (string, error) {
	win, err := exec.Command("xdotool getactivewindow getwindowname").Output()
	if err != nil {
		return "", err
	}
	return string(win), nil
}
