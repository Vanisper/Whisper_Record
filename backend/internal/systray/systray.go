package systray

import (
	"Whisper_Record/backend/config"
	"embed"
	"io/fs"
	"time"

	"github.com/energye/systray"
)

var AddMenuItem func(title string, tooltip string) *systray.MenuItem
var AddMenuItemCheckbox func(title string, tooltip string, checked bool) *systray.MenuItem
var AddSeparator func()
var Quit func()

type sysTray struct {
	title   string
	icon    []byte
	OnClick func()
	onReady func()
	onExit  func()
}

func NewSysTray(
	title string,
	icon []byte,
	OnClick func(),
	onReady func(),
	onExit func(),
) *sysTray {
	AddMenuItem = systray.AddMenuItem
	AddMenuItemCheckbox = systray.AddMenuItemCheckbox
	AddSeparator = systray.AddSeparator
	Quit = systray.Quit
	return &sysTray{
		title,
		icon,
		OnClick,
		onReady,
		onExit,
	}
}

func (s *sysTray) Run() {
	systray.Run(func() {
		systray.SetTitle(s.title)
		systray.SetTooltip(s.title)
		systray.SetIcon(s.icon)
		systray.SetOnClick(s.OnClick)
		s.onReady()
	}, s.onExit)
}

func ChangeIcon(iconBytes []byte) {
	systray.SetIcon(iconBytes)
}

type iconPlayer struct {
	icons [][]byte
	stop  chan bool // 创建一个bool类型的通道，用来控制定时器和select语句
}

func NewIconPlayer(bashPath string, iconsFs embed.FS) *iconPlayer {
	var icons [][]byte
	err := fs.WalkDir(config.AppIcons, bashPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			imageBytes, err := config.AppIcons.ReadFile(path)
			if err != nil {
				return err
			}
			icons = append(icons, imageBytes)
		}
		return err
	})
	if err != nil {
		panic(err)
	}
	return &iconPlayer{
		icons: icons,
		stop:  make(chan bool),
	}
}

func (p *iconPlayer) Play(interval int) {
	go func() {
		// 创建定时器，每interval毫秒触发一次
		ticker := time.NewTicker(time.Duration(interval) * time.Millisecond)
		index := 0
		// goroutine中使用select监听定时器的到期和停止通道的信号
		defer ticker.Stop() // 安全地停止定时器
		for range ticker.C {
			select {
			case <-p.stop:
				return
			default:
				if index == int(len(p.icons)) {
					index = 0
				}
				ChangeIcon(p.icons[index])
				index++
			}
		}
	}()
}

func (p *iconPlayer) Stop() {
	p.stop <- true
}
