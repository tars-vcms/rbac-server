package main

import "github.com/TarsCloud/TarsGo/tars"

func main() {
	// Get server cfg
	cfg := tars.GetServerConfig()

	// New servant imp
	imp := new(logic.ConfigReaderImpl)
	// New servant
	app := new(vcms.ConfigReader)
	// Register Servant
	app.AddServant(imp, cfg.App+"."+cfg.Server+".configReaderObj")

	// Run application
	tars.Run()
}