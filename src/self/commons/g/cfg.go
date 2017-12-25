/**
 * Created by shiyi on 2017/10/1.
 * Email: shiyi@fightcoder.com
 */

package g

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Run   RunConfig   `toml:"run"`
	Log   LogConfig   `toml:"log"`
	Mysql MysqlConfig `toml:"mysql"`
	Nsq   NsqConfig   `toml:"nsq"`
	Minio MinioConfig `toml:"minio"`
	Judge JudgeConfig `toml:"judge"`
}

type RunConfig struct {
	WaitTimeout int    `toml:"waitTimeout"`
	HTTPPort    int    `toml:"httpPort"`
	Mode        string `toml:"mode"`
	DataPath    string `toml:"dataPath"`
}

type LogConfig struct {
	Enable    bool   `toml:"enable"`
	Path      string `toml:"path"`
	Level     string `toml:"level"`
	RotatTime int    `toml:"rotatTime"`
	MaxAge    int    `toml:"maxAge"`
}

type MysqlConfig struct {
	MaxIdle int    `toml:"maxIdle"`
	MaxOpen int    `toml:"maxOpen"`
	Debug   bool   `toml:"debug"`
	WebAddr string `toml:"webAddr"`
}

type NsqConfig struct {
	Lookupds     []string `toml:"lookupds"`
	JudgeTopic   string   `toml:"judgeTopic"`
	JudgeChannel string   `toml:"judgeChannel"`
	HandlerCount int      `toml:"handlerCount"`
}

type MinioConfig struct {
	Endpoint        string `toml:"endpoint"`
	AccessKeyID     string `toml:"accessKeyID"`
	SecretAccessKey string `toml:"secretAccessKey"`
	Secure          bool   `toml:"secure"`
}

type JudgeConfig struct {
	Path string `toml:"path"`
}

func Conf() *Config {
	return config
}

var (
	ConfigFile string
	config     *Config
	configLock = new(sync.RWMutex)
)

func InitConfig(cfgFile string) {
	//配置文件路径是否为空
	if cfgFile == "" {
		log.Fatalln("config file not specified: use -c $filename")
	}

	//配置文件是否存在
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		log.Fatalln("config file specified not found:", cfgFile)
	}

	ConfigFile = cfgFile

	if bs, err := ioutil.ReadFile(cfgFile); err != nil {
		log.Fatalf("read config file failed: %s\n", err.Error())
	} else {
		if _, err := toml.Decode(string(bs), &config); err != nil {
			log.Fatalf("decode config file failed: %s\n", err.Error())
		} else {
			log.Printf("load config from %s\n", cfgFile)
		}
	}

	fmt.Printf("配置文件内容：%#v\n", config)
}
