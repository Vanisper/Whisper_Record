package backend

import (
	"Whisper_Record/backend/config"
	"Whisper_Record/backend/internal"
	"Whisper_Record/util"
	"Whisper_Record/util/file"
	"Whisper_Record/util/image/ico"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/image/bmp"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
)

type AppInit struct {
	Log     *logrus.Logger
	LogFile string
}

// NewAppInit creates a new AppInit application struct
func NewAppInit() *AppInit {
	return &AppInit{}
}

// Init 初始化操作
func (init *AppInit) Init() {
	// 日志
	init.LogFile = fmt.Sprintf(config.LogFile, util.MkDir(util.GetLogPath()))
	init.Log = internal.NewLogger(init.LogFile)
	init.Log.Info("Init begin")
	//初始化软件图标
	message, err := initAppIcon()
	if err != nil {
		init.Log.Errorf("Init Icon err: %v", err)
	}
	init.Log.Infof("Init Icon: %+v", message)
	//初始化软件配置文件
	message, err = initAppConfig()
	if err != nil {
		init.Log.Errorf("Init Config err: %v", err)
	}
	init.Log.Infof("Init Config: %+v", message)
	//初始化数据库
	message, err = initAppDB()
	if err != nil {
		init.Log.Errorf("Init DB err: %v", err)
	}
	init.Log.Infof("Init DB: %+v", message)

	// Init End
	init.Log.Info("Init end")
}

func initAppIcon() (string, error) {
	if file.IsFile(config.AppIconName + ".ico") {
		return "已存在图标文件(icon.ico)", nil
	}
	// 获取文件二进制(会自动关闭文件)
	fs, err := os.Open(config.AppIconPath)
	var fSrc []byte
	var fileType string
	if err != nil {
		//如果可变图标丢失的话，获取内预置的图标
		fSrc = config.AppIcon
		//并且写入文件
		fileType = file.GetFileType(fSrc)
		fileName := config.AppIconName + "." + fileType
		fileDir := util.MkDir(util.GetCurrPath() + "\\" + config.AssetsPath)
		filePath := fileDir + "\\" + fileName
		err := os.WriteFile(filePath, fSrc, 0666)
		if err != nil {
			return "写出预设图标文件失败", err
		}
	} else {
		fSrc, _ = io.ReadAll(fs)
		fileType = file.GetFileType(fSrc)
	}

	//println(fileType)

	//再次读取文件，获得fs(因为上一次的fs关闭了)
	fs, err = os.Open(config.AssetsPath + config.AppIconName + "." + fileType)
	if err != nil {
		return "打开图标文件失败(获取文件fs时)", err
	}
	defer func(fs *os.File) {
		err = fs.Close()
		if err != nil {
			panic(err)
		}
	}(fs)
	//创建文件fs
	var newIcon *os.File
	//在根目录下创建软件图标
	if newIcon, err = os.Create(filepath.Join(util.GetCurrPath() + "\\icon.ico")); err != nil {
		panic(err)
	}
	defer func(newIcon *os.File) {
		err = newIcon.Close()
		if err != nil {
			panic(err)
		}
	}(newIcon)

	switch fileType {
	case "ico":
		img, _ := ico.Decode(fs)
		err = ico.Encode(newIcon, img)
		if err != nil {
			panic(err)
		}
	case "png":
		img, _ := png.Decode(fs)
		err = ico.Encode(newIcon, img)
		if err != nil {
			panic(err)
		}
	case "jpg":
		img, _ := jpeg.Decode(fs)
		err = ico.Encode(newIcon, img)
		if err != nil {
			panic(err)
		}
	case "gif":
		img, _ := gif.Decode(fs)
		err = ico.Encode(newIcon, img)
		if err != nil {
			panic(err)
		}
	case "bmp":
		// 第三方库
		img, _ := bmp.Decode(fs)
		err = ico.Encode(newIcon, img)
		if err != nil {
			panic(err)
		}
	default:
		panic("软件图标文件格式不在预设列表中: " + fileType)
	}
	return "创建图标文件成功", nil
}

func initAppConfig() (string, error) {
	return "", nil
}

func initAppDB() (string, error) {
	return "", nil
}
