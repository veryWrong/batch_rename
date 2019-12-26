package gui

import (
	"errors"
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"image/color"
	"os"
)

var (
	pwd, _ = os.Getwd()
)

func DirBar() fyne.CanvasObject {
	folder := widget.NewButtonWithIcon(pwd, theme.FolderIcon(), nil)
	box := widget.NewHBox(folder)
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(box, nil, nil, nil), box)
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
	entry.SetText(pwd)
	entry.SetPlaceHolder("输入要重命名文件的所在文件夹路径")
	box := widget.NewScrollContainer(widget.NewHBox(entry))
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(box, nil, nil, nil), box)
}

func RuleBox(w fyne.Window) fyne.CanvasObject {
	var reType = "加前缀"
	box := widget.NewVBox()
	fixBox := setFixBox(reType, w)
	renameType := widget.NewRadio([]string{"加前缀", "加后缀", "重命名"}, func(s string) {
		switch s {
		case "加前缀":
			fallthrough
		case "加后缀":
			fixBox.Hide()
			fixBox = setFixBox(s, w)
		case "重命名":
			fixBox.Hide()
			box.Children = append(box.Children[:0], box.Children[0:]...)
		default:
			//box.Children = box.Children[:1]
			err := errors.New("请选择一种重命名方式")
			dialog.ShowError(err, w)
		}
		box.Children = append(box.Children[:1], box.Children[2:]...)
		box.Append(fixBox)
	})
	renameType.SetSelected("加前缀")
	renameType.Horizontal = true
	form := widget.NewForm()
	fileType := widget.NewEntry()
	fileType.SetPlaceHolder("example: .jpg,.png")
	scroll := widget.NewScrollContainer(widget.NewHBox(fileType))
	contain := fyne.NewContainerWithLayout(layout.NewBorderLayout(scroll, nil, nil, nil), scroll)
	form.Append("文件类型:", contain)
	form.Append("重命名方式:", renameType)
	box.Append(form)
	box.Append(fixBox)
	return box
}

func setFixBox(fix string, w fyne.Window) fyne.CanvasObject {
	text1 := widget.NewLabel("间隔")
	entry1 := widget.NewEntry()
	entry1.SetText("1")
	s1 := widget.NewScrollContainer(widget.NewHBox(entry1))
	con1 := fyne.NewContainerWithLayout(layout.NewBorderLayout(s1, nil, nil, nil), s1)
	text2 := widget.NewLabel("个文件添" + fix)
	radio := widget.NewRadio([]string{"数字", "字母", "随机"}, func(s string) {
		rect := canvas.NewRectangle(&color.RGBA{R: 255, G: 255, B: 255, A: 0})
		rect.SetMinSize(fyne.NewSize(150, 0))
		content := widget.NewEntry()
		switch s {
		case "数字":
			content.SetPlaceHolder("example: 0-9")
			content.OnChanged = func(text string) {
				//fmt.Println("Entered", text)
			}
			scroll := widget.NewScrollContainer(widget.NewHBox(content))
			contain := fyne.NewContainerWithLayout(layout.NewBorderLayout(scroll, nil, nil, nil), scroll, rect)
			dialog.ShowCustom("数字范围", "确认", contain, w)
		case "字母":
			content.SetPlaceHolder("example: a-z/A-Z")
			content.OnChanged = func(text string) {
				//fmt.Println("Entered", text)
			}
			scroll := widget.NewScrollContainer(widget.NewHBox(content))
			contain := fyne.NewContainerWithLayout(layout.NewBorderLayout(scroll, nil, nil, nil), scroll, rect)
			dialog.ShowCustom("字母范围", "确认", contain, w)
		case "随机":
			content.SetPlaceHolder("example: 012...abcd...xyz")
			content.OnChanged = func(text string) {
				//fmt.Println("Entered", text)
			}
			scroll := widget.NewScrollContainer(widget.NewHBox(content))
			contain := fyne.NewContainerWithLayout(layout.NewBorderLayout(scroll, nil, nil, nil), scroll, rect)
			dialog.ShowCustom("随机范围", "确认", contain, w)
		default:
			err := errors.New("请选择一种格式")
			dialog.ShowError(err, w)
		}
	})
	radio.Horizontal = true
	box := widget.NewHBox(text1, con1, text2, radio)
	return box
}
