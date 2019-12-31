package gui

import (
	"errors"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"image/color"
	"os"
	"strconv"
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
	entry.SetText(conf.workDir)
	entry.SetPlaceHolder("输入要重命名文件的所在文件夹路径")
	entry.OnChanged = func(text string) {
		conf.setWorkDir(text)
	}
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
			conf.setRenameType(s)
		case rename:
			fixBox.Hide()
			fixBox = renameBox()
			conf.setRenameType(rename)
			conf.setRenameLength("5")
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
	fileType.OnChanged = func(text string) {
		conf.setFileType(text)
	}
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
	entry1.SetText(strconv.Itoa(conf.fixConf.lapse))
	entry1.OnChanged = func(text string) {
		conf.setFixLapse(text)
	}
	s1 := widget.NewScrollContainer(widget.NewHBox(entry1))
	con1 := fyne.NewContainerWithLayout(layout.NewBorderLayout(s1, nil, nil, nil), s1)
	text2 := widget.NewLabel("个文件添" + fix)
	radio := widget.NewRadio([]string{numMold, wordMold, randMold}, func(s string) {
		rect := canvas.NewRectangle(&color.RGBA{R: 255, G: 255, B: 255, A: 0})
		rect.SetMinSize(fyne.NewSize(150, 0))
		content := widget.NewEntry()
		switch s {
		case numMold:
			content.SetText("1-12")
			content.SetPlaceHolder("example: 0-9")
			content.OnChanged = func(text string) {
				conf.setFixConfine(text)
			}
			conf.setFixMold(numMold)
			conf.setFixConfine("1-12")
			scroll := widget.NewScrollContainer(widget.NewHBox(content))
			contain := fyne.NewContainerWithLayout(layout.NewBorderLayout(scroll, nil, nil, nil), scroll, rect)
			dialog.ShowCustom("数字范围", "确认", contain, w)
		case wordMold:
			content.SetText("a-Z")
			content.SetPlaceHolder("example: a-z/A-Z/a-Z")
			content.OnChanged = func(text string) {
				conf.setFixConfine(text)
			}
			conf.setFixMold(wordMold)
			conf.setFixConfine("a-Z")
			scroll := widget.NewScrollContainer(widget.NewHBox(content))
			contain := fyne.NewContainerWithLayout(layout.NewBorderLayout(scroll, nil, nil, nil), scroll, rect)
			dialog.ShowCustom("字母范围", "确认", contain, w)
		case randMold:
			content.SetText(defaultRandStr)
			content.SetPlaceHolder("example: 012...abcd...xyz")
			content.OnChanged = func(text string) {
				conf.setFixConfine(text)
			}
			conf.setFixMold(randMold)
			conf.setFixConfine(defaultRandStr)
			scroll := widget.NewScrollContainer(widget.NewHBox(content))
			contain := fyne.NewContainerWithLayout(layout.NewBorderLayout(scroll, nil, nil, nil), scroll, rect)
			dialog.ShowCustom("随机范围", "确认", contain, w)
		default:
			err := errors.New("请选择一种格式")
			dialog.ShowError(err, w)
		}
	})
	radio.SetSelected(numMold)
	radio.Horizontal = true
	box := widget.NewHBox(text1, con1, text2, radio)
	return box
}

func renameBox() fyne.CanvasObject {
	length := widget.NewEntry()
	length.SetText(strconv.Itoa(conf.renameConf.length))
	length.OnChanged = func(text string) {
		conf.setRenameLength(text)
	}
	scroll := widget.NewScrollContainer(widget.NewHBox(length))
	contain := fyne.NewContainerWithLayout(layout.NewBorderLayout(scroll, nil, nil, nil), scroll)
	form := widget.NewForm()
	form.Append("文件名长度:", contain)
	content := widget.NewEntry()
	content.SetText(conf.renameConf.confine)
	content.SetPlaceHolder("example: 012...abcd...xyz")
	content.OnChanged = func(text string) {
		conf.setRenameConfine(text)
	}
	sc := widget.NewScrollContainer(widget.NewHBox(content))
	con := fyne.NewContainerWithLayout(layout.NewBorderLayout(sc, nil, nil, nil), sc)
	form.Append("随机内容范围:", con)
	return form
}

func StartButton(w fyne.Window) fyne.CanvasObject {
	button := widget.NewButtonWithIcon("开始", theme.MailSendIcon(), func() {
		fmt.Println(conf)
		conf.do(w)
	})
	return button
}
