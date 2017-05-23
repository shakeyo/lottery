package main

import (
	"server/conf"
	"server/modules/api"
	"server/modules/chat"
	"server/modules/dashboard"
	"server/modules/game"
	"server/modules/gate"
	"server/modules/world"

	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		gate.Module,
		world.Module,
		chat.Module,
		game.Module,
		dashboard.Module,
		api.Module,
	)
}
