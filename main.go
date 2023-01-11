package main

import (
	"Whisper_Record/backend"
	"Whisper_Record/backend/config"
	"Whisper_Record/util"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"log"
	"net/http"
	"strings"
)

//go:embed all:frontend/dist
var assets embed.FS

type FileLoader struct {
	http.Handler
}

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}

func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	requestedFilename := strings.TrimPrefix(req.URL.Path, "/") // 获取文件名称
	log.Println(requestedFilename)
	if requestedFilename == "" {
		// 判断是否是读取index.html
		requestedFilename = "index.html"
	}
	// 静态资源基础路径设置为程序运行位置
	assetBasePath := util.GetCurrPath() + "\\"
	assetFullPath := assetBasePath + requestedFilename
	log.Println(assetFullPath)
	http.ServeFile(res, req, assetFullPath)
	//fileData, err := os.ReadFile( assetFullPath)
	///*
	//	读取程序的运行路径下的对应资源文件
	//	这里的资源前缀可以自定义，如果是本地资源的话就使用os.ReadFile的方式获取文件，
	//	如果是网络资源的话，可以使用 http.ServeFile() 方法获取资源文件。
	//	当然，也可以统一成http.ServeFile，就是在本地以某个文件夹起一个server即可
	//	---
	//	该方法是在 assets 中获取不到资源的情况下执行的，这种场景一般出现在：前端网页中有些网络资源可能失效了，
	//	所以在assets出现404的情况下，就会执行handler方法，用于可能的补救措施，
	//	或者说是前端有意为之的，变相地将本地资源“转发”到前端所处端口下
	//*/
	//if err != nil {
	//	println(err.Error())
	//	res.WriteHeader(http.StatusBadRequest)
	//	res.Write([]byte(fmt.Sprintf("无法加载文件 %s", requestedFilename)))
	//}
	//res.Write(fileData)
}

func main() {
	app := backend.NewApp()

	err := wails.Run(&options.App{
		Title: "微记", // 标题
		//Width:             1100,  // 启动宽度
		//Height:            768,   // 启动高度
		MinWidth:          config.Width,  // 最小宽度
		MinHeight:         config.Height, // 最小高度
		HideWindowOnClose: false,         // 关闭的时候隐藏窗口
		StartHidden:       false,         // 启动的时候隐藏窗口 （建议生产环境关闭此项，开发环境开启此项，原因自己体会）
		AlwaysOnTop:       false,         // 窗口固定在最顶层
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: NewFileLoader(),
		},
		BackgroundColour: &options.RGBA{
			R: 27, G: 38, B: 54, A: 1,
		},
		LogLevel:      logger.DEBUG,      // 日志级别
		OnStartup:     app.OnStartup,     // 程序启动回调
		OnDomReady:    app.OnDomReady,    // 前端 dom 加载完成回调
		OnBeforeClose: app.OnBeforeClose, // 关闭应用程序之前回调
		OnShutdown:    app.OnShutdown,    // 程序退出回调
		//Menu: app.Menu(),
		//无边框
		Frameless: true,
		//指定html中的可拖动标签
		CSSDragProperty: "--wails-draggable",
		CSSDragValue:    "drag",
		Bind: []interface{}{
			app,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  true,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.DefaultAppearance,
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			About: &mac.AboutInfo{
				Title:   "微记",
				Message: "© 2022 Vangogh",
				Icon:    config.AppIcon,
			},
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
