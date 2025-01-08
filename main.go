package main

import (
	"voice-assistant-manager/initialize"
)

func main() {
	initialize.InitConfig()
	initialize.GinInit()
}
