package server

import (
	"Whisper_Record/util"
	"Whisper_Record/util/file"
	_ "embed"
	"log"
	"path"
)

//go:embed test.md
var testMD []byte

type Posts struct {
	Path string
}

func NewPosts() *Posts {
	return &Posts{}
}

func (p *Posts) Init() {
	p.Path = "./posts"
	if file.IsExist(p.Path) && file.IsDir(p.Path) {
		log.Println("文件夹存在")
	} else {
		util.MkDir(p.Path)
	}
	p.outputTestMD()
}

func (p *Posts) outputTestMD() {
	util.WriteToFile(string(testMD), path.Join(p.Path, "_test.md"), true)
}
