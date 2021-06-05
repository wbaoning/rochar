package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//定义配置结构听

var Conf = new(App)

type App struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Port      int32  `mapstructure:"port"`
	StartTime string `mapstructure:"startTime"`
	MachineID int64  `mapstructure:"machineID"`
	Log       *Log   `mapstructure:"log"`
	Mysql     *Mysql `mapstructure:"mysql"`
	Redis     *Redis `mapstructure:"Redis"`
}

type Log struct {
	Lever      string `mapstructure:"lever"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxAge     int    `mapstructure:"maxAge"`
	MaxBackups int    `mapstructure:"maxBackups"`
}

type Mysql struct {
	Host         string `mapstructure:"host"`
	Port         int32  `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Database     string `mapstructure:"databaase"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int32  `mapstructure:"port"`
	Pb       int    `mapstructure:"pb"`
	PoolSize int    `mapstructure:"poolsize"`
}

//加载配置项目
func Init() error {
	//viper.SetConfigName("conf")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath("./config")
	//viper.AddConfigPath(".")

	viper.SetConfigFile("./config/conf.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("load config init() %v", err)
	}

	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("序列号失败%v", err)
	}
	fmt.Println(Conf)
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("load change conf.port :%d", Conf.Port)

		if err := viper.Unmarshal(Conf); err != nil {
			//fmt.Printf("load change fail %v", err)
			fmt.Printf("load change conf.port :%d", Conf.Port)
		}
	})

	return err
}
