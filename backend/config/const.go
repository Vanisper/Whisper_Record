package config

import (
	"embed"
)

//go:embed icon.png
var AppIcon []byte

//go:embed icon.ico
var AppIconICO []byte

//go:embed ff.ico
var AppIconICOFF []byte

//go:embed test.ico
var AppIconICOTest []byte

//go:embed all:icons
var AppIcons embed.FS

// 应用
const (
	App         = "微记"
	AssetsPath  = "assets/"
	AppIconName = "icon"
	AppIconPath = AssetsPath + AppIconName + ".ico"
	Version     = "v0.1.0"
)

const (
	Title          = App + " " + Version
	Description    = "一个以富文本编辑器为主体的，个人内容管理系统。"
	VersionNewMsg  = "当前已经是最新版本!"
	VersionOldMsg  = "最新版本: %s, 是否立即更新?"
	BtnConfirmText = "确定"
	BtnCancelText  = "取消"
)

// 窗口尺寸
const (
	Width    = 875 / 1.25
	Height   = 625 / 1.25
	Maximise = false
	OnTop    = false
	Hide     = false
)

// 文件配置
var (
	CfgFile = "%s/config.yaml"
	LogFile = "%s/app.log"
	DBFile  = "%s/WR.db"
)
