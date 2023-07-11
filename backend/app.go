package backend

import (
	"Whisper_Record/backend/config"
	"Whisper_Record/backend/internal"
	"Whisper_Record/backend/internal/systray"
	"Whisper_Record/backend/server"
	"Whisper_Record/util"
	"Whisper_Record/util/file"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"

	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx           context.Context
	init          *AppInit
	Log           *logrus.Logger
	Git           *internal.Git
	DB            *gorm.DB
	ServerPreview *server.Preview
	ServerPosts   *server.Posts
	PostsPath     string
	CfgFile       string
	LogFile       string
	DBFile        string
	onTop         bool
	isMaximised   bool
	isHide        bool
	//hide 是否启用关闭至托盘的模式
	hide bool
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) showWindow() {
	runtime.WindowShow(a.ctx)
	a.isHide = false
}

func (a *App) hiddenWindow() {
	runtime.WindowHide(a.ctx)
	a.isHide = true
}

func (a *App) toggleWindow() {
	if a.isHide {
		a.showWindow()
		return
	}
	a.hiddenWindow()
}

// OnStartup 应用启动
func (a *App) OnStartup(ctx context.Context) {
	//透明背景会存在最大化无法点击顶部状态栏的问题
	// hwnd := win.FindWindow(nil, syscall.StringToUTF16Ptr("微记"))
	// win.SetWindowLong(hwnd, win.GWL_EXSTYLE, win.GetWindowLong(hwnd, win.GWL_EXSTYLE)|win.WS_EX_LAYERED)
	// 初始化
	a.init = NewAppInit()
	a.init.Init()
	// 初始化博客文件夹
	a.ServerPosts = server.NewPosts()
	a.ServerPosts.Init()
	a.PostsPath = a.ServerPosts.Path
	// 初始化markdown-blog
	a.ServerPreview = server.NewPreview()
	a.ServerPreview.Init()
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
	a.isHide = config.Hide
	if a.isHide {
		a.hiddenWindow()
	}

	cfgPath := util.GetCfgPath()

	// 托盘
	sysTray := systray.NewSysTray("微记", config.AppIconICO, func() {
		a.toggleWindow()
		if a.isHide {
			systray.ChangeIcon(config.AppIconICOFF)
			return
		}
		systray.ChangeIcon(config.AppIconICO)
	}, func() {
		mCheck := systray.AddMenuItemCheckbox("切换", "切换", false)
		iconPlayer := systray.NewIconPlayer("icons", config.AppIcons)
		mCheck.Click(func() {
			if !mCheck.Checked() {
				mCheck.Check()
				iconPlayer.Play(50)
			} else {
				mCheck.Uncheck()
				iconPlayer.Stop()
			}
		})
		systray.AddSeparator()
		mQuit := systray.AddMenuItem("退出", "退出程序")
		mQuit.SetIcon(config.AppIconICO)
		mQuit.Click(func() {
			systray.Quit()
		})
	}, func() { a.CloseWindow(false) })
	sysTray.Run()

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

/**
goBindingsDir := filepath.Join(wailsjsbasedir, "go")
	err = os.RemoveAll(goBindingsDir)
	if err != nil {
		return err
	}
	_ = fs.MkDirs(goBindingsDir)

	err = bindings.GenerateGoBindings(goBindingsDir)
	if err != nil {
		return err
	}
*/
// CloseWindow 关闭窗口,
// isHide 传参用来标记是否: 不退出程序，仅隐藏桌面窗口
func (a *App) CloseWindow(isHide bool) {
	a.hide = isHide
	runtime.Quit(a.ctx)
}

// 关闭预览进程
func (a *App) KillPreview() {
	a.ServerPreview.Kill()
	log.Println("Kill ServerPreview")
}

// 重启预览进程
func (a *App) ReStartPreview() {
	a.ServerPreview.ReStart()
	log.Println("ReStart ServerPreview")
	a.OpenPreviewURL()
}

// 打开预览网页
func (a *App) OpenPreviewURL() {
	url := "http://127.0.0.1:" + (a.ServerPreview.Port)
	runtime.BrowserOpenURL(a.ctx, url)
	log.Println("Server Run In:", url)
}

type FileInfo struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Pathname    string      `json:"pathname"`
	BirthTime   interface{} `json:"birthTime"`
	IsFile      bool        `json:"isFile"`
	IsDirectory bool        `json:"isDirectory"`
	IsMarkdown  bool        `json:"isMarkdown"`
}

type FilesTree struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Pathname    string       `json:"pathname"`
	IsFile      bool         `json:"isFile"`
	IsDirectory bool         `json:"isDirectory"`
	IsMarkdown  bool         `json:"isMarkdown"`
	IsCollapsed bool         `json:"isCollapsed"`
	Files       []*FileInfo  `json:"files"`
	Folders     []*FilesTree `json:"folders"`
}

func getFilesTree(directory string) (*FilesTree, error) {
	filesTree := &FilesTree{
		Id:          filepath.Base(directory),
		Name:        filepath.Base(directory),
		Pathname:    directory,
		IsDirectory: true,
		IsFile:      false,
		IsMarkdown:  false,
		IsCollapsed: false,
	}

	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		filePathname := filepath.Join(directory, file.Name())

		if file.IsDir() {
			folder, err := getFilesTree(filePathname)
			if err != nil {
				return nil, err
			}
			filesTree.Folders = append(filesTree.Folders, folder)
		} else {
			fileExtname := filepath.Ext(file.Name())
			_fileInfo, _ := file.Info()
			fileInfo := &FileInfo{
				Id:          _fileInfo.Name(),
				Name:        _fileInfo.Name()[0 : len(_fileInfo.Name())-len(fileExtname)],
				Pathname:    filePathname,
				BirthTime:   _fileInfo.ModTime(),
				IsFile:      true,
				IsDirectory: false,
				IsMarkdown:  fileExtname == ".md",
			}
			filesTree.Files = append(filesTree.Files, fileInfo)
		}
	}

	sort.Slice(filesTree.Files, func(i, j int) bool {
		f1IsDir, f2IsDir := filesTree.Files[i].IsDirectory, filesTree.Files[j].IsDirectory
		if f1IsDir != f2IsDir {
			return !f1IsDir
		}

		return filesTree.Files[i].Name < filesTree.Files[j].Name
	})

	sort.Slice(filesTree.Folders, func(i, j int) bool {
		return filesTree.Folders[i].Name < filesTree.Folders[j].Name
	})

	return filesTree, nil
}

// 获取文章的文件列表
func (a *App) GetPostsList() string {
	fmt.Printf("a.PostsPath: %v\n", a.PostsPath)
	list, err := getFilesTree(a.PostsPath)
	if err != nil {
		panic(err)
	}
	list_json, _ := json.Marshal(list)
	a.Log.Info("获取文章的文件列表:", string(list_json))
	return string(list_json)
}

// 获取文章的文件内容
func (a *App) GetPostsContent(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(content)
}

// 写入内容至文件
func (a *App) WritePostsContent(path string, content string) bool {
	err := os.WriteFile(path, []byte(content), 0644)
	return err == nil
}

// 添加文件
func (a *App) AddFile(path string) bool {
	err := os.WriteFile(path, []byte(""), 0644)
	return err == nil
}

// 添加文件夹
func (a *App) AddFolder(path string) string {
	return util.MkDir(path)
}

// 删除文件/夹
func (a *App) RemoveFile(path string) bool {
	err := os.RemoveAll(path)
	return err == nil
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
		a.hiddenWindow()
	} else {
		a.KillPreview()
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
