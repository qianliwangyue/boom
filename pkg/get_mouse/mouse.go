package get_mouse

import (
	"errors"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	procGetCursorPos = user32.NewProc("GetCursorPos")
)

type point struct {
	X, Y int32
}

func findMousePosWindows() (x, y int, err error) {
	var pt point
	ret, _, _ := procGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return 0, 0, errors.New("查询鼠标位置失败")
	}
	return int(pt.X), int(pt.Y), nil
}

func findMousePosLinux() (x, y int, err error) {
	conn, err := xgb.NewConn()
	if err != nil {
		return 0, 0, errors.New("连接 X server 失败: " + err.Error())
	}
	defer conn.Close()

	setup := xproto.Setup(conn)
	screen := setup.DefaultScreen(conn)
	root := screen.Root

	reply, err := xproto.QueryPointer(conn, root).Reply()
	if err != nil {
		return 0, 0, errors.New("查询鼠标位置失败: " + err.Error())
	}

	return int(reply.RootX), int(reply.RootY), nil
}

// GetOS 获取当前操作系统（linux / windows / darwin ...）
func GetOS() string {
	return runtime.GOOS
}

// FindMousePos 自动根据操作系统选择实现，获取鼠标屏幕坐标
func FindMousePos() (x, y int, err error) {
	switch runtime.GOOS {
	case "windows":
		return findMousePosWindows()
	case "linux":
		return findMousePosLinux()
	default:
		return 0, 0, errors.New("当前操作系统暂不支持: " + runtime.GOOS)
	}
}
