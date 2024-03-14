package send_sms

import (
	"github.com/go-vgo/robotgo"
	"time"
)

// AutoSendMessage 自动发送轰炸信息
func AutoSendMessage(message string, x, y, n int) {
	robotgo.MoveClick(x, y)
	time.Sleep(2 * time.Second)
	for range n {
		robotgo.TypeStr(message)
		time.Sleep(time.Millisecond * 500)
		_ = robotgo.KeyTap(robotgo.Enter)
	}
}
