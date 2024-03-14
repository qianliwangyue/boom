package pages

import (
	"boom/pkg/get"
	"boom/pkg/send_sms"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"time"
)

var (
	Y int
	X int
)

func PageBoom(app fyne.App) {
	window := app.NewWindow("信息轰炸")
	t1 := widget.NewEntry()
	t1.SetPlaceHolder("请输入信息内容...")
	t2 := widget.NewEntry()
	l1 := widget.NewLabel("将鼠标移动到输入框处，点击Esc键，\n获取信息输入框的坐标位置")
	l2 := widget.NewLabel("")
	l3 := widget.NewLabel("请输入轰炸次数:")
	var b2Clicked bool
	b2 := widget.NewButton("开始轰炸", func() {
		b2Clicked = true
	})
	content := container.NewVBox(
		container.NewGridWrap(fyne.NewSize(300, 120), t1),
		l1, l2,
		container.New(layout.NewGridLayoutWithColumns(2), l3, t2),

		container.NewCenter(b2),
	)
	window.SetContent(content)
	//
	go func() {
		for {
			if b2Clicked {
				b2.Disable() // 如果 clicked 为 true，禁用按钮
				//
				n, err := strconv.Atoi(t2.Text)
				if err != nil {
					//fmt.Println("=====", 1)
					t2.SetPlaceHolder("输入错误,请输入正确的数字")
					b2Clicked = false
					b2.Enable()
					window.Canvas().Refresh(b2) // 刷新按钮状态
					continue
				}
				send_sms.AutoSendMessage(t1.Text, X, Y, n)
				//fmt.Println("=======", X, Y)
				//
				b2Clicked = false
				b2.Enable()
				window.Canvas().Refresh(b2) // 刷新按钮状态
			}
			time.Sleep(time.Second) // 每秒检查一次 clicked 变量
		}
	}()

	window.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) {
		//fmt.Println("==========", e.Name)
		if e.Name == fyne.KeyEscape {
			var err error
			X, Y, err = get.FindMousePos()

			if err != nil {
				l2.SetText("获取坐标失败:" + err.Error())
			} else {
				l2.SetText("输入框坐标:" + strconv.Itoa(X) + "," + strconv.Itoa(Y))
			}
		}
	})

	//
	window.Resize(fyne.NewSize(120, 350))
	window.Show()
}
