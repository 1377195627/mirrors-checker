package main

import (
	"github.com/1377195627/mirrors-checker/core"
	"log"
)

func main() {
	if err := core.Init();err!=nil{
		log.Panicln(err)
	}

	if err := core.InitPackage();err!=nil{
		log.Panicln(err)
	}
}



