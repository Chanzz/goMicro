package common

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/go-ini/ini"
)

func GetConfigMap(section string) map[string]string {
	var appPath string
	var err error
	if appPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	}

	absPath := path.Join(appPath, CONFIG_FILE)
	// 读取配置 .ini 文件
	cfg, err := ini.Load(absPath)
	if err != nil {
		log.Print(err)
	}

	// 读取指定 section 内容，放回 map 类型的数据
	sec, err := cfg.GetSection(section)
	if err != nil {
		log.Print(err)
	}
	return sec.KeysHash()
}
