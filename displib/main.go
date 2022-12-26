package main

import "C"

import (
	"fmt"
	"syscall"
)

const (
	SM_CXSCREEN = 0
	SM_CYSCREEN = 1
)

var (
	user32           = syscall.NewLazyDLL("User32.dll")
	getSystemMetrics = user32.NewProc("GetSystemMetrics")
)

func GetSystemMetrics(nIndex int) int {
	index := uintptr(nIndex)
	ret, _, _ := getSystemMetrics.Call(index)
	return int(ret)
}

//export DispSizeX
func DispSizeX() int {
	return GetSystemMetrics(SM_CXSCREEN)
}

//export DispSizeY
func DispSizeY() int {
	return GetSystemMetrics(SM_CYSCREEN)
}

func main() {
	width := DispSizeX()
	height := DispSizeY()

	fmt.Printf("%dx%d\n", width, height)
}
