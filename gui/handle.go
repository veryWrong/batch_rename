package gui

import (
	"errors"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type fixConf struct {
	lapse   int
	mold    string
	confine string
}

type renameConf struct {
	length  int
	confine string
}

type config struct {
	workDir    string
	fileType   []string
	renameType string
	fixConf    *fixConf
	renameConf *renameConf
}

const (
	prefix   = "加前缀"
	suffix   = "加后缀"
	rename   = "重命名"
	numMold  = "数字"
	wordMold = "字母"
	randMold = "随机"
)

func defaultConf() *config {
	return &config{
		workDir:    pwd,
		fileType:   []string{".jpg", ".png", ".jpeg"},
		renameType: prefix,
		fixConf: &fixConf{
			lapse:   6,
			mold:    numMold,
			confine: "1-12",
		},
		renameConf: &renameConf{
			length:  5,
			confine: defaultRandStr,
		},
	}
}

func (c *config) setWorkDir(workDir string) {
	c.workDir = workDir
}

func (c *config) setFileType(fileType string) {
	c.fileType = strings.Split(fileType, ",")
}

func (c *config) setRenameType(renameType string) {
	c.renameType = renameType
}

func (c *config) setFixLapse(lapse string) {
	c.fixConf.lapse, _ = strconv.Atoi(lapse)
}

func (c *config) setFixMold(mold string) {
	c.fixConf.mold = mold
}

func (c *config) setFixConfine(confine string) {
	c.fixConf.confine = confine
}

func (c *config) setRenameLength(length string) {
	c.renameConf.length, _ = strconv.Atoi(length)
}

func (c *config) setRenameConfine(confine string) {
	c.renameConf.confine = confine
}

func (c *config) do(w fyne.Window) {
	fs, err := ioutil.ReadDir(c.workDir)
	if err != nil {
		dialog.ShowError(errors.New("工作目录不存在"), w)
		return
	}
	files := make(map[string][]string, 0)
	readDir(fs, files, c.workDir)
	if len(files) == 0 {
		dialog.ShowInformation("", "未找到符合的文件", w)
		return
	}
	for dir, value := range files {
		switch c.renameType {
		case prefix, suffix:
			if c.fixConf.mold == numMold || c.fixConf.mold == wordMold {
				fix := parseConfine(w)
				for len(fix) < len(value)-1/c.fixConf.lapse {
					fix = append(fix, fix...)
				}
				for i, val := range value {
					if c.renameType == prefix {
						os.Rename(filepath.Join(dir, val), filepath.Join(dir, fix[i/c.fixConf.lapse]+"-"+val))
					} else {
						ext := filepath.Ext(val)
						name := strings.Replace(val, ext, "", 1)
						os.Rename(filepath.Join(dir, val), filepath.Join(dir, name+"-"+fix[i/c.fixConf.lapse])+ext)
					}
				}
			} else {
				for _, val := range value {
					if c.renameType == prefix {
						os.Rename(filepath.Join(dir, val), filepath.Join(dir, randomString(1, c.fixConf.confine)+"-"+val))
					} else {
						os.Rename(filepath.Join(dir, val), filepath.Join(dir, val+"-"+randomString(1, c.fixConf.confine)))
					}
				}
			}
		case rename:
			for _, val := range value {
				ext := filepath.Ext(val)
				filename := randomString(c.renameConf.length, c.renameConf.confine)
				os.Rename(filepath.Join(dir, val), filepath.Join(dir, filename+ext))
			}
		}
	}
	dialog.ShowInformation("success", "重命名完成", w)
}

func readDir(fs []os.FileInfo, files map[string][]string, dir string) {
	for _, item := range fs {
		if item.IsDir() {
			fs, err := ioutil.ReadDir(filepath.Join(dir, item.Name()))
			if err != nil {
				continue
			}
			readDir(fs, files, filepath.Join(dir, item.Name()))
		} else {
			ext := filepath.Ext(item.Name())
			for _, ft := range conf.fileType {
				if ft == ext {
					if val, ok := files[dir]; ok {
						val = append(val, item.Name())
						files[dir] = val
					} else {
						files[dir] = []string{item.Name()}
					}
					break
				}
			}
		}
	}
}

func parseConfine(w fyne.Window) []string {
	var (
		formatErr = errors.New("范围格式不正确")
		arr       = make([]string, 0)
	)
	strs := strings.Split(conf.fixConf.confine, "-")
	if len(strs) != 2 {
		dialog.ShowError(formatErr, w)
		return nil
	}
	if conf.fixConf.mold == numMold {
		min, err := strconv.Atoi(strs[0])
		if err != nil {
			dialog.ShowError(formatErr, w)
			return nil
		}
		max, err := strconv.Atoi(strs[1])
		if err != nil {
			dialog.ShowError(formatErr, w)
			return nil
		}
		for i := min; i <= max; i++ {
			arr = append(arr, strconv.Itoa(i))
		}
	} else {
		start := []rune(strings.ToUpper(strs[0]))[0]
		end := []rune(strings.ToLower(strs[1]))[0]
		if start > end {
			start, end = end, start
		}
		for i := start; i <= end; i++ {
			if 91 <= i && i <= 96 {
				continue
			}
			arr = append(arr, string(i))
		}
	}
	return arr
}

func randomString(length int, confine string) string {
	byts := []byte(confine)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, byts[r.Intn(len(byts))])
	}
	time.Sleep(time.Millisecond * 500)
	return string(result)
}
