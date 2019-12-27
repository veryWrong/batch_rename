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
	"strings"
)

const (
	WinFontPath    = `C:\Windows\Fonts\SIMKAI.TTF`
	LinuxFontPath  = `/usr/share/fonts/truetype/arphic/ukai.ttc`
	defaultRandStr = "0123456789abcdefghijklmnopqrstuvwxyz!@$*_-"
)

var (
	pwd, _ = os.Getwd()
	conf   = defaultConf()
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
	entry.SetText(conf.wordDir)
	entry.SetPlaceHolder("输入要重命名文件的所在文件夹路径")
	box := widget.NewScrollContainer(widget.NewHBox(entry))
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(box, nil, nil, nil), box)
}

func RuleBox(w fyne.Window) fyne.CanvasObject {
	box := widget.NewVBox()
	fixBox := setFixBox(prefix, w)
	renameType := widget.NewRadio([]string{prefix, suffix, rename}, func(s string) {
		switch s {
		case prefix:
			fallthrough
		case suffix:
			fixBox.Hide()
			fixBox = setFixBox(s, w)
		case rename:
			fixBox.Hide()
			fixBox = renameBox()
			//box.Children = append(box.Children[:0], box.Children[0:]...)
		default:
			err := errors.New("请选择一种重命名方式")
			dialog.ShowError(err, w)
		}
		box.Children = append(box.Children[:1], box.Children[2:]...)
		box.Append(fixBox)
	})
	renameType.SetSelected(prefix)
	renameType.Horizontal = true
	fileType := widget.NewEntry()
	fileType.SetPlaceHolder("example: .jpg,.png")
	fileType.SetText(strings.Join(conf.fileType, ","))
	scroll := widget.NewScrollContainer(widget.NewHBox(fileType))
	contain := fyne.NewContainerWithLayout(layout.NewBorderLayout(scroll, nil, nil, nil), scroll)
	form := widget.NewForm()
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
	radio := widget.NewRadio([]string{numMold, wordMold, randMold}, func(s string) {
		rect := canvas.NewRectangle(&color.RGBA{R: 255, G: 255, B: 255, A: 0})
		rect.SetMinSize(fyne.NewSize(150, 0))
		content := widget.NewEntry()
		switch s {
		case numMold:
			content.SetPlaceHolder("example: 0-9")
			content.OnChanged = func(text string) {
				//fmt.Println("Entered", text)
			}
			scroll := widget.NewScrollContainer(widget.NewHBox(content))
			contain := fyne.NewContainerWithLayout(layout.NewBorderLayout(scroll, nil, nil, nil), scroll, rect)
			dialog.ShowCustom("数字范围", "确认", contain, w)
		case wordMold:
			content.SetPlaceHolder("example: a-z/A-Z")
			content.OnChanged = func(text string) {
				//fmt.Println("Entered", text)
			}
			scroll := widget.NewScrollContainer(widget.NewHBox(content))
			contain := fyne.NewContainerWithLayout(layout.NewBorderLayout(scroll, nil, nil, nil), scroll, rect)
			dialog.ShowCustom("字母范围", "确认", contain, w)
		case randMold:
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

func renameBox() fyne.CanvasObject {
	length := widget.NewEntry()
	length.SetText("5")
	scroll := widget.NewScrollContainer(widget.NewHBox(length))
	contain := fyne.NewContainerWithLayout(layout.NewBorderLayout(scroll, nil, nil, nil), scroll)
	form := widget.NewForm()
	form.Append("文件名长度:", contain)
	content := widget.NewEntry()
	content.SetText(defaultRandStr)
	content.SetPlaceHolder("example: 012...abcd...xyz")
	sc := widget.NewScrollContainer(widget.NewHBox(content))
	con := fyne.NewContainerWithLayout(layout.NewBorderLayout(sc, nil, nil, nil), sc)
	form.Append("随机内容范围:", con)
	return form
}

func StartButton() fyne.CanvasObject {
	button := widget.NewButtonWithIcon("开始", theme.MailSendIcon(), func() {

	})
	return button
}
