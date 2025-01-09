package main

import (
	"fmt"
	"os"

	"github.com/osamikoyo/void-init/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage - void-init [command] [package name]")
		return
	}
	switch os.Args[1]{
	case "init":
		utils.Init(os.Args[2])
	case "install":
		utils.Install(os.Args[2])
	}
}