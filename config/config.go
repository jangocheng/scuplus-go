package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Mysql 配置
type Mysql struct {
	Host     string
	User     string
	Password string
	DB       string
	Port     string
}

// CourseTask 任务配置文件
type CourseTask struct {
	StudentID int    `toml:"student_id"`
	Password  string `toml:"password"`
	PageNO    int    `toml:"page_no"`
}

// Wechat 微信配置
type Wechat struct {
	Appid                  string `toml:"appid"`
	Secret                 string `toml:"secret"`
	TemplateGrade          string `toml:"template_grade"`
	TemplateBook           string `toml:"template_book"`
	TemplateExam           string `toml:"template_exam"`
	TemplateFeedback       string `toml:"template_feedback"`
	TemplateLostFind       string `toml:"template_lost_find"`
	TemplateLostFindStatus string `toml:"template_lost_find_status"`
}

// Github 配置文件
type Github struct {
	AccessToken   string `toml:"access_token"`
	Repo          string `toml:"repo"`
	OwnerUser     string `toml:"owner_user"`
	WebhookSecret string `toml:"webhook_secret"`
}

// Redis redis配置
type Redis struct {
	IP   string `toml:"ip"`
	Port string `toml:"port"`
}

// Youtu 腾讯优图配置
type Youtu struct {
	AppID     string `toml:"app_id"`
	SecretID  string `toml:"secret_id"`
	SecretKey string `toml:"secret_key"`
	QQ        string `toml:"qq"`
}

// COS cos配置
type COS struct {
	AppID     string `toml:"app_id"`
	SecretID  string `toml:"secret_id"`
	SecretKey string `toml:"secret_key"`
	Bucket    string `toml:"bucket"`
	Folder    string `toml:"folder"`
}

// Config 对应config.yml文件的位置
type Config struct {
	Debug      bool `toml:"debug"`
	Port       string
	Secret     string
	JobWorkers int    `toml:"job_workers"`
	JwtSecret  string `toml:"jwt_secret"`
	Mysql      `toml:"mysql"`
	CourseTask `toml:"course_task"`
	Wechat     `toml:"wechat"`
	Github     `toml:"github"`
	Redis      `toml:"redis"`
	Youtu      `toml:"youtu"`
	COS        `toml:"cos"`
}

// config
var config Config

// 配置文件路径
var configFile = ""

func env(key, val string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}
	return val
}

// Get 获取config
func Get() Config {
	if config.Host == "" {
		// 默认配置文件在同级目录
		filepath := getPath(configFile)

		// 解析配置文件
		if _, err := toml.DecodeFile(filepath, &config); err != nil {
			log.Fatal("配置文件读取失败！", err)
		}
	}

	return config
}

// SetPath 设置Config文件的路径
func SetPath(path string) {
	configFile = path
}

// 获取文件路径
func getPath(path string) string {
	if path != "" {
		return path
	}

	path = os.Getenv("SCUPLUS_CONF")
	if path != "" {
		return path
	}

	// 获取当前环境
	env := os.Getenv("SCUPLUS_ENV")
	if env == "" {
		env = "develop"
	}

	// 默认配置文件在同级目录
	filepath := "config.toml"

	// 根据环境变量获取配置文件目录
	switch env {
	case "test":
		filepath = os.Getenv("GOPATH") + "/src/github.com/mohuishou/scuplus-go/config/" + filepath
	}
	return filepath
}
