package git

import (
	"os/exec"
	"log"
	"os"
)

func Clone(url, path string) error {
	cmd := exec.Command("git","clone",url,path)
	log.Println("git clone "+url+"...")
	if err := cmd.Run();err!=nil{
		return err
	}
	return nil
}

func Pull(path string) error {
	pwd,err:=os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(pwd)

	os.Chdir(path)
	cmd := exec.Command("git","pull")
	log.Println("git pull...")
	if err := cmd.Run();err!=nil{
		return err
	}
	return nil
}