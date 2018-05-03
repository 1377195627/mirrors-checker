package core

import (
	"errors"
	"github.com/1377195627/mirrors-checker/checker"
	"github.com/1377195627/mirrors-checker/git"
	"log"
	"os"
)

var (
	github *checker.Github
)

func Init() error {
	if err := InitConfig(); err != nil {
		return err
	}

	github = checker.NewGithub()

	return nil
}

func InitPackage() error {
	if _, err := os.Stat(Conf.PackageDir); err != nil {
		if os.IsNotExist(err) {
			return git.Clone(Conf.PackageRemote, Conf.PackageDir)
		}
		return err
	}

	if _, err := os.Stat(Conf.PackageDir + "/.git"); err != nil {
		if os.IsNotExist(err) {
			return errors.New("package dir is exist but not has .git")
		}
		return err
	}

	log.Println("update package dir")
	return git.Pull(Conf.PackageDir)
}

func CheckAll() error {
	return nil
}
