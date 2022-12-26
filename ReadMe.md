# gogetwindow


This simple module allows go programs to get the name of the currently active window. 

This is useful for creating application specific macros or tracking application usage.

For Windows it is a wrapper around the GetForegroundWindow() function in the user32.dll library, in Linux it uses xdotool.

It supports both Windows and Linux. Linux requires xdotool, Wayland based DEs are a serious PITA to support but I will take patches for it.


See example/ for how to use it.

## Help needed

xdotool is basically a hack so I welcome contributions to replace it with a better solution. I also welcome a Mac implementation.

This module is AGPLv3 licensed.