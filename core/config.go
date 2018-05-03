package core

import (
	"os"
	"encoding/json"
)

var Conf *Config

type Config struct {
	PackageRemote string `json:"packageRemote"`
	PackageDir string `json:"packageDir"`
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
	err = json.NewDecoder(conf).Decode(&c)

	Conf = &c
	return err
}
