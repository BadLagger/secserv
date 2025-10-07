package main

import "secserv/utils"

func main() {
	log := utils.GlobalLogger().SetLevel(utils.Debug)
	defer log.Log(utils.Info, "App DONE!!!")

	appCfg := utils.CfgLoad("SecServ")

	log.Debug("AppName: %s", appCfg.AppName)

	log.Log(utils.Info, "App run!!!")
}
