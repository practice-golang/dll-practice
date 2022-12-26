package main // import "hello"

import (
	"fmt"
	"syscall"
)

var ()

func main() {
	dllDISP := syscall.NewLazyDLL("displib.dll")
	procDispSizeX := dllDISP.NewProc("DispSizeX")
	procDispSizeY := dllDISP.NewProc("DispSizeY")

	dllHello := syscall.NewLazyDLL("hellolib.dll")
	procHello := dllHello.NewProc("Get1234")

	r, _, err := procHello.Call()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Hello:", r)

	// https://stackoverflow.com/a/48610796/8964990
	rx, _, err := procDispSizeX.Call()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("X:", rx)
	ry, _, err := procDispSizeY.Call()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Y:", ry)
}
