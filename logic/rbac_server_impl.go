package logic

import (
	"github.com/TarsCloud/TarsGo/tars/protocol/res/basef"
	"github.com/tars-vcms/rbac-server/entity"
	"github.com/tars-vcms/rbac-server/repo/mysql"
	"github.com/tars-vcms/rbac-server/tars-protocol/vcms"
)

type RbacServerImpl struct {}

func (RbacServerImpl) GetAccessAbilityByIdent(req string, rsp *vcms.AccessAbilityRsp) (ret int32, err error) {
	var thisRole mysql.AccessRoleModel
	var rowCount int
	thisRole, rowCount, err = mysql.AccessRoleSqlImpl.FindByIdent(req)
	if err == nil && rowCount == 1 {
		rsp.Name = thisRole.Name
		rsp.Ident = thisRole.Ident
		rsp.AccessList = thisRole.Access
		rsp.Comment = thisRole.Comment
		//TODO: 错误返回
		return basef.TARSSERVERSUCCESS, nil
	}else{
		//TODO: 错误返回
		rsp.Name = entity.InvalidRoleName
		rsp.Ident = entity.InvalidRoleIdent
		return basef.TARSSERVERSUCCESS, nil
	}
}

func (RbacServerImpl) GetAccessPointInfo(reqAccessPointId int32, rsp *vcms.AccessPointInfo) (ret int32, err error) {
	var thisPoint mysql.AccessPointModel
	var rowCount int
	thisPoint, rowCount, err = mysql.AccessPointSqlImpl.Find(int(reqAccessPointId))
	if err == nil && rowCount == 1 {
		rsp.Comment = thisPoint.Comment
		rsp.AccessIdent = thisPoint.Ident
		rsp.AccessId = int32(thisPoint.ID)
		//TODO: 错误返回
		return basef.TARSSERVERSUCCESS, nil
	}else{
		//TODO: 错误返回
		rsp.AccessIdent = entity.InvalidPointIdent
		return basef.TARSSERVERSUCCESS, nil
	}
}

func (RbacServerImpl) IsValidAccess(req *vcms.IsValidAccessReq, rsp *bool) (ret int32, err error) {
	//为后续扩展，本方法实现三步走。
	var thisPoint mysql.AccessPointModel
	var rowCount int
	thisPoint, rowCount, err = mysql.AccessPointSqlImpl.FindByIdent(req.AccessIdent)
	if err != nil || rowCount != 1 { //找不到指定ident的接入点
		*rsp = false
		return basef.TARSSERVERSUCCESS, nil
	}
	var thisRole mysql.AccessRoleModel
	thisRole, rowCount, err = mysql.AccessRoleSqlImpl.FindByIdent(req.RoleIdent)
	if err != nil || rowCount != 1 { //找不到指定ident的角色
		*rsp = false
		return basef.TARSSERVERSUCCESS, nil
	}
	//判断AccessRole的ID是否在AccessList中
	var isValid bool
	isValid, rowCount, err = mysql.AccessRoleSqlImpl.HaveAccessPointId(thisRole.ID, thisPoint.ID)
	if err == nil && rowCount == 1 {
		*rsp = isValid
		return basef.TARSSERVERSUCCESS, nil
	}else{
		*rsp = false
		return basef.TARSSERVERSUCCESS, nil
	}

}

