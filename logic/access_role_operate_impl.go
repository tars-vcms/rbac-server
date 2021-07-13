package logic

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/basef"
	"github.com/tars-vcms/rbac-server/repo/mysql"
	"github.com/tars-vcms/rbac-server/tars-protocol/vcms"
)

type AccessRoleOperateImpl struct {}

func (a AccessRoleOperateImpl) CreateAccessRole(req *vcms.CreateAccessRoleReq, rsp *vcms.AccessRoleInfo) (ret int32, err error) {
	thisRole, rowCount, err := mysql.AccessRoleSqlImpl.Create(req.RoleIdent, req.RoleName, req.Access, req.Comment)
	if rowCount != 1{
		rsp.RoleId = int32(thisRole.ID)
		rsp.RoleIdent = thisRole.Ident
		rsp.RoleName = thisRole.Name
		rsp.Comment = thisRole.Comment
		rsp.Access = thisRole.Access
		return basef.TARSSERVERSUCCESS, nil
	}else{
		return basef.TARSSERVERSUCCESS, err
	}
}

func (a AccessRoleOperateImpl) GetAccessRole(req *vcms.GetAccessRoleReq, rsp *[]vcms.AccessRoleInfo) (ret int32, err error) {
	whereConditionString := fmt.Sprintf("`id` = %d AND `ident` LIKE %%%s%% AND `name` LIKE %%%s%% AND `access` LIKE %%%s%% AND `comment` LIKE %%%s%%", req.WhereCondition.RoleId, req.WhereCondition.RoleIdent, req.WhereCondition.RoleName, req.WhereCondition.Access, req.WhereCondition.Comment)
	thisRoles, _, err := mysql.AccessRoleSqlImpl.Read(whereConditionString, int(req.Limit.PageSize), int(req.Limit.Offset))
	if err != nil {
		for _, v := range thisRoles {
			*rsp = append(*rsp, vcms.AccessRoleInfo{
				RoleId:    int32(v.ID),
				RoleIdent: v.Ident,
				RoleName:  v.Name,
				Access:    v.Access,
				Comment:   v.Comment,
			})
		}
		return basef.TARSSERVERSUCCESS, nil
	}else{
		return basef.TARSSERVERSUCCESS, err
	}
}

func (a AccessRoleOperateImpl) UpdateAccessRole(req *vcms.AccessRoleInfo, rsp *vcms.AccessRoleInfo) (ret int32, err error) {
	thisRole, rowCount, err := mysql.AccessRoleSqlImpl.Update(int(req.RoleId), req.RoleIdent, req.RoleName, req.Access, req.Comment)
	if rowCount != 1{
		rsp.RoleId = int32(thisRole.ID)
		rsp.RoleIdent = thisRole.Ident
		rsp.RoleName = thisRole.Name
		rsp.Comment = thisRole.Comment
		rsp.Access = thisRole.Access
		return basef.TARSSERVERSUCCESS, nil
	}else{
		return basef.TARSSERVERSUCCESS, err
	}
}

func (a AccessRoleOperateImpl) DeleteAccessRole(req *vcms.AccessRoleInfo, rsp *vcms.AccessRoleInfo) (ret int32, err error) {
	thisRole, rowCount, err := mysql.AccessRoleSqlImpl.Delete(int(req.RoleId))
	if rowCount != 1{
		rsp.RoleId = int32(thisRole.ID)
		rsp.RoleIdent = thisRole.Ident
		rsp.RoleName = thisRole.Name
		rsp.Comment = thisRole.Comment
		rsp.Access = thisRole.Access
		return basef.TARSSERVERSUCCESS, nil
	}else{
		return basef.TARSSERVERSUCCESS, err
	}
}