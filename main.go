package main

import (
	"batch_rename/gui"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"os"
	"runtime"
)

func init() {
	if runtime.GOOS == `windows` {
		os.Setenv("FYNE_FONT", gui.WinFontPath)
	}
}

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("批量重命名工具")
	spaceBox := widget.NewHBox(widget.NewToolbarSeparator().ToolbarObject())
	labelBox := widget.NewVBox(
		widget.NewLabel("主题:"),
		widget.NewLabel("当前路径:"), spaceBox,
		widget.NewLabel("工作目录:"),
	)
	contentBox := widget.NewVBox(
		gui.ThemeRadio(a),
		gui.DirBar(),
		gui.DirInput(),
	)
	content := widget.NewHBox(labelBox, contentBox)
	w.SetContent(content)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}