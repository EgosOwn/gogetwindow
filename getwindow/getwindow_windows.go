//go:build windows
// +build windows

package getwindow

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

var (
	mod               = windows.NewLazyDLL("user32.dll")
	procGetClassNameW = mod.NewProc("GetClassNameW")
)

type (
	HANDLE uintptr
	HWND   HANDLE
)

func getClassName(hwnd HWND) (name string, err error) {
	n := make([]uint16, 256)
	p := &n[0]
	r0, _, e1 := syscall.Syscall(procGetClassNameW.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(p)), uintptr(len(n)))
	if r0 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
		return
	}
	name = syscall.UTF16ToString(n)
	return
}

func getWindow(funcName string) uintptr {
	proc := mod.NewProc(funcName)
	hwnd, _, _ := proc.Call()
	return hwnd
}

func GetWindow() (string, error) {

	hwnd := getWindow("GetForegroundWindow")
	if hwnd != 0 {
		cn, _ := getClassName(HWND(hwnd))
		return cn, nil
	}
	return "", errors.New("Could not get window")

}
