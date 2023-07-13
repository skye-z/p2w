package main

import (
	"p2w/cmd"
	"p2w/global"
)

func main() {
	global.InitConfig()
	cmd.Execute()
}
