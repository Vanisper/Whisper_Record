package backend

import (
	"Whisper_Record/backend/internal"
	"Whisper_Record/config"
	"Whisper_Record/util/file"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os/exec"

	"Whisper_Record/util"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx     context.Context
	Log     *logrus.Logger
	Git     *internal.Git
	DB      *gorm.DB
	CfgFile string
	LogFile string
	DBFile  string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// OnStartup 应用启动
func (a *App) OnStartup(ctx context.Context) {
	a.ctx = ctx
	cfgPath := util.GetCfgPath()
	logPath := util.GetLogPath()
	// 日志
	a.LogFile = fmt.Sprintf(config.LogFile, logPath)
	a.Log = internal.NewLogger(a.LogFile)
	a.Log.Info("OnStartup begin")
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

// Menu 应用菜单
func (a *App) Menu() *menu.Menu {
	return menu.NewMenuFromItems(
		menu.SubMenu(config.App, menu.NewMenuFromItems(
			menu.Text("关于", nil, func(_ *menu.CallbackData) {
				a.diag(config.Description)
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
			menu.Separator(),
			menu.Text("退出", keys.CmdOrCtrl("Q"), func(_ *menu.CallbackData) {
				runtime.Quit(a.ctx)
			}),
		)),
		menu.EditMenu(),
		menu.SubMenu("Help", menu.NewMenuFromItems(
			menu.Text(
				"打开配置文件",
				keys.Combo("C", keys.CmdOrCtrlKey, keys.ShiftKey),
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
				keys.Combo("L", keys.CmdOrCtrlKey, keys.ShiftKey),
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
				keys.Combo("H", keys.CmdOrCtrlKey, keys.ShiftKey),
				func(_ *menu.CallbackData) {
					runtime.BrowserOpenURL(a.ctx, config.GitRepoURL)
				},
			),
		)),
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
	// 返回 true 将阻止程序关闭
	return false
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
	})
}
