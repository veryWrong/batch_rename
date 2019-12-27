package gui

type fixConf struct {
	lapse   uint
	mold    string
	confine string
}

type renameConf struct {
	length  uint
	confine string
}

type config struct {
	wordDir    string
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
		wordDir:    pwd,
		fileType:   []string{".jpg", ".png", ".jpeg"},
		renameType: prefix,
		fixConf:    nil,
		renameConf: nil,
	}
}
