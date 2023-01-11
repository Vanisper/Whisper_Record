package config

var (
	Owner       = "Vanisper"
	Repo        = "Whisper_Record"
	Email       = "273266469@qq.com"
	AccessToken = ""
)

// Git 配置
// https://docs.github.com/cn/rest/reference/repos#contents
const (
	GitApiURL   = "https://api.github.com/repos/%s/%s/contents/%s"
	GitTagURL   = "https://api.github.com/repos/%s/%s/tags"
	GitFileURL  = "https://cdn.jsdelivr.net/gh/%s/%s/%s"
	GitDBFile   = "resource/WR.db"
	GitFilePath = "resource/%s/%s%s"
	GitMarkFile = "mark"
	GitMessage  = "upload by vangogh"
	GitRepoURL  = "https://github.com/%s/%s"
	GitAppURL   = GitRepoURL + "/releases/tag/%s"
)
