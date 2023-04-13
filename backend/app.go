package backend

import (
	"Whisper_Record/backend/config"
	"Whisper_Record/backend/internal"
	"Whisper_Record/util"
	"Whisper_Record/util/file"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"os/exec"
)

// App struct
type App struct {
	ctx         context.Context
	init        *AppInit
	Log         *logrus.Logger
	Git         *internal.Git
	DB          *gorm.DB
	CfgFile     string
	LogFile     string
	DBFile      string
	onTop       bool
	isMaximised bool
	//hide 是否启用关闭至托盘的模式
	hide bool
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// OnStartup 应用启动
func (a *App) OnStartup(ctx context.Context) {
	//透明背景会存在最大化无法点击顶部状态栏的问题
	//hwnd := win.FindWindow(nil, syscall.StringToUTF16Ptr("微记"))
	//win.SetWindowLong(hwnd, win.GWL_EXSTYLE, win.GetWindowLong(hwnd, win.GWL_EXSTYLE)|win.WS_EX_LAYERED)
	// 初始化
	a.init = NewAppInit()
	a.init.Init()
	// 日志
	a.Log = a.init.Log
	a.Log.Info("OnStartup begin")
	// 获取上下文
	a.ctx = ctx
	a.onTop = config.OnTop

	if a.onTop {
		runtime.WindowSetAlwaysOnTop(a.ctx, a.onTop)
	}
	a.isMaximised = config.Maximise
	if a.isMaximised {
		runtime.WindowMaximise(a.ctx)
	}
	a.hide = config.Hide

	cfgPath := util.GetCfgPath()
	////创建系统托盘
	//trayOptions := options.SystemTray{
	//	LightModeIcon: nil,
	//	DarkModeIcon:  nil,
	//	Title:         "微记",
	//	Tooltip:       "微记",
	//	StartHidden:   false,
	//	Menu:          nil,
	//	OnLeftClick: func() {
	//		runtime.Show(a.ctx)
	//	},
	//	OnRightClick:       nil,
	//	OnLeftDoubleClick:  nil,
	//	OnRightDoubleClick: nil,
	//	OnMenuClose:        nil,
	//	OnMenuOpen:         nil,
	//}
	//tray := wails.Application.NewSystemTray(&trayOptions)
	////托盘图标
	//tray.SetIcons(&options.SystemTrayIcon{Data: config.AppIcon}, &options.SystemTrayIcon{Data: config.AppIcon})
	////托盘菜单
	//tray.SetMenu(a.Menu())
	//err := tray.Run()
	//if err != nil {
	//	return
	//}

	//-----------------
	// Git
	a.CfgFile = fmt.Sprintf(config.CfgFile, cfgPath)
	a.Git = &internal.Git{}
	if err := yaml.Unmarshal([]byte(file.Read(a.CfgFile)), a.Git); err != nil {
		a.Log.Errorf("OnStartup cfgfile err: %v", err)
	}
	a.Log.Infof("OnStartup cfg: %+v", a.Git)
	// 数据库文件
	a.DBFile = fmt.Sprintf(config.DBFile, cfgPath)
	// 如果无 Git 配置, 不处理数据库相关
	if a.Git.Repo == "" {
		return
	}
	//// 数据库处理
	//a.database()
	return
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// OnTopWindow 置顶窗口
func (a *App) OnTopWindow() {
	a.onTop = !a.onTop
	runtime.WindowSetAlwaysOnTop(a.ctx, a.onTop)
}

func (a *App) WindowIsOnToped() bool {
	return a.onTop
}

// MaximiseWindow 最大化窗口
func (a *App) MaximiseWindow() {
	if runtime.WindowIsMaximised(a.ctx) {
		runtime.WindowUnmaximise(a.ctx)
	} else {
		runtime.WindowMaximise(a.ctx)
	}

}

// CloseWindow 关闭窗口,
// isHide 传参用来标记是否: 不退出程序，仅隐藏桌面窗口
func (a *App) CloseWindow(isHide bool) {
	a.hide = isHide
	runtime.Quit(a.ctx)
}

// Menu 应用菜单
func (a *App) Menu() *menu.Menu {
	return menu.NewMenuFromItems(
		menu.Separator(),
		menu.SubMenu(config.App, menu.NewMenuFromItems(
			menu.Text("关于", nil, func(_ *menu.CallbackData) {
				_, err := a.diag(config.Description)
				if err != nil {
					return
				}
			}),
			menu.Text("检查更新", nil, func(_ *menu.CallbackData) {
				lastVersion := a.Git.GetLastVersion()
				needUpdate := config.Version < lastVersion
				msg := config.VersionNewMsg
				btns := []string{config.BtnConfirmText}
				if needUpdate {
					msg = fmt.Sprintf(config.VersionOldMsg, lastVersion)
					btns = []string{config.BtnConfirmText, config.BtnCancelText}
				}
				selection, err := a.diag(msg, btns...)
				if err != nil {
					return
				}
				if needUpdate && selection == config.BtnConfirmText {
					url := fmt.Sprintf(config.GitAppURL, lastVersion)
					runtime.BrowserOpenURL(a.ctx, url)
				}
			}),
		)),
		menu.SubMenu("帮助", menu.NewMenuFromItems(
			menu.Text(
				"打开配置文件",
				nil,
				func(_ *menu.CallbackData) {
					if !file.IsExist(a.CfgFile) {
						a.diag("文件不存在, 请先更新配置")
						return
					}
					_, err := exec.Command("open", a.CfgFile).Output()
					if err != nil {
						a.diag("操作失败: " + err.Error())
						return
					}
				},
			),
			menu.Text(
				"打开日志文件",
				nil,
				func(_ *menu.CallbackData) {
					if !file.IsExist(a.LogFile) {
						a.diag("文件不存在, 请先更新配置")
						return
					}
					_, err := exec.Command("open", a.LogFile).Output()
					if err != nil {
						a.diag("操作失败: " + err.Error())
						return
					}
				},
			),
			menu.Separator(),
			menu.Text(
				"打开应用主页",
				nil,
				func(_ *menu.CallbackData) {
					runtime.BrowserOpenURL(a.ctx, config.GitRepoURL)
				},
			),
			menu.Separator(),
			menu.Text("修复窗口位置异常", nil, func(_ *menu.CallbackData) {
				runtime.WindowCenter(a.ctx)
				runtime.Show(a.ctx)
			}),
		)),
		menu.Separator(),
		menu.Text("退出", nil, func(_ *menu.CallbackData) {
			a.CloseWindow(false)
		}),
	)
}

// OnDomReady ...
func (a *App) OnDomReady(ctx context.Context) {
	a.Log.Info("OnDomReady")
	return
}

// OnShutdown ...
func (a *App) OnShutdown(ctx context.Context) {
	a.Log.Info("OnShutdown")
	return
}

// OnBeforeClose ...
func (a *App) OnBeforeClose(ctx context.Context) bool {
	a.Log.Info("OnBeforeClose")

	if a.hide {
		runtime.Hide(a.ctx)
	}

	// 返回 true 将阻止程序关闭
	return a.hide
}

// ----------------------------------------------------------------

// migrate 数据同步
func (a *App) database() {
	a.Log.Info("OnStartup migrate begin")

	// 1. 如果 .db 存在, 初始化, 返回
	if file.IsExist(a.DBFile) {
		a.DB = internal.NewDB(a.DBFile)
		return
	}

	// 2. 校验远程 .db 是否存在, 存在直接同步后初始化, 返回
	dbContent := a.Git.GetContent(config.DBFile)
	if dbContent != "" {
		if err := file.Write(a.DBFile, dbContent); err != nil {
			a.Log.Errorf("OnStartup migrate sqlite err: %v", err)
			return
		}
		a.DB = internal.NewDB(a.DBFile)
		a.Log.Info("OnStartup migrate sqlite success")
		return
	}

	// 3. 校验远程 database.json 是否存在, 存在则迁移数据
	a.DB = internal.NewDB(a.DBFile)
	jsonContent := a.Git.GetContent("resource/database.json")
	if jsonContent != "" {
		list := make([]internal.File, 0)
		if err := json.Unmarshal([]byte(jsonContent), &list); err != nil {
			a.Log.Errorf("OnStartup migrate json err: %v", err)
			return
		}
		success := 0
		for i := len(list) - 1; i >= 0; i-- {
			res := a.DB.Create(&list[i])
			if res.Error == nil {
				success++
			}
		}
		a.Log.Infof("OnStartup migrate json success: %d", success)
	}

	return
}

// diag ...
func (a *App) diag(message string, buttons ...string) (string, error) {
	if len(buttons) == 0 {
		buttons = []string{
			config.BtnConfirmText,
		}
	}
	return runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          runtime.InfoDialog,
		Title:         config.Title,
		Message:       message,
		CancelButton:  config.BtnConfirmText,
		DefaultButton: config.BtnConfirmText,
		Buttons:       buttons,
		Icon:          config.AppIcon, // 这个是个虚设, Wails 内部代码没有对这个图标作实现
	})
}
