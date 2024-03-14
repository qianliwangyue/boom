package main

import (
	"boom/pages"
	"boom/theme"
	"fyne.io/fyne/v2/app"
)

func main() {
	App := app.New()
	App.Settings().SetTheme(&theme.MyTheme{})
	pages.PageBoom(App)
	App.Run()
}
