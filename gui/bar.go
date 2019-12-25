package gui

import (
	"github.com/golang/freetype"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

const WinFontPath = `C:\Windows\Fonts\SIMKAI.TTF`
const LinuxFontPath = `/usr/share/fonts/truetype/arphic/ukai.ttc`

type PwdBar struct {
	content []byte
}

func (p *PwdBar) Name() string {
	return "PwdBar"
}

func (p *PwdBar) Content() []byte {
	return p.content
}

func NewPwdBar(pwd string) *PwdBar {
	if err := pwdPng(pwd); err != nil {
		log.Print(err)
		return nil
	}
	file, err := os.Open("pwd.png")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer file.Close()
	content, _ := ioutil.ReadAll(file)
	return &PwdBar{content: content}
}

func pwdPng(pwd string) error {
	file, err := os.Create("pwd.png")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer file.Close()
	img := image.NewNRGBA(image.Rect(0, 0, 300, 28))
	for y := 0; y < 28; y++ {
		for x := 0; x < 300; x++ {
			//设置一块 白色(255,255,255)透明的背景
			img.Set(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 0})
		}
	}
	//读取字体数据
	fontBytes, err := ioutil.ReadFile(WinFontPath)
	if err != nil {
		return err
	}
	//载入字体数据
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return err
	}
	f := freetype.NewContext()
	//设置分辨率
	f.SetDPI(72)
	//设置字体
	f.SetFont(font)
	//设置尺寸
	f.SetFontSize(20)
	f.SetClip(img.Bounds())
	//设置输出的图片
	f.SetDst(img)
	//设置字体颜色(红色)
	f.SetSrc(image.NewUniform(color.RGBA{R: 255, A: 255}))
	//设置字体的位置
	pt := freetype.Pt(5, 20)
	if _, err := f.DrawString(pwd, pt); err != nil {
		return err
	}
	return png.Encode(file, img)
}
