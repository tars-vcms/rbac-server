package main

import (
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/tars-vcms/rbac-server/logic"
	"github.com/tars-vcms/rbac-server/tars-protocol/vcms"
)

func main() {

	cfg := tars.GetServerConfig()
	// RBAC服务实例
	imp := new(logic.RbacServerImpl)
	app := new(vcms.RbacServer)
	app.AddServant(imp, cfg.App+"."+cfg.Server+".RbacServerObj")



	tars.Run()
}