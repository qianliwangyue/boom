package get

/*
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>
*/
import "C"
import (
	"errors"
)

// FindMousePos 获取所有进程中的鼠标位置
func FindMousePos() (x, y int, err error) {
	display := C.XOpenDisplay(nil) //打开显示连接
	if display == nil {
		return 0, 0, errors.New("当前打开的显示设备连接失败")
	}
	defer C.XCloseDisplay(display) //关闭显示连接

	var rootWindow C.Window
	rootWindow = C.XRootWindow(display, 0)

	var rootX, rootY C.int
	var winX, winY C.int
	var mask C.uint
	success := C.XQueryPointer(
		display,
		rootWindow,
		&rootWindow,
		&rootWindow,
		&rootX,
		&rootY,
		&winX,
		&winY,
		&mask,
	)
	if success == 0 {
		return 0, 0, errors.New("查询鼠标位置失败")
	}
	//fmt.Printf("Mouse position: (%d, %d)\n", int(rootX), int(rootY))
	return int(rootX), int(rootY), nil
}
