package server

import (
	"Whisper_Record/util"
	"bufio"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

//go:embed all:lib/markdown-blog
var assets embed.FS

type Preview struct {
	PID        int
	Cmd        *exec.Cmd
	CommandExe string
	CommandArg [3]string
	Port       string
}

func NewPreview() *Preview {
	return &Preview{}
}

func (p *Preview) Init() {
	path := "lib/markdown-blog"
	libExe := fmt.Sprintf("%s\\%s\\%s\\markdown-blog", path, runtime.GOOS, runtime.GOARCH)
	if runtime.GOOS == "windows" {
		libExe += ".exe"
	}
	libConfig := fmt.Sprintf("%s\\config.yml", path)
	p.CommandExe = util.GetCurrPath() + "\\" + libExe
	p.CommandArg = [...]string{"web", "--config", (util.GetCurrPath() + "\\" + libConfig)}
	// 判断可执行文件是否存在
	_, err := os.Stat(util.GetCurrPath() + "\\" + libExe)
	if err == nil { // 文件存在
		fmt.Println("文件已存在")
	} else if os.IsNotExist(err) { // 文件不存在
		err = os.RemoveAll(util.GetCurrPath() + "\\" + path)
		if err != nil {
			log.Println("删除文件失败", err)
			err = nil
			err = os.Chmod(util.GetCurrPath()+"\\"+path, 0777)
			if err != nil {
				panic(err)
			}
			err = os.RemoveAll(util.GetCurrPath() + "\\" + path)
			if err != nil {
				panic(err)
			}
		}
		folderData, err := deepWalk(path, assets)
		if err != nil {
			panic(err)
		}
		jsonData, err := json.Marshal(folderData)
		if err != nil {
			panic(err)
		}
		log.Println(string(jsonData))
	} else { // 其他错误
		fmt.Println("判断文件夹是否存在时发生错误：", err)
		panic(err)
	}
	p.start()
}

func (p *Preview) Kill() error {
	return p.Cmd.Process.Kill()
}

func (p *Preview) ReStart() {
	p.Kill()
	p.start()
}

func (p *Preview) start() {
	p.Cmd = exec.Command(p.CommandExe, p.CommandArg[0], p.CommandArg[1], p.CommandArg[2])

	go func() {
		stdout, err := p.Cmd.StdoutPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, "error=>", err.Error())
		}
		if runtime.GOOS == "windows" {
			p.Cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // 隐藏命令行窗口
		}
		if err := p.Cmd.Start(); err != nil {
			log.Fatalf("exe 启动失败: %v", err)
		}
		p.Port = "5006"
		go func() {
			// 输出命令信息
			scanner := bufio.NewScanner(stdout)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}()

		log.Println("Process started with PID:", p.Cmd.Process.Pid, p.Cmd.String())
		go func() {
			// 等待命令执行完成
			if err := p.Cmd.Wait(); err != nil {
				log.Printf("命令执行失败: %v", err)
			}
			log.Println("命令执行成功。")
		}()
	}()
}

type FileData struct {
	Path    string `json:"Path"`
	Content string `json:"Content"`
}

type FolderData struct {
	Path string `json:"Path"`
}

type FoldersData struct {
	Path    string
	Files   []*FileData
	Folders []*FolderData
}

func deepWalk(path string, assets embed.FS) (FoldersData, error) {
	var folderData FoldersData
	folderData.Path = path

	err := fs.WalkDir(assets, path, func(fpath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			// 创建文件夹(深度)
			util.MkDir(util.GetCurrPath() + "\\" + fpath)
			folderData.Folders = append(folderData.Folders, &FolderData{
				Path: fpath,
			})
		} else {
			fContent, err := assets.ReadFile(fpath)
			if err != nil {
				return err
			}
			err = os.WriteFile(util.GetCurrPath()+"\\"+fpath, fContent, 0666)
			if err != nil {
				return err
			}

			folderData.Files = append(folderData.Files, &FileData{
				Path:    fpath,
				Content: "",
			})
		}
		return nil
	})

	return folderData, err
}
