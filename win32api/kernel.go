package win32api

import (
	"syscall"
	"unsafe"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll") //user32.dll
	systemParametersInfo = user32.NewProc("SystemParametersInfoW")
)

// SetWallpaper can set windows wallpaper
func SetWallpaper(imageLoc string) bool {
	imageLocPtr, err := syscall.UTF16PtrFromString(imageLoc)
	if err != nil {
		return false
	}
	ret, _, _ := systemParametersInfo.Call(
		20,
		0,
		uintptr(unsafe.Pointer(imageLocPtr)), 1)
	if ret != 0 {
		return true
	}
	return false
}
