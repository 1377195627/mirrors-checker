package core

import (
	"os/exec"
	"log"
	"strconv"
)

func Vercmp(ver1,ver2 string) int {
	cmd := exec.Command("vercmp",ver1,ver2)
	out,err:=cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return -1
	}

	if result,err:=strconv.Atoi(string(out));err!=nil{
		log.Println(err)
		return -1
	} else {
		return result
	}
}
