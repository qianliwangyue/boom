package send_sms

import (
	"github.com/go-vgo/robotgo"
	"github.com/go-vgo/robotgo/clipboard"
	"time"
)

// AutoSendMessage 自动发送轰炸信息
func AutoSendMessage(message string, x, y, n int) {
	robotgo.MoveClick(x, y)
	time.Sleep(1 * time.Second)
	_ = clipboard.WriteAll(message)
	for range n {
		_ = robotgo.KeyToggle(robotgo.Ctrl, "down")
		_ = robotgo.KeyTap(robotgo.KeyV)
		_ = robotgo.KeyToggle(robotgo.Ctrl, "up")
		time.Sleep(10 * time.Millisecond)
		_ = robotgo.KeyTap(robotgo.Enter)
	}

}
