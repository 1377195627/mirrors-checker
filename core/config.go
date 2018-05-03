package core

import (
	"os"
	"encoding/json"
)

var Conf *Config

type Config struct {
	PackageRemote string `json:"packageRemote"`
	PackageDir string `json:"packageDir"`
	PackageTmp string `json:"packageTmp"`
	Packages []Package `json:"package"`
}

type Package struct {
	Name string `json:"name"`
	IgnoreVer []string `json:"ignoreVer"`
	UpstreamURL string `json:"upstreamURL"`
}


func InitConfig() error {
	var c Config

	conf, err := os.Open("config.json")
	if err != nil {
		return err
	}
	if err = json.NewDecoder(conf).Decode(&c);err!=nil{
		return err
	}

	Conf = &c

	if _,err := os.Stat(c.PackageTmp);err!=nil{
		if os.IsNotExist(err) {
			os.MkdirAll(c.PackageTmp,0755)
		}
	}

	return nil
}
