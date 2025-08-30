package conf

import (
	"fmt"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
	Port      int    `mapstructure:"port"`

	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"dbname"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
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
		mode = "develop" // 默认为开发环境
	}

	var configFile string
	switch mode {
	case "production":
		configFile = "./resources/config.production.yaml"
	case "develop":
		configFile = "./resources/config.yaml"
	default:
		configFile = "./resources/config.yaml"
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
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return err
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		return err
	}

	// 设置默认模式
	if Conf.Mode == "" {
		Conf.Mode = mode
	}

	// 设置默认日志路径
	if Conf.LogConfig.Filename == "" {
		Conf.LogConfig.Filename = "./logs/app.log"
	}
	if Conf.LogConfig.MaxSize == 0 {
		Conf.LogConfig.MaxSize = 100
	}
	if Conf.LogConfig.MaxAge == 0 {
		Conf.LogConfig.MaxAge = 30
	}
	if Conf.LogConfig.MaxBackups == 0 {
		Conf.LogConfig.MaxBackups = 7
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return
}
