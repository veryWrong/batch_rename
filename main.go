package main

import (
	"batch_rename/gui"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"os"
	"runtime"
)

func init() {
	if runtime.GOOS == `windows` {
		os.Setenv("FYNE_FONT", gui.WinFontPath)
	} else if runtime.GOOS == `linux` {
		os.Setenv("FYNE_FONT", gui.LinuxFontPath)
	}
}

// go build -ldflags="-H windowsgui"
func main() {
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	theme.TextSize()
	w := a.NewWindow("批量重命名工具")
	spaceBox := widget.NewHBox(widget.NewToolbarSeparator().ToolbarObject())
	labelBox := widget.NewVBox(
		widget.NewLabel("主题:"),
		widget.NewLabel("当前路径:"), spaceBox, spaceBox,
		widget.NewLabel("工作目录:"),
	)
	contentBox := widget.NewVBox(
		gui.ThemeRadio(a),
		gui.DirBar(),
		gui.DirInput(),
	)
	group := widget.NewGroup("重命名规则", gui.RuleBox(w))
	iBox := widget.NewHBox(labelBox, contentBox)
	content := widget.NewVBox(iBox, group, gui.StartButton(w))
	w.SetContent(content)
	//w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
