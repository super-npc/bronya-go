package conf

import (
	"log"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Settings = new(AppConfig)

type AppConfig struct {
	Name string `mapstructure:"name"`
	Mode string `mapstructure:"mode"`
	Port int    `mapstructure:"port"`

	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*Etcd        `mapstructure:"etcd"`
}

type Etcd struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenCount int    `mapstructure:"max_open_count"`
	MaxIdleCount int    `mapstructure:"max_idle_count"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"dbname"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleCount int    `mapstructure:"min_idle_count"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

func InitSettings() (err error) {
	// 获取运行环境
	mode := os.Getenv("APP_MODE")
	if mode == "" {
		mode = "dev" // 默认为开发环境
	}

	var configFile string
	switch mode {
	case "prd":
		configFile = "./resources/application-prd.yml"
	case "dev":
		configFile = "./resources/application-dev.yml"
	default:
		configFile = "./resources/application.yml"
	}

	// 设置配置文件路径
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	// 设置环境变量前缀，支持环境变量覆盖配置文件
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()

	// 替换配置中的点分隔符为下划线，使环境变量更易读
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		log.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return err
	}

	// 把读取到的配置信息反序列化到 Settings 变量中
	if err := viper.Unmarshal(Settings); err != nil {
		log.Printf("viper.Unmarshal failed, err:%v\n", err)
		return err
	}

	// 设置默认模式
	if Settings.Mode == "" {
		Settings.Mode = mode
	}

	// 设置默认日志路径
	if Settings.LogConfig.Filename == "" {
		Settings.LogConfig.Filename = "./logs/app.log"
	}
	if Settings.LogConfig.MaxSize == 0 {
		Settings.LogConfig.MaxSize = 100
	}
	if Settings.LogConfig.MaxAge == 0 {
		Settings.LogConfig.MaxAge = 30
	}
	if Settings.LogConfig.MaxBackups == 0 {
		Settings.LogConfig.MaxBackups = 7
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("配置文件已修改")
		if err := viper.Unmarshal(Settings); err != nil {
			log.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return
}
