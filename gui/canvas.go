package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"os"
)

var (
	pwd, _      = os.Getwd()
	copyContent string
)

func DirBar() fyne.CanvasObject {
	folder := widget.NewButtonWithIcon(pwd, theme.FolderIcon(), nil)
	replicate := widget.NewButtonWithIcon("复制", theme.ContentCopyIcon(), func() {
		copyContent = pwd
	})
	box := widget.NewHBox(folder, replicate)
	return fyne.NewContainerWithLayout(
		layout.NewBorderLayout(box, nil, nil, nil),
		box,
	)
}

func ThemeRadio(a fyne.App) fyne.CanvasObject {
	radio := widget.NewRadio([]string{"Light", "Dark"}, func(s string) {
		if s == "Dark" {
			a.Settings().SetTheme(theme.DarkTheme())
		} else {
			a.Settings().SetTheme(theme.LightTheme())
		}
	})
	radio.SetSelected("Light")
	radio.Horizontal = true
	return radio
}

func DirInput() fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("输入要重命名文件的所在文件夹路径")
	paste := widget.NewButtonWithIcon("粘贴", theme.ContentPasteIcon(), func() {
		entry.SetText(copyContent)
	})
	box := widget.NewHBox(entry, paste)
	return fyne.NewContainerWithLayout(
		layout.NewBorderLayout(box, nil, nil, nil),
		box,
	)
}
